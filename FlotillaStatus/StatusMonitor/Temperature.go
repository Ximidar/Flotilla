/*
* @Author: Ximidar
* @Date:   2018-12-18 16:10:53
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-28 14:30:26
 */

package StatusMonitor

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

// Temperature will hold the value of the temperature at a certain time
type Temperature struct {
	Temp   map[string]float64 `json:"current_temp"`
	Target map[string]float64 `json:"target_temp"`
	Time   time.Time          `json:"timecode"`
}

// NewTemperature constructs a new temperature
func NewTemperature() *Temperature {
	temp := new(Temperature)
	temp.Temp = make(map[string]float64)
	temp.Target = make(map[string]float64)
	temp.Time = time.Now()
	return temp
}

// NewTemperatureFromMSG will construct a temperature from a Nats MSG
func NewTemperatureFromMSG(msg *nats.Msg) (*Temperature, error) {
	bTemp := new(Temperature)
	err := json.Unmarshal(msg.Data, bTemp)
	if err != nil {
		return nil, err
	}
	return bTemp, nil
}

// PackageJSON will package the temperature struct into JSON bytes
func (t *Temperature) PackageJSON() ([]byte, error) {
	tempByte, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return tempByte, nil
}

// OrganizeByTime will sort a temperature slice by the time it was taken
type OrganizeByTime []Temperature

func (al OrganizeByTime) Len() int {
	return len(al)
}
func (al OrganizeByTime) Swap(i, j int) {
	al[i], al[j] = al[j], al[i]
}
func (al OrganizeByTime) Less(i, j int) bool {
	return al[i].Time.Unix() < al[j].Time.Unix()
}

//////////////////////////////////////////////////////////
//					TemperatureMonitor                  //
//////////////////////////////////////////////////////////

// PublishTemperature is an outside function given to TemperatureMonitor
// to publish a temperature update when it gets one
type PublishTemperature func(t *Temperature)

// TemperatureMonitor will monitor the output of the comm for any temperature updates
type TemperatureMonitor struct {
	CurrentTemperature *Temperature
	TemperatureHistory []*Temperature
	PublishTemperature PublishTemperature

	// Precompiled Regex
	FindT       *regexp.Regexp
	FindTExtra  *regexp.Regexp
	FindB       *regexp.Regexp
	FindBExtra  *regexp.Regexp
	FindAt      *regexp.Regexp
	FindAtExtra *regexp.Regexp
	FindNums    *regexp.Regexp
}

// NewTemperatureMonitor will construct a new TemperatureMonitor object
func NewTemperatureMonitor(pubtemp PublishTemperature) (*TemperatureMonitor, error) {
	tm := new(TemperatureMonitor)
	tm.PublishTemperature = pubtemp
	err := tm.CompileRegex()
	if err != nil {
		return nil, err
	}

	tm.ResetTemps()

	return tm, nil
}

// ResetTemps will empty the temperature history and current temp
func (tm *TemperatureMonitor) ResetTemps() {
	tm.CurrentTemperature = NewTemperature()
	tm.TemperatureHistory = make([]*Temperature, 0, 100)
}

// CompileRegex will compile all functions used to find info
func (tm *TemperatureMonitor) CompileRegex() error {
	var err error
	tm.FindT, err = regexp.Compile("T:([0-9.-]*[0-9]+)([/ ]+)([0-9.-]*[0-9]+)")
	tm.FindTExtra, err = regexp.Compile("T([0-9]+):([0-9.-]*[0-9]+)([/ ]+)([0-9.-]*[0-9]+)")
	tm.FindB, err = regexp.Compile("B:([0-9.-]*[0-9]+)([/ ]+)([0-9.-]*[0-9]+)")
	tm.FindBExtra, err = regexp.Compile("B([0-9]+):([0-9.-]*[0-9]+)([/ ]+)([0-9.-]*[0-9]+)")

	tm.FindNums, err = regexp.Compile("([0-9.-]*[0-9]+)")

	// TODO add the regex for finding the @
	return err
}

// UpdateTemperature will be given a raw temperature line and it will extract the current temperature
func (tm *TemperatureMonitor) UpdateTemperature(rawTemperatureLine string) {
	//example input: ok T:21.25 /0.00 B:22.50 /0.00 @:0 B@:0
	temporaryTemp := NewTemperature()

	// Find the T temps
	Ttemp, err := tm.grabCurrentAndSet(tm.FindT, rawTemperatureLine)
	if err != nil {
		//fmt.Println(err)
	} else {
		// Populate the Ttemps
		temporaryTemp.Temp["T"] = Ttemp[0]
		temporaryTemp.Target["T"] = Ttemp[1]
	}

	// find the Btemps
	Btemp, err := tm.grabCurrentAndSet(tm.FindB, rawTemperatureLine)
	if err != nil {
		//fmt.Println(err)
	} else {
		// Populate the Btemps
		temporaryTemp.Temp["B"] = Btemp[0]
		temporaryTemp.Target["B"] = Btemp[1]
	}
	//TODO expand for multiple heads and beds

	// Update Temps
	tm.TemperatureHistory = tm.AppendFixed(tm.TemperatureHistory, tm.CurrentTemperature)
	tm.CurrentTemperature = temporaryTemp
	go tm.PublishTemperature(tm.CurrentTemperature)
}

// AppendFixed will append a value to a slice and shift all of the slices down
// if there is not enough room in the slice
func (tm *TemperatureMonitor) AppendFixed(slicer []*Temperature, data ...*Temperature) []*Temperature {
	caps := cap(slicer)

	for _, d := range data {
		if len(slicer) < caps {

			slicer = append(slicer, d)
		} else {
			// Shift down and put in last place
			slicer = slicer[:copy(slicer[0:], slicer[1:])]
			slicer = append(slicer, d)
		}
	}
	return slicer
}

// grabCurrentAndSet will attempt to use the matcher to pull out the current and set temps for the selected matcher
func (tm *TemperatureMonitor) grabCurrentAndSet(matcher *regexp.Regexp, rawline string) ([]float64, error) {
	if matcher.MatchString(rawline) {
		found := matcher.Find([]byte(rawline))
		nums := tm.FindNums.FindAll(found, -1)
		// We should only find 2 numbers at this point
		if len(nums) == 2 {
			current, err := strconv.ParseFloat(string(nums[0]), 64)
			target, err := strconv.ParseFloat(string(nums[1]), 64)
			if err == nil {
				return []float64{current, target}, nil

			}
			return nil, err

		}
		return nil, errors.New("more than 2 nums")
	}
	return nil, errors.New("no matches")
}

// GetTempCommand will return the correct command to set a temperature for the designated heater
func (tm *TemperatureMonitor) GetTempCommand(heater string, temp uint64) (string, error) {
	// Make sure heater string is close to valid
	if len(heater) > 3 {
		return "", errors.New("heater does not exist")
	}

	// Select heater
	if strings.Contains(heater, "T") {
		// Create command
		command := fmt.Sprintf("%v %v S%v", "M104", heater, temp)
		return command, nil
	} else if strings.Contains(heater, "B") {
		// Create command
		command := fmt.Sprintf("%v S%v", "M140", temp)
		return command, nil
	}
	return "", errors.New("heater does not exist")

}

// GetTempHistory will return the full history of temperature
func (tm *TemperatureMonitor) GetTempHistory() []*Temperature {
	return tm.TemperatureHistory
}

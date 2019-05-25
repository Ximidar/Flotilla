/*
* @Author: Ximidar
* @Date:   2019-02-16 20:34:58
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-26 12:57:50
 */

// Package Config will be used to load the main configuration file and
// send configurations over NATS
package Config

import (
	"encoding/json"
	"fmt"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/ximidar/Flotilla/DataStructures/ConfigStructures"
	"github.com/ximidar/Flotilla/NodeLauncher/RootFolder"
)

// Loader will use variables from the RootFolder to then
// Load a the config and share it with other nodes.
type Loader struct {
	RootFolder *RootFolder.RootFolder
	Config     *viper.Viper
	mux        sync.Mutex
}

// NewLoader will return a Loader object
func NewLoader(rf *RootFolder.RootFolder) (*Loader, error) {
	loader := new(Loader)
	loader.RootFolder = rf

	// Create a new Viper Config
	loader.Config = viper.New()
	err := loader.setupViper()
	if err != nil {
		return nil, err
	}
	err = loader.AddRootVars()
	if err != nil {
		return nil, err
	}

	return loader, nil
}

func (loader *Loader) setupViper() error {

	// Split Config path into name and directory
	basename := path.Base(loader.RootFolder.ConfigFile)
	noExt := strings.TrimSuffix(basename, filepath.Ext(basename))
	configPath := strings.TrimSuffix(loader.RootFolder.ConfigFile, basename)

	// Debug Message
	fmt.Printf("Adding File: %v from Path: %v to Viper\n", basename, configPath)

	// Add Basename to viper
	loader.Config.SetConfigName(noExt)

	// Add Config Path
	loader.Config.AddConfigPath(configPath)

	// Read in the config
	err := loader.Config.ReadInConfig()
	if err != nil {
		return err
	}

	// Setup Watch For Changes
	loader.Config.WatchConfig()
	loader.Config.OnConfigChange(loader.onConfigChange)

	return nil
}

func (loader *Loader) onConfigChange(fsnotify.Event) {
	fmt.Println("Config Changed!")
}

// GetVars will be used to get values from the config
func (loader *Loader) GetVars(configReply ConfigStructures.ConfigReply) ([]byte, error) {
	// Figure out what data to query
	if configReply.ToJSON {
		varmap := make(map[string]interface{})

		for _, value := range configReply.VarNames {
			loader.mux.Lock()
			temp := loader.Config.Get(value)
			loader.mux.Unlock()

			varmap[value] = temp
		}

		mData, err := json.Marshal(varmap)
		if err != nil {
			// respond with no data
			return []byte(nil), err
		}

		return mData, nil
	}

	// Respond with a map to bytes
	varmap := make(map[string][]byte)
	for _, value := range configReply.VarNames {
		loader.mux.Lock()
		temp := loader.Config.Get(value)
		loader.mux.Unlock()

		val, err := ConfigStructures.EncodeBytes(temp)
		if err != nil {
			fmt.Println("Couldn't Package", value, "Because", err)
		}

		varmap[value] = val
	}

	bvarmap, err := ConfigStructures.EncodeBytes(varmap)
	if err != nil {
		// respond with no data
		return []byte(nil), err
	}

	return bvarmap, nil
}

// AddRootVars will take in the root folder and write the paths to the path variable
func (loader *Loader) AddRootVars() error {

	loader.mux.Lock()
	fmt.Println("Writing path to Config!")
	loader.Config.Set("path", make(map[string]string))
	loader.Config.Set("path.root", loader.RootFolder.RootPath)
	loader.Config.Set("path.gcode", loader.RootFolder.GCODEPath)
	loader.Config.Set("path.bin", loader.RootFolder.BinPath)
	loader.Config.Set("path.extra_bins", loader.RootFolder.ExtraBins)
	loader.Config.Set("path.core_bins", loader.RootFolder.CoreBins)
	loader.Config.Set("path.workspace", loader.RootFolder.Workspace)
	loader.Config.Set("path.log", loader.RootFolder.Logs)

	err := loader.Config.WriteConfig()
	if err != nil {
		fmt.Println("Could not write config")
		return err
	}

	err = loader.Config.ReadInConfig()
	if err != nil {
		fmt.Println("Could not read in updated config")
		return err
	}
	loader.mux.Unlock()

	return nil
}

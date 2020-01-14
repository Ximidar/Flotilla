/*
* @Author: Ximidar
* @Date:   2019-02-05 16:23:20
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-28 17:35:26
 */

package FlotillaInstance

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	"github.com/Ximidar/Flotilla/CommonTools/NatsConnect"
	version "github.com/Ximidar/Flotilla/CommonTools/versioning"
	"github.com/Ximidar/Flotilla/DataStructures/ConfigStructures"
	"github.com/Ximidar/Flotilla/NodeLauncher/Config"
	"github.com/Ximidar/Flotilla/NodeLauncher/Monitor"
	"github.com/Ximidar/Flotilla/NodeLauncher/RootFolder"
	nats "github.com/nats-io/go-nats"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// CoreNodes is a map of all core nodes and the arguments that go with them
var CoreNodes = map[string][]string{
	"gnatsd":              []string{"-m", "5070"},
	"Commango":            []string{},
	"FlotillaFileManager": []string{},
	"FlotillaStatus":      []string{},
}

// TermChannel will monitor for an exit signal
var TermChannel chan os.Signal

// FlotillaInstance will start an instance for flotilla
type FlotillaInstance struct {
	ActiveCoreNodes map[string]*Monitor.Monitor
	RootFolder      *RootFolder.RootFolder

	setupLogs  bool
	config     *Config.Loader
	nc         *nats.Conn
	configNats *Config.Nats
	startNats  bool
	tlsEnabled bool
	tlsKey     string
	tlsCert    string
}

// NewFlotillaInstance will create an instance for a Flotilla Package Path
func NewFlotillaInstance(rootFolderPath string, startNatsflag bool, tlsEnabled bool) (*FlotillaInstance, error) {
	var err error
	Flotilla := new(FlotillaInstance)
	Flotilla.RootFolder, err = RootFolder.NewRootFolder(rootFolderPath)
	if err != nil {
		return nil, err
	}
	Flotilla.startNats = startNatsflag
	Flotilla.tlsEnabled = tlsEnabled

	// Setup the values for TLS
	err = Flotilla.configureTLS()
	if err != nil {
		return nil, err
	}

	err = Flotilla.PopulateCoreNodes()
	if err != nil {
		return nil, err
	}

	// Setup config
	Flotilla.config, err = Config.NewLoader(Flotilla.RootFolder)
	if err != nil {
		return nil, err
	}

	Flotilla.setupLogging()

	return Flotilla, nil
}

func (flotilla *FlotillaInstance) setupLogging() {

	// limit the log size and age.
	log.SetOutput(&lumberjack.Logger{
		Filename:   path.Clean(flotilla.RootFolder.Logs + "/flotilla.log"),
		MaxSize:    15, // megabytes
		MaxBackups: 2,
		MaxAge:     30, //days
	})
	flotilla.setupLogs = true
}

// WriteLogFromMonitor will write to the log as well as print to the Console
func (flotilla *FlotillaInstance) WriteLogFromMonitor(name string, message string) {
	go fmt.Print(name+" ", message)
	go log.Print(name+" ", message)
}

// PopulateCoreNodes will use the RootFolder to figure out where all the core node executables are
func (flotilla *FlotillaInstance) PopulateCoreNodes() error {

	// Walk the Core Nodes Directory and Make all nodes
	coreNodes, err := ioutil.ReadDir(flotilla.RootFolder.CoreBins)
	if err != nil {
		return err
	}

	// ForEach File found
	for _, file := range coreNodes {
		// Get the basename and compare to known Core List
		basename := path.Base(file.Name())
		val, ok := CoreNodes[basename]
		// If we successfully grabbed a path to a known Core executable, then make a new Monitor
		if ok {
			nodePath := flotilla.RootFolder.CoreBins + "/" + file.Name()
			node, err := Monitor.NewMonitor(nodePath, flotilla.WriteLogFromMonitor, val...)
			if err != nil {
				fmt.Println("Could not make a monitor for:", file.Name(), nodePath)
				return err
			}
			// Add Node to Active Node list
			if flotilla.ActiveCoreNodes == nil {
				flotilla.ActiveCoreNodes = make(map[string]*Monitor.Monitor)
			}
			flotilla.ActiveCoreNodes[file.Name()] = node
		} else {
			// Report anomolies
			fmt.Println("Couldn't find entry for", basename, file.Name())
		}
	}

	return nil
}

func (flotilla *FlotillaInstance) configureTLS() error {
	if !flotilla.tlsEnabled {
		return nil
	}

	// Attempt to setup TLS
	fmt.Println("Configuring TLS")

	// Find the key and certificate
	files, err := ioutil.ReadDir(flotilla.RootFolder.TLS)
	if err != nil {
		fmt.Println("TLS Folder is not available")
		return err
	}

	// Check that there are at least 2 files
	if len(files) < 2 {
		fmt.Println("TLS directory does not contain any Certificates")
		return errors.New("tls directory does not contain certificates")
	}

	// Find the Key and Cert
	keyPos, err := flotilla.findNamePos(files, ".key")
	if err != nil {
		fmt.Println("Could not find TLS Key")
		return err
	}
	// Populate Key
	flotilla.tlsKey = files[keyPos].Name()
	flotilla.tlsKey = path.Clean(flotilla.RootFolder.TLS + "/" + flotilla.tlsKey)

	certPos, err := flotilla.findNamePos(files, ".pem")
	if err != nil {
		fmt.Println("Could not find PEM Cert")
		return err
	}
	// Populate Cert
	flotilla.tlsCert = files[certPos].Name()
	flotilla.tlsCert = path.Clean(flotilla.RootFolder.TLS + "/" + flotilla.tlsCert)

	// Add arguments to the core nodes
	_, ok := CoreNodes["gnatsd"]
	if !ok {
		fmt.Println("gnatsd is not on the CoreNode map")
		return errors.New("gnatsd is not on the corenode map")
	}
	/*
		--tls                        Enable TLS, do not verify clients (default: false)
		--tlscert FILE               Server certificate file
		--tlskey FILE                Private key for server certificate
	*/
	CoreNodes["gnatsd"] = append(CoreNodes["gnatsd"], "--tls",
		"--tlscert", flotilla.tlsCert,
		"--tlskey", flotilla.tlsKey)

	return nil
}

func (flotilla *FlotillaInstance) findNamePos(files []os.FileInfo, suffix string) (int, error) {
	for index, f := range files {
		if strings.HasSuffix(f.Name(), suffix) {
			return index, nil
		}
	}

	return -1, errors.New("could not find certs")
}

func (flotilla *FlotillaInstance) startConfigNats() error {
	var err error
	flotilla.nc, err = NatsConnect.DefaultConn(NatsConnect.DockerNATS, ConfigStructures.Name)
	if err != nil {
		return err
	}

	flotilla.configNats, err = Config.NewNats(flotilla.config)
	if err != nil {
		return err
	}

	return flotilla.configNats.StartNatsInteraction(flotilla.nc)

}

// Serve will start the Flotilla Instance
func (flotilla *FlotillaInstance) Serve() error {

	flotilla.Logo()

	// Start Nats Server First if we are supposed to
	if flotilla.startNats {
		val, ok := flotilla.ActiveCoreNodes["gnatsd"]

		if ok {
			val.StartProcessInGoroutine()
		} else {
			return errors.New("could not find gnatsd executable")
		}

		<-time.After(200 * time.Millisecond)
	}

	err := flotilla.startConfigNats()
	if err != nil {
		return err
	}

	startOrder := []string{"FlotillaStatus", "FlotillaFileManager", "Commango"}

	for _, progname := range startOrder {
		// Start all Nodes in order
		flotilla.ActiveCoreNodes[progname].StartProcessInGoroutine()
		<-time.After(200 * time.Millisecond)
	}

	// Function for waiting for exit on the main loop
	// Wait for termination
	TermChannel = make(chan os.Signal)
	signal.Notify(TermChannel, os.Interrupt, syscall.SIGTERM)
	select {
	case <-TermChannel:
		fmt.Println("Recieved Interrupt Sig, Now Exiting.")
		flotilla.Stop()
		os.Exit(1)
	}

	return nil
}

// Stop will tell all nodes to exit
func (flotilla *FlotillaInstance) Stop() {
	for _, node := range flotilla.ActiveCoreNodes {
		err := node.Exit()
		if err != nil {
			flotilla.WriteLogFromMonitor(node.Name, err.Error())
		}
	}
}

// Logo will show the Logo of flotilla
func (flotilla *FlotillaInstance) Logo() {
	logo := "\n" +
		"   ___ _       _   _ _ _       \n" +
		"  / __\\ | ___ | |_(_) | | __ _ \n" +
		" / _\\ | |/ _ \\| __| | | |/ _` |\n" +
		"/ /   | | (_) | |_| | | | (_| |\n" +
		"\\/    |_|\\___/ \\__|_|_|_|\\__,_|\n\n" +
		"Author: " + version.CompiledBy + "\n" +
		"Compiled On: " + version.CompiledDate + "\n" +
		"Version: " + version.Version + "\n" +
		"Commit: " + version.CommitHash
	fmt.Println(logo)

	// Just in case someone is just printing the logo
	if flotilla.setupLogs {
		log.Println(logo)
	}

}

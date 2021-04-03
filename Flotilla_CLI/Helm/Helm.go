package Helm

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	mapset "github.com/deckarep/golang-set"
	"github.com/spf13/viper"
)

// Helm is a tool to run a flotilla instance.
// Here are some features I am shooting for:
/*
	- Hot Reload: Changes will cause a recompile and restart
	- Monitor: Any process that goes down will be recovered
	- Container/Process: Processes are either run locally or in a container
	- Smooth start/stop
	- Configurable: easy yaml config
*/

func StartHelm(confPath string) {
	// Figure out configuration path
	if confPath == "" {
		newpath, err := FindConfiguration()
		if err != nil {
			fmt.Println("Cannot Find FlotHelm.yaml")
			os.Exit(1)
		}
		confPath = newpath
		fmt.Println("Found Config:", confPath)
	}

	// Load Configuration
	err := ConfigureHelm(confPath)
	if err != nil {
		fmt.Println("Configuration could not load:", err)
		os.Exit(1)
	}

	// Parse Configuration into exec rules
	rules := ParseConfig()

	// Execute Rules
	Executor := NewExecutor(rules)
	Executor.Start()
}

func FindConfiguration() (string, error) {
	fmt.Println("No Configuration given. Searching locally for config")
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot Access current directory", err)
		os.Exit(1)
	}

	files, _ := ioutil.ReadDir(pwd)

	for _, info := range files {
		if info.Name() == "FlotHelm.yaml" {
			return filepath.Join(pwd, info.Name()), nil
		}
	}

	return "", errors.New("cannot find configuration locally")

}

func ConfigureHelm(confPath string) error {
	name := filepath.Base(confPath)
	dir := filepath.Dir(confPath)

	viper.SetConfigName(name)
	viper.AddConfigPath(dir)
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}

func ParseConfig() []*ExecRule {
	// get top level keys
	tlk := GetTopLevelKeys()

	// give the tlk to a constructor which will make
	// exec rules for each key

	rules := make([]*ExecRule, 0)
	for _, key := range tlk {
		rule := ExecRuleCreator(key)
		rules = append(rules, rule)
	}

	return rules

}

func GetTopLevelKeys() []string {
	vkeys := viper.AllKeys()
	topLevelKeys := mapset.NewSet()
	for _, key := range vkeys {
		fmt.Println("parse:", key)
		split := strings.Split(key, ".")
		topLevelKeys.Add(split[0])
	}

	tlk := make([]string, 0)
	iter := topLevelKeys.Iterator()
	for key := range iter.C {
		tlk = append(tlk, key.(string))
	}
	return tlk
}

package Helm

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

type ExecRule struct {
	Name           string
	Command        string
	Args           []string
	WorkDir        string
	Build          string
	Container      bool
	ContainerName  string
	ContainerImage string
	Dockerfile     string

	LogChannel  chan string
	ControlChan chan string

	Info interface{} // TODO change to whatever the Command Type is
}

func ExecRuleCreator(TopLevelKey string) *ExecRule {

	// get configuration keys
	command := viper.GetString(fmt.Sprintf("%s.command", TopLevelKey))
	args := viper.GetStringSlice(fmt.Sprintf("%s.args", TopLevelKey))
	workDir := viper.GetString(fmt.Sprintf("%s.path", TopLevelKey))
	build := viper.GetString(fmt.Sprintf("%s.build", TopLevelKey))
	container := viper.GetBool(fmt.Sprintf("%s.container", TopLevelKey))
	containerName := viper.GetString(fmt.Sprintf("%s.container_name", TopLevelKey))
	image := viper.GetString(fmt.Sprintf("%s.image", TopLevelKey))
	dockerfile := viper.GetString(fmt.Sprintf("%s.dockerfile", TopLevelKey))

	// Parse Paths
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot Detirmine current directory", err)
		os.Exit(1)
	}

	workDir = path.Join(pwd, workDir)
	workDir = path.Clean(workDir)

	if container {
		dockerfile = path.Join(pwd, dockerfile)
		dockerfile = path.Clean(dockerfile)
	}

	// Create Exec Rule
	var exec *ExecRule
	if container {
		exec = NewExecContainerRule(
			TopLevelKey,
			args,
			containerName,
			image,
			dockerfile,
		)
	} else {
		exec = NewExecRule(
			TopLevelKey,
			command,
			build,
			args,
			workDir,
		)
	}

	return exec
}

// NewExecRule will set up a new rule
func NewExecRule(Name string, Command string, Build string, Args []string, WD string) *ExecRule {
	rule := new(ExecRule)
	rule.Build = Build
	rule.Command = Command
	rule.Args = Args
	rule.WorkDir = WD
	rule.Name = Name
	return rule
}

func NewExecContainerRule(Name string, Args []string, ContainerName string, Image string, Dockerfile string) *ExecRule {
	rule := new(ExecRule)
	rule.Name = Name
	rule.Args = Args
	rule.ContainerName = ContainerName
	rule.ContainerImage = Image
	rule.Dockerfile = Dockerfile

	return rule
}

// Start will begin execution of the rule
func (rule *ExecRule) Start() error {
	return nil
}

// Stop will stop the rule
func (rule *ExecRule) Stop(force bool) error {
	return nil
}

func (rule *ExecRule) BuildRule() error {
	if !rule.Container {

	}

	return nil
}

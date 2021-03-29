package Helm

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/go-connections/nat"
	"github.com/spf13/viper"
)

var ErrRuleNotContainer = errors.New("rule is not a container")
var ErrNoContainerImage = errors.New("no container image")

type ExecRule struct {
	Name    string
	Command string
	Args    []string
	WorkDir string
	Build   string

	Container        bool
	ContainerName    string
	ContainerImage   string
	Dockerfile       string
	ContainerTag     string
	ContainerVolumes []string
	ContainerPorts   []string
	ContainerId      string

	LogChannel chan string

	Exec *exec.Cmd
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
	volumes := viper.GetStringSlice(fmt.Sprintf("%s.volumes", TopLevelKey))
	ports := viper.GetStringSlice(fmt.Sprintf("%s.ports", TopLevelKey))

	fmt.Printf("Command: %s\nArgs: %s\nWorkDir: %s\nBuild: %s\nContainer: %t\nContainerName: %s\nImage: %s\nDockerFile: %s\nVolumes: %s\nPorts: %s\n",
		command, args, workDir, build, container, containerName, image, dockerfile, volumes, ports)

	// Parse Paths
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot Detirmine current directory", err)
		os.Exit(1)
	}

	workDir = path.Join(pwd, workDir)
	workDir = path.Clean(workDir)

	if container && dockerfile != "" {
		dockerfile = path.Join(pwd, dockerfile)
		dockerfile = path.Clean(dockerfile)
	}

	// Parse any environment variables in volumes
	if len(volumes) != 0 {
		var re = regexp.MustCompile(`(?m)[$]{(\S+)}`)
		for index, volume := range volumes {
			newVol := volume
			for _, match := range re.FindAllStringSubmatch(volume, -1) {
				val := os.Getenv(match[1])
				newVol = strings.Replace(newVol, match[0], val, 1)
			}

			if strings.HasPrefix(newVol, "~/") {
				homedir, _ := os.UserHomeDir()
				newVol = strings.Replace(newVol, "~", homedir, 1)
			}

			volumes[index] = newVol

		}
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
			volumes,
			ports,
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

	fmt.Println(exec.String())
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

func NewExecContainerRule(Name string, Args []string, ContainerName string, Image string, Dockerfile string,
	Volumes []string, Ports []string) *ExecRule {
	rule := new(ExecRule)
	rule.Name = Name
	rule.Args = Args
	rule.Container = true
	rule.ContainerName = ContainerName
	rule.ContainerImage = Image
	rule.Dockerfile = Dockerfile
	rule.ContainerTag = fmt.Sprintf("flot/%s:latest", rule.Name)
	rule.ContainerVolumes = Volumes
	rule.ContainerPorts = Ports

	return rule
}

func (rule *ExecRule) String() string {
	return fmt.Sprintf("Command: %s\nArgs: %s\nWorkDir: %s\nBuild: %s\nContainer: %t\nContainerName: %s\nImage: %s\nDockerFile: %s\nVolumes: %s\nPorts: %s\n",
		rule.Command, rule.Args, rule.WorkDir, rule.Build, rule.Container, rule.ContainerName, rule.ContainerImage, rule.Dockerfile, rule.ContainerVolumes, rule.ContainerPorts)
}

// Start will begin execution of the rule
func (rule *ExecRule) Start() error {
	if rule.Container {
		return rule.StartContainer()
	}

	err := rule.RunCommand(
		rule.Command,
		rule.Args,
		rule.WorkDir,
	)
	return err
}

// Stop will stop the rule
func (rule *ExecRule) Stop(force bool) error {
	if rule.Container {
		return rule.StopContainer()
	}

	rule.Exec.Process.Signal(os.Kill)
	return rule.Exec.Wait()
}

func (rule *ExecRule) BuildRule() error {
	if rule.Container {
		return rule.BuildContainerRule()
	}
	err := rule.RunCommand(
		rule.Build,
		[]string{},
		rule.WorkDir,
	)
	if err != nil {
		fmt.Println("Could not build", rule.Name)
		return err
	}
	return nil

}

func (rule *ExecRule) BuildContainerRule() error {
	if !rule.Container {
		return ErrRuleNotContainer
	}

	if rule.Dockerfile == "" {
		rule.LogChannel <- "No Dockerfile, Attempting to pull"
		return rule.PullContainer()
	}

	cli, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}

	// create build config
	dockerfileName := path.Base(rule.Dockerfile)
	config := types.ImageBuildOptions{
		Dockerfile: dockerfileName,
		Tags:       []string{rule.ContainerTag},
		Remove:     true,
	}

	// archive the folder
	if rule.WorkDir == "" {
		rule.WorkDir = path.Dir(rule.Dockerfile)
	}
	rule.LogChannel <- fmt.Sprintf("Attempting to archive %s", rule.WorkDir)
	archive, _ := archive.TarWithOptions(
		rule.WorkDir,
		&archive.TarOptions{},
	)

	// make the context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// build the image
	build, err := cli.ImageBuild(ctx, archive, config)
	if err != nil {
		rule.LogChannel <- fmt.Sprintf("Image Build Failed for %s Err: %s\n", rule.Name, err)
		return err
	}

	defer build.Body.Close()

	rule.ScanReader(build.Body)

	return nil
}

func (rule *ExecRule) PullContainer() error {
	if !rule.Container {
		return ErrRuleNotContainer
	}

	if rule.ContainerImage == "" {
		rule.LogChannel <- "No container image name to pull"
		return ErrNoContainerImage
	}

	// client
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}

	// make the context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// pull the image
	log, err := cli.ImagePull(ctx, rule.ContainerImage, types.ImagePullOptions{})
	if err != nil {
		rule.LogChannel <- fmt.Sprintf("Cannot Pull Image %s", rule.ContainerImage)
	}
	defer log.Close()

	rule.ScanReader(log)

	return nil
}

func (rule *ExecRule) StartContainer() error {
	if !rule.Container {
		return ErrRuleNotContainer
	}

	// make the context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// client
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}

	// ports
	portMap := nat.PortMap{}
	for _, rawport := range rule.ContainerPorts {
		prts := strings.Split(rawport, ":")
		local := prts[0]
		container := prts[1]
		protoString, portString := nat.SplitProtoPort(container)

		cport, err := nat.NewPort(protoString, portString)
		if err != nil {
			rule.LogChannel <- fmt.Sprintf("Could not create port for %s", rawport)
			return err
		}

		entry := []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: local,
			},
		}

		portMap[cport] = entry

	}

	// Volumes (mounts)
	volumeBind := make([]mount.Mount, 0)
	for _, volume := range rule.ContainerVolumes {
		vols := strings.Split(volume, ":")
		local := vols[0]
		cont := vols[1]
		mountPoint := mount.Mount{
			Type:   mount.TypeBind,
			Source: local,
			Target: cont,
		}
		volumeBind = append(volumeBind, mountPoint)
	}

	// Config
	contConfig := container.Config{
		Image: rule.ContainerImage,
	}

	hostConfig := container.HostConfig{
		Mounts:       volumeBind,
		PortBindings: portMap,
		AutoRemove:   true,
	}

	container, err := cli.ContainerCreate(
		ctx,
		&contConfig,
		&hostConfig,
		nil, nil, rule.Name)
	if err != nil {
		rule.LogChannel <- fmt.Sprintf("RULE: %s Could not build image. Err: %s", rule.Name, err)
		return err
	}

	rule.ContainerId = container.ID
	rule.LogChannel <- fmt.Sprintf("RULE: %s Created Container ID: %s", rule.ContainerName, rule.ContainerId)

	err = cli.ContainerStart(ctx, rule.ContainerId, types.ContainerStartOptions{})
	if err != nil {
		rule.LogChannel <- fmt.Sprintf("RULE: %s Cannot Start Container ID: %s Err: %s", rule.ContainerName, rule.ContainerId, err)
		return err
	}
	rule.ContainerLogger()
	return nil
}

func (rule *ExecRule) StopContainer() error {
	if !rule.Container {
		return ErrRuleNotContainer
	}

	// make the context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// client
	cli, err := client.NewClientWithOpts()
	if err != nil {
		rule.LogChannel <- fmt.Sprintf("Cannot stop %s Err: %s", rule.Name, err)
		return err
	}

	rule.LogChannel <- fmt.Sprintf("Stopping %s", rule.Name)
	duration, _ := time.ParseDuration("5m")
	err = cli.ContainerStop(ctx, rule.ContainerId, &duration)
	if err != nil {
		rule.LogChannel <- fmt.Sprintf("Cannot stop %s Err: %s", rule.Name, err)
	}

	rule.LogChannel <- fmt.Sprintf("Removing %s ID: %s", rule.Name, rule.ContainerId)
	err = cli.ContainerRemove(ctx, rule.ContainerId, types.ContainerRemoveOptions{})
	if err != nil {
		rule.LogChannel <- fmt.Sprintf("Cannot Remove %s Err: %s", rule.Name, err)
	}
	return err
}

func (rule *ExecRule) ContainerLogger() {
	if !rule.Container {
		return
	}
	if rule.ContainerId == "" {
		return
	}

	logConfig := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Details:    true,
	}

	// make the context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// client
	cli, err := client.NewClientWithOpts()
	if err != nil {
		rule.LogChannel <- "Cannot form client"
		rule.LogChannel <- err.Error()
		return
	}

	logReader, err := cli.ContainerLogs(ctx, rule.ContainerId, logConfig)

	if err != nil {
		rule.LogChannel <- "Cannot get logs for container"
		rule.LogChannel <- err.Error()
		return
	}

	defer logReader.Close()

	rule.ScanReader(logReader)

}

func (rule *ExecRule) ScanReader(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		rule.LogChannel <- scanner.Text()
	}
}

func (rule *ExecRule) RunCommand(command string, args []string, workDir string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = workDir

	// get stdout
	outPipe, _ := cmd.StdoutPipe()
	go rule.ScanReader(outPipe)

	// get stderr
	errPipe, _ := cmd.StderrPipe()
	go rule.ScanReader(errPipe)

	rule.Exec = cmd

	cmd.Start()
	err := cmd.Wait()
	return err
}

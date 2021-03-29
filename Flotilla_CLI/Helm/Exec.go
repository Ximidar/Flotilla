package Helm

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
	Executor will monitor all individual rules and stream all their logs into one channel
*/
type Executor struct {
	Rules []*ExecRule

	LogChan  chan string
	StopChan []chan string
}

func NewExecutor(rules []*ExecRule) *Executor {
	exec := new(Executor)
	exec.Rules = rules
	exec.LogChan = make(chan string, 1000)
	return exec
}

func (exec *Executor) Start() error {
	go exec.LogMon()
	for _, rule := range exec.Rules {
		// Give all the rules a single channel to output to
		fmt.Println("Attempting to start:", rule.Name)
		rule.LogChannel = exec.LogChan
		rule.Start()

	}

	// Monitor for exit signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	for sig := range stop {

		fmt.Println("Got Signal:", sig)

		exec.Stop()
		fmt.Println("Finished Quitting!")
		break

	}
	return nil
}

func (exec *Executor) Stop() error {

	for _, rule := range exec.Rules {
		rule.Stop(false)
	}

	for _, logger := range exec.StopChan {
		logger <- "Exit"
	}

	return nil
}

func (exec *Executor) LogMon() {
	escape := make(chan string, 10)
	exec.StopChan = append(exec.StopChan, escape)
	for {
		select {
		case mess := <-exec.LogChan:
			fmt.Println(mess)
		case mess := <-escape:
			fmt.Println("Escaping Due to", mess)
			return
		}
	}
}

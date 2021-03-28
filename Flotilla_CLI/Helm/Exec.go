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
	exec.LogChan = make(chan string, 100)
	return exec
}

func (exec *Executor) Start() error {

	for _, rule := range exec.Rules {

		go exec.LogMon(rule.LogChannel)
		go rule.Start()

	}

	// Monitor for exit signal
	stop := make(chan os.Signal, 10)
	signal.Notify(stop)
	for {
		select {
		case sig := <-stop:
			fmt.Println("Got Signal:", sig)
			if sig == syscall.SIGTERM {
				exec.Stop()
				fmt.Println("Finished Quitting!")
				return nil
			}
		}
	}
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

func (exec *Executor) LogMon(logChan chan string) {
	escape := make(chan string, 10)
	exec.StopChan = append(exec.StopChan, escape)
	for {
		select {
		case mess := <-logChan:
			exec.LogChan <- mess
		case mess := <-escape:
			fmt.Println("Escaping Due to", mess)
			break
		}
	}
}

package app

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"syscall"
	"time"

	daemon "github.com/tyranron/daemonigo"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func Run() (exitCode int) {

	switch isDaemon, err := daemon.Daemonize(); {
	case !isDaemon:
		return
	case err != nil:
		fmt.Errorf("could not start daemon, reason -> %v", err.Error())
		return 1
	}

	var err error
	if err = prepareTemplates(); err != nil {
		fmt.Errorf("could not parse html templates, reason -> %v", err.Error())
		return 1
	}

	httpPort := os.Getenv("_GO_HTTP")
	if httpPort == "" {
		httpPort = DefaultHttpPort
	}

	bindHttpHandlers()
	canExit, httpErr := make(chan sig, 1), make(chan error, 1)
	go func() {
		defer close(canExit)
		if err := http.ListenAndServe(httpPort, nil); err != nil {
			httpErr <- fmt.Errorf(
				"creating HTTP server on port '%s' FAILED, reason: %v",
				httpPort, err.Error(),
			)
		}
	}()
	select {
	case err = <-httpErr:
		fmt.Errorf("Http Error recived thogth the channel: %v", err.Error())
		return 1
	case <-time.After(300 * time.Millisecond):
	}

	notifyParentProcess()

	<-canExit
	return
}

// Notifies parent process that everything is OK.
func notifyParentProcess() {
	if err := syscall.Kill(os.Getppid(), syscall.SIGUSR1); err != nil {
		fmt.Errorf(
			"Notifying parent process FAILED, reason -> %v", err.Error(),
		)
	} else {
		fmt.Println("Notifying parent process SUCCEED")
	}
}

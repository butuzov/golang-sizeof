package app

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Run() error {

	var err error
	if err = prepareTemplates(); err != nil {
		return fmt.Errorf("%+v: while parsing templates", err)
	}

	bindHttpHandlers()

	var errCh = make(chan error)

	go func() {
		defer close(errCh)

		if err := http.ListenAndServe(defaultHttpServer, nil); err != nil {
			errCh <- err
		}
	}()

	time.AfterFunc(time.Second, func() {
		if err := Open("http://" + defaultHttpServer); err != nil {
			log.Panic(err)
		} else {
			log.Println("Running at http://" + defaultHttpServer)
		}
	})

	return <-errCh
}

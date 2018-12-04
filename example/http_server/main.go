package main

import (
	"fmt"
	"github.com/chaitin/service"
	"io"
	"log"
	"net/http"
	"os"
)

type program struct{}

func ensure(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (p *program) Start(s service.Service) error {
	go func() {
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			_, err := io.WriteString(writer, fmt.Sprintf("Hello World!"))
			ensure(err)
		})
		err := http.ListenAndServe(":20147", nil)
		ensure(err)
	}()
	return nil
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	serviceConfig := &service.Config{
		Name:        "GoServiceExampleHTTPServer",
		DisplayName: "Go Service Example for HTTP Server",
		Description: "This is an example Go service that response 'Hello World'",
	}
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		ensure(err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "update":
			err := service.UpdateServicePath(s)
			if err != nil {
				ensure(err)
			}
		default:
			ensure(service.Control(s, os.Args[1]))
		}
	} else {
		ensure(s.Run())
	}
}

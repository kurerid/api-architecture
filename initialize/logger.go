package initialize

import (
	"github.com/execaus/exloggo"
	"log"
)

func Logger() error {
	if err := exloggo.SetParameters(&exloggo.Parameters{
		Mode:          exloggo.DevelopmentMode,
		ServerVersion: "development",
		Directory:     "/logs",
	}); err != nil {
		log.Fatal(err.Error())
	}
	return nil
}

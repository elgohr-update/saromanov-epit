package main

import (
	"fmt"

	"github.com/saromanov/epit/pkg/epit"
	"go.uber.org/zap"
)

func main() {
	log, err := zap.NewProduction()
	if err != nil {
		panic("unable to init logging")
	}
	if err := epit.LoadConfig("./examples/basic/.epit.yml"); err != nil {
		log.Fatal(fmt.Sprintf("unable to parse config: %v", err))
	}
}
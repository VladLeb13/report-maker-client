package main

import (
	"config"
	"context"
	"fmt"
	"log"
	"report"
	"testing"
	"tools"
)

func TestPrseConfig(t *testing.T) {
	ctx := tools.AppContex{
		Context: context.Background(),
	}

	err := config.Parse(&ctx)
	if err != nil {
		log.Fatal("configuration read error detected: ", err)
	}

	cnf := ctx.Context.Value("configuration").(*config.Config)

	fmt.Println("folder: ", cnf.SavingFolder)
	fmt.Println("server: ", cnf.ServerAddress)
}

func TestCreateReport(t *testing.T) {
	ctx := tools.AppContex{
		Context: context.Background(),
	}
	report.Create(&ctx)
}

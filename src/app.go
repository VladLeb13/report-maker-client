package main

import (
	"config"
	"context"
	"log"
	"os"
	"report"
	"report/converting"
	"tools"
)

//go build -ldflags "-s -H windowsgui" main.go -без консоли

func main() {

	ctx := tools.AppContex{
		Context: context.Background(),
	}

	err := config.Parse(&ctx)
	if err != nil {
		log.Println("configuration read error detected: ", err)
		os.Exit(1)
	}

	report.Create(&ctx)
	converting.HTML(&ctx)
	report.Save(&ctx)

	return
}

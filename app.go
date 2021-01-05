package main

import (
	"context"
	"log"
	"os"
	"report-maker-client/config"
	"report-maker-client/report"
	"report-maker-client/report/converting"
	"report-maker-client/tools"
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
	err = converting.JSON(&ctx)

	report.Save(&ctx)

	if err == nil {
		report.Send(&ctx)
	}

	return
}

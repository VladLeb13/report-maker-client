package main

import (
	"context"
	"fmt"
	"log"
	"report-maker-client/config"
	"report-maker-client/report"
	"report-maker-client/report/converting"
	"report-maker-client/tools"
	"testing"
)

func TestPrseConfig(t *testing.T) {
	ctx := tools.AppContex{
		Context: context.Background(),
	}

	err := config.Parse(&ctx)
	if err != nil {
		log.Fatal("configuration read error detected: ", err)
	}

	cnf := ctx.Context.Value("configuration").(*tools.Config)

	fmt.Println("folder: ", cnf.SavingFolder)
	fmt.Println("server: ", cnf.ServerAddress)
	fmt.Println("login: ", cnf.Auth.Login)
	fmt.Println("password: ", cnf.Auth.Password)
}

func TestCreateReport(t *testing.T) {
	ctx := tools.AppContex{
		Context: context.Background(),
	}
	report.Create(&ctx)
}

func TestSave(t *testing.T) {
	ctx := tools.AppContex{
		Context: context.Background(),
	}

	conf := &tools.Config{
		SavingFolder: "D:\\Reports\\",
	}
	ctx.Context = context.WithValue(ctx.Context, "configuration", conf)

	ctx.Context = context.WithValue(ctx.Context, "HTMLReport", "report string")

	report.Save(&ctx)
}

func TestConvert(t *testing.T) {
	ctx := tools.AppContex{
		Context: context.Background(),
	}
	report.Create(&ctx)
	converting.HTML(&ctx)

}

func TestApp(t *testing.T) {
	main()
}

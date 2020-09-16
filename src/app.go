package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"gui"
	"manager"
)

var (
	isCLI            = flag.Bool("cli", false, "use \"-cli\" flag to disable gui")
	isGenerateReport = flag.Bool("report", false, "use the \"-r\" flag to generate the report")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	mgr := new(manager.AppManager)
	go func() {
		mgr.Start()
	}()
	ctx = context.WithValue(ctx, "manager", mgr)

	if *isCLI == false {
		go func() {
			gui.Start(ctx)

		}()
		go func() {
			//запускаем саму апку
			fmt.Println("fffff")
		}()
	}

	if *isCLI && *isGenerateReport {
		go func() {
			//запускаем саму апку
			fmt.Println("fffff")
		}()
		//даем команду на старт отчета
	}

	time.Sleep(10 * time.Second)
	cancel()
	if err := ctx.Err(); err != nil {
		fmt.Println(err)
		time.Sleep(40 * time.Second)
	}

	return
}

package report

import (
	"context"
	"report-maker-client/report/fill"
	"report-maker-client/tools"
	"runtime"
	"sync"

	"github.com/VladLeb13/report-maker-lib/datalib"
)

func Create(ctx *tools.AppContex) {
	report := datalib.Report{}

	var wg sync.WaitGroup
	wg.Add(1)

	runtime.GOMAXPROCS(2)

	go func() {

		report.Hardware = fill.Hardware()
		report.Perfomance = fill.Perfomance()

		wg.Done()
	}()

	report.Software = fill.Software()
	report.Events = fill.Events()

	wg.Wait()

	ctx.Context = context.WithValue(ctx.Context, "report", report)
}

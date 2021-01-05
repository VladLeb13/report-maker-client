package converting

import (
	"context"
	"encoding/json"
	"log"
	"report-maker-client/tools"

	"github.com/VladLeb13/report-maker-lib/datalib"
)

func JSON(ctx *tools.AppContex) (err error) {
	report := ctx.Context.Value("report").(datalib.Report)

	b, err := json.Marshal(report)
	if err != nil {
		log.Println("Error marshal json: ", err)
	}

	ctx.Context = context.WithValue(ctx.Context, "JSONReport", b)

	return
}

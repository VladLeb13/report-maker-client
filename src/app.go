package main

import (
	"context"
	"flag"
	"fmt"
)

var (
	isCLI            = flag.Bool("cli", false, "use \"-cli\" flag to disable gui")
	isGenerateReport = flag.Bool("report", false, "use the \"-r\" flag to generate the report")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	if err := ctx.Err(); err != nil {
		fmt.Println(err)
	}

	return
}

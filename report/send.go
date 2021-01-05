package report

import (
	"bytes"
	"log"
	"net/http"
	"report-maker-client/tools"
)

func Send(ctx *tools.AppContex) {
	cnf := ctx.Context.Value("configuration").(*tools.Config)
	b := ctx.Context.Value("JSONReport").([]byte)

	req, err := http.NewRequest("POST", cnf.ServerAddress, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	if cnf.Auth.Login == "" && cnf.Auth.Password == "" {
		log.Panicln("No auth record to config")
	}

	req.SetBasicAuth(cnf.Auth.Login, cnf.Auth.Password)

	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		log.Println("Response error", err)
		log.Println("DATA :", string(b))
		return
	}
	defer resp.Body.Close()
	log.Println("response Status:", resp.Status)

}

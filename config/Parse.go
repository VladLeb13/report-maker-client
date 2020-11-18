package config

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"report-maker-client/tools"
)

func Parse(ctx *tools.AppContex) (err error) {
	configfile, err := os.Open("config/config.json")
	if err != nil {
		return err
	}
	defer configfile.Close()

	data, err := ioutil.ReadAll(configfile)
	if err != nil {
		return err
	}

	cnf := &tools.Config{}
	err = json.Unmarshal(data, &cnf)
	if err != nil {
		return err
	}

	ctx.Context = context.WithValue(ctx.Context, "configuration", cnf)

	return nil
}

package config

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"tools"
)

type (
	Config struct {
		ServerAddress string `json:"address"`
		SavingFolder  string `json:"folder"`
	}
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

	cnf := &Config{}
	err = json.Unmarshal(data, &cnf)
	if err != nil {
		return err
	}

	ctx.Context = context.WithValue(ctx.Context, "configuration", cnf)

	return nil
}

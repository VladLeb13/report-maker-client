package tools

import "context"

type (
	AppContex struct {
		Context context.Context
	}

	Config struct {
		ServerAddress string `json:"address"`
		SavingFolder  string `json:"folder"`
	}
)

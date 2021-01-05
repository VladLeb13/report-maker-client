package tools

import "context"

type (
	AppContex struct {
		Context context.Context
	}

	Config struct {
		ServerAddress string `json:"address"`
		SavingFolder  string `json:"folder"`
		Auth          Auth   `json:"auth"`
	}

	Auth struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
)

package env

import (
	"github.com/codingconcepts/env"
)

type Env struct {
	Port int `env:"PORT" required:"true"`
}

func LoadEnv() (*Env, error) {
	var e Env
	if err := env.Set(&e); err != nil {
		return nil, err
	}
	return &e, nil
}

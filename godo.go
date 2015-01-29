package godo

import (
	"log"

	"github.com/pkieltyka/godo-app/data"

	"github.com/pkieltyka/jwtauth"
)

var App *Godo

type Godo struct {
	Config    *Config
	TokenAuth *jwtauth.JwtAuth
}

func NewGodo(conf *Config) *Godo {
	t := &Godo{Config: conf}
	conf.Apply()

	// Data
	t.ConnectDB()

	// Auth
	t.TokenAuth = jwtauth.New("HS256", []byte(conf.Jwt.Secret), nil)

	return t
}

func (t *Godo) ConnectDB() {
	_, err := data.NewDBSession(t.Config.DB)
	if err != nil {
		log.Fatal(err)
	}
}

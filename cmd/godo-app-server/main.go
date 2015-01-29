package main

import (
	"flag"
	"log"
	"os"
	"syscall"

	"github.com/pkieltyka/godo-app"
	"github.com/pkieltyka/godo-app/api"
	"github.com/zenazn/goji/graceful"
)

var (
	flags      = flag.NewFlagSet("godo-server", flag.ExitOnError)
	configFile = flags.String("config", "", "path to config file")

	bind     = flags.String("bind", "0.0.0.0:3333", "<addr>:<port> to bind HTTP server")
	maxProcs = flags.Int("max-procs", 0, "GOMAXPROCS, default is NumCpu()")
)

func main() {
	flags.Parse(os.Args[1:])

	var err error
	var conf *godo.Config

	// Load config file from flag or env (if specified)
	conf, err = godo.NewConfigFromFile(*configFile, os.Getenv("CONFIG"))
	if err != nil {
		log.Fatal(err)
	}

	// Load new config with defaults (when no config file)
	if conf == nil {
		conf = godo.NewConfig()
		conf.Bind = *bind
		conf.MaxProcs = *maxProcs
	}

	app := godo.NewGodo(conf)
	godo.App = app

	graceful.AddSignal(syscall.SIGINT, syscall.SIGTERM)
	err = graceful.ListenAndServe(conf.Bind, api.New())
	if err != nil {
		log.Fatal(err)
	}
	graceful.Wait()
}

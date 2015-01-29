package godo

import (
	"errors"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/pkieltyka/godo-app/data"
)

var ErrNoConfigFile = errors.New("no configuration file specified")

type Config struct {
	Bind     string `toml:"bind"`
	MaxProcs int    `toml:"max_procs"`

	// [db]
	DB data.DBConf `toml:"db"`

	// [jwt]
	Jwt struct {
		Secret string `toml:"secret"`
	} `toml:"jwt"`

	// [webapp]
	Webapp struct {
		Path string `toml:"path"`
	} `toml:"webapp"`
}

func NewConfig() *Config {
	return &Config{}
}

func NewConfigFromFile(confFile string, confEnv string) (*Config, error) {
	cf := &Config{}
	if confFile == "" {
		confFile = confEnv
	}

	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		return nil, ErrNoConfigFile
	}

	if _, err := toml.DecodeFile(confFile, &cf); err != nil {
		return nil, err
	}
	return cf, nil
}

func (cf *Config) Apply() {
	if cf.MaxProcs <= 0 {
		cf.MaxProcs = runtime.NumCPU()
	}
	runtime.GOMAXPROCS(cf.MaxProcs)
}

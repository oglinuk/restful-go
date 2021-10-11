package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	cfgName = "cfg.yaml"

	_cfg *Config

	defaultCfg = &Config{
		cfgName,
		&ServerConfig{
			Host: "0.0.0.0",
			Port: 9001,
			Endpoints: []string{
				"/",
				"/books",
			},
		},
		&DatabaseConfig{
			Driver: "sqlite3",
			Name: "library",
			File: "library.db",
		},
	}
)

type Config struct {
	Name string
	Server *ServerConfig
	Database *DatabaseConfig
}

type ServerConfig struct {
	Host string
	Port int
	Endpoints []string
}

type DatabaseConfig struct {
	Driver string
	Name string
	File string
}

// load checks if cfgName exists, if not it will call save using
// defaultCfg. It then opens cfgName and calls yaml.Unmarshal, passing
// the (b)ytes read and _cfg
func load() {
	if _, err := os.Stat(cfgName); err != nil {
		save(defaultCfg)
	}

	f, err := os.Open(cfgName)
	if err != nil {
		log.Fatalf("config::Load::os.Open: %s\n", err.Error())
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("config::Load::io.ReadAll: %s\n", err.Error())
	}

	err = yaml.Unmarshal(b, &_cfg)
	if err != nil {
		log.Fatalf("config::Load::yaml.Unmarshal: %s\n", err.Error())
	}
}

// save cfg to a file called cfgName
func save(cfg *Config) {
	f, err := os.Create(cfgName)
	if err != nil {
		log.Fatalf("config::Save::os.Create: %s\n", err.Error())
	}
	defer f.Close()

	d, err := yaml.Marshal(&cfg)
	if err != nil {
		log.Fatalf("config::Save::yaml.Marshal: %s\n", err.Error())
	}

	f.Write(d)
}

// Get _cfg if not nil, otherwise call load before returning _cfg
func Get() *Config {
	if _cfg == nil {
		load()
	}

	return _cfg
}

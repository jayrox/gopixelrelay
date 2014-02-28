package config

// Just an entry point to ./something.json
// Init it first in application, then
// Any modules can call LoadInto() into their own struct, that they know what to do.

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	file string
}

var cfg *Config

func Init(FileName string) {
	if cfg == nil {
		cfg = &Config{file: FileName}
	}
}

func LoadInto(in interface{}) {
	if cfg == nil {
		panic("Config is not initialized. Run config.Init('somefile.json') first.")
	}

	data, err := ioutil.ReadFile(cfg.file)
	if err != nil {
		log.Println("Unable to read file:", err)
		return
	}

	err = json.Unmarshal(data, in)
	if err != nil {
		log.Println("Unable to read config.", err)
	}
}

func Get() *Config {
	return cfg
}

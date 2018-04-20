package downloader

import (
	"encoding/json"
	io "io/ioutil"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (self *Config) Load(filename string, v interface{}) {
	data, err := io.ReadFile(filename)

	if err != nil {
		return
	}

	dataJson := []byte(data)
	err = json.Unmarshal(dataJson, v)

	if err != nil {
		return
	}
}

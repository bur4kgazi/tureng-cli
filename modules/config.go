package modules

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Lang string
	Res  int
}

func Get(configFile string) Config{
	if configFile == "" {
		configFile = "conf.json"
	}
	var myConfig Config

	file, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &myConfig)
	if err != nil {
		panic(err)
	}

	return myConfig
}

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

type ColorTempLimitOverride struct {
	Min	int
	Max	int
}

type PolybarLampConfig struct {
	ApiBaseUrl string	`yaml:"apiBaseUrl"`
	ApiKey     string	`yaml:"apiKey"`
	LampEntity string	`yaml:"lampEntity"`
	ColorTempLimitOverride *ColorTempLimitOverride	`yaml:"colorTempLimitOverride,omitempty"`
}

func WriteInitialConfig() {
	config := PolybarLampConfig{}
	filebytes, err := yaml.Marshal(config)

	if err != nil {
		panic(config)
	}

	err = ioutil.WriteFile(path.Join(filepath.Dir(os.Args[0]), "config.yaml"), filebytes, 0644)

	if err != nil {
		panic(err)
	}
}

func GetPolybarLampConfig() PolybarLampConfig {
	file, err := ioutil.ReadFile(path.Join(filepath.Dir(os.Args[0]), "config.yaml"))

	if err != nil {
		panic(err)
	}

	config := PolybarLampConfig{}

	err = yaml.Unmarshal(file, &config)

	if err != nil {
		panic(err)
	}

	return config
}

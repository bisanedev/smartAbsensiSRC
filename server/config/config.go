package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Addr 			string		`yaml:"addr"`		
	Timezone 		string		`yaml:"timezone"`
	Sekolah			string		`yaml:"sekolah"`
	Warna			string		`yaml:"warna"`
	UsernameDB		string		`yaml:"usernameDB"`
	PasswordDB		string		`yaml:"passwordDB"`
	HostDB			string		`yaml:"hostDB"`
	NamaDB			string		`yaml:"namaDB"`
}

var config *Config

func Load(path string) error {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(result, &config)
}

func Get() *Config {
	return config
}
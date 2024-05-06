package config

import (
	"clash_config/log"
	"github.com/BurntSushi/toml"
	"math"
)

var Config = new(config)
var ConfigPath string = "./config/config.toml"

func ConfigInit() {
	_, err := toml.DecodeFile(ConfigPath, Config)
	if err != nil {
		log.Log.Fatalln(err.Error())
	}
	if Config.General.DefaultTraffic < 0 || Config.General.DefaultTraffic > float64(math.MaxUint64) {
		log.Log.Fatalln("DefaultTraffic is illegal")
	}
	log.Log.Println("config init")
}

type config struct {
	General general
}

type general struct {
	DefaultTraffic  float64
	Offset          float64
	NetworkCardName string
	//StartTime       string
	StartTraffic uint64
	ExpireTime   string
	//Location        string
	ClashPath    string
	ResetMonthly bool
}

package method

import (
	"clash_config/config"
	"clash_config/log"
	"github.com/BurntSushi/toml"
	"github.com/shirou/gopsutil/net"
	"os"
	"time"
)

func GetRemainingTraffic() (uint64, error) {
	usedTraffic, err := GetUsedTraffic()
	if err != nil {
		return 0, err
	}
	defaultTraffic := uint64(config.Config.General.DefaultTraffic * 1024 * 1024 * 1024)
	return defaultTraffic - usedTraffic, nil
}

func GetUsedTraffic() (uint64, error) {
	nowTraffic, err := GetAllTraffic(config.Config.General.NetworkCardName)
	if err != nil {
		return 0, err
	}
	return nowTraffic - config.Config.General.StartTraffic + uint64(config.Config.General.Offset*1024*1024*1024), nil
}

func GetAllTraffic(networkCardName string) (uint64, error) {
	stats, err := net.IOCounters(true)
	if err != nil {
		return 0, err
	}

	var totalBytes uint64

	for _, stat := range stats {
		if stat.Name == networkCardName {
			totalBytes += stat.BytesSent
			totalBytes += stat.BytesRecv
		}
	}

	return totalBytes, nil
}

func ResetConfig(isRenewExpireTime bool) error {

	config.Config.General.StartTraffic, _ = GetAllTraffic(config.Config.General.NetworkCardName)
	if isRenewExpireTime {
		expireTime := time.Now().AddDate(0, 1, 0)
		config.Config.General.ExpireTime = expireTime.String()
	}

	//file.Seek(0, 0)
	err := RewriteConfigFile()
	if err != nil {
		return err
	}
	return nil
}

func RewriteConfigFile() error {
	file, err := os.Create(config.ConfigPath)
	if err != nil {
		return err
	}
	defer file.Close()
	err = toml.NewEncoder(file).Encode(config.Config)
	if err != nil {
		return err
	}
	return nil
}

func ResetMonthly() {
	if !config.Config.General.ResetMonthly {
		return
	}
	expireTime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", config.Config.General.ExpireTime)
	if err != nil {
		log.Log.Fatalln("unknown expire time")
	}
	for true {
		now := time.Now()
		if now.Year() == expireTime.Year() && now.Month() == expireTime.Month() && now.Day() == expireTime.Day() {
			config.Config.General.Offset = 0
			config.Config.General.StartTraffic, err = GetAllTraffic(config.Config.General.NetworkCardName)
			if err != nil {
				log.Log.Fatalln(err.Error())
			}
			config.Config.General.ExpireTime = now.AddDate(0, 1, 0).String()
			expireTime = now.AddDate(0, 1, 0)
			err = RewriteConfigFile()
			if err != nil {
				log.Log.Fatalln(err.Error())
			}
		}
		time.Sleep(time.Second)
	}
}

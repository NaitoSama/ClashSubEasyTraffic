package method

import (
	"clash_config/config"
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
	return nowTraffic - config.Config.General.StartTraffic, nil
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
	file, err := os.OpenFile(config.ConfigPath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	config.Config.General.StartTraffic, err = GetAllTraffic(config.Config.General.NetworkCardName)
	if err != nil {
		return err
	}
	if isRenewExpireTime {
		expireTime := time.Now().AddDate(0, 1, 0)
		config.Config.General.ExpireTime = expireTime.String()
	}

	file.Seek(0, 0)
	err = toml.NewEncoder(file).Encode(config.Config)
	if err != nil {
		return err
	}

	return nil
}

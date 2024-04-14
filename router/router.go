package router

import (
	"clash_config/config"
	"clash_config/method"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"time"
)

func Router() {
	r := gin.Default()

	r.GET("/", clashConfig)

	r.Run(":8080")
}

func clashConfig(c *gin.Context) {
	configPath := config.Config.General.ClashPath
	configName := filepath.Base(configPath)
	usedTraffic, err := method.GetUsedTraffic()
	if err != nil {
		c.String(http.StatusInternalServerError, "can not get used traffic")
		return
	}
	faultTraffic := uint64(config.Config.General.DefaultTraffic * 1024 * 1024 * 1024)
	expireTime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", config.Config.General.ExpireTime)
	if err != nil {
		c.String(http.StatusInternalServerError, "can not format expire time")
		return
	}
	timestamp := expireTime.Unix()
	c.Header("Content-Disposition", "attachment; filename="+configName)
	c.Header("Subscription-Userinfo", fmt.Sprintf("upload=0; download=%v; total=%v; expire=%d", usedTraffic, faultTraffic, timestamp))
	c.File(configPath)
}

package method

import (
	"clash_config/log"
	"flag"
)

func Flags() {
	var flagReset bool
	var flagResetNoExpire bool

	flag.BoolVar(&flagReset, "r", false, "r Reset traffic count")
	flag.BoolVar(&flagResetNoExpire, "rn", false, "r Reset traffic count without renew expire time")
	flag.Parse()

	if flagReset {
		err := ResetConfig(true)
		if err != nil {
			log.Log.Fatalln(err)
		}
	}
	if flagResetNoExpire {
		err := ResetConfig(false)
		if err != nil {
			log.Log.Fatalln(err)
		}
	}
	go ResetMonthly()
}

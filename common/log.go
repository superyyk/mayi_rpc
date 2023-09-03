package common

import (
	log "github.com/sirupsen/logrus"
)

var Log log.Logger

func init() {
	Log.Warnf("init warnf:%s", "yyk_2012")
	Log.Warningln("dddd", "ffff")
}

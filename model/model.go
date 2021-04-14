package model

import (
	"terra/util"
)

var mlog *util.MLogger

func InitModel(env string) {
	mlog, _ = util.InitLog("model", env)
}

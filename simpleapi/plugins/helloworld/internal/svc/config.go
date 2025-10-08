package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
)

func (svcCtx *ServiceContext) SetConfigListener() {
	svcCtx.ConfigCenter.AddListener(func() {
		v := svcCtx.ConfigCenter.MustGetConfig()

		logx.Infof("reload config successfully")
		switch v.Log.Level {
		case "debug":
			logx.SetLevel(logx.DebugLevel)
		case "info":
			logx.SetLevel(logx.InfoLevel)
		case "error":
			logx.SetLevel(logx.ErrorLevel)
		case "severe":
			logx.SetLevel(logx.SevereLevel)
		}

		// add custom logic here
	})
}

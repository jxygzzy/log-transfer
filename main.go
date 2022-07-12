package main

import (
	"logtransfer/es"
	"logtransfer/kafka"
	"logtransfer/model"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

func main() {
	var config = new(model.Config)
	err := ini.MapTo(config, "./config/logtransfer.ini")
	if err != nil {
		logrus.Errorf("load config failed, err:%v", err)
		return
	}
	logrus.Info("load config success!")

	err = es.Init(config.EsConfig.Address, config.EsConfig.Index, config.EsConfig.MaxChanSize, config.EsConfig.GoroutineNum)
	if err != nil {
		logrus.Errorf("init elastic search failed, err:%v", err)
		return
	}
	logrus.Info("init elastic search success!")

	err = kafka.Init(strings.Split(config.KafkaConfig.Address, ","), config.KafkaConfig.Topic)
	if err != nil {
		logrus.Errorf("init kafka failed, err:%v", err)
		return
	}
	logrus.Info("init kafka success!")

	run()
}

func run() {
	select {}
}

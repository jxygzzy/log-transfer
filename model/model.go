package model

type Config struct {
	KafkaConfig `ini:"kafka"`
	EsConfig    `ini:"es"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type EsConfig struct {
	Address      string `ini:"address"`
	Index        string `ini:"index"`
	MaxChanSize  int    `ini:"max_chan_size"`
	GoroutineNum int    `ini:"goroutine_num"`
}

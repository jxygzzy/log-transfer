package es

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type EsClient struct {
	client      *elastic.Client
	logDataChan chan interface{}
	index       string
}

var (
	esClient = new(EsClient)
)

func Init(address string, index string, maxChanSize int, goroutine_num int) (err error) {
	client, err := elastic.NewClient(elastic.SetURL("http://" + address))
	if err != nil {
		logrus.Errorf("connect to es failed, err:%v", err)
		return
	}
	esClient.client = client
	esClient.index = index
	esClient.logDataChan = make(chan interface{}, maxChanSize)
	for i := 0; i < goroutine_num; i++ {
		go sendToEs()
	}
	return
}

func sendToEs() {
	for msg := range esClient.logDataChan {
		logrus.Infof("msg:%#v", msg)
		put, err := esClient.client.Index().
			Index(esClient.index).
			BodyJson(msg).
			Do(context.Background())
		if err != nil {
			logrus.Errorf("fail to put msg to es, err:%v", err)
			continue
		}
		logrus.Infof("Indexed msg %s to index %s, type %s", put.Id, put.Index, put.Type)
	}
}

func ToLogData(msg interface{}) {
	esClient.logDataChan <- msg
}

package kafka

import (
	"encoding/json"
	"logtransfer/es"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

func Init(address []string, topic string) (err error) {
	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		return
	}
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		logrus.Errorf("fail to get list of partition, err:%v", err)
		return
	}
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			logrus.Errorf("fail to start consumer for partition %d, err:%v", partition, err)
			return err
		}
		// defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				// 发送至es通道
				var msgJson map[string]interface{}
				err := json.Unmarshal(msg.Value, &msgJson)
				if err != nil {
					logrus.Errorf("json unmarshal failed, err:%v", err)
					continue
				}
				es.ToLogData(msgJson)
			}
		}(pc)
	}

	return
}

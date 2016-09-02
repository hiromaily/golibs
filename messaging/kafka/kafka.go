package kafka

import (
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	lg "github.com/hiromaily/golibs/log"
	"sync"
)

type ChReceive struct {
	ChWait chan bool
	ChCMsg chan *sarama.ConsumerMessage
}

//-----------------------------------------------------------------------------
// function
//-----------------------------------------------------------------------------
func createConfig() *sarama.Config {
	lg.Info("createConfig()")

	config := sarama.NewConfig()
	config.ClientID = "saramaId"
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	return config
}

//-----------------------------------------------------------------------------
// Consumer
//-----------------------------------------------------------------------------
func CreateConsumer(host string, port int) (sarama.Consumer, error) {
	lg.Info("CreateConsumer()")
	//config := sarama.NewConfig()
	config := createConfig()
	config.Producer.Partitioner = sarama.NewManualPartitioner

	// Create new consumer client. Kafka broker is running at localhost:9092
	//consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	consumer, err := sarama.NewConsumer([]string{fmt.Sprintf("%s:%d", host, port)}, config)
	//consumer, err := sarama.NewConsumer([]string{fmt.Sprintf("%s:%d", host, port)}, nil)
	if err != nil {
		//panic(err)
		return nil, errors.New(fmt.Sprintf("Failed to start consumer:%s", err))
	}
	return consumer, err
}

func Consumer(c sarama.Consumer, topic string, ch ChReceive) {
	lg.Info("Reveiver()")

	defer c.Close()

	// Create consumer object for TestTopic
	pc, err := c.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		//return errors.New(fmt.Sprintf("Failed to get ConsumePartition:%s", err))
		panic(fmt.Sprintf("Failed to get ConsumePartition:%s", err))
	}
	lg.Debug("Reveiver() Connected to kafka broker")

	defer pc.Close()

	go func() {
		//infinite loop
		lg.Debug("Reveiver() -> pc.Messages() Start")
		ch.ChWait <- true //notification for being ready.
		for m := range pc.Messages() {
			//send message
			ch.ChCMsg <- m

			//*sarama.ConsumerMessage
			//fmt.Printf("%+v\n", m)
			//fmt.Printf("Key: %v\n", string(m.Key))
			//fmt.Printf("Offset: %v\n", m.Offset)
			//fmt.Printf("Partition: %v\n", m.Partition)
			//fmt.Printf("Timestamp: %v\n", m.Timestamp)
			//fmt.Printf("Topic: %v\n", m.Topic)
			//fmt.Printf("Value: %v\n", string(m.Value))
		}
	}()

	//wait for notification from caller
	<-ch.ChWait

	defer func() {
		ch.ChWait <- true
	}()
	lg.Debug("Reveiver() finish")

	return
}

//TODO: Work in progress. Not checked yet.
func ReveiverOnMultiplePartitions(c sarama.Consumer, topic string) error {
	lg.Info("ReveiverOnMultiplePartitions()")

	var wg sync.WaitGroup

	defer c.Close()

	//when there are multiple partitions
	partitionList, err := c.Partitions(topic)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to get the list of partitions:%s", err))
	}
	lg.Debug("Connected to kafka broker")

	for partition := range partitionList {
		pc, err := c.ConsumePartition("topic.ops.falcon", int32(partition), sarama.OffsetNewest)
		if err != nil {
			return errors.New(fmt.Sprintf("Failed to start consumer for partition %d: %s\n", partition, err))
		}

		wg.Add(1)

		go func(pc sarama.PartitionConsumer) {
			defer func() {
				pc.AsyncClose()
				wg.Done()
			}()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()

	return nil
}

//-----------------------------------------------------------------------------
// Producer
//-----------------------------------------------------------------------------
func CreateMsg(topic, key, val string) *sarama.ProducerMessage {
	lg.Info("createMsg()")

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	//msg.Partition = int32(-1)
	msg.Partition = 1

	msg.Key = sarama.StringEncoder(key)
	msg.Value = sarama.ByteEncoder(val)

	return msg
}

func CreateProducer(host string, port int) (sarama.SyncProducer, error) {
	lg.Info("CreateProducer()")

	config := createConfig()

	producer, err := sarama.NewSyncProducer([]string{fmt.Sprintf("%s:%d", host, port)}, config)
	if err != nil {
		//
		return nil, errors.New(fmt.Sprintf("Failed to produce message: %s", err))
	}
	return producer, nil
}

func Producer(producer sarama.SyncProducer, msg *sarama.ProducerMessage) error {
	lg.Info("Sender()")

	//log
	//sarama.Logger = logger

	//defer producer.Close()

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to produce message: %s", err))
	}
	lg.Debugf("Sender() partition=%d, offset=%d\n", partition, offset)

	return nil
}

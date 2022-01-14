package main

import (
	"context"
	"flag"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var brokers = flag.String("brokers", "", "broker地址")

func main() {
	//监听一下kill
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)

	//consumer设置
	c := sarama.NewConfig()
	//从最新开始消费
	c.Consumer.Offsets.Initial = sarama.OffsetOldest
	//设置版本号
	c.Version = sarama.V2_8_1_0
	cg, err := sarama.NewConsumerGroup([]string{*brokers}, "kafka-yes", c)
	if err != nil {
		panic(err)
	}
	consum := NewConsum()
	go func() {
		for {
			//此方法阻塞；需要放到goroutine中执行，且永不退出
			err = cg.Consume(context.Background(), []string{"kafka-yes"}, consum)
			if err != nil {
				logrus.Error(err)
			}
		}

	}()

	s := <-sigs
	logrus.Infof("get sign %v", s)
	//停止消费
	consum.done <- struct{}{}
	//停止处理
	close(consum.pipe)
	logrus.Info("closed pipe")
	//关闭session
	err = cg.Close()
	if err != nil {
		panic(err)
	}
	logrus.Info("closed NewConsumerGroup")

}

func NewConsum() *MyConsumer {
	pipe := make(chan string, 8)
	done := make(chan struct{}, 1)
	c := MyConsumer{
		done: done,
		pipe: pipe,
	}
	//在这里写自定义的任务处理即可
	go func() {
		defer func() {
			logrus.Info("pipe finished ")
		}()
		for msg := range c.pipe {
			logrus.Info(msg)
		}
	}()

	return &c
}

type MyConsumer struct {
	done chan struct{}
	pipe chan string
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (mc *MyConsumer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (mc *MyConsumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (mc *MyConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
loop:
	for {
		select {
		case _ = <-mc.done:
			break loop
		case msg, ok := <-claim.Messages():
			if ok && msg != nil {
				session.MarkMessage(msg, "")
				logrus.Infof("offset %d", msg.Offset)
				//也可以在这里处理业务逻辑，
				mc.pipe <- string(msg.Value)
			}
		}
	}
	logrus.Info("ConsumeClaim finished")
	return nil
}

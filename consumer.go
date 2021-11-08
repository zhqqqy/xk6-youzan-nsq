package xk6_youzan_nsq

/**
* @author zhaohq
* @date 2021/11/5 11:35 上午
 */

import (
	"context"
	"github.com/youzan/go-nsq"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/lib"
	"log"
	"os"
	"time"
)

/**
* @author zhaohq
* @date 2021/11/5 11:34 上午
 */

func init() {
	modules.Register("k6/x/nsq", new(Nsq))
}

type Nsq struct{}
type MyTestHandler struct {
	ctx     context.Context
	message chan nsq.Message
}

var recieved = make(chan nsq.Message)

func (*Nsq) Consume(lookups []string, topic, channel string, maxInFlight, partition int) *nsq.Consumer {
	cfg := nsq.NewConfig()

	if maxInFlight == 0 {
		cfg.MaxInFlight = 100
	}
	if partition != 0 {
		cfg.EnableOrdered = true
	}
	consumer, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		ReportError(err, " new consumer object failed topic "+topic+" channel: "+channel)
		return nil
	}
	consumer.SetLogger(log.New(os.Stderr, "", log.LstdFlags), nsq.LogLevelError)
	log.SetPrefix("[bench_reader] ")
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
		recieved <- *m
		return nil
	}), 10)
	err = consumer.ConnectToNSQLookupds(lookups)
	if err != nil {
		ReportError(err, "Connect to nsq failed")
		return nil
	}
	return consumer
}

func (*Nsq) Received(ctx context.Context, timeout uint) nsq.Message {
	state := lib.GetState(ctx)
	if timeout == 0 {
		timeout = 500
	}
	if state == nil {
		common.Throw(common.GetRuntime(ctx), ErrorState)

	}
	select {
	case msg := <-recieved:
		return msg
	case <-time.After(time.Duration(timeout) * time.Millisecond):
		return nsq.Message{}

	}

}

func (*Nsq) Close(ctx context.Context, client *nsq.Consumer, timeout uint) {
	state := lib.GetState(ctx)
	if state == nil {
		common.Throw(common.GetRuntime(ctx), ErrorState)
		return
	}
	client.Stop()
	return

}

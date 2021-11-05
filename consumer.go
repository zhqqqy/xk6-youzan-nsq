package xk6_youzan_nsq

/**
* @author zhaohq
* @date 2021/11/5 11:35 上午
 */

import (
	"context"
	"errors"
	"github.com/youzan/go-nsq"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/lib"
	"go.k6.io/k6/stats"
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
	ctx context.Context
}

func (*Nsq) Reader(topic, channel string, maxInFlight, partition int) *nsq.Consumer {
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

	return consumer
}
func (*Nsq) Consume(
	ctx context.Context, lookups []string, reader *nsq.Consumer, limit, rfor int) string {
	return ConsumeInternal(ctx, lookups, reader, limit, rfor)
}

func ConsumeInternal(ctx context.Context, lookups []string, reader *nsq.Consumer, concurrency, rfor int) string {
	handler := MyTestHandler{
		ctx: ctx,
	}
	reader.AddConcurrentHandlers(&handler, concurrency)
	err := reader.ConnectToNSQLookupds(lookups)
	if err != nil {
		ReportError(err, "Connect to nsq failed")
		return ""
	}
	time.Sleep(time.Duration(rfor) * time.Minute)
	return ""
}

func (h *MyTestHandler) HandleMessage(message *nsq.Message) error {
	return ReportReaderStats(h.ctx, message)

}

func ReportReaderStats(ctx context.Context, currentStats *nsq.Message) error {
	state := lib.GetState(ctx)
	err := errors.New("state is nil")

	if state == nil {
		ReportError(err, "Cannot determine state")
		return err
	}

	tags := make(map[string]string)
	tags["nsqdAddress"] = currentStats.NSQDAddress
	tags["partition"] = currentStats.Partition

	now := time.Now()

	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Time:   now,
		Metric: ReaderAttempt,
		Tags:   stats.IntoSampleTags(&tags),
		Value:  float64(currentStats.Attempts),
	})

	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Time:   now,
		Metric: ReaderMessageID,
		Tags:   stats.IntoSampleTags(&tags),
		Value:  float64(nsq.GetNewMessageID(currentStats.ID[:])),
	})

	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Time:   now,
		Metric: ReaderTimestamp,
		Tags:   stats.IntoSampleTags(&tags),
		Value:  float64(currentStats.Timestamp),
	})
	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Time:   now,
		Metric: ReaderBytes,
		Tags:   stats.IntoSampleTags(&tags),
		Value:  float64(len(currentStats.Body)),
	})
	return nil
}

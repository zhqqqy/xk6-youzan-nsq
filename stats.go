package xk6_youzan_nsq

import "go.k6.io/k6/stats"

var (
	ReaderAttempt      = stats.New("nsq.reader.attempt.count", stats.Counter)
	ReaderMessageID    = stats.New("nsq.reader.messageID.count", stats.Counter)
	ReaderTimestamp  = stats.New("nsq.reader.message.Timestamp", stats.Counter,stats.Time)
	ReaderBytes      = stats.New("nsq.reader.message.bytes", stats.Counter, stats.Data)
	ReaderRebalances = stats.New("nsq.reader.rebalance.count", stats.Counter)
	ReaderTimeouts   = stats.New("nsq.reader.timeouts.count", stats.Counter)
	ReaderErrors     = stats.New("nsq.reader.error.count", stats.Counter)

	ReaderDialTime   = stats.New("nsq.reader.dial.seconds", stats.Trend, stats.Time)
	ReaderReadTime   = stats.New("nsq.reader.read.seconds", stats.Trend, stats.Time)
	ReaderWaitTime   = stats.New("nsq.reader.wait.seconds", stats.Trend, stats.Time)
	ReaderFetchSize  = stats.New("nsq.reader.fetch.size", stats.Counter)
	ReaderFetchBytes = stats.New("nsq.reader.fetch.bytes", stats.Counter, stats.Data)

	ReaderOffset        = stats.New("nsq.reader.offset", stats.Gauge)
	ReaderLag           = stats.New("nsq.reader.lag", stats.Gauge)
	ReaderMinBytes      = stats.New("nsq.reader.fetch_bytes.min", stats.Gauge)
	ReaderMaxBytes      = stats.New("nsq.reader.fetch_bytes.max", stats.Gauge)
	ReaderMaxWait       = stats.New("nsq.reader.fetch_wait.max", stats.Gauge, stats.Time)
	ReaderQueueLength   = stats.New("nsq.reader.queue.length", stats.Gauge)
	ReaderQueueCapacity = stats.New("nsq.reader.queue.capacity", stats.Gauge)

	WriterDials      = stats.New("nsq.writer.dial.count", stats.Counter)
	WriterWrites     = stats.New("nsq.writer.write.count", stats.Counter)
	WriterMessages   = stats.New("nsq.writer.message.count", stats.Counter)
	WriterBytes      = stats.New("nsq.writer.message.bytes", stats.Counter, stats.Data)
	WriterRebalances = stats.New("nsq.writer.rebalance.count", stats.Counter)
	WriterErrors     = stats.New("nsq.writer.error.count", stats.Counter)

	WriterDialTime   = stats.New("nsq.writer.dial.seconds", stats.Trend, stats.Time)
	WriterWriteTime  = stats.New("nsq.writer.write.seconds", stats.Trend, stats.Time)
	WriterWaitTime   = stats.New("nsq.writer.wait.seconds", stats.Trend, stats.Time)
	WriterRetries    = stats.New("nsq.writer.retries.count", stats.Counter)
	WriterBatchSize  = stats.New("nsq.writer.batch.size", stats.Counter)
	WriterBatchBytes = stats.New("nsq.writer.batch.bytes", stats.Counter, stats.Data)

	WriterMaxAttempts       = stats.New("nsq.writer.attempts.max", stats.Gauge)
	WriterMaxBatchSize      = stats.New("nsq.writer.batch.max", stats.Gauge)
	WriterBatchTimeout      = stats.New("nsq.writer.batch.timeout", stats.Gauge, stats.Time)
	WriterReadTimeout       = stats.New("nsq.writer.read.timeout", stats.Gauge, stats.Time)
	WriterWriteTimeout      = stats.New("nsq.writer.write.timeout", stats.Gauge, stats.Time)
	WriterRebalanceInterval = stats.New("nsq.writer.rebalance.interval", stats.Gauge, stats.Time)
	WriterRequiredAcks      = stats.New("nsq.writer.acks.required", stats.Gauge)
	WriterAsync             = stats.New("nsq.writer.async", stats.Rate)
	WriterQueueLength       = stats.New("nsq.writer.queue.length", stats.Gauge)
	WriterQueueCapacity     = stats.New("nsq.writer.queue.capacity", stats.Gauge)
)

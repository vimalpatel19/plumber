package rstreams

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/batchcorp/plumber/cli"
	"github.com/batchcorp/plumber/dproxy"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func Dynamic(opts *cli.Options) error {
	log := logrus.WithField("pkg", "rstreams/dynamic")

	// Start up writer
	writer, err := NewStreamsClient(opts)
	if err != nil {
		return errors.Wrap(err, "unable to connect to Redis Streams")
	}

	defer writer.Close()

	// Start up dynamic connection
	grpc, err := dproxy.New(opts, "redis-streams")
	if err != nil {
		return errors.Wrap(err, "could not establish connection to Batch")
	}

	go grpc.Start()

	// Continually loop looking for messages on the channel.
	for {
		select {
		case outbound := <-grpc.OutboundMessageCh:
			for _, streamName := range opts.RedisStreams.Streams {
				_, err := writer.XAdd(context.Background(), &redis.XAddArgs{
					Stream: streamName,
					ID:     opts.RedisStreams.WriteID,
					Values: map[string]interface{}{
						opts.RedisStreams.WriteKey: outbound.Blob,
					},
				}).Result()
				if err != nil {
					log.Errorf("Failed to publish message to stream '%s': %s", streamName, err)
					continue
				}

				log.Debugf("Replayed message to Redis stream '%s' for replay '%s'", streamName, outbound.ReplayId)
			}
		}
	}
}

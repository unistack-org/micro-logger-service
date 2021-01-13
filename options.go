package service

import (
	"github.com/unistack-org/micro/v3/client"
	"github.com/unistack-org/micro/v3/logger"
	"github.com/unistack-org/micro/v3/store"
)

type clientKey struct{}

// Client to call logger service
func Client(c client.Client) logger.Option {
	return logger.SetOption(clientKey{}, c)
}

type storeKey struct{}

// Store to hold messages while logger service is unavailable
func Store(s store.Store) logger.Option {
	return logger.SetOption(storeKey{}, s)
}

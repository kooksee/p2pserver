package config

import (
	"sync"

	"github.com/patrickmn/go-cache"
	"github.com/dgraph-io/badger"
	"github.com/kooksee/log"
)

var (
	once     sync.Once
	instance *Config
	logger   = log.New("package", "config")
)

type Config struct {
	Debug    bool
	LogLevel string
	Home     string
	UdpPort  int
	UdpHost  string

	HttpPort int
	HttpHost string

	AdvertiseUdpAddr  string
	AdvertiseHttpAddr string

	Cache *cache.Cache
	Db    *badger.DB

	Seeds []string
	PriV string
}

package config

import (
	"sync"

	"github.com/patrickmn/go-cache"
	"github.com/dgraph-io/badger"
	"github.com/kooksee/log"
	"time"
	"path"
)

var (
	once     sync.Once
	instance *Config
)

type Config struct {
	l        log.Logger
	Cache    *cache.Cache `mapstructure:"_" yaml:"-"`
	Db       *badger.DB   `mapstructure:"_" yaml:"-"`
	Home     string       `mapstructure:"_" yaml:"_"`
	DbPath   string       `mapstructure:"db_path" yaml:"db_path"`

	LogPath  string       `mapstructure:"log_path" yaml:"log_path"`
	LogLevel string       `mapstructure:"log_level" yaml:"log_level"`

	UdpPort  int          `mapstructure:"udp_port" yaml:"udp_port"`
	UdpHost  string       `mapstructure:"udp_host" yaml:"udp_host"`

	HttpPort int    `mapstructure:"http_port" yaml:"http_port"`
	HttpHost string `mapstructure:"http_host" yaml:"http_host"`

	ExtIP string `mapstructure:"ext_ip" yaml:"ext_ip"`

	AdvertiseUdpAddr  string `mapstructure:"advertise_udp_addr" yaml:"advertise_udp_addr"`
	AdvertiseHttpAddr string `mapstructure:"advertise_http_addr" yaml:"advertise_http_addr"`

	Seeds []string `mapstructure:"seeds" yaml:"seeds"`
	PriV  string   `mapstructure:"priv" yaml:"priv"`
}

func GetCfg() *Config {
	once.Do(func() {
		instance = &Config{
			UdpPort:           8081,
			UdpHost:           "0.0.0.0",
			HttpHost:          "0.0.0.0",
			HttpPort:          8080,
			AdvertiseHttpAddr: "0.0.0.0:8080",
			AdvertiseUdpAddr:  "0.0.0.0:8081",
			LogLevel:          "debug",
			Cache:             cache.New(time.Minute, 5*time.Minute),
			Home:              "kdata",
			DbPath:            "",
			LogPath:           "",
		}
		if instance.DbPath == "" {
			instance.DbPath = path.Join(instance.Home, "db")
		}

		if instance.LogPath == "" {
			instance.LogPath = path.Join(instance.Home, "log")
		}

	})
	return instance
}

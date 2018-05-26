package config

import (
	"github.com/dgraph-io/badger"
	"time"
	"os"
	log "github.com/inconshreveable/log15"
	"github.com/patrickmn/go-cache"
	"path"
	"github.com/kooksee/sp2p"
)

func GetLog(ctx ...interface{}) log.Logger {
	return Log().New(ctx...)
}

func GetDb() *badger.DB {
	cfg := GetCfg()
	if cfg.db == nil {
		panic("please init db")
	}
	return cfg.db
}

func GetP2p() *sp2p.SP2p {
	cfg := GetCfg()
	if cfg.p2p == nil {
		panic("please init sp2p")
	}
	return cfg.p2p
}

func Log() log.Logger {
	cfg := GetCfg()
	if cfg.l == nil {
		panic("please init log")
	}
	return cfg.l
}

func GetCfg() *Config {
	if instance == nil {
		panic("please init config")
	}
	return instance
}

func GetCache() *cache.Cache {
	return GetCfg().cache
}

func NewCfg(defaultHomeDir string) *Config {
	defaultHomeDir = GetHomeDir(defaultHomeDir)
	instance = &Config{
		UdpPort:           8081,
		UdpHost:           "0.0.0.0",
		HttpHost:          "0.0.0.0",
		HttpPort:          8080,
		AdvertiseHttpAddr: "",
		AdvertiseUdpAddr:  "",
		LogLevel:          "debug",
		cache:             cache.New(time.Minute, 5*time.Minute),
		home:              defaultHomeDir,
		Version:           "1.0.0",
	}

	if instance.DbPath == "" {
		instance.DbPath = path.Join(instance.home, "db")
	}

	if instance.LogPath == "" {
		instance.LogPath = path.Join(instance.home, "log")
	}

	if instance.PriV == "" {
		instance.PriV = path.Join(instance.home, "private.key")
	}

	if instance.configPath == "" {
		instance.configPath = path.Join(instance.home, "config.yaml")
	}

	return instance
}

func GetHomeDir(defaultHome string) string {
	if len(os.Args) > 2 && os.Args[len(os.Args)-2] == "--home" {
		defaultHome = os.Args[len(os.Args)-1]
		os.Args = os.Args[:len(os.Args)-2]
	}
	return defaultHome
}

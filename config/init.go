package config

import (
	"sync"

	"github.com/patrickmn/go-cache"
	"github.com/dgraph-io/badger"
	log "github.com/inconshreveable/log15"
	"time"
	"os"
	"net/http"
	"io/ioutil"
	"bytes"
	"github.com/spf13/viper"
	"github.com/kooksee/sp2p"
	"net"
)

var (
	once     sync.Once
	instance *Config
)

type Config struct {
	l          log.Logger
	configPath string
	cache      *cache.Cache
	db         *badger.DB
	home       string
	p2p        *sp2p.SP2p

	Version  string `mapstructure:"version" yaml:"version"`
	DbPath   string `mapstructure:"db_path" yaml:"db_path"`
	LogPath  string `mapstructure:"log_path" yaml:"log_path"`
	LogLevel string `mapstructure:"log_level" yaml:"log_level"`

	UdpPort int    `mapstructure:"udp_port" yaml:"udp_port"`
	UdpHost string `mapstructure:"udp_host" yaml:"udp_host"`

	HttpPort int    `mapstructure:"http_port" yaml:"http_port"`
	HttpHost string `mapstructure:"http_host" yaml:"http_host"`

	ExtIP string `mapstructure:"ext_ip" yaml:"ext_ip"`

	AdvertiseUdpAddr  string `mapstructure:"advertise_udp_addr" yaml:"advertise_udp_addr"`
	AdvertiseHttpAddr string `mapstructure:"advertise_http_addr" yaml:"advertise_http_addr"`

	Seeds []string `mapstructure:"seeds" yaml:"seeds"`
	PriV  string   `mapstructure:"priv" yaml:"priv"`
}

func (t *Config) InitConfigFile() {
	if _, err := os.Stat(t.configPath); os.IsNotExist(err) {
		return
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath(t.home)

	if err := v.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	if err := v.Unmarshal(t); err != nil {
		panic(err.Error())
	}
}

// 获取外网地址
func (t *Config) InitExtIp() {
	logger := Log()
	for {
		resp, err := http.Get("http://ipinfo.io/ip")
		if err != nil {
			logger.Error("获取外网地址失败", "err", err)
			time.Sleep(time.Second * 2)
			continue
		}
		extIp, _ := ioutil.ReadAll(resp.Body)
		t.ExtIP = string(bytes.TrimSpace(extIp))
		if t.ExtIP == "" {
			logger.Error("获取不到外网IP")
			time.Sleep(time.Second * 2)
			continue
		}
		logger.Info("获取外网IP", "ip", t.ExtIP)
		break
	}
}
func (t *Config) InitP2pConfig(seeds ... string) {
	kcfg := sp2p.DefaultKConfig()
	kcfg.InitDb(GetDb())
	kcfg.InitLog(Log())

	kcfg.Host = t.UdpHost
	kcfg.Port = t.UdpPort

	addr, err := net.ResolveUDPAddr("udp", t.AdvertiseUdpAddr)
	if err != nil {
		Log().Error(err.Error())
		panic(err.Error())
	}
	kcfg.AdvertiseAddr = addr
	kcfg.Seeds = seeds

	t.p2p = sp2p.NewSP2p()
}

func (t *Config) InitDb() {
	opts := badger.DefaultOptions
	opts.Dir = t.DbPath
	opts.ValueDir = t.DbPath
	db, err := badger.Open(opts)
	if err != nil {
		Log().Error(err.Error())
		panic(err.Error())
	}
	t.db = db
}

func (t *Config) InitLog() {
	t.l = log.New()
	if t.LogLevel != "error" {
		ll, err := log.LvlFromString(t.LogLevel)
		if err != nil {
			panic(err.Error())
		}
		t.l.SetHandler(log.LvlFilterHandler(ll, log.StreamHandler(os.Stdout, log.TerminalFormat())))
	} else {
		h, err := log.FileHandler(t.LogPath, log.LogfmtFormat())
		if err != nil {
			t.l.Error(err.Error())
			panic(err.Error())
		}
		log.MultiHandler(
			log.LvlFilterHandler(log.LvlError, log.StreamHandler(os.Stderr, log.LogfmtFormat())),
			h,
		)
	}
}

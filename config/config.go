package config

import (
	"github.com/spf13/viper"
	"encoding/json"
	"net/http"
	"time"
	"io/ioutil"
	"bytes"
	"os"
	"github.com/kooksee/log"
	"github.com/dgraph-io/badger"
	"github.com/kooksee/sp2p"
	"github.com/kooksee/crypt"
)

func (t *Config) LoadConfigFile() {
	if _, err := ioutil.ReadFile(t.configPath); os.IsNotExist(err) {
		return
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("config")

	v.AddConfigPath(t.Home)

	if err := v.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	if err := v.Unmarshal(t); err != nil {
		panic(err.Error())
	}

	if t.LogLevel != "error" {
		d, _ := json.Marshal(t)
		t.l.Debug("config")
		t.l.Debug(string(d))
	}
}

// 获取外网地址
func (t *Config) GetExtIp() {
	logger := t.l
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
func (t *Config) InitP2pConfig() {
	kcfg := sp2p.DefaultKConfig()
	priv, err := crypto.LoadECDSA(t.PriV)
	if err != nil {
		panic(err.Error())
	}

	kcfg.PriV = priv
	kcfg.Db = t.Db
	kcfg.Host = t.UdpHost
	kcfg.Port = t.UdpPort
	kcfg.LogLevel = t.LogLevel
	sp2p.SetCfg(kcfg)
}

func (t *Config) InitDb() {
	opts := badger.DefaultOptions
	opts.Dir = t.DbPath
	opts.ValueDir = t.DbPath
	db, err := badger.Open(opts)
	if err != nil {
		panic(err.Error())
	}
	t.Db = db
}

func (t *Config) Log() log.Logger {
	if t.l == nil {
		panic("please init log")
	}
	return t.l
}

func (t *Config) GetLog(ctx ...interface{}) log.Logger {
	if t.l == nil {
		panic("please init log")
	}
	return t.l.New(ctx...)
}

func (t *Config) InitLog() {
	t.l = log.New()
	if t.LogLevel != "error" {
		ll, err := log.LvlFromString(t.LogLevel)
		if err != nil {
			panic(err.Error())
		}
		t.l.SetHandler(log.LvlFilterHandler(ll, log.StreamHandler(os.Stdout, log.TerminalFormat(true))))
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

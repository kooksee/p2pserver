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
)

func (t *Config) LoadConfigFile() {

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
		t.l.Info(string(d))
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

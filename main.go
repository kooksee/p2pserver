package main

import (
	"flag"
	"runtime"
	"github.com/dgraph-io/badger"
	"time"
	"github.com/patrickmn/go-cache"
	"github.com/kooksee/p2pserver/config"
	"github.com/kooksee/p2pserver/node"
)

const Version = "1.0"

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GC()

	cfg := config.GetCfg()
	flag.StringVar(&cfg.Home, "home", cfg.Home, "app config home")
	flag.StringVar(&cfg.LogLevel, "log", cfg.LogLevel, "app log level")
	flag.IntVar(&cfg.UdpPort, "udp", cfg.UdpPort, "app udp port")
	flag.IntVar(&cfg.HttpPort, "http", cfg.HttpPort, "app http port")
	flag.StringVar(&cfg.Pv, "pv", cfg.Pv, "app priV")

	flag.Parse()

	cfg.Cache = cache.New(time.Minute, 5*time.Minute)

	opts := badger.DefaultOptions
	opts.Dir = cfg.Home
	opts.ValueDir = cfg.Home
	db, err := badger.Open(opts)
	if err != nil {
		panic(err.Error())
	}
	cfg.Db = db

	node.SetCfg(cfg)

	cfg.Seeds = append(cfg.Seeds,
		"enode://e48df98fe81f7fddc7d481ce92d2f1559143d7ca7bda77f363778d1f4ac8aa9542ebbf41b79949c6acbd7efd571bb3199dd531ad2cbfa9860c25395a623bfff8@127.0.0.1:9091",
		"enode://1573401e0ce91c823d543bc8f13acf4654684d2f37929f0fee32720d03cd5d9facf4a934a2a56063804be8f7bef69485fff67c30ad4bb12be7e6244ffc89c9b1@127.0.0.1:9092",
		"enode://cd68adcb908e550875fb3c7e35c9313d616841d1d1a093bc03f72067c1437bfb16242a5e6ea305b1ea1b2b7a1dc90f662a384ebccc71c57d4814fdcbbb703bbc@127.0.0.1:9093",
		"enode://0c2ec16892f65e2b38771f95fa4b75ebe8e0352e89cd8faa003105b34471e015af935ca18b3e9bda8fb09b54a70839a57a3feebf8af1e331d5b98b4fad6405da@127.0.0.1:9094",
		"enode://3caab4cc114adb8975a7ac58bb70d40a590bd88dddcfb9cc611b0d8914e39f911afe1b4d930c86deb70aae44b53bc70903d8046926fe3ec0b7ca179af40cef54@127.0.0.1:9095",
	)

	node.New(cfg.Seeds).RunHttpServer()
	select {}
}

package node

import (
	"github.com/json-iterator/go"
	"github.com/kooksee/log"
	"github.com/kooksee/sp2p"
	"github.com/kooksee/p2pserver/config"
)

var (
	json   = jsoniter.ConfigCompatibleWithStandardLibrary
	logger = log.New("package", "node")
	hm     = sp2p.GetHManager()
	cfg    *config.Config
)

func SetCfg(cfg1 *config.Config) {
	cfg = cfg1
}

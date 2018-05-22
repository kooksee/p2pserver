package node

import (
	"github.com/json-iterator/go"
	"github.com/kooksee/sp2p"
	"github.com/kooksee/p2pserver/config"
	"github.com/kooksee/log"
)

var (
	json   = jsoniter.ConfigCompatibleWithStandardLibrary
	logger log.Logger
	hm     *sp2p.HandleManager
	cfg    *config.Config
)

func SetCfg(c *config.Config) {
	cfg = c
	hm = sp2p.GetHManager()
	logger = cfg.GetLog("package", "node")
}

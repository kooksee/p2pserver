package node

import (
	"github.com/json-iterator/go"
	"github.com/kooksee/sp2p"
	"github.com/kooksee/p2pserver/config"
	"github.com/inconshreveable/log15"
)

var (
	json   = jsoniter.ConfigCompatibleWithStandardLibrary
	logger log15.Logger
	hm     *sp2p.HandleManager
	cfg    *config.Config
	p2p    *sp2p.SP2p
)

func Init() {
	cfg = config.GetCfg()
	hm = sp2p.GetHManager()
	logger = config.GetLog("package", "node")
	p2p = config.GetP2p()
}

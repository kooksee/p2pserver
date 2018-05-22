package node

import (
	"net/http"
	"fmt"
	"bytes"
	"github.com/kooksee/sp2p"
	"github.com/gin-gonic/gin"
	"github.com/kooksee/crypt"
)

func New(seeds []string) *KVNode {
	kcfg := sp2p.DefaultKConfig()
	priv, err := crypto.LoadECDSA(cfg.PriV)
	if err != nil {
		panic(err.Error())
	}

	kcfg.PriV = priv
	kcfg.Db = cfg.Db
	kcfg.Host = cfg.UdpHost
	kcfg.Port = cfg.UdpPort
	kcfg.LogLevel = cfg.LogLevel
	sp2p.SetCfg(kcfg)

	p2p := sp2p.NewSP2p(seeds)
	fmt.Println(p2p.GetTable().GetNode().String())
	return &KVNode{SP2p: p2p}
}

type KVNode struct {
	*sp2p.SP2p
}

func (n *KVNode) RunHttpServer() {
	router := gin.Default()
	router.POST("/kv", n.indexPost)
	router.GET("/kv/:id", n.indexGet)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", "0.0.0.0", cfg.HttpPort), router); err != nil {
		panic(err)
	}
}

func (n *KVNode) indexPost(c *gin.Context) {

	message, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	message = bytes.TrimSpace(message)
	logger.Debug("message data", "data", string(message))

	msg := &sp2p.KVSetReq{}
	if err := json.Unmarshal(message, msg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	n.Write(&sp2p.KMsg{Data: msg})
	c.JSON(http.StatusOK, "ok")
}

func (n *KVNode) indexGet(c *gin.Context) {
	sid := c.GetString("id")
	d, _ := cfg.Cache.Get(sid)
	if d != nil {
		c.JSON(http.StatusOK, string(d.([]byte)))
		return
	}
	c.JSON(http.StatusOK, "not found")
}

package node

import "github.com/kooksee/sp2p"

type WSReq struct {
	N int `json:"n,omitempty"`
}

func (t *WSReq) T() byte               { return WSReqT }
func (t *WSReq) String() string        { return WSReqS }
func (t *WSReq) Create() sp2p.IMessage { return &WSReq{} }
func (t *WSReq) OnHandle(p *sp2p.SP2p, msg *sp2p.KMsg) {
}

type WSResp struct {
	Nodes []string `json:"nodes,omitempty"`
}

func (t *WSResp) T() byte               { return WSRespT }
func (t *WSResp) String() string        { return WSRespS }
func (t *WSResp) Create() sp2p.IMessage { return &WSResp{} }
func (t *WSResp) OnHandle(p *sp2p.SP2p, msg *sp2p.KMsg) {
}

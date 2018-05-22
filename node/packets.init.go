package node

import "github.com/kooksee/sp2p"

func init() {

	// 注册数据类型
	hm := sp2p.GetHManager()
	sp2p.MustNotErr(hm.Registry(WSReqT, &WSReq{}))
	sp2p.MustNotErr(hm.Registry(WSRespT, &WSResp{}))
}

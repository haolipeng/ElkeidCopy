package pool

import (
	"github.com/patrickmn/go-cache"
)

type GRPCPool struct {
	//connPool is the global agent data Cache.
	//the key is AgentID,which used to identify a agent,must be unique.
	//the value is *Connection
	connPool *cache.Cache
}

//连接信息
type Connection struct {
	AgentID    string `json:"agent_id"`
	SourceAddr string `json:"addr"`
}

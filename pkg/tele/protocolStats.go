package tele

import (
	"sync"

	"github.com/silverton-io/honeypot/pkg/protocol"
)

type ProtocolStats struct {
	vmu     sync.Mutex
	imu     sync.Mutex
	Invalid map[string]map[string]int64 `json:"invalid"`
	Valid   map[string]map[string]int64 `json:"valid"`
}

func (ps *ProtocolStats) Build() {
	var vProtoStat = make(map[string]map[string]int64)
	var invProtoStat = make(map[string]map[string]int64)
	ps.Valid = vProtoStat
	ps.Invalid = invProtoStat
	for _, protocol := range protocol.GetIntputProtocols() {
		var vEventStat = make(map[string]int64)
		var invEventStat = make(map[string]int64)
		ps.Valid[protocol] = vEventStat
		ps.Invalid[protocol] = invEventStat
	}
}

func (ps *ProtocolStats) IncrementValid(protocol string, event string, count int64) {
	i := ps.Valid[protocol][event]
	ps.vmu.Lock()
	defer ps.vmu.Unlock()
	ps.Valid[protocol][event] = i + count
}

func (ps *ProtocolStats) IncrementInvalid(protocol string, event string, count int64) {
	i := ps.Invalid[protocol][event]
	ps.imu.Lock()
	defer ps.imu.Unlock()
	ps.Invalid[protocol][event] = i + count
}

func BuildProtocolStats() *ProtocolStats {
	ps := ProtocolStats{}
	ps.Build()
	return &ps
}
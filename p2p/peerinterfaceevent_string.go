// Code generated by "stringer -type=peerInterfaceEvent"; DO NOT EDIT.

package p2p

import "strconv"

const _peerInterfaceEvent_name = "ActivePeerMessagePeerDisconnectedPeer"

var _peerInterfaceEvent_index = [...]uint8{0, 10, 21, 37}

func (i peerInterfaceEvent) String() string {
	if i < 0 || i >= peerInterfaceEvent(len(_peerInterfaceEvent_index)-1) {
		return "peerInterfaceEvent(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _peerInterfaceEvent_name[_peerInterfaceEvent_index[i]:_peerInterfaceEvent_index[i+1]]
}
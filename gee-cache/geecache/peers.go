/*
@Copyright : Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
@File : peers.go
@Time : 2023/6/2 16:36
@Author : junerwang@tencent.com
@Desc: this is a simple description...
*/

package geecache

import pb "geecache/geecachepb"

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key.
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}

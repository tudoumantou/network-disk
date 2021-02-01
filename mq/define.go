package mq

import "network-disk/common"

type TransferData struct {
	FileHash string
	CurLocation string
	DestLocation string
	DestStoreType common.StoreType
}

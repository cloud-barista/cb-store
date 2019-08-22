// CB-Store is a common repository for managing Meta Info of Cloud-Barista.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
//
// by powerkim@etri.re.kr, 2019.08.

package cbstore

import (
	"github.com/cloud-barista/cb-store/config"
	icbs "github.com/cloud-barista/cb-store/interfaces"
	lfsdrv "github.com/cloud-barista/cb-store/store-drivers/lfs-driver"
	etcddrv "github.com/cloud-barista/cb-store/store-drivers/etcd-driver"
)

var configInfo *config.CBSTORECONFIG

func init() {
	config.Cblogger.Info("calling!")
	configInfo = config.GetConfigInfos()
}

// initialize db
func InitStore() {
	if configInfo.STORETYPE == "LFS" {
		// 1. remove path: rm -rf ./meta_store/*
		// @todo init lfs metainfo
	}
	if configInfo.STORETYPE == "ETCD" {
		// @todo init etcd metainfo
	}
}

func GetStore() icbs.Store {
	if configInfo.STORETYPE == "LFS" {
		store := lfsdrv.LFSDriver{}
		return &store
	}
	if configInfo.STORETYPE == "ETCD" {
		store := etcddrv.ETCDDriver{}
		return &store
	}
	config.Cblogger.Errorf("STORETYPE:" + configInfo.STORETYPE +" is not supported!!")

	return nil
}


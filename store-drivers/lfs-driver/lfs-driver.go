// CB-Store is a common repository for managing Meta Info of Cloud-Barista.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
//
// by powerkim@etri.re.kr, 2019.07.

package cbstore

import (
	icbs "github.com/cloud-barista/cb-store/interfaces"
)

type LFSDriver struct{}

func (lfsDriver *LFSDriver) Put(key string, value string) error {

	return nil
}

func (lfsDriver *LFSDriver) Get(key string) (*icbs.KeyValue, error) {

	return nil, nil
}

func (lfsDriver *LFSDriver) GetAll(key string, sortAscend bool) ([]*icbs.KeyValue, error) {

	return nil, nil
}

func (lfsDriver *LFSDriver)Delete(key string) error {

	return nil
}


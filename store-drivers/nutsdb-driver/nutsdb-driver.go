// CB-Store is a common repository for managing Meta Info of Cloud-Barista.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
//
// by powerkim@etri.re.kr, 2019.07.

package cbstore

import (
	_ "io/ioutil"
	_ "os"

	"github.com/cloud-barista/cb-store/config"
	"github.com/xujiajun/nutsdb"
	icbs "github.com/cloud-barista/cb-store/interfaces"
)

type NUTSDBDriver struct{}

var (
        db     *nutsdb.DB
        bucket string
)

func init() {
	fileDir := config.GetConfigInfos().NUTSDB.DBPATH
        //fileDir := "/tmp/nutsdb_example"

        opt := nutsdb.DefaultOptions
/* clean db files @todo make this into initDB() by powerkim, 2019.09.09
        files, _ := ioutil.ReadDir(fileDir)
        for _, f := range files {
                name := f.Name()
                if name != "" {
                        //fmt.Println(fileDir + "/" + name)
                        err := os.RemoveAll(fileDir + "/" + name)
                        if err != nil {
                                panic(err)
                        }
                }
        }
*/
        opt.Dir = fileDir
        opt.SegmentSize = config.GetConfigInfos().NUTSDB.SEGMENTSIZE

        //opt.SegmentSize = 1024 * 1024 // 1MB
        db, _ = nutsdb.Open(opt)
        bucket = "bucketForString"
}


func (nutsdbDriver *NUTSDBDriver) Put(key string, value string) error {
	config.Cblogger.Info("Key:" + key  + ", value:" + value)

        if err := db.Update(
                func(tx *nutsdb.Tx) error {
                        key := []byte(key)
                        val := []byte(value)
                        return tx.Put(bucket, key, val, 0)
                }); err != nil {
			config.Cblogger.Error(err)
			return err
        }

	return nil
}

func (nutsdbDriver *NUTSDBDriver) Get(key string) (*icbs.KeyValue, error) {
	config.Cblogger.Info("Key:" + key)

	var keyValue icbs.KeyValue
        if err := db.View(
                func(tx *nutsdb.Tx) error {
                        key := []byte(key)
                        e, err := tx.Get(bucket, key)
                        if err != nil {
				config.Cblogger.Error(err)
                                return err
                        }
			keyValue = icbs.KeyValue{string(key), string(e.Value)}
			return nil
                }); err != nil {
			config.Cblogger.Error(err)
			return nil, err
		}

	return &keyValue, nil
}

func (nutsdbDriver *NUTSDBDriver) GetList(key string, sortAscend bool) ([]*icbs.KeyValue, error) {
        config.Cblogger.Info("Key:" + key)

        var keyValueList []*icbs.KeyValue
        if err := db.View(
                func(tx *nutsdb.Tx) error {
                        key := []byte(key)
                        entries, err := tx.PrefixScan(bucket, key, 100)
                        if err != nil {
                                config.Cblogger.Error(err)
                                return err
                        }
			keyValueList = make([]*icbs.KeyValue, len(entries))
			if sortAscend {
				for k, entry := range entries {
					tmpOne := icbs.KeyValue{string(entry.Key), string(entry.Value)}
					keyValueList[k] = &tmpOne
				}
			} else {
				for k, entry := range entries {
					tmpOne := icbs.KeyValue{string(entry.Key), string(entry.Value)}
					keyValueList[len(entries)-1-k] = &tmpOne
				}
			}
                        return nil
                }); err != nil {
                        config.Cblogger.Error(err)
                        return nil, err
                }

        return keyValueList, nil
}

func (nutsdbDriver *NUTSDBDriver)Delete(key string) error {
	config.Cblogger.Info("Key:" + key)

	if err := db.Update(
		func(tx *nutsdb.Tx) error {
		key := []byte(key)
		if err := tx.Delete(bucket, key); err != nil {
			config.Cblogger.Error(err)
			return err
		}
		return nil
	}); err != nil {
		config.Cblogger.Error(err)
		return err
	}
	return nil
}


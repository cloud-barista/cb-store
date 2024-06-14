// CB-Log: Logger for Cloud-Barista.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// load and set config file
//
// ref) https://github.com/go-yaml/yaml/tree/v3
//	https://godoc.org/gopkg.in/yaml.v3
//
// by powerkim@powerkim.co.kr, 2019.08.

package config

import (
	"io/ioutil"
	"os"
	"strings"

	cblog "github.com/cloud-barista/cb-log"
	"github.com/sirupsen/logrus"

	"gopkg.in/yaml.v3"
)

type CBSTORECONFIG struct {
	STORETYPE string // option: NUTSDB | ETCD

	NUTSDB struct {
		DBPATH      string
		SEGMENTSIZE int64
	}

	ETCD struct {
		ETCDSERVERPORT string
	}
}

var Cblogger *logrus.Logger
var configInfo *CBSTORECONFIG

func init() {
	// cblog is a global variable.
	//Cblogger = cblog.GetLogger("CB-STORE")
	Cblogger = cblog.GetLogger("CLOUD-BARISTA") // by powerkim, 2019.09.24
}

func NewCBSTORECONFIG() CBSTORECONFIG {
    config := CBSTORECONFIG{}

    config.STORETYPE = "NUTSDB" // default store type

    config.NUTSDB.DBPATH = "./meta_db/dat"
    config.NUTSDB.SEGMENTSIZE = 1048576  // 1048576 1024*1024 (1MB)

    config.ETCD.ETCDSERVERPORT = "localhost:2379"

    return config
}

func load(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	return data, err
}

func GetConfigInfos() *CBSTORECONFIG {
	if configInfo != nil {
		return configInfo
	}

	cbstoreRootPath := os.Getenv("CBSTORE_ROOT")
	if cbstoreRootPath == "" {
	    Cblogger.Info("CBSTORE_ROOT is not set. Using default configuration")
	    config := NewCBSTORECONFIG()
	    return &config
	}

	data, err := load(cbstoreRootPath + "/conf/store_conf.yaml")

	if err != nil {
		Cblogger.Error(err)
		panic(err)
	}

	configInfo = new(CBSTORECONFIG)
	err = yaml.Unmarshal([]byte(data), &configInfo)
	if err != nil {
		Cblogger.Error(err)
		panic(err)
	}

	configInfo.NUTSDB.DBPATH = replaceEnvPath(configInfo.NUTSDB.DBPATH)
	return configInfo
}

// $ABC/def ==> /abc/def
func replaceEnvPath(str string) string {
	if strings.Index(str, "$") == -1 {
		return str
	}

	// ex) input "$CBSTORE_ROOT/meta_db/dat"
	strList := strings.Split(str, "/")
	for n, one := range strList {
		if strings.Index(one, "$") != -1 {
			cbstoreRootPath := os.Getenv(strings.Trim(one, "$"))
			if cbstoreRootPath == "" {
				Cblogger.Error(one + " is not set!")
			}
			strList[n] = cbstoreRootPath
		}
	}

	var resultStr string
	for _, one := range strList {
		resultStr = resultStr + one + "/"
	}
	// ex) "/root/go/src/github.com/cloud-barista/cb-spider/meta_db/dat/"
	resultStr = strings.TrimRight(resultStr, "/")
	resultStr = strings.ReplaceAll(resultStr, "//", "/")
	return resultStr
}

func GetConfigString(configInfos *CBSTORECONFIG) string {
	d, err := yaml.Marshal(configInfos)
	if err != nil {
		Cblogger.Error(err)
		//panic(err)
	}
	return string(d)
}

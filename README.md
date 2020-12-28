# cb-store
cb-store is a common repository for managing Meta Info of Cloud-Barista.
You can choose NUTSDB or ETCD for repository of cb-store.

  A.	NUTSDB: Embedded Key-Value Store on the Local Filesystem.
      - https://github.com/xujiajun/nutsdb
  
  B.	ETCD(Client V3.0): Distributed Key-Value Store
      - https://github.com/etcd-io/etcd

# 1.	install cb-store library pkg
  A.	$ go get github.com/cloud-barista/cb-store  
 
  B.  $ export CBSTORE_ROOT=~/go/src/github.com/cloud-barista/cb-store
    
  C.  $ vi conf/store_conf.yaml # set up storetype(NUTSDB|ETCD)
  
# 2.	example & test
  A.  example: https://github.com/cloud-barista/cb-store/blob/master/test/test.go
  
  B. install ETCD (Client V3.0): When using ETCD
    
  C.	$ cd test  
    
  D.	$ go run test.go 

      …
      =========================== Put(...)
      </> root
      </key1> value
      </key1> value1
      </key1/> value2
      </key1/%> value3%
      </key1/key2/key3> value4
      </space key> space value5
      </newline
       key> newline
       value6
      </a/b/c/123/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u> value/value/value
      ===========================
      =========================== Get("/")
      </> root
      ===========================
      =========================== Get("space key")
      </space key> space value5
      ===========================
      =========================== GetList("/", Ascending)
      </> root
      </a/b/c/123/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u> value/value/value
      </key1> value1
      </key1/> value2
      </key1/%> value3%
      </key1/key2/key3> value4
      </newline
       key> newline
       value6
      </space key> space value5
      ===========================
      =========================== GetList("/", Descending)
      </space key> space value5
      </newline
       key> newline
       value6
      </key1/key2/key3> value4
      </key1/%> value3%
      </key1/> value2
      </key1> value1
      </a/b/c/123/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u> value/value/value
      </> root

      …

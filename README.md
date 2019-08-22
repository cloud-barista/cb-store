# cb-store
CB-Store is a common repository for managing Meta Info of Cloud-Barista.

# 1.	install CB-Store library pkg
  A.	$ go get github.com/cloud-barista/cb-store
  
  B.  $ vi setup.env
  
  C.  $ source setup.env
  
# 2.	example with ETCD V3.0
  A.	https://github.com/cloud-barista/cb-store/blob/master/test/test_etcd.go

# 3.	test example
  
  A. install ETCD
  
  B.  $ vi conf/store_conf.yaml
  
  C.	$ cd test  
  
  D.	$ go run test_etcd.go

      …
      =========================== Get("/")
      </> root
      ===========================
      =========================== Get("space key")
      </space key> space value5
      ===========================
      =========================== GetAll("/", Ascending)
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
      =========================== GetAll("/", Descending)
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
      ===========================
      …

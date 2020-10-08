module github.com/cloud-barista/cb-store

go 1.12

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3
	// ETCD 환경이 없을때 Dial timeout 적용용
	github.com/coreos/etcd v3.3.18+incompatible => github.com/coreos/etcd v2.3.8+incompatible
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

	// build 시에 오류 제거용 (go mod 사용 버전 기준)
	// # github.com/coreos/etcd/clientv3/balancer/picker
	// go/pkg/mod/github.com/coreos/etcd@v3.3.18+incompatible/clientv3/balancer/picker/err.go:37:44: undefined: balancer.PickOptions
	// go/pkg/mod/github.com/coreos/etcd@v3.3.18+incompatible/clientv3/balancer/picker/roundrobin_balanced.go:55:54: undefined: balancer.PickOptions
	// # github.com/coreos/etcd/clientv3/balancer/resolver/endpoint
	// go/pkg/mod/github.com/coreos/etcd@v3.3.18+incompatible/clientv3/balancer/resolver/endpoint/endpoint.go:114:78: undefined: resolver.BuildOption
	// go/pkg/mod/github.com/coreos/etcd@v3.3.18+incompatible/clientv3/balancer/resolver/endpoint/endpoint.go:182:31: undefined: resolver.ResolveNowOption
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/cloud-barista/cb-log v0.2.0-cappuccino.0.20201008023843-31002c0a088d
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/etcd-io/etcd v2.3.8+incompatible
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/xujiajun/nutsdb v0.5.1-0.20200320023740-0cc84000d103
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200603094226-e3079894b1e8
)

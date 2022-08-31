module github.com/cloud-barista/cb-store

go 1.19

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt v3.2.1+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	gopkg.in/yaml.v3 => gopkg.in/yaml.v3 v3.0.0
)

require (
	github.com/cloud-barista/cb-log v0.6.3
	github.com/etcd-io/etcd v3.3.27+incompatible
	github.com/sirupsen/logrus v1.9.0
	github.com/xujiajun/nutsdb v0.10.0
	google.golang.org/grpc v1.49.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/coreos/bbolt v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/etcd v3.3.27+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20220810130054-c7d1c02cb6cf // indirect
	github.com/dgrijalva/jwt-go v0.0.0-00010101000000-000000000000 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/snowzach/rotatefilehook v0.0.0-20220211133110-53752135082d // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20201229170055-e5319fda7802 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	github.com/xujiajun/mmap-go v1.0.1 // indirect
	github.com/xujiajun/utils v0.0.0-20190123093513-8bf096c4f53b // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/net v0.0.0-20220826154423-83b083e8dc8b // indirect
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11 // indirect
	google.golang.org/genproto v0.0.0-20220829175752-36a9c930ecbf // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

retract (
	v0.3.13
	v0.3.12
	v0.3.11
)

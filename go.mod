module github.com/cloud-barista/cb-store

go 1.19

require (
	github.com/cloud-barista/cb-log v0.8.2
	github.com/sirupsen/logrus v1.9.0
	github.com/xujiajun/nutsdb v0.10.0
	go.etcd.io/etcd/client/v3 v3.5.4
	google.golang.org/grpc v1.56.3
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/snowzach/rotatefilehook v0.0.0-20220211133110-53752135082d // indirect
	github.com/xujiajun/mmap-go v1.0.1 // indirect
	github.com/xujiajun/utils v0.0.0-20190123093513-8bf096c4f53b // indirect
	go.etcd.io/etcd/api/v3 v3.5.4 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.4 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

retract (
	v0.3.13
	v0.3.12
	v0.3.11
)

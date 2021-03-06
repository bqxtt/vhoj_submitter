module github.com/ecnuvj/vhoj_submitter

go 1.14

require (
	github.com/aws/aws-sdk-go v1.36.33 // indirect
	github.com/ecnuvj/vhoj_common v0.0.0-20210204125811-c22717ad12a6
	github.com/ecnuvj/vhoj_db v0.0.0-00010101000000-000000000000
	github.com/ecnuvj/vhoj_rpc v0.0.0-20210318154530-31bf2dcf5cc5
	github.com/gojuukaze/YTask/v2 v2.3.1
	github.com/golang/mock v1.4.4 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/golang/snappy v0.0.2 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/streadway/amqp v1.0.0 // indirect
	github.com/tidwall/gjson v1.6.7 // indirect
	go.mongodb.org/mongo-driver v1.4.5 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/ecnuvj/vhoj_common => ../vhoj_common

replace github.com/ecnuvj/vhoj_db => ../vhoj_db

replace github.com/ecnuvj/vhoj_rpc => ../vhoj_rpc

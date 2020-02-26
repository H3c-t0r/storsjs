module storj.io/storj/cmd/gateway

go 1.13

// force specific versions for minio
require github.com/garyburd/redigo v1.0.1-0.20170216214944-0d253a66e6e1 // indirect

replace storj.io/storj => ../../

require storj.io/gateway v0.0.0-20200221161528-36f00ab51fa7

exclude gopkg.in/olivere/elastic.v5 v5.0.72 // buggy import, see https://github.com/olivere/elastic/pull/869

replace google.golang.org/grpc => github.com/storj/grpc-go v1.27.2-0.20200225082019-bd19b647a81c

require (
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/cheggaaa/pb v1.0.5-0.20160713104425-73ae1d68fe0b // indirect
	github.com/djherbis/atime v1.0.0 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/eclipse/paho.mqtt.golang v1.1.1 // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.0 // indirect
	github.com/fatih/structs v1.0.0 // indirect
	github.com/go-ini/ini v1.52.0 // indirect
	github.com/gorilla/handlers v1.4.0 // indirect
	github.com/gorilla/rpc v1.1.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-msgpack v0.5.3 // indirect
	github.com/hashicorp/raft v1.0.0 // indirect
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c // indirect
	github.com/inconshreveable/go-update v0.0.0-20160112193335-8152e7eb6ccf // indirect
	github.com/klauspost/cpuid v0.0.0-20180405133222-e7e905edc00e // indirect
	github.com/klauspost/reedsolomon v0.0.0-20180704173009-925cb01d6510 // indirect
	github.com/minio/dsync v0.0.0-20180124070302-439a0961af70 // indirect
	github.com/minio/highwayhash v0.0.0-20180501080913-85fc8a2dacad // indirect
	github.com/minio/lsync v0.0.0-20180328070428-f332c3883f63 // indirect
	github.com/minio/mc v0.0.0-20180926130011-a215fbb71884 // indirect
	github.com/minio/sio v0.0.0-20180327104954-6a41828a60f0 // indirect
	github.com/nats-io/gnatsd v1.3.0 // indirect
	github.com/nats-io/go-nats v1.6.0 // indirect
	github.com/nats-io/go-nats-streaming v0.4.2 // indirect
	github.com/nats-io/nats v1.6.0 // indirect
	github.com/nats-io/nats-streaming-server v0.12.2 // indirect
	github.com/nats-io/nuid v1.0.0 // indirect
	github.com/pascaldekloe/goe v0.0.0-20180627143212-57f6aae5913c // indirect
	github.com/pkg/profile v1.2.1 // indirect
	github.com/prometheus/procfs v0.0.0-20190517135640-51af30a78b0e // indirect
	github.com/rs/cors v1.5.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/streadway/amqp v0.0.0-20180806233856-70e15c650864 // indirect
	github.com/tidwall/gjson v1.1.3 // indirect
	github.com/tidwall/match v0.0.0-20171002075945-1731857f09b1 // indirect
	gopkg.in/Shopify/sarama.v1 v1.18.0 // indirect
	gopkg.in/cheggaaa/pb.v1 v1.0.25 // indirect
	gopkg.in/ini.v1 v1.52.0 // indirect
	gopkg.in/olivere/elastic.v5 v5.0.76 // indirect
)

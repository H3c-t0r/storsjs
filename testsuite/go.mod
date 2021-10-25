module storj.io/storj/testsuite

go 1.16

replace storj.io/storj => ../

require (
	github.com/go-rod/rod v0.101.8
	github.com/spacemonkeygo/monkit/v3 v3.0.17
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.17.0
	storj.io/common v0.0.0-20211021003554-f155346ccc11
	storj.io/gateway-mt v1.14.4-0.20211015103214-01eddbc864fb
	storj.io/private v0.0.0-20211022095408-ad8a59190f12
	storj.io/storj v0.12.1-0.20210916114455-b2d724962c24
)

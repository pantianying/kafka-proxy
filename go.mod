module github.com/grepplabs/kafka-proxy

go 1.14

require (
	github.com/Shopify/sarama v1.26.4
	github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/elazarl/goproxy v0.0.0-20200426045556-49ad98f6dac1
	github.com/elazarl/goproxy/ext v0.0.0-20200426045556-49ad98f6dac1
	github.com/fsnotify/fsnotify v1.4.9
	github.com/golang/protobuf v1.4.2
	github.com/golang/snappy v0.0.1 // indirect
	github.com/hashicorp/go-hclog v0.14.1
	github.com/hashicorp/go-multierror v1.1.0
	github.com/hashicorp/go-plugin v1.3.0
	github.com/klauspost/cpuid v1.3.0
	github.com/oklog/oklog v0.3.2
	github.com/pierrec/lz4 v2.5.2+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.6.0
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c
	golang.org/x/net v0.0.0-20200528225125-3c3fba18258b
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/api v0.28.0
	google.golang.org/grpc v1.30.0
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/ldap.v2 v2.5.1
)

replace github.com/Shopify/sarama => /Users/mac/github/pantianying/sarama

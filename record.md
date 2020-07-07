
```
./kafka-proxy server --bootstrap-server-mapping "172.17.40.166:9092,127.0.0.1:32400" \
                   --dynamic-listeners-disable
```

./kafka-proxy server --bootstrap-server-mapping "172.17.40.166:9092,0.0.0.0:32400" 

# 服务器上
./kafka-proxy server --bootstrap-server-mapping "kafka-dev-hb2.dian.so:9092,0.0.0.0:32400,kafka-dev-hb2.dian.so:32400"
nohup ./kafka-proxy server --bootstrap-server-mapping "kafka-dev-hb2.dian.so:9092,0.0.0.0:32400,kafka-dev-hb2.dian.so:32400" >> kafka-proxy.log 2>&1 &

```
/home/admin/kafka_2.12-2.0.0/bin/kafka-topics.sh --create --zookeeper localhost:2181 --topic test-topic-yangchun --replication-factor 1 --partitions 24
```

```
copy 逻辑在 copyThenClose,myCopyN

```

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

python -m SimpleHTTPServer
```
                   
              
                   
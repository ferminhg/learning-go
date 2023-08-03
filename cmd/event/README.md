# Introducing Apache Kafka: consumers & producers

## Ad Consumer

```shell
go run cmd/event/consumer/main.go -brokers localhost:9092 -topics topic.ads.1 -group example
```


## Examples

**How to run:**

Create & run kafka broker.
```shell
docker compose build
docker compose run
```

Create two topics `topic.access_log.1` and `topic.important.1`:

```shell
go run cmd/event/broker/broker.go -brokers localhost:9092 -topics topic.access_log.1,topic.important.1
```

Run http server producer:
```shell
go run cmd/event/producer/http_server.go -addr localhost:8090 -brokers localhost:9092 -verbose
```

Run consumer:
```shell
go run cmd/event/consumer/main.go -brokers localhost:9092 -topics topic.important.1,topic.access_log.1 -group example
```

Generate event:
```shell
curl --location --request GET 'http://localhost:8090/'
```


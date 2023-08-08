# Introducing Apache Kafka: consumers & producers

## Ad Consumer

```shell
go run cmd/event/consumer/main.go -brokers localhost:9092 -topics muken.ads.1.ad.publised,muken.ads.1.ads.removed -group example
```


## Examples

**How to run:**

Create & run kafka broker.
```shell
docker compose build
docker compose run
```

Create two topics `muken.ads.1.ad.publised` and `muken.ads.1.ads.removed`:

```shell
go run cmd/event/broker/broker.go -brokers localhost:9092 -topics muken.ads.1.ad.publised,muken.ads.1.ads.removed
```

Run http server producer:
```shell
go run cmd/event/producer/http_server.go -addr localhost:8090 -brokers localhost:9092 -verbose
```

Run consumer:
```shell
go run cmd/event/consumer/main.go -brokers localhost:9092 -topics muken.ads.1.ad.publised,muken.ads.1.ads.removed -group example
```

Generate event:
```shell
curl --location --request GET 'http://localhost:8090/'
```


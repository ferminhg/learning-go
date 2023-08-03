<p align="center">
    <img alt="&quot;a random gopher created by gopherize.me&quot;" src="../../img/gopher-challenge-6.png" width="200px" style="display: block; margin: 0 auto"/>
</p>

<h1 align="center" style="text-align: center;">
  Challenge #6. Introducing Apache Kafka: consumers & producers
</h1>

At this point we already have an application that exposes an HTTP API and persist information in a real database. The
next step is to publish/consume messages to a streaming platform, Apache Kafka in our case.

## Instructions

In this challenge we will be using `shopify/sarama`, which is the most used third party library to connect to a kafka
cluster/broker. Furthermore, for the sake of simplicity, we are going to produce/consume plain messages without any kind
of validation.

In addition, to validate and test your solution you will need a Kafka broker, you can use this minimal docker-compose file to 
dockerize a single-broker cluster: 

```yaml
version: '3.4'

services:
  zookeeper:
    image: confluentinc/cp-zookeeper:7.0.3
    ports:
      - "2181:2181"
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000

  kafka:
    image: confluentinc/cp-kafka:7.0.3
    ports:
      - "9092:9092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: 'zookeeper:2181'
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka:29092,PLAINTEXT_HOST://localhost:9092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: '1'
      KAFKA_LOG_RETENTION_MS: 20000
      KAFKA_LOG_RETENTION_CHECK_INTERVAL_MS: 2000
      KAFKA_MIN_INSYNC_REPLICAS: '1'
      KAFKA_AUTO_CREATE_TOPICS_ENABLE: "false"
```
If your app is dockerized and you want to make sure that Kakfa is available when it starts, you can use a healthcheck:
```yaml
  your-app:
    build: .
    depends_on:
      kafka:
        condition: service_healthy
  kafka:
    image: confluentinc/cp-kafka:7.0.3
    healthcheck:
      test: ["CMD-SHELL", "kafka-topics --bootstrap-server localhost:9092 --list"]
      interval: 5s
      timeout: 10s
      retries: 5
```



Finally, we are going to need some topic to be created. You can use the following snippet to take a look at how to create 
topics using `shopify/sarama`:

```go
func main() {
    brokerAddrs := []string{"localhost:9092"}
    config := sarama.NewConfig()
    config.Version = sarama.V2_1_0_0
    admin, err := sarama.NewClusterAdmin(brokerAddrs, config)
    if err != nil {
        log.Fatal("Error while creating cluster admin: ", err.Error())
    }
    defer func() { _ = admin.Close() }()
    err = admin.CreateTopic("topic.test.1", &sarama.TopicDetail{
        NumPartitions:     1,
        ReplicationFactor: 1,
    }, false)
    if err != nil {
        log.Fatal("Error while creating topic: ", err.Error())
    }
}
```

### Produce messages after posting an ad.

In this first part of the challenge we have to produce a message after posting an ad. You can choose the format of the 
message, but be sure that it contains at least the ad. You can find an example of a Kafka producer using Sarama right 
here: https://github.com/Shopify/sarama/tree/main/examples/http_server

In order to test this producer you have two options:
* To implement an integration test suite that relies on a Docker container: https://fedorov.dev/posts/2020-09-26-go-kafka-integration-testing-gnomock/
* To mock Kafka communication using the mocks that Sarama provides: https://github.com/Shopify/sarama/tree/main/mocks

### Consume the messages

Now we can start consuming the messages we produced earlier. To do this we are going to create a simple consumer that
just logs every message. You can find an example of a Kafka consumer (with a consumer group)
right here: https://github.com/Shopify/sarama/tree/main/examples/consumergroup

Whenever you have the consumer implemented (and tested), try to set up the environment and, by posting new ads, try to
answer the following questions:

* Can you log also the partition of the consumed message?

Yes. SendMessage method return the partition


* Try to post ads with different IDs so the messages ends in different partitions. Is the application able to consume from the all partitions? Note: to answer this question you will need to create the topic with at least 2 partitions.

Yes.

* What happens if you run several instances of the application? Are they sharing the different partitions?
They share broker + partitions.

## Resources
1. Example of a Kafka producer using `shopify/sarama`: https://github.com/Shopify/sarama/tree/main/examples/http_server
2. Example of integration test: https://fedorov.dev/posts/2020-09-26-go-kafka-integration-testing-gnomock/
3. `shopify/sarama` mocks to unit test consumers/producers: https://github.com/Shopify/sarama/tree/main/mocks
4. Example of a Kafka consumer using `shopify/sarama`: https://github.com/Shopify/sarama/tree/main/examples/consumergroup

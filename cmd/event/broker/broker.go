package main

import (
	"flag"
	"github.com/IBM/sarama"
	"log"
	"strings"
)

var (
	topics = ""
)

func main() {
	brokerAddrs := []string{"localhost:9092"}

	flag.StringVar(&topics, "topics", "", "Kafka topics to be created, as a comma separated list")
	flag.Parse()
	if len(topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}

	// Kafka config
	config := sarama.NewConfig()
	admin, err := sarama.NewClusterAdmin(brokerAddrs, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer func() { _ = admin.Close() }()

	// execute topic creator
	for _, t := range strings.Split(topics, ",") {
		createTopic(t, admin)
	}

}

func createTopic(t string, admin sarama.ClusterAdmin) {
	log.Printf("Creating topic: %s\n", t)
	if err := admin.CreateTopic(t, &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, false); err != nil {
		log.Println("Error while creating topic: ", err.Error())
	}
}

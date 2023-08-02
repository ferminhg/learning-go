package main

import (
	"flag"
	"github.com/IBM/sarama"
	"log"
	"strings"
	"sync"
)

var (
	brokers = ""
	topics  = ""
)

func main() {
	flag.StringVar(&brokers, "brokers", "", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.StringVar(&topics, "topics", "", "Kafka topics to be created, as a comma separated list")
	flag.Parse()

	if len(brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}

	if len(topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}

	// Kafka config
	config := sarama.NewConfig()
	admin, err := sarama.NewClusterAdmin(strings.Split(brokers, ","), config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer func() { _ = admin.Close() }()

	topicsToCreate := strings.Split(topics, ",")

	var wg sync.WaitGroup
	wg.Add(len(topicsToCreate))

	for _, t := range topicsToCreate {
		createTopic(t, admin, &wg)
	}
	wg.Wait()
}

func createTopic(t string, admin sarama.ClusterAdmin, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Creating topic: %s\n", t)
	if err := admin.CreateTopic(t, &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, false); err != nil {
		log.Println("Error while creating topic: ", err.Error())
	}
}

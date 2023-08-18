package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	//Mapeando a configuração do kafka cosumer
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka-container:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group",
		"auto.offset.reset": "earliest", // Caso queira que o cosumer leias todas as mensagens incluido as antigas declare esse comando
	}

	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("Erro consumer menssage:", err.Error())
	}

	// Declarando o topic
	topics := []string{"payments"}
	c.SubscribeTopics(topics, nil)

	//Permanece em um loop lendo as mensagens publicadas no topico
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf(string(msg.Value), msg.TopicPartition)
		}
	}
}

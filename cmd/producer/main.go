package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	producer := NewKafkaProducer()
	// Publicando mensagem
	Publish("Pagamento Aprovado", "payments", producer, nil)
	/*
	   Esvazie e aguarde as mensagens e solicitações pendentes para concluir a entrega. Inclui mensagens no ProduceChannel. Executa até que o valor chegue a zero ou em timeoutMs. Retorna o número de eventos pendentes ainda não liberados.
	*/
	producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	//Mapeando a configuração do kafka producer
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "kafka-container:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true",
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

//Publica mensagem
func Publish(msg string, topic string, producer *kafka.Producer, key []byte) error {
	messaging := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}

	err := producer.Produce(messaging, nil)
	if err != nil {
		return err
	}
	return nil
}

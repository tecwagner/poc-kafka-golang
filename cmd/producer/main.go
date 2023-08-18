package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	//Evento do kafka que publica a entrega do kafla
	deliveryChannel := make(chan kafka.Event)

	producer := NewKafkaProducer()
	// Publicando mensagem
	//Ao definir uma key para publicação de mensagem é uma forma de garantir que as mesmas seja ordenadas e cairam sempre na mesma partição
	Publish("Pagamento Em Analise", "payments", producer, []byte("transfer"), deliveryChannel)

	// gol routine
	go DeliveryReport(deliveryChannel)
	producer.Flush(1000)
	// //Evento está recebendo dados do canal
	// e := <-deliveryChannel
	// // Evento esta recebendo os dados de mensagem assincrono
	// msg := e.(*kafka.Message)
	// if msg.TopicPartition.Error != nil {
	// 	fmt.Println("Erro ao enviar a mensagem")
	// } else {
	// 	// Retorna os dados da mensagem e a particao que está publicada
	// 	fmt.Println("Mensagem enviada:", msg.TopicPartition)
	// }
	// /*
	//    Esvazie e aguarde as mensagens e solicitações pendentes para concluir a entrega. Inclui mensagens no ProduceChannel. Executa até que o valor chegue a zero ou em timeoutMs. Retorna o número de eventos pendentes ainda não liberados.
	// */

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
func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChannel chan kafka.Event) error {
	messaging := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}

	err := producer.Produce(messaging, deliveryChannel) // Sempre publicar mensagem que envie para o canal deliveryChannel
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChannel chan kafka.Event) {
	// evento de mensagem assincona
	for e := range deliveryChannel {
		switch event := e.(type) {
		case *kafka.Message:
			// Evento esta recebendo os dados de mensagem assincrono
			if event.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar a mensagem")
			} else {
				// Retorna os dados da mensagem e a particao que está publicada
				fmt.Println("Mensagem enviada:", event.TopicPartition)
				// anotar no banco de dados que a mensagem foi processada
				// ex: confirma que uma transferencia bancaria ocorreu
			}
		}
	}
}

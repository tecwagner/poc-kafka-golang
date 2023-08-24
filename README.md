## poc-kafka-golang

# Para fins de uso e configuração do host.dokcer.internal ao seu docker

    - No windows o caminho: D:\Windows\System32\drivers\etc
    - No Linux o caminho: cat /etc/hosts

# Desenvolvendo um aplicação em Go que será consumer e producer

    - Criando o Dockerfile utilizando da biblioteca: 'librdkafka-dev -y', para fazer a comunicação com kafka

    - Verificando os logs em Dockerfile comunicando com kafka: docker logs kafka-container

# Criando projeto go e acessando o docker container

    - docker exec -it gokafka bash
        - comando para criação do projeto: go mod init github.com/tecwagner/go-kafka-golang

    - Entrar no container do kafka:
        - docker exec -it kafka-container bash

    - Criando um topic para aplicação do kafka:
        - criar um topic por meio de linha de comando: kafka-topics --create --bootstrap-server=localhost:9092 --topic=payments --partitions=3

# Publicando a primera mensagem.

    - No diretorio cmd/producer
        - Foi implementado o metodo para o envio de mensagem.

        - O Metodo NewKafkaProducer(): Recebe o mapeamento de configuracão do kafka.

        - O Metodo Publish(): Publica as mensagens que estão sendo produzidas

# Delivery Repor

    - Devivery channels, é um canal de comunicação do kafka entre a publicação da mensagen e uma função ou uma rotina que foi executada.
    - Envia o retorno para um canal de comunicação.
    - Teremos um metodo que ficará lendo as mensagens que passaram por essse canal.
    - Quando chegar uma mensagens nesse canal, teremos uma mensagen de retorno do que aconteceu com envio.

# Consumindo mensagens

    - No diretorio cmd/consumer
        - Foi implementado o metodo para que seja feito o consumo das mensagens publicadas no kafka

    - Foi criado três partições de grupo para ler as mensagens.

    - O Grupo apontado para o consumer deve ser o que foi informado no mapeamento do configMap

    - O comando é para listar as partições e os grupos de consumir de mensagens do group goapp-group: kafka-consumer-groups --bootstrap-server=localhost:9092 --describe --group=goapp-group

    - É permitido que seja criado novos consumer para o mesmo groups, para que seja processado as mensagens mais rapidamente

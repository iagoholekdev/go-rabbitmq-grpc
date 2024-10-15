
# Go RabbitMQ gRPC

Este projeto é uma implementação de um servidor gRPC em Go que publica mensagens em uma fila RabbitMQ. O cliente gRPC pode enviar mensagens que são publicadas no RabbitMQ, permitindo comunicação assíncrona entre diferentes serviços.

## Tecnologias Usadas

- Go (versão 1.23.2)
- gRPC
- RabbitMQ
- Protocol Buffers

## Estrutura do Projeto

```
go-rabbitmq-grpc/
│
├── grpc-client/            # Cliente gRPC
│   └── client.go             # Código do cliente
│
├── grpc-server/            # Servidor gRPC
│   ├── server.go             # Código do servidor
│   └── publisher/          # Código do publicador RabbitMQ
│       └── publisher.go
│
├── rabbitmq/               # Configuração do RabbitMQ
│   ├── config/             # Arquivos de configuração
│   │   └── amqp_builder.go  # Conexão e criação de canais
│   └── publisher/          # Publicador de mensagens
│       └── publisher.go
│
├── .gitignore              # Arquivo para ignorar arquivos desnecessários
├── go.mod                  # Gerenciador de dependências
└── README.md               # Este arquivo
```

## Como Executar o Projeto

### 1. Configurar o RabbitMQ

Certifique-se de que o RabbitMQ está instalado e em execução na sua máquina. Você pode seguir as [instruções de instalação do RabbitMQ](https://www.rabbitmq.com/download.html) para configurar o RabbitMQ.

### 2. Clonar o Repositório

Clone este repositório para sua máquina local:

```bash
git clone https://github.com/iagoholekdev/go-rabbitmq-grpc.git
cd go-rabbitmq-grpc
```

### 3. Instalar Dependências

Instale as dependências do projeto utilizando o Go modules:

```bash
go mod tidy
```

### 4. Executar o Servidor gRPC

Abra um terminal e navegue até a pasta do servidor, depois execute:

```bash
cd grpc-server
go run server.go
```

### 5. Executar o Cliente gRPC

Em outro terminal, navegue até a pasta do cliente e execute:

```bash
cd grpc-client
go run client.go
```

## Como Funciona

1. O servidor gRPC escuta por solicitações na porta `50051`.
2. Quando o cliente envia uma mensagem, o servidor recebe essa mensagem e a publica na fila do RabbitMQ.
3. Você pode usar um consumidor RabbitMQ para ler as mensagens da fila.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

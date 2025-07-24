# Buscador de CEP com Multithreading

Este projeto implementa um buscador de CEP que faz requisições simultâneas para duas APIs diferentes e retorna o resultado da API mais rápida, seguindo as melhores práticas de arquitetura de software.

## 🏗️ Arquitetura

O projeto foi estruturado seguindo os padrões aprendidos no curso:

```
cep-finder-multithreading/
├── cmd/
│   └── cep-finder/
│       └── main.go              # Ponto de entrada da aplicação
├── configs/
│   └── config.go                # Configurações da aplicação
├── internal/
│   ├── dto/
│   │   └── cep_dto.go           # Data Transfer Objects
│   ├── entity/
│   │   ├── cep.go               # Entidade de domínio
│   │   └── cep_test.go          # Testes da entidade
│   └── service/
│       ├── cep_service.go       # Lógica de negócio
│       └── cep_service_test.go  # Testes do serviço
├── pkg/                         # Pacotes reutilizáveis
├── test/
│   └── cep.http                 # Testes HTTP para IDE
├── go.mod                       # Dependências
└── README.md                    # Documentação
```

## 🚀 Funcionalidades

- ✅ Requisições simultâneas para BrasilAPI e ViaCEP
- ✅ Retorna o resultado da API mais rápida
- ✅ Timeout configurável (padrão: 1 segundo)
- ✅ Exibe qual API retornou o resultado
- ✅ Mostra o tempo de resposta da API vencedora
- ✅ Arquitetura limpa com separação de responsabilidades
- ✅ Testes unitários completos
- ✅ Configuração via Viper

## 🛠️ Tecnologias Utilizadas

- **Go 1.19**: Linguagem principal
- **Viper**: Gerenciamento de configurações
- **Testify**: Framework de testes
- **Context**: Controle de timeout e cancelamento
- **Goroutines & Channels**: Concorrência

## 📦 APIs Utilizadas

1. **BrasilAPI**: `https://brasilapi.com.br/api/cep/v1/{cep}`
2. **ViaCEP**: `http://viacep.com.br/ws/{cep}/json/`

## 🏃‍♂️ Como Executar

### Pré-requisitos
- Go 1.23 ou superior

### Execução
```bash
# Clone o repositório
git clone <repository-url>
cd cep-finder-multithreading

# Execute o programa
go run cmd/cep-finder/main.go 01153000
```

### Exemplo de Uso
```bash
$ go run cmd/cep-finder/main.go 01153000

Buscando informações para o CEP: 01153000
Fazendo requisições simultâneas para BrasilAPI e ViaCEP...

✅ Resultado obtido da API: BrasilAPI
⏱️  Tempo de resposta: 118.628671ms

📋 Dados do endereço:
   CEP: 01153000
   Logradouro: Rua Vitorino Carmilo
   Bairro: Barra Funda
   Cidade: São Paulo
   Estado: SP
```

## 🧪 Testes

### Executar Todos os Testes
```bash
go test ./...
```

### Executar Testes com Coverage
```bash
go test -cover ./...
```

### Executar Testes Específicos
```bash
# Testes da entidade
go test ./internal/entity

# Testes do serviço
go test ./internal/service
```

### Testes HTTP
O arquivo `test/cep.http` contém testes para as APIs que podem ser executados diretamente na IDE (VS Code, IntelliJ, etc.).

## ⚙️ Configuração

O projeto usa Viper para gerenciamento de configurações. Por padrão:

- **Timeout**: 1 segundo

Para alterar configurações, você pode:
1. Usar variáveis de ambiente
2. Criar arquivo de configuração
3. Modificar o código

## 🏛️ Padrões de Arquitetura Aplicados

### Clean Architecture
- **Entities**: Regras de negócio puras (`internal/entity`)
- **Use Cases**: Casos de uso da aplicação (`internal/service`)
- **Interface Adapters**: DTOs e configurações (`internal/dto`, `configs`)
- **Frameworks & Drivers**: Ponto de entrada (`cmd/cep-finder`)

### SOLID Principles
- **Single Responsibility**: Cada pacote tem uma responsabilidade específica
- **Open/Closed**: Fácil extensão sem modificação
- **Dependency Inversion**: Dependências através de interfaces

### Design Patterns
- **Factory Pattern**: `NewCEPService()`
- **Strategy Pattern**: Diferentes APIs de CEP
- **Repository Pattern**: Abstração para acesso a dados

## 🔍 Conceitos de Multithreading Aplicados

### Goroutines
- Execução simultânea das requisições HTTP
- Cada API roda em uma goroutine separada

### Channels
- `resultChan`: Comunicação entre goroutines
- `errorChan`: Tratamento de erros assíncrono

### Select
- Aguarda o primeiro resultado ou timeout
- Implementa o padrão "race condition" de forma controlada

### Context
- Controle de timeout e cancelamento
- Propagação de cancelamento para goroutines filhas

## 📊 Métricas e Performance

O programa mede e exibe:
- Tempo de resposta de cada API
- Qual API foi mais rápida
- Total de tempo de execução

## 🐛 Tratamento de Erros

- **Timeout**: Se nenhuma API responder no tempo limite
- **CEP Inválido**: Tratamento de CEPs inexistentes
- **Erro de Rede**: Falhas de conectividade
- **Erro de API**: Status codes de erro das APIs

## 🔧 Possíveis Melhorias Futuras

- [ ] Adicionar mais APIs de CEP
- [ ] Implementar cache de resultados
- [ ] Adicionar métricas de performance
- [ ] Criar API REST para o serviço
- [ ] Implementar rate limiting
- [ ] Adicionar logs estruturados
- [ ] Dockerização da aplicação

## 📝 Licença

Este projeto foi criado como parte do curso de Pós Go Expert da FullCycle.

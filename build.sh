#!/bin/bash

echo "🔨 Buildando o Buscador de CEP..."
echo "=================================="

# Verifica se o Go está instalado
if ! command -v go &> /dev/null; then
    echo "❌ Go não está instalado. Por favor, instale o Go 1.23.0 ou superior."
    exit 1
fi

# Executa os testes
echo "🧪 Executando testes..."
go test -v ./...

if [ $? -ne 0 ]; then
    echo "❌ Testes falharam. Abortando build."
    exit 1
fi

# Build da aplicação
echo "🏗️  Compilando aplicação..."
go build -o bin/cep-finder cmd/cep-finder/main.go

if [ $? -eq 0 ]; then
    echo "✅ Build concluído com sucesso!"
    echo "📁 Executável criado em: bin/cep-finder"
    echo ""
    echo "🚀 Para usar:"
    echo "   ./bin/cep-finder 01153000"
else
    echo "❌ Erro no build."
    exit 1
fi

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/m4rcelotoledo/cep-finder-multithreading/configs"
	"github.com/m4rcelotoledo/cep-finder-multithreading/internal/service"
)

func main() {
	// Carrega configurações
	config, err := configs.LoadConfig()
	if err != nil {
		fmt.Printf("Erro ao carregar configurações: %v\n", err)
		os.Exit(1)
	}

	// Valida argumentos
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run cmd/cep-finder/main.go <CEP>")
		fmt.Println("Exemplo: go run cmd/cep-finder/main.go 45807000")
		os.Exit(1)
	}

	cep := os.Args[1]
	fmt.Printf("Buscando informações para o CEP: %s\n", cep)
	fmt.Println("Fazendo requisições simultâneas para BrasilAPI e ViaCEP...")
	fmt.Println()

	// Cria o serviço
	cepService := service.NewCEPService(config.Timeout)

	// Busca o CEP
	ctx := context.Background()
	result, err := cepService.SearchCEP(ctx, cep)
	if err != nil {
		fmt.Printf("❌ Erro: %v\n", err)
		os.Exit(1)
	}

	// Exibe o resultado
	fmt.Printf("✅ Resultado obtido da API: %s\n", result.API)
	fmt.Printf("⏱️  Tempo de resposta: %v\n", result.ResponseTime)
	fmt.Println()
	fmt.Println("📋 Dados do endereço:")
	fmt.Printf("   CEP: %s\n", result.CEP)
	fmt.Printf("   Logradouro: %s\n", result.Logradouro)
	fmt.Printf("   Bairro: %s\n", result.Bairro)
	fmt.Printf("   Cidade: %s\n", result.Cidade)
	fmt.Printf("   Estado: %s\n", result.Estado)
}

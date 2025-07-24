package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCEPService(t *testing.T) {
	timeout := 2 * time.Second
	service := NewCEPService(timeout)

	assert.NotNil(t, service)
}

func TestCEPService_SearchCEP_Timeout(t *testing.T) {
	// Teste com timeout muito baixo para forçar timeout
	service := NewCEPService(1 * time.Millisecond)

	ctx := context.Background()
	result, err := service.SearchCEP(ctx, "01153000")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timeout")
}

func TestCEPService_SearchCEP_ValidCEP(t *testing.T) {
	// Teste com timeout adequado
	service := NewCEPService(5 * time.Second)

	ctx := context.Background()
	result, err := service.SearchCEP(ctx, "01153000")

	// Pode falhar se as APIs estiverem fora do ar, mas não deve dar timeout
	if err != nil {
		t.Logf("Erro esperado se APIs estiverem fora do ar: %v", err)
		return
	}

	assert.NotNil(t, result)
	assert.NotEmpty(t, result.CEP)
	assert.NotEmpty(t, result.Logradouro)
	assert.NotEmpty(t, result.Cidade)
	assert.NotEmpty(t, result.Estado)
	assert.NotEmpty(t, result.API)
	assert.True(t, result.IsValid())
	assert.True(t, result.ResponseTime > 0)
}

func TestCEPService_SearchCEP_InvalidCEP(t *testing.T) {
	service := NewCEPService(5 * time.Second)

	ctx := context.Background()
	result, err := service.SearchCEP(ctx, "00000000")

	// CEP inválido deve resultar em timeout ou erro
	if err != nil {
		t.Logf("Erro esperado para CEP inválido: %v", err)
		return
	}

	// Se não deu erro, o resultado deve ser válido
	if result != nil {
		assert.True(t, result.IsValid())
	}
}

func TestCEPService_SearchCEP_ContextCancelled(t *testing.T) {
	service := NewCEPService(5 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancela imediatamente

	result, err := service.SearchCEP(ctx, "01153000")

	assert.Nil(t, result)
	assert.Error(t, err)
}

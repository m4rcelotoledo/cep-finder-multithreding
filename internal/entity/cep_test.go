package entity

import (
	"testing"
	"time"
)

func TestNewCEP(t *testing.T) {
	cep := "01153000"
	logradouro := "Rua Vitorino Carmilo"
	bairro := "Barra Funda"
	cidade := "São Paulo"
	estado := "SP"
	api := "BrasilAPI"
	responseTime := 100 * time.Millisecond

	result := NewCEP(cep, logradouro, bairro, cidade, estado, api, responseTime)

	if result.CEP != cep {
		t.Errorf("CEP esperado %s, obtido %s", cep, result.CEP)
	}
	if result.Logradouro != logradouro {
		t.Errorf("Logradouro esperado %s, obtido %s", logradouro, result.Logradouro)
	}
	if result.Bairro != bairro {
		t.Errorf("Bairro esperado %s, obtido %s", bairro, result.Bairro)
	}
	if result.Cidade != cidade {
		t.Errorf("Cidade esperada %s, obtida %s", cidade, result.Cidade)
	}
	if result.Estado != estado {
		t.Errorf("Estado esperado %s, obtido %s", estado, result.Estado)
	}
	if result.API != api {
		t.Errorf("API esperada %s, obtida %s", api, result.API)
	}
	if result.ResponseTime != responseTime {
		t.Errorf("ResponseTime esperado %v, obtido %v", responseTime, result.ResponseTime)
	}
}

func TestCEP_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		cep      *CEP
		expected bool
	}{
		{
			name:     "CEP válido",
			cep:      NewCEP("01153000", "Rua Vitorino Carmilo", "Barra Funda", "São Paulo", "SP", "BrasilAPI", 100*time.Millisecond),
			expected: true,
		},
		{
			name:     "CEP sem logradouro",
			cep:      NewCEP("01153000", "", "Barra Funda", "São Paulo", "SP", "BrasilAPI", 100*time.Millisecond),
			expected: false,
		},
		{
			name:     "CEP sem cidade",
			cep:      NewCEP("01153000", "Rua Vitorino Carmilo", "Barra Funda", "", "SP", "BrasilAPI", 100*time.Millisecond),
			expected: false,
		},
		{
			name:     "CEP sem estado",
			cep:      NewCEP("01153000", "Rua Vitorino Carmilo", "Barra Funda", "São Paulo", "", "BrasilAPI", 100*time.Millisecond),
			expected: false,
		},
		{
			name:     "CEP vazio",
			cep:      NewCEP("", "Rua Vitorino Carmilo", "Barra Funda", "São Paulo", "SP", "BrasilAPI", 100*time.Millisecond),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.cep.IsValid()
			if result != tt.expected {
				t.Errorf("IsValid() = %v, esperado %v", result, tt.expected)
			}
		})
	}
}

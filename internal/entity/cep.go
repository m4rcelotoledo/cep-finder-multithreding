package entity

import "time"

// CEP representa uma entidade de endereço baseada em CEP
type CEP struct {
	CEP          string
	Logradouro   string
	Bairro       string
	Cidade       string
	Estado       string
	API          string
	ResponseTime time.Duration
}

// NewCEP cria uma nova instância de CEP
func NewCEP(cep, logradouro, bairro, cidade, estado, api string, responseTime time.Duration) *CEP {
	return &CEP{
		CEP:          cep,
		Logradouro:   logradouro,
		Bairro:       bairro,
		Cidade:       cidade,
		Estado:       estado,
		API:          api,
		ResponseTime: responseTime,
	}
}

// IsValid verifica se o CEP é válido
func (c *CEP) IsValid() bool {
	return c.CEP != "" && c.Logradouro != "" && c.Cidade != "" && c.Estado != ""
}

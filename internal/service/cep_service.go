package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/m4rcelotoledo/cep-finder-multithreading/internal/dto"
	"github.com/m4rcelotoledo/cep-finder-multithreading/internal/entity"
)

// CEPService interface define os métodos para busca de CEP
type CEPService interface {
	SearchCEP(ctx context.Context, cep string) (*entity.CEP, error)
}

// cepService implementa CEPService
type cepService struct {
	brasilAPIURL string
	viaCEPURL    string
	timeout      time.Duration
}

// NewCEPService cria uma nova instância do serviço de CEP
func NewCEPService(timeout time.Duration) CEPService {
	return &cepService{
		brasilAPIURL: "https://brasilapi.com.br/api/cep/v1/%s",
		viaCEPURL:    "http://viacep.com.br/ws/%s/json/",
		timeout:      timeout,
	}
}

// SearchCEP busca informações de CEP nas duas APIs simultaneamente
func (s *cepService) SearchCEP(ctx context.Context, cep string) (*entity.CEP, error) {
	// Canal para receber o primeiro resultado
	resultChan := make(chan *entity.CEP, 2)
	errorChan := make(chan error, 2)

	// Contexto com timeout
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	// Goroutine para BrasilAPI
	go func() {
		result, err := s.fetchBrasilAPI(ctx, cep)
		if err != nil {
			errorChan <- err
			return
		}
		if result != nil {
			resultChan <- result
		}
	}()

	// Goroutine para ViaCEP
	go func() {
		result, err := s.fetchViaCEP(ctx, cep)
		if err != nil {
			errorChan <- err
			return
		}
		if result != nil {
			resultChan <- result
		}
	}()

	// Aguarda o primeiro resultado ou timeout
	select {
	case result := <-resultChan:
		return result, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout: nenhuma API respondeu em %v", s.timeout)
	}
}

// fetchBrasilAPI busca informações na BrasilAPI
func (s *cepService) fetchBrasilAPI(ctx context.Context, cep string) (*entity.CEP, error) {
	start := time.Now()

	url := fmt.Sprintf(s.brasilAPIURL, cep)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("brasilapi retornou status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var brasilAPI dto.BrasilAPIResponse
	if err := json.Unmarshal(body, &brasilAPI); err != nil {
		return nil, err
	}

	responseTime := time.Since(start)

	return entity.NewCEP(
		brasilAPI.CEP,
		brasilAPI.Street,
		brasilAPI.Neighborhood,
		brasilAPI.City,
		brasilAPI.State,
		"BrasilAPI",
		responseTime,
	), nil
}

// fetchViaCEP busca informações na ViaCEP
func (s *cepService) fetchViaCEP(ctx context.Context, cep string) (*entity.CEP, error) {
	start := time.Now()

	url := fmt.Sprintf(s.viaCEPURL, cep)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("viacep retornou status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var viaCEP dto.ViaCEPResponse
	if err := json.Unmarshal(body, &viaCEP); err != nil {
		return nil, err
	}

	// Verifica se a resposta não contém erro
	if viaCEP.CEP == "" {
		return nil, fmt.Errorf("viacep retornou CEP vazio")
	}

	responseTime := time.Since(start)

	return entity.NewCEP(
		viaCEP.CEP,
		viaCEP.Logradouro,
		viaCEP.Bairro,
		viaCEP.Localidade,
		viaCEP.UF,
		"ViaCEP",
		responseTime,
	), nil
}

package dto

// BrasilAPIResponse representa a resposta da BrasilAPI
type BrasilAPIResponse struct {
	CEP          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
	Location     struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"location"`
}

// ViaCEPResponse representa a resposta da ViaCEP
type ViaCEPResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

// CEPInput representa a entrada para busca de CEP
type CEPInput struct {
	CEP string
}

// CEPOutput representa a saída formatada do CEP
type CEPOutput struct {
	CEP          string `json:"cep"`
	Logradouro   string `json:"logradouro"`
	Bairro       string `json:"bairro"`
	Cidade       string `json:"cidade"`
	Estado       string `json:"estado"`
	API          string `json:"api"`
	ResponseTime string `json:"response_time"`
}

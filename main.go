type (
	Address struct {
		CEP string `json:"cep"`
		Logradouro string `json:"logradouro"`
		Bairro     string `json:"bairro"`
		Cidade     string `json:"cidade"`
		Estado     string `json:"estado"`
		APIOrigem  string `json:"api_origem"`
	}
	BrasilAPIResponse struct {
		CEP         string `json:"cep"`
		State       string `json:"state"`
		City        string `json:"city"`
		Neighborhood string `json:"neighborhood"`
		Street      string `json:"street"`
	}
	ViaCEPResponse struct {
		CEP        string `json:"cep"`
		Logradouro string `json:"logradouro"`
		Bairro     string `json:"bairro"`
		Localidade string `json:"localidade"`
		UF         string `json:"uf"`
	}
)

func fetchAddress(url string, apiName string, ch chan<- Address) {
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var addr Address

	if apiName == "BrasilAPI" {
		var brasilResp BrasilAPIResponse
		if err := json.NewDecoder(resp.Body).Decode(&brasilResp); err != nil {
			return
		}
		addr = Address{
			CEP:        brasilResp.CEP,
			Logradouro: brasilResp.Street,
			Bairro:     brasilResp.Neighborhood,
			Cidade:     brasilResp.City,
			Estado:     brasilResp.State,
			APIOrigem:  "BrasilAPI",
		}
	} else if apiName == "ViaCEP" {
		var viaCEPResp ViaCEPResponse
		if err := json.NewDecoder(resp.Body).Decode(&viaCEPResp); err != nil {
			return
		}
		addr = Address{
			CEP:        viaCEPResp.CEP,
			Logradouro: viaCEPResp.Logradouro,
			Bairro:     viaCEPResp.Bairro,
			Cidade:     viaCEPResp.Localidade,
			Estado:     viaCEPResp.UF,
			APIOrigem:  "ViaCEP",
		}
	}

	ch <- addr
}

func handleCEPRequest(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP é obrigatório", http.StatusBadRequest)
		return
	}

	ch := make(chan Address)

	api1 := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	api2 := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	go fetchAddress(api1, "BrasilAPI", ch)
	go fetchAddress(api2, "ViaCEP", ch)

	select {
	case result := <-ch:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	case <-time.After(1 * time.Second):
		http.Error(w, "Timeout! Nenhuma API respondeu a tempo.", http.StatusGatewayTimeout)
	}
}

func main() {
	http.HandleFunc("/buscar", handleCEPRequest)

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
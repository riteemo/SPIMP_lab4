package converter

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

const apiBaseURL = "https://v6.exchangerate-api.com/v6"

type APIResponse struct {
    Result          string             
    BaseCode        string             
    ConversionRates map[string]float64 
    ErrorType       string             
}

func FetchRates(apiKey, baseCurrency string) (map[string]float64, error) {

    url := fmt.Sprintf("%s/%s/latest/%s", apiBaseURL, apiKey, baseCurrency)

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Get(url)
    if err != nil {
        return nil, fmt.Errorf("Ошибка соединения: %w", err)
    }
    defer resp.Body.Close()

    var apiResp APIResponse
    if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
        return nil, fmt.Errorf("Ошибка парсинга: %w", err)
    }

    if apiResp.Result != "success" {
        return nil, fmt.Errorf("Ошибка API: %s", apiResp.ErrorType)
    }

    return apiResp.ConversionRates, nil
}
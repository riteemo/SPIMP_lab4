package converter

import (
    "fmt"
    "math"
    "strings"
)

func Convert(rates map[string]float64, from, to string, amount float64) (float64, error) {
    from = strings.ToUpper(strings.TrimSpace(from))
    to = strings.ToUpper(strings.TrimSpace(to))

    if amount <= 0 {
        return 0, fmt.Errorf("Введите сумму >0")
    }

    fromRate, ok := rates[from]
    if !ok {
        return 0, fmt.Errorf("Неизвестная валюта: %s", from)
    }

    toRate, ok := rates[to]
    if !ok {
        return 0, fmt.Errorf("Неизвестная валюта: %s", to)
    }

    result := (amount / fromRate) * toRate
    return math.Round(result*10000) / 10000, nil
}

func FormatResult(amount float64, from string, result float64, to string) string {
    return fmt.Sprintf("%.2f %s = %.4f %s", amount, from, result, to)
}
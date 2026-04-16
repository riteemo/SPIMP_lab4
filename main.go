package main

import (
    "bufio"
    "converter/config"
    "converter/converter_dir"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        fmt.Printf("Ошибка конфига: %v\n", err)
        os.Exit(1)
    }

    args := os.Args[1:]

    switch {
    case len(args) == 0:
        runInteractive(cfg)
    case args[0] == "list":
        runList(cfg)
    case len(args) == 3:
        runDirect(cfg, args[0], args[1], args[2])
    default:
        printUsage()
    }
}

func runDirect(cfg *config.Config, amountStr, from, to string) {
    amount, err := strconv.ParseFloat(amountStr, 64)
    if err != nil {
        fmt.Printf("Неверная сумма: %s\n", amountStr)
        os.Exit(1)
    }

    rates, err := converter.FetchRates(cfg.APIKey, "USD")
    if err != nil {
        fmt.Printf("Ошибка: %v\n", err)
        os.Exit(1)
    }

    result, err := converter.Convert(rates, from, to, amount)
    if err != nil {
        fmt.Printf("%v\n", err)
        os.Exit(1)
    }
  
    fmt.Println(converter.FormatResult(amount, strings.ToUpper(from), result, strings.ToUpper(to)))
}

func runInteractive(cfg *config.Config) {
    fmt.Println("Введите валюту для конвертации в формате <сумма из в>")
    rates, err := converter.FetchRates(cfg.APIKey, "USD")
    if err != nil {
        fmt.Printf("%v\n", err)
        os.Exit(1)
    }
    
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("\n> ")
        if !scanner.Scan() {
            break
        }

        input := strings.TrimSpace(scanner.Text())
        if input == "" {
            break
        }

        parts := strings.Fields(input)
        if len(parts) != 3 {
            fmt.Println("Формат: сумма из в (пример: 100 USD EUR)")
            continue
        }

        amount, err := strconv.ParseFloat(parts[0], 64)
        if err != nil {
            fmt.Println("Неверная сумма")
            continue
        }

        result, err := converter.Convert(rates, parts[1], parts[2], amount)
        if err != nil {
            fmt.Printf("%v\n", err)
            continue
        }
      
        fmt.Println(converter.FormatResult(amount, strings.ToUpper(parts[1]), result, strings.ToUpper(parts[2])))
    }
}

func runList(cfg *config.Config) {
    rates, err := converter.FetchRates(cfg.APIKey, "USD")
    if err != nil {
        fmt.Printf(" %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Доступные валюты:")
    for currency := range rates {
        fmt.Printf("  %s\n", currency)
    }
}

func printUsage() {
    fmt.Println(`Использование:
  currency-converter                 # интерактивный режим
  currency-converter 100 USD EUR     # прямая конвертация
  currency-converter list            # список валют`)
}
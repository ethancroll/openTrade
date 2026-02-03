package finance

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// GlobalQuote represents the Alpha Vantage Global Quote response
type GlobalQuote struct {
	GlobalQuote QuoteData `json:"Global Quote"`
}

type QuoteData struct {
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	Change           string `json:"09. change"`
	ChangePercent    string `json:"10. change percent"`
}

// GetStockQuote fetches the current stock quote from Alpha Vantage API
func GetStockQuote(symbol string, apiKey string) (*QuoteData, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", symbol, apiKey)
	
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	
	var quote GlobalQuote
	if err := json.Unmarshal(body, &quote); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	
	// Check if the response is empty (rate limit or error)
	if quote.GlobalQuote.Symbol == "" {
		return nil, fmt.Errorf("empty response - possible rate limit or invalid symbol")
	}
	
	return &quote.GlobalQuote, nil
}

// GetAPIKey reads the Alpha Vantage API key from environment
func GetAPIKey() string {
	return os.Getenv("ALPHA_VANTAGE_API_KEY")
}

// FetchBasics fetches stock data for a given ticker and prints it
func FetchBasics(tickr string) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	// Get API key
	apiKey := GetAPIKey()
	if apiKey == "" || apiKey == "your_api_key_here" {
		log.Fatal("Please set ALPHA_VANTAGE_API_KEY in .env file")
	}

	// Fetch stock data
	fmt.Printf("Fetching %s stock data from Alpha Vantage...\n", tickr)
	quote, err := GetStockQuote(tickr, apiKey)
	if err != nil {
		log.Println("Error fetching stock data:", err)
		return
	}

	// Print the stock data
	fmt.Printf("\n%s\n", tickr)
	fmt.Printf("Symbol: %s\n", quote.Symbol)
	fmt.Printf("Price: $%s\n", quote.Price)
	fmt.Printf("Open: $%s\n", quote.Open)
	fmt.Printf("High: $%s\n", quote.High)
	fmt.Printf("Low: $%s\n", quote.Low)
	fmt.Printf("Volume: %s\n", quote.Volume)
	fmt.Printf("Latest Trading Day: %s\n", quote.LatestTradingDay)
	fmt.Printf("Previous Close: $%s\n", quote.PreviousClose)
	fmt.Printf("Change: %s (%s)\n", quote.Change, quote.ChangePercent)
	fmt.Println("========================\n")
	
	// Rate limiting: wait 1.5 seconds between requests
	time.Sleep(1500 * time.Millisecond)
}

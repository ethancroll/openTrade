package finance

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

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

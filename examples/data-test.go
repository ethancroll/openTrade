package main

import (
	"github.com/ethancroll/openTrade/pkg/finance"
)

func main() {
	finance.FetchBasics("NVDA")
	finance.FetchBasics("GOOG")
}

package http

import (
	"github.com/samuelbeyeler/goarch-example/domain/prices"
	"net/http"
)

// PriceService ...
type PriceService interface {
	GetPricesForAll() []prices.Price
}

// PriceHandler ...
type PriceHandler struct {
	prices PriceService
}

// NewPriceHandler Since the PriceService handles it's own dependencies, all we need is the service itself
func NewPriceHandler(prices PriceService) PriceHandler {
	return PriceHandler{prices: prices}
}

// GetAllPrices ...
func (h PriceHandler) GetAllPrices(w http.ResponseWriter, r *http.Request) {
	h.prices.GetPricesForAll()
}

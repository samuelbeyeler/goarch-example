// Package main will compile all of the various services/handlers and throw a compile-time error if there are cyclical dependencies.
package main

import (
	"fmt"
	"net/http"

	"github.com/samuelbeyeler/goarch-example/db/sql"
	"github.com/samuelbeyeler/goarch-example/domain/prices"
	"github.com/samuelbeyeler/goarch-example/domain/product"
	handlers "github.com/samuelbeyeler/goarch-example/transport/http"
)

// NOTE: Dependencies are only ever moving upwards in complexity of service.
func main() {
	// Repos
	priceRepo := sql.NewPriceRepo()
	prodRepo := sql.NewProductRepo()

	// Services
	prodService := product.NewService(prodRepo)
	priceService := prices.NewService(priceRepo, prodService)

	// Handlers
	productHandler := handlers.NewProductHandler(prodService)
	priceHandler := handlers.NewPriceHandler(priceService)

	// Server
	http.HandleFunc("/products", productHandler.GetAllProducts)
	http.HandleFunc("/prices", priceHandler.GetAllPrices)
	fmt.Print("No circular deps!\n")
	http.ListenAndServe(":8090", nil)
}

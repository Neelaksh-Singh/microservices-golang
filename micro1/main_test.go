package main

import (
	"fmt"
	"micro1/sdk/client"
	"micro1/sdk/client/products"
	"testing"
)

func TestClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()
	// params.
	prods, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v", prods.GetPayload()[0])
	t.Fail()
}

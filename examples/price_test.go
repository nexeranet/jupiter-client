package examples

import (
	"context"
	"testing"

	jupiter "github.com/nexeranet/jupiter-client/client"
)

func TestGetPrice(t *testing.T) {
	ctx := context.Background()
	client, err := jupiter.NewClientWithResponses("https://lite-api.jup.ag")
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.PriceGetWithResponse(ctx, &jupiter.PriceGetParams{
		Ids: &jupiter.PriceMintAddresses{"JUPyiwrYJFskUPiHa7hkeR8VUtAeFoSYbKedZNsDvCN"},
	})
	if response.JSON200 == nil {
		t.Fatal("Error not 200")
	}
	for key, value := range *response.JSON200 {
		t.Log(key, value.BlockId)
	}

}

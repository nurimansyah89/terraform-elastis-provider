package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/nurimansyah89/terraform-elastis-provider/elastis"
)

func main() {
	providerserver.Serve(context.Background(), elastis.New, providerserver.ServeOpts{
		Address: "elastis.id/provider/elastis",
	})
}

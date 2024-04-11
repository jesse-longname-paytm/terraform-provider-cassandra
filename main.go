package main

import (
	"github.com/jesse-longname-paytm/terraform-provider-cassandra/cassandra"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: cassandra.Provider})
}

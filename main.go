package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sriddell/datasource-template-gz/templategz"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: templategz.Provider,
	})
}

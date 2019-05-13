package main

import "fmt"

func main() {

	// plugin.Serve(&plugin.ServeOpts{
	// 	ProviderFunc: func() terraform.ResourceProvider {
	// 		return Provider()
	// 	},
	// })

	c := NewClient("fff")
	fmt.Print(c)
}

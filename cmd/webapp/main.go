package main

import (
	"fmt"
	"gosample/cmd/webapp/routes"

	"github.com/MakMoinee/go-mith/pkg/goserve"
)

func main() {
	fmt.Println("Starting the Service ...")

	httpService := goserve.NewService(":8443")
	routes.Set(httpService)
	fmt.Println("The Service Started")
	if err := httpService.Start(); err != nil {
		panic(err)
	}

}

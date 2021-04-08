package main

import ( 
	"fmt"
	router "github.com/codeedu/imersaofsfc2-simulator/application/router"
)

//primeira function a ser executada do programa
func main() {
	route := router.Route{
		ID: "1",
		ClientID: "1",
	}

	route.LoadPositions()

	stringjson, err := route.ExportJsonPositions()

	if err != nil {
		fmt.Println(err)
	}
	//imprimir os dados
	fmt.Println(stringjson[0])
}
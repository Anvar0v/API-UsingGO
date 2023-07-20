package main

import (
	"fmt"
	"movies_api/services"
)

func main(){
	services := services.GetMovies();
	fmt.Print(services[len(services)-1])
}
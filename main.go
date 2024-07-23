package main

import (
	"golang/service"
	"os"
)

func main() {
	inputFile := "input.txt"
	outputFile := "output.txt"

	if len(os.Args) != 1 {
		args := os.Args[1:]
		inputFile = args[0]
		outputFile = args[1]
	}

	prod := service.NewProducer(inputFile)
	pres := service.NewPresenter(outputFile)

	serv := service.NewService(prod, pres)

	err := serv.Run()
	if err != nil {
		return
	}
}

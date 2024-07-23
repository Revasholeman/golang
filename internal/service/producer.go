package service

import (
	"bufio"
	"fmt"
	"os"
)

type Producer struct {
	filePath string
}

func NewProducer(filePath string) *Producer {
	return &Producer{filePath: filePath}
}

func (producer *Producer) produce() ([]string, error) {
	file, err := os.Open(producer.filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	var text []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text, nil
}

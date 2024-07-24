package service

import (
	"fmt"
	"os"
)

type Presenter struct {
	filePath string
}

func NewPresenter(filePath string) *Presenter {
	return &Presenter{filePath}
}

func (p *Presenter) present(text []string) error {
	file, err := os.Create(p.filePath)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for _, line := range text {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}

	}

	return nil
}

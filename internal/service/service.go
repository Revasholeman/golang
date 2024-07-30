package service

import (
	"sync"
)

//go:generate mockgen -source=service.go -destination=../internal/mock/service.go -package=mock
type producer interface {
	Produce() ([]string, error)
}

//go:generate mockgen -source=service.go -destination=../internal/mock/service.go -package=mock
type presenter interface {
	Present([]string) error
}

type Service struct {
	prod producer
	pres presenter
}

func NewService(prod producer, pres presenter) *Service {
	return &Service{pres: pres, prod: prod}
}

func (s *Service) Run() error {
	file, err := s.prod.Produce()
	wg := sync.WaitGroup{}
	if err != nil {
		return err
	}

	limit := 10

	fanIn := make(chan string)  // пишем результат SpamMasker
	fanOut := make(chan string) // извлекаем результат из канала в spamMasked

	semaphore := make(chan struct{}, limit) // ограничивает кол-во горутин по limit

	var spamMasked []string

	for i := 0; i < limit; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for _, text := range file {
				result := s.SpamMasker(text)
				fanIn <- result
				<-semaphore
			}
		}()
	}
	go func() {
		for text := range fanIn {
			semaphore <- struct{}{}
			fanOut <- text
		}
		close(fanIn)
	}()

	go func() {
		wg.Wait()
		close(fanOut)
	}()

	for val := range fanOut {
		spamMasked = append(spamMasked, val)
	}

	err = s.pres.Present(spamMasked)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SpamMasker(message string) string {
	http := [7]rune{'h', 't', 't', 'p', ':', '/', '/'}
	count := 0
	result := make([]rune, 0, len(message))
	for _, w := range message {
		switch {
		case w == ' ':
			result = append(result, w)
			count = 0
		case count == 7 && w != ' ':
			result = append(result, '*')
		case w == http[count]:
			result = append(result, w)
			count++
		default:
			result = append(result, w)
		}
	}
	return string(result)
}

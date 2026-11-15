package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает методы разбора данных и получения информации
// об активности. Его реализуют типы тренировок и дневных прогулок.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info последовательно обрабатывает записи об активности через интерфейс DataParser.
// Ошибочные записи журналируются и не прерывают обработку остальных.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("ошибка разбора данных: %v", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка получения информации: %v", err)
			continue
		}
		fmt.Print(info)
	}
}

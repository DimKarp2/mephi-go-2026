package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps хранит дневную активность и встроенные данные пользователя.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку с количеством шагов и длительностью прогулки.
// При ошибке входных данных возвращает её без изменения полей структуры.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат прогулки")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("неверное количество шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("неверная продолжительность: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// ActionInfo формирует строку с итогами прогулки.
// Если рассчитать калории не удалось, возвращает ошибку.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}

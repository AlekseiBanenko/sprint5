package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
    parts := strings.Split(datastring, ",")
    if len(parts) != 3 {
        return errors.New("некорректный формат строки")
    }

    steps, err := strconv.Atoi(parts[0])
    if err != nil {
        return err
    }
    **if steps <= 0 {
        return errors.New("некорректное количество шагов")
    }**
    t.Steps = steps

    t.TrainingType = parts[1]

    t.Duration, err = time.ParseDuration(parts[2])
    if err != nil {
        return err
    }
    **if t.Duration <= 0 {
        return errors.New("некорректная продолжительность")
    }**

    return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Steps == 0 || t.Duration == 0 {
		return "", errors.New("не хватает данных о тренировке")
	}

	dist := spentenergy.Distance(t.Steps, t.Height)        
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration) 

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%s\n"+           
			"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f",
		t.Name,       
		t.TrainingType,
		t.Duration.Hours(),
		dist,
		speed,
		calories,
	), nil
}


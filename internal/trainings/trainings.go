package trainings

import (
	"errors"
	"fmt"
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

	steps, err := personaldata.Atoi(parts[0])
	if err != nil {
		return err
	}
	t.Steps = steps

	t.TrainingType = parts[1]

	t.Duration, err = time.ParseDuration(parts[2])
	if err != nil {
		return err
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Steps == 0 || t.Duration == 0 {
		return "", errors.New("не хватает данных о тренировке")
	}

	dist := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
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
		t.Personal.Name,
		t.TrainingType,
		t.Duration.Hours(),
		dist,
		speed,
		calories,
	), nil
}

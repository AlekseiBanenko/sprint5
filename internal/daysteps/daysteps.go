package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"internal/personaldata"
	"internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	Personal personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("некорректный формат строки")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	ds.Steps = steps

	ds.Duration, err = time.ParseDuration(parts[1])
	if err != nil {
		return err
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 || ds.Duration <= 0 {
		return "", errors.New("некорректные входные данные")
	}

	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		ds.Steps,
		distance,
		calories,
	), nil
}

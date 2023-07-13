package dependPractice

import (
	"github.com/Esilahic/dog"
)

func Practice() string {
	return "Nice!"
}

func EvenMore() string {
	return "Good Job!"
}

func OldPractice() string {
	return dog.WhenGrownUp(Practice())
}

func OldMore() string {
	return dog.WhenGrownUp(EvenMore())
}

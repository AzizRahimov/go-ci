package main

import "fmt"

// % (Number of promotes) - % (Number of detractors)
// 100 respondents:
// - max of promoters - 100%, min - 0%
// - max of detractors - 100%, min - 0%
// граничные случаи
// 100% - 0 = 100%
// 0 - 100% = -100%
func main() {
	// problem 1: как хранить много данных
	score1 := 10
	score2 := 7
	score3 := 10



	// nps = 100
	promoters := 0
	detractors := 0
	// neutrals? - practice vs theory

	// if'ы - условия
	// boolean - тип данных
	// >, >=, <, <= - очень опасны (определяют границы)
	// проблема валидности данных: 0-10
	// problem 2: дублирование кода
	// problem 3: magic values - очень дурным тоном
	// 60, 12, 7 - magic values -> кол-во секунд 12 месяцев 7 дней

	// refactoring: улучшение структуры код без модификации поведения
	// ctrl + alt + v с выделением - позволяет создать локальную переменную
	// shift + f6 - переименование имени
	promotersLowerBound := 9
	detractorsUpperBound := 6
	// problem x: auto-testing

	if score1 >= promotersLowerBound {
		promoters = promoters + 1
	}

	if score1 <= detractorsUpperBound {
		detractors = detractors + 1
	}

	// ctrl + alt + shift + левый клик мыши (много курсоров)
	if score2 >= promotersLowerBound {
		promoters = promoters + 1
	}

	if score2 <= detractorsUpperBound {
		detractors = detractors + 1
	}
	if score3 >= promotersLowerBound {
		promoters = promoters + 1
	}

	if score3 <= detractorsUpperBound {
		detractors = detractors + 1
	}

	// nps := (2 - 0) / 3 * 100
	// 2 / 3 * 100 -> 0 * 100 -> 0
	// 2 * 100 / 3 -> 200 / 3 -> 66
	// nps := (promoters - detractors) / 3 * 100
	// 3 - magic values
	nps := (promoters - detractors) * 100 / 3
	fmt.Println(nps)
}

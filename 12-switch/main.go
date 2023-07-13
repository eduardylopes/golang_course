package main

import "fmt"

func dayOfWeek(number int) string {
	// When the first condition is true, golang will escape from switch after process internal code
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid number"
	}
}

func dayOfWeek2(number int) string {
	switch {
	case number == 1:
		return "Sunday"
	case number == 2:
		return "Monday"
	case number == 3:
		return "Tuesday"
	case number == 4:
		return "Wednesday"
	case number == 5:
		return "Thursday"
	case number == 6:
		return "Friday"
	case number == 7:
		return "Saturday"
	default:
		return "Invalid number"
	}
}

func dayOfWeek3(number int) string {
	var dayOfWeek string

	switch {
	case number == 1:
		dayOfWeek = "Sunday"
		fallthrough // This is used to goto inside next case if this condition matches
	case number == 2:
		dayOfWeek = "Monday"
	case number == 3:
		dayOfWeek = "Tuesday"
	case number == 4:
		dayOfWeek = "Wednesday"
	case number == 5:
		dayOfWeek = "Thursday"
	case number == 6:
		dayOfWeek = "Friday"
	case number == 7:
		dayOfWeek = "Saturday"
	default:
		dayOfWeek = "Invalid number"
	}

	return dayOfWeek
}

func main() {
	number := 1

	day := dayOfWeek(number)
	fmt.Println(day)

	day2 := dayOfWeek2(number)
	fmt.Println(day2)

	day3 := dayOfWeek3(number)
	fmt.Println(day3)
}
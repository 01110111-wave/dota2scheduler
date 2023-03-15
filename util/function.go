package util

import (
	"time"
	"unicode"
)

func intToCapitalChar(num int) rune {
	return rune(unicode.ToUpper(rune(num + 64)))
}

func IntToCapitalString(num int) string {
	return string(intToCapitalChar(num))
}

func RoundUpToHour(t time.Time) time.Time {
	if t.Minute() >= 50 {
		return t.Truncate(time.Hour).Add(time.Hour)
	}
	return t.Truncate(time.Hour)
}

func RegionIntToRegionString(number int) string {
	var region string

	switch number {
	case 0:
		region = "LAN"
	case 1:
		region = "NA"
	case 2:
		region = "SA"
	case 3:
		region = "WEU"
	case 4:
		region = "EEU"
	case 5:
		region = "CN"
	case 6:
		region = "SEA"
	default:
		region = "Invalid"
	}

	return region
}

func NodeTypeToBO(input int) int {
	switch input {
	case 1:
		return 1
	case 2:
		return 3
	case 3:
		return 5
	case 4:
		return 2
	default:
		return 0
	}
}

func MonthIntToString(monthNum int) string {
	switch monthNum {
	case 1:
		return "January"
	case 2:
		return "February"
	case 3:
		return "March"
	case 4:
		return "April"
	case 5:
		return "May"
	case 6:
		return "June"
	case 7:
		return "July"
	case 8:
		return "August"
	case 9:
		return "September"
	case 10:
		return "October"
	case 11:
		return "November"
	case 12:
		return "December"
	default:
		return "Invalid month number"
	}
}

func MonthToDays(month int) int {
	switch month {
	case 1:
		return 31
	case 2:
		return 28
	case 3:
		return 31
	case 4:
		return 30
	case 5:
		return 31
	case 6:
		return 30
	case 7:
		return 31
	case 8:
		return 31
	case 9:
		return 30
	case 10:
		return 31
	case 11:
		return 30
	case 12:
		return 31
	default:
		return 0
	}
}

func WeekdayToColor(weekday string) string {
	switch weekday {
	case "Monday":
		return "#FFFF00" // yellow
	case "Tuesday":
		return "#FF69B4" // pink
	case "Wednesday":
		return "#008000" // green
	case "Thursday":
		return "#FFA500" // orange
	case "Friday":
		return "#0000FF" // blue
	case "Saturday":
		return "#7600BC" // violet
	case "Sunday":
		return "#FF0000" // red
	default:
		return "#000000"
	}
}

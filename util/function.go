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

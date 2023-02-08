package writexls

import (
	"dotascheduler/util"
	"fmt"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func Writexls(leagues []util.League) excelize.File {

	scheculeTable := excelize.NewFile()
	sheet := "Sheet1"
	scheculeTable.SetCellValue(sheet, "A1", "dummy")
	for _, league := range leagues {
		for i := 2; i <= 26; i++ {
			fmt.Println(util.IntToCapitalString(i))
			if util.IntToCapitalString(i) == "B" {
				fmt.Println(util.IntToCapitalString(i) + "1")
				scheculeTable.SetCellValue(sheet, util.IntToCapitalString(i)+"1", "Time")
			} else {
				scheculeTable.SetCellValue(sheet, util.IntToCapitalString(i)+"1", strconv.Itoa(i-3)+":00")
			}
		}

		scheculeTable.SetCellValue(sheet, "A2", "date")
		for i := 1; i <= 31; i++ {
			scheculeTable.SetCellValue(sheet, "A"+strconv.Itoa(i+2), strconv.Itoa(i))
		}

		current_month := time.January
		style, err := scheculeTable.NewStyle(`{"fill":{"type":"pattern","color":["#0acbf1"],"pattern":1},"alignment":{"horizontal":"center"}}`)
		if err != nil {
			fmt.Println(err)
		}
		for _, match := range league.Matchs {
			if match.StartTime.Month() == current_month {
				fmt.Println("working")
				daterow := strconv.Itoa(match.StartTime.Day() + 2)
				timecolumn := util.IntToCapitalString(util.RoundUpToHour(match.StartTime).Hour() + 3)
				totimecolumn := util.IntToCapitalString(util.RoundUpToHour(match.StartTime).Hour() + 5)

				if match.BO > 1 {
					fmt.Println("merge from", timecolumn+daterow, "to", totimecolumn+daterow)
					scheculeTable.MergeCell(sheet, timecolumn+daterow, totimecolumn+daterow)
					scheculeTable.SetCellStyle(sheet, timecolumn+daterow, totimecolumn+daterow, style)
				}

				if match.BO == 1 {
					fmt.Println("tie break", match.Team1, match.Team2)
				}
				scheculeTable.SetCellValue(sheet, timecolumn+daterow, match.Team1.Tag+" vs "+match.Team2.Tag)
			}
		}
	}

	return *scheculeTable
}

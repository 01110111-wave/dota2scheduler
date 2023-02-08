package writexls

import (
	"dotascheduler/util"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/olekukonko/tablewriter"
)

func Writexls(leagues []util.League, month int) excelize.File {

	scheculeTable := excelize.NewFile()
	sheet := "Sheet1"
	scheculeTable.SetCellValue(sheet, "A1", "dummy")

	for _, league := range leagues {
		for i := 2; i <= 26; i++ {
			//fmt.Println(util.IntToCapitalString(i))
			if util.IntToCapitalString(i) == "B" {
				//fmt.Println(util.IntToCapitalString(i) + "1")
				scheculeTable.SetCellValue(sheet, util.IntToCapitalString(i)+"1", "Time")
			} else {
				scheculeTable.SetCellValue(sheet, util.IntToCapitalString(i)+"1", strconv.Itoa(i-3)+":00")
			}
		}
		timeToStream := map[time.Time][]util.Match{}

		for _, match := range league.Matchs {
			if int(match.StartTime.Month()) != month {
				continue
			}
			startTime := match.StartTime
			if _, exists := timeToStream[startTime]; !exists {
				timeToStream[startTime] = []util.Match{match}
			} else {
				timeToStream[startTime] = append(timeToStream[startTime], match)
			}
		}

		scheculeTable.SetCellValue(sheet, "A2", "date")
		for i := 1; i <= 31; i++ {
			scheculeTable.SetCellValue(sheet, "A"+strconv.Itoa(i+2), strconv.Itoa(i))
		}

		style, err := scheculeTable.NewStyle(`{"fill":{"type":"pattern","color":["#0acbf1"],"pattern":1},"alignment":{"horizontal":"center"}}`)
		switch league.Region {
		case "EEU":
			{
				style, err = scheculeTable.NewStyle(`{"fill":{"type":"pattern","color":["#ce7e00"],"pattern":1},"alignment":{"horizontal":"center"}}`)
			}
		case "WEU":
			{
				style, err = scheculeTable.NewStyle(`{"fill":{"type":"pattern","color":["#2986cc"],"pattern":1},"alignment":{"horizontal":"center"}}`)
			}
		case "NA":
			{
				style, err = scheculeTable.NewStyle(`{"fill":{"type":"pattern","color":["#c90076"],"pattern":1},"alignment":{"horizontal":"center"}}`)
			}
		case "SA":
			{
				style, err = scheculeTable.NewStyle(`{"fill":{"type":"pattern","color":["#c27ba0"],"pattern":1},"alignment":{"horizontal":"center"}}`)
			}
		case "SEA":
			{
				style, err = scheculeTable.NewStyle(`{"fill":{"type":"pattern","color":["#8fce00"],"pattern":1},"alignment":{"horizontal":"center"}}`)
			}
		case "CN":
			{
				style, err = scheculeTable.NewStyle(`{"fill":{"type":"pattern","color":["#ff5700"],"pattern":1},"alignment":{"horizontal":"center"}}`)
			}
		case "LAN":
			{
				style, err = scheculeTable.NewStyle(`{"fill":{"type":"pattern","color":["#ffe599"],"pattern":1},"alignment":{"horizontal":"center"}}`)
			}
		}
		if err != nil {
			fmt.Println(err)
		}

		for k := range timeToStream {
			daterow := strconv.Itoa(k.Day() + 2)
			timecolumn := util.IntToCapitalString(util.RoundUpToHour(k).Hour() + 3)
			matchend := util.RoundUpToHour(k).Hour() + 3 + timeToStream[k][0].BO - 1
			totimecolumn := util.IntToCapitalString(matchend)
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Index", "Time", "Name", "League"})
			if len(timeToStream[k]) > 1 {
				for i, v := range timeToStream[k] {
					table.Append([]string{strconv.Itoa(i), k.String(), v.Team1.Name + " vs " + v.Team2.Name, league.Name})
				}
				table.Render()
				var n int
				for {
					fmt.Print("please choose match watch from index : ")
					fmt.Scanln(&n)
					if n < len(timeToStream[k]) {
						break
					}
				}
				match := timeToStream[k][n]
				fmt.Println("choose to watch " + match.Team1.Tag + " vs " + match.Team2.Tag)
				if matchend > 24 {
					totimecolumn = "Z"
				}
				if match.BO > 1 {
					scheculeTable.MergeCell(sheet, timecolumn+daterow, totimecolumn+daterow)
					scheculeTable.SetCellStyle(sheet, timecolumn+daterow, totimecolumn+daterow, style)
				}

				if match.BO == 1 {
					fmt.Println("tie break", match.Team1, match.Team2)
				}
				scheculeTable.SetCellValue(sheet, timecolumn+daterow, match.Team1.Tag+" vs "+match.Team2.Tag)
			} else {
				match := timeToStream[k][0]
				if matchend > 24 {
					totimecolumn = "Z"
				}
				if match.BO > 1 {
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

package writexls

import (
	"dotascheduler/util"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	excelize "github.com/xuri/excelize/v2"

	"github.com/olekukonko/tablewriter"
)

func Writexls(leagues []util.League, month int) *excelize.File {

	scheculeTable := excelize.NewFile()
	sheet := "Sheet1"
	scheculeTable.SetCellValue(sheet, "A1", util.MonthIntToString(month))
	_ = scheculeTable.SetColWidth(sheet, "AA", "AB", 20)
	presentcolumn := 2
	for i := 2; i <= 26; i++ {
		if util.IntToCapitalString(i) == "B" {
			scheculeTable.SetCellValue(sheet, util.IntToCapitalString(i)+"1", "Time")
		} else {
			scheculeTable.SetCellValue(sheet, util.IntToCapitalString(i)+"1", strconv.Itoa(i-3)+":00")
		}
	}
	var startofmonth time.Time
	if month >= int(time.Now().Month()) {
		year, _, _ := time.Now().Date()
		startofmonth = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	} else {
		year, _, _ := time.Now().Date()
		startofmonth = time.Date(year+1, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	}

	for i := 1; i <= util.MonthToDays(month); i++ {
		scheculeTable.SetCellValue(sheet, "A"+strconv.Itoa(i+2), strconv.Itoa(i)+" "+startofmonth.Month().String()[0:3])
		style, _ := scheculeTable.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Color: []string{util.WeekdayToColor(startofmonth.Weekday().String())}, Pattern: 1}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"}})
		scheculeTable.SetCellStyle(sheet, "B"+strconv.Itoa(i+2), "B"+strconv.Itoa(i+2), style)
		//scheculeTable.SetCellValue(sheet, "B"+strconv.Itoa(i+2), startofmonth.Weekday().String())
		startofmonth = startofmonth.AddDate(0, 0, 1)
	}

	for _, league := range leagues {

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

		style, err := scheculeTable.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Color: []string{"#E0EBF5"}, Pattern: 1}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"}})
		switch league.Region {
		case "EEU":
			{
				style, err = scheculeTable.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Color: []string{"#f50989"}, Pattern: 1}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"}})
				scheculeTable.SetCellStyle(sheet, "AB"+strconv.Itoa(presentcolumn), "AB"+strconv.Itoa(presentcolumn), style)
				name := strings.Split(league.Name, " ")
				scheculeTable.SetCellValue(sheet, "AB"+strconv.Itoa(presentcolumn), name[0]+name[1]+name[2])
				scheculeTable.SetCellHyperLink(sheet, "AB"+strconv.Itoa(presentcolumn), league.URL, "External")
			}
		case "WEU":
			{
				style, err = scheculeTable.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Color: []string{"#0acbf1"}, Pattern: 1}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"}})
				scheculeTable.SetCellStyle(sheet, "AB"+strconv.Itoa(presentcolumn), "AB"+strconv.Itoa(presentcolumn), style)
				name := strings.Split(league.Name, " ")
				scheculeTable.SetCellValue(sheet, "AB"+strconv.Itoa(presentcolumn), name[0]+name[1]+name[2])
				scheculeTable.SetCellHyperLink(sheet, "AB"+strconv.Itoa(presentcolumn), league.URL, "External")
			}
		case "NA":
			{
				style, err = scheculeTable.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Color: []string{"#f2e20e"}, Pattern: 1}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"}})
				scheculeTable.SetCellStyle(sheet, "AB"+strconv.Itoa(presentcolumn), "AB"+strconv.Itoa(presentcolumn), style)
				name := strings.Split(league.Name, " ")
				scheculeTable.SetCellValue(sheet, "AB"+strconv.Itoa(presentcolumn), name[0]+name[1]+name[2])
				scheculeTable.SetCellHyperLink(sheet, "AB"+strconv.Itoa(presentcolumn), league.URL, "External")
			}
		case "SA":
			{
				style, err = scheculeTable.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Color: []string{"#f44336"}, Pattern: 1}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"}})
				scheculeTable.SetCellStyle(sheet, "AB"+strconv.Itoa(presentcolumn), "AB"+strconv.Itoa(presentcolumn), style)
				name := strings.Split(league.Name, " ")
				scheculeTable.SetCellValue(sheet, "AB"+strconv.Itoa(presentcolumn), name[0]+name[1]+name[2])
				scheculeTable.SetCellHyperLink(sheet, "AB"+strconv.Itoa(presentcolumn), league.URL, "External")
			}
		case "SEA":
			{
				style, err = scheculeTable.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Color: []string{"#8fce00"}, Pattern: 1}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"}})
				scheculeTable.SetCellStyle(sheet, "AB"+strconv.Itoa(presentcolumn), "AB"+strconv.Itoa(presentcolumn), style)
				name := strings.Split(league.Name, " ")
				scheculeTable.SetCellValue(sheet, "AB"+strconv.Itoa(presentcolumn), name[0]+name[1]+name[2])
				scheculeTable.SetCellHyperLink(sheet, "AB"+strconv.Itoa(presentcolumn), league.URL, "External")
			}
		case "CN":
			{
				style, err = scheculeTable.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Color: []string{"#2986cc"}, Pattern: 1}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"}})
				scheculeTable.SetCellStyle(sheet, "AB"+strconv.Itoa(presentcolumn), "AB"+strconv.Itoa(presentcolumn), style)
				name := strings.Split(league.Name, " ")
				scheculeTable.SetCellValue(sheet, "AB"+strconv.Itoa(presentcolumn), name[0]+name[1]+name[2])
				scheculeTable.SetCellHyperLink(sheet, "AB"+strconv.Itoa(presentcolumn), league.URL, "External")
			}
		case "LAN":
			{
				style, err = scheculeTable.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Color: []string{"#c90076"}, Pattern: 1}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"}})
				scheculeTable.SetCellStyle(sheet, "AB"+strconv.Itoa(presentcolumn), "AB"+strconv.Itoa(presentcolumn), style)
				name := strings.Split(league.Name, " ")
				scheculeTable.SetCellValue(sheet, "AB"+strconv.Itoa(presentcolumn), name[0]+name[1]+name[2])
				scheculeTable.SetCellHyperLink(sheet, "AB"+strconv.Itoa(presentcolumn), league.URL, "External")
			}
		}
		presentcolumn++
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
				if matchend > 26 {
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
				if matchend > 26 {
					nextdaterow := strconv.Itoa(k.Day() + 2 + 1)
					Extimecolumn := util.IntToCapitalString(3)
					Extotimecolumn := util.IntToCapitalString(3 + matchend - 26 - 1)
					// fmt.Println(matchend, match.StartTime, match.Team1.Tag+" vs "+match.Team2.Tag)
					// fmt.Println(Extimecolumn+nextdaterow, Extotimecolumn+nextdaterow)
					scheculeTable.MergeCell(sheet, timecolumn+daterow, "Z"+daterow)
					scheculeTable.MergeCell(sheet, Extimecolumn+nextdaterow, Extotimecolumn+nextdaterow)
					scheculeTable.SetCellStyle(sheet, timecolumn+daterow, "Z"+daterow, style)
					scheculeTable.SetCellStyle(sheet, Extimecolumn+nextdaterow, Extotimecolumn+nextdaterow, style)
					scheculeTable.SetCellValue(sheet, Extimecolumn+nextdaterow, match.Team1.Tag+" vs "+match.Team2.Tag)
				} else {
					scheculeTable.MergeCell(sheet, timecolumn+daterow, totimecolumn+daterow)
					scheculeTable.SetCellStyle(sheet, timecolumn+daterow, totimecolumn+daterow, style)
					scheculeTable.SetCellValue(sheet, timecolumn+daterow, match.Team1.Tag+" vs "+match.Team2.Tag)
				}
			}
		}
	}

	return scheculeTable
}

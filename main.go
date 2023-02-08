package main

import (
	"dotascheduler/callapi"
	"dotascheduler/util"
	"dotascheduler/writexls"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var leagueIDs string
var month int

func init() {
	flag.StringVar(&leagueIDs, "id", "none", "LeagueID to create schedule.Separate each ID with comma")
	flag.IntVar(&month, "month", int(time.Now().Month()), "Month to create schedule.Default is the current month.")
}

func main() {
	flag.Parse()
	if leagueIDs == "none" {
		fmt.Print("please input leagueID separate with comma (Ex.14859,14927,14921 ): ")
		fmt.Scanln(&leagueIDs)
	}
	stringIDs := strings.Split(leagueIDs, ",")
	IDs := []int{}

	for _, v := range stringIDs {
		ID, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		IDs = append(IDs, ID)
	}

	schedule := callapi.GetAllLeagueSchedule(IDs)
	//fmt.Println(schedule[0].Matchs)

	// for _, match := range schedule[0].Matchs {
	// 	fmt.Println(match.BO, match.Team1.Name, match.Team2.Name, match.StartTime)
	// }

	scheduleTable := writexls.Writexls(schedule, month)

	err := scheduleTable.SaveAs("./" + util.MonthIntToString(month) + "Dota2schedule.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success")
	fmt.Scanln()
	//fmt.Println(util.IntToCapitalString(2))
}

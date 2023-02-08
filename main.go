package main

import (
	"dotascheduler/callapi"
	"dotascheduler/util"
	"dotascheduler/writexls"
	"fmt"
)

func main() {
	schedule := callapi.GetAllLeagueSchedule([]int{14927, 14892})
	fmt.Println(schedule[0].Matchs)

	scheduleTable := writexls.Writexls(schedule)

	err := scheduleTable.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(util.IntToCapitalString(2))
}

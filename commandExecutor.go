package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Command string

const (
	AddDriverCmd = "ADD_DRIVER"
	AddRiderCmd  = "ADD_RIDER"
	MatchCmd     = "MATCH"
	StartRideCmd = "START_RIDE"
	StopRideCmd  = "STOP_RIDE"
	BillCmd      = "BILL"
)

func CommandExecutor(argList []string, Allmatches map[string][]string) {

	cmd := argList[0]
	switch cmd {
	case AddDriverCmd:

		AddDriver(argList[1], argList[2], argList[3])
		return

	case AddRiderCmd:

		AddRider(argList[1], argList[2], argList[3])
		return

	case MatchCmd:

		matches := Match(argList[1])
		Allmatches[argList[1]] = matches
		if len(matches) == 0 {
			fmt.Println("NO_DRIVERS_AVAILABLE")
			return
		}
		fmt.Printf("DRIVERS_MATCHED %s\n", strings.Join(matches, " "))
		return

	case StartRideCmd:

		if x, found := Allmatches[argList[3]]; found {
			idx, err := strconv.Atoi(argList[2])
			if err != nil {
				fmt.Println("invalid number", err)
				return
			}
			if len(x) < idx {
				fmt.Println("INVALID_RIDE")
				return
			}
			StartRide(argList[1], x[idx-1], argList[3])
		}
		return

	case StopRideCmd:

		StopRide(argList[1], argList[2], argList[3], argList[4])
		return

	case BillCmd:

		Bill(argList[1])
		return

	default:
		return

	}

}

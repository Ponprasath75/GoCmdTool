package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	BaseFare    = 50
	ServiceTax  = 0.2
	CostPerKM   = 6.5
	CostPerMin  = 2
	SquarePower = 2
)

func main() {

	cliArgs := os.Args[1:]
	Allmatches := make(map[string][]string)

	testDriver := []Driver{}

	for i := 0; i < 1000000; i++ {
		driver := Driver{
			DriverId:    "hello",
			DriverXCord: 1,
			DriverYCord: 3,
			IsAvailable: true,
		}

		testDriver = append(testDriver, driver)
	}

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")

		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")

		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		args := scanner.Text()
		argList := strings.Fields(args)
		CommandExecutor(argList, Allmatches)
	}

	fmt.Println(testDriver[56666], "Line 58")
	nextTest := testDriver
	summa(nextTest)
	fmt.Println(testDriver[56666], "Line 60")
	fmt.Println(nextTest[56666], "Line 60")
}

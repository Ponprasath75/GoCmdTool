package main

import (
	"fmt"
	"math"
	"strconv"
)

func getRiderInfo(riderId string) *Rider {
	var rider *Rider

	for _, riderInfo := range AllRiders {
		if riderInfo.RiderId == riderId {
			rider = riderInfo
			return rider
		}
	}

	return nil

}

func getRideInfo(rideId string) *Ride {
	var ride *Ride

	for _, rideInfo := range AllRides {
		if rideInfo.RideId == rideId {
			ride = rideInfo
			return ride
		}
	}

	return nil
}

func getDriverInfo(driverId string) *Driver {
	var driver *Driver

	for _, driverInfo := range AllDrivers {
		if driverInfo.DriverId == driverId {
			driver = driverInfo
			return driver
		}
	}

	return nil
}

func calcDistance(x1, y1, x2, y2 int) float64 {
	xdifsqrd := math.Pow((float64(x2) - float64(x1)), SquarePower)
	ydifsqrd := math.Pow((float64(y2) - float64(y1)), SquarePower)

	distance := math.Sqrt((xdifsqrd + ydifsqrd))

	distance = math.Round((distance * 100)) / 100

	return distance
}

func getIntFromString(strVal string) int {

	strValInt, err := strconv.Atoi(strVal)
	if err != nil {
		fmt.Println("passed string cannot be parsed to a valid int value returning 0", err)
		return 0
	}

	return strValInt
}

func summa(drivers []Driver) {
	drivers[56666].DriverId = "I changed it"
}

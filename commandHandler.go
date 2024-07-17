package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Driver struct {
	DriverId    string
	DriverXCord int
	DriverYCord int
	IsAvailable bool
}

type Rider struct {
	RiderId    string
	RiderXCord int
	RiderYCord int
}

type Ride struct {
	RideId      string
	DriverId    string
	RiderId     string
	StartPoint  string
	EndPoint    string
	IsCompleted bool
	Distance    float64
	Bill        float64
	Time        int
}

var AllDrivers []*Driver
var AllRiders []*Rider
var AllRides []*Ride

func AddDriver(driverId string, xcord string, ycord string) {

	driver := Driver{}
	//set driver id
	driver.DriverId = driverId
	//set drive x coord
	xcordInt := getIntFromString(xcord)
	driver.DriverXCord = xcordInt
	//set driver y coord
	ycordInt := getIntFromString(ycord)
	driver.DriverYCord = ycordInt
	//set driver availability
	driver.IsAvailable = true
	//push the driver instance to alldrivers slice
	AllDrivers = append(AllDrivers, &driver)
}

func AddRider(riderId string, xcord string, ycord string) {

	rider := Rider{}
	//set rider id
	rider.RiderId = riderId
	//set rider x coord
	xcordInt := getIntFromString(xcord)
	rider.RiderXCord = xcordInt
	//set rider y coord
	ycordInt := getIntFromString(ycord)
	rider.RiderYCord = ycordInt
	//push the rider instance to allriders slice
	AllRiders = append(AllRiders, &rider)
}

func Match(riderId string) []string {
	var rider *Rider
	var distanceDriverIdSlice []string
	var matches []string

	allowedDistance := 5

	rider = getRiderInfo(riderId)

	if rider == nil {
		fmt.Println("INVALID_RIDE")
	}

	for _, driverInfo := range AllDrivers {
		if driverInfo.IsAvailable {
			distance := calcDistance(rider.RiderXCord, rider.RiderYCord, driverInfo.DriverXCord, driverInfo.DriverYCord)
			if distance <= float64(allowedDistance) {
				s := fmt.Sprintf("%.2f-%s", distance, driverInfo.DriverId)
				distanceDriverIdSlice = append(distanceDriverIdSlice, s)
			}
		}
	}

	sort.Strings(distanceDriverIdSlice)

	for _, driverId := range distanceDriverIdSlice {
		matches = append(matches, strings.Split(driverId, "-")[1])
	}

	return matches

}

func StartRide(rideId string, driverId string, riderId string) {

	var ride *Ride
	ride = getRideInfo(rideId)
	if ride != nil {
		fmt.Println("INVALID_RIDE")
		return
	}
	ride = &Ride{}

	rider := getRiderInfo(riderId)
	if rider == nil {
		fmt.Println("INVALID_RIDE")
		return
	}

	driver := getDriverInfo(driverId)
	if driver == nil {
		fmt.Println("INVALID_RIDE")
		return
	}
	if !driver.IsAvailable {
		fmt.Println("INVALID_RIDE")
		return
	}
	driver.IsAvailable = false

	ride.RiderId = riderId
	ride.DriverId = driverId
	ride.IsCompleted = false
	ride.RideId = rideId
	ride.StartPoint = fmt.Sprintf("%d,%d", rider.RiderXCord, rider.RiderYCord)
	fmt.Printf("RIDE_STARTED %s\n", rideId)
	AllRides = append(AllRides, ride)

}

func StopRide(rideId, destXCord, destYCord, timeTakenInMin string) {

	var driver *Driver
	var ride *Ride

	destXCordInt := getIntFromString(destXCord)
	destYCordInt := getIntFromString(destYCord)
	timeTakenInMinInt := getIntFromString(timeTakenInMin)

	ride = getRideInfo(rideId)
	if ride == nil {
		fmt.Println("INVALID_RIDE")
		return
	}
	if ride.IsCompleted {
		fmt.Println("INVALID_RIDE")
		return
	}

	startCords := strings.Split(ride.StartPoint, ",")
	startXCordInt := getIntFromString(startCords[0])
	startYCordInt := getIntFromString(startCords[1])
	rideDistance := calcDistance(startXCordInt, startYCordInt, destXCordInt, destYCordInt)
	ride.IsCompleted = true
	ride.Time = timeTakenInMinInt
	ride.Distance = rideDistance
	ride.EndPoint = fmt.Sprintf("%d,%d", destXCordInt, destYCordInt)
	driver = getDriverInfo(ride.DriverId)
	driver.IsAvailable = true
	driver.DriverXCord = destXCordInt
	driver.DriverYCord = destYCordInt

	fmt.Printf("RIDE_STOPPED %s\n", rideId)

}

func Bill(rideId string) {

	ride := getRideInfo(rideId)

	if ride == nil {
		fmt.Println("INVALID_RIDE")
		return
	}

	if !ride.IsCompleted {
		fmt.Println("RIDE_NOT_COMPLETED")
		return
	}

	distanceCharge := ride.Distance * CostPerKM
	distanceCharge = math.Round(distanceCharge*100) / 100
	timeCharge := ride.Time * CostPerMin

	bill := distanceCharge + float64(timeCharge) + float64(BaseFare)
	totalBill := bill + ServiceTax*(bill)

	totalBill = math.Round(totalBill*100) / 100

	ride.Bill = totalBill

	fmt.Printf("BILL %s %s %.2f\n", ride.RideId, ride.DriverId, ride.Bill)

}

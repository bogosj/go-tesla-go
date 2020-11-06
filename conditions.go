package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/bogosj/tesla"
	"github.com/umahmood/haversine"
)

func checkConditions(flags flags, chs *tesla.ChargeState, ds *tesla.DriveState, cls *tesla.ClimateState) {
	if flags.ifBlockedDates != "" {
		now := time.Now()
		for _, d := range strings.Split(flags.ifBlockedDates, ",") {
			t, err := time.Parse("2006-01-02", d)
			if err != nil {
				panic(err)
			}
			if t.Year() == now.Year() && t.YearDay() == now.YearDay() {
				log.Printf("running on %s, exiting\n", d)
				os.Exit(1)
			}
		}
	}

	if isFlagPassed("if_speed_above") {
		log.Printf("car speed is %v\n", ds.Speed)
		if ds.Speed > flags.ifSpeedAbove {
			log.Printf("speed is above %v\n", flags.ifSpeedAbove)
		} else {
			log.Printf("speed is below %v, exiting\n", flags.ifSpeedAbove)
			os.Exit(1)
		}
	}

	if isFlagPassed("if_speed_below") {
		log.Printf("car speed is %v\n", ds.Speed)
		if ds.Speed < flags.ifSpeedBelow {
			log.Printf("speed is below %v\n", flags.ifSpeedBelow)
		} else {
			log.Printf("speed is above %v, exiting\n", flags.ifSpeedBelow)
			os.Exit(1)
		}
	}

	if flags.ifPluggedIn {
		cs := chs.ChargingState
		log.Printf("charge state: %v\n", cs)
		if cs == "Charging" || cs == "Complete" {
			log.Println("car is plugged in, continuing")
		} else {
			log.Println("car is not plugged in, exiting")
			os.Exit(1)
		}
	}

	if flags.ifInGeofence {
		testPoint := haversine.Coord{Lat: flags.lat, Lon: flags.lon}
		carPoint := haversine.Coord{Lat: ds.Latitude, Lon: ds.Longitude}
		miles, _ := haversine.Distance(testPoint, carPoint)
		log.Printf("car is %v miles from test point\n", miles)
		if miles > flags.miles {
			log.Printf("car is too far from test point, exiting")
			os.Exit(1)
		}
	}

	if flags.ifInsideTempOver != 0 {
		checkTemp := ftoc(flags.ifInsideTempOver)
		if checkTemp > cls.InsideTemp {
			log.Printf("temp of %v is greater than inside temp of %v, exiting\n", checkTemp, cls.InsideTemp)
			os.Exit(1)
		}
	}

	if flags.ifOutsideTempOver != 0 {
		checkTemp := ftoc(flags.ifOutsideTempOver)
		if checkTemp > cls.OutsideTemp {
			log.Printf("temp of %v is greater than outside temp of %v, exiting\n", checkTemp, cls.OutsideTemp)
			os.Exit(1)
		}
	}

	if flags.ifInsideTempUnder != 0 {
		checkTemp := ftoc(flags.ifInsideTempUnder)
		if checkTemp < cls.InsideTemp {
			log.Printf("temp of %v is less than inside temp of %v, exiting\n", checkTemp, cls.InsideTemp)
			os.Exit(1)
		}
	}

	if flags.ifOutsideTempUnder != 0 {
		checkTemp := ftoc(flags.ifOutsideTempUnder)
		if checkTemp < cls.OutsideTemp {
			log.Printf("temp of %v is less than outside temp of %v, exiting\n", checkTemp, cls.OutsideTemp)
			os.Exit(1)
		}
	}
}

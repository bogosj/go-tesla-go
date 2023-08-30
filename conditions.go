package main

import (
	"strings"
	"time"

	"github.com/bogosj/tesla"
	log "github.com/sirupsen/logrus"
	"github.com/umahmood/haversine"
)

func checkConditions(flags flags, sr *tesla.VehicleData) {
	ds := sr.Response.DriveState
	chs := sr.Response.ChargeState
	cls := sr.Response.ClimateState
	if flags.ifBlockedDates != "" {
		now := time.Now()
		for _, d := range strings.Split(flags.ifBlockedDates, ",") {
			t, err := time.Parse("2006-01-02", d)
			if err != nil {
				log.Fatal(err)
			}
			if t.Year() == now.Year() && t.YearDay() == now.YearDay() {
				log.Fatalf("running on %s, exiting", d)
			}
		}
	}

	if isFlagPassed("if_battery_above") {
		log.Printf("battery level is %v", chs.BatteryLevel)
		if chs.BatteryLevel < flags.ifBatteryAbove {
			log.Fatalf("battery level is below %v, exiting", flags.ifBatteryAbove)
		}
		log.Printf("battery level is above %v", flags.ifBatteryAbove)
	}

	if isFlagPassed("if_battery_below") {
		log.Printf("battery level is %v", chs.BatteryLevel)
		if chs.BatteryLevel > flags.ifBatteryBelow {
			log.Fatalf("battery level is above %v, exiting", flags.ifBatteryBelow)
		}
		log.Printf("battery level is below %v", flags.ifBatteryBelow)
	}

	if isFlagPassed("if_speed_above") {
		log.Printf("car speed is %v", ds.Speed)
		if ds.Speed < flags.ifSpeedAbove {
			log.Fatalf("speed is below %v, exiting", flags.ifSpeedAbove)
		}
		log.Printf("speed is above %v", flags.ifSpeedAbove)
	}

	if isFlagPassed("if_speed_below") {
		log.Printf("car speed is %v", ds.Speed)
		if ds.Speed > flags.ifSpeedBelow {
			log.Fatalf("speed is above %v, exiting", flags.ifSpeedBelow)
		}
		log.Printf("speed is below %v", flags.ifSpeedBelow)
	}

	if flags.ifPluggedIn {
		cs := chs.ChargingState
		log.Printf("charge state: %v", cs)
		if cs == "Disconnected" {
			log.Fatalf("car is not plugged in, exiting")
		}
		log.Println("car is plugged in, continuing")
	}

	if flags.ifInGeofence {
		testPoint := haversine.Coord{Lat: flags.lat, Lon: flags.lon}
		carPoint := haversine.Coord{Lat: ds.Latitude, Lon: ds.Longitude}
		miles, _ := haversine.Distance(testPoint, carPoint)
		log.Printf("car is %v miles from test point", miles)
		if miles > flags.miles {
			log.Fatalf("car is too far from test point, exiting")
		}
	}

	if isFlagPassed("if_inside_temp_over") {
		checkTemp := ftoc(flags.ifInsideTempOver)
		if checkTemp > cls.InsideTemp {
			log.Fatalf("temp of %v is greater than inside temp of %v, exiting", checkTemp, cls.InsideTemp)
		}
		log.Printf("temp of %v is less then inside temp of %v", checkTemp, cls.InsideTemp)
	}

	if isFlagPassed("if_outside_temp_over") {
		checkTemp := ftoc(flags.ifOutsideTempOver)
		if checkTemp > cls.OutsideTemp {
			log.Fatalf("temp of %v is greater than outside temp of %v, exiting", checkTemp, cls.OutsideTemp)
		}
		log.Printf("temp of %v is less than outside temp of %v", checkTemp, cls.OutsideTemp)
	}

	if isFlagPassed("if_inside_temp_under") {
		checkTemp := ftoc(flags.ifInsideTempUnder)
		if checkTemp < cls.InsideTemp {
			log.Fatalf("temp of %v is less than inside temp of %v, exiting", checkTemp, cls.InsideTemp)
		}
		log.Printf("temp of %v is greater than inside temp of %v", checkTemp, cls.InsideTemp)
	}

	if isFlagPassed("if_outside_temp_under") {
		checkTemp := ftoc(flags.ifOutsideTempUnder)
		if checkTemp < cls.OutsideTemp {
			log.Fatalf("temp of %v is less than outside temp of %v, exiting", checkTemp, cls.OutsideTemp)
		}
		log.Printf("temp of %v is greater than outside temp of %v", checkTemp, cls.OutsideTemp)
	}
}

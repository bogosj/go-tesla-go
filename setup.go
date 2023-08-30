package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/bogosj/go-tesla-go/config"
	"github.com/bogosj/tesla"
	log "github.com/sirupsen/logrus"
)

type flags struct {
	// Actions
	spew, startAC, stopAC bool
	setChargeLimit        int
	setTemp               float64

	// Conditions
	ifInGeofence    bool
	lat, lon, miles float64

	ifPluggedIn bool

	ifInsideTempOver, ifInsideTempUnder   float64
	ifOutsideTempOver, ifOutsideTempUnder float64

	ifSpeedBelow, ifSpeedAbove float64

	ifBatteryAbove, ifBatteryBelow int

	ifBlockedDates string
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func setupFlags() flags {
	h1 := flag.Bool("help", false, "display this message")
	h2 := flag.Bool("h", false, "display this message")

	// Actions
	spew := flag.Bool("spew", false, "emit the data retrieved from the API. emits regardless of conditions")
	scl := flag.Int("set_charge_limit", 0, "set to an integer for the percent SOC desired")
	st := flag.Float64("set_temp", 0, "sets the interior temperature (driver and passenger), in F")
	sac := flag.Bool("start_ac", false, "turns the A/C on")
	stac := flag.Bool("stop_ac", false, "turns the A/C off")

	// Conditions
	iigf := flag.Bool("if_in_geofence", false, "set to test car position. must set lat/lon/miles")
	lat := flag.Float64("lat", 0, "latitude of the center of the geofence")
	lon := flag.Float64("lon", 0, "longitude of the center of the geofence")
	miles := flag.Float64("miles", 0, "distance in miles of the radius of the geofence")

	ipi := flag.Bool("if_plugged_in", true, "execute commands only if the car is plugged in")

	iito := flag.Float64("if_inside_temp_over", 0, "set to test the inside temp, in F")
	ioto := flag.Float64("if_outside_temp_over", 0, "set to test the outside temp, in F")
	iitu := flag.Float64("if_inside_temp_under", 0, "set to test the inside temp, in F")
	iotu := flag.Float64("if_outside_temp_under", 0, "set to test the outside temp, in F")

	isa := flag.Float64("if_speed_above", 0, "set to test the speed of the car")
	isb := flag.Float64("if_speed_below", 0, "set to test the speed of the car")

	iba := flag.Int("if_battery_above", 0, "set to test the battery level")
	ibb := flag.Int("if_battery_below", 0, "set to test the battery level")

	ibd := flag.String("if_blocked_dates", "", "dates (YYYY-MM-DD) comma separated to not execute on")

	flag.Parse()
	if *h1 || *h2 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	return flags{
		spew:               *spew,
		setChargeLimit:     *scl,
		setTemp:            *st,
		startAC:            *sac,
		stopAC:             *stac,
		ifInGeofence:       *iigf,
		lat:                *lat,
		lon:                *lon,
		miles:              *miles,
		ifPluggedIn:        *ipi,
		ifInsideTempOver:   *iito,
		ifInsideTempUnder:  *iitu,
		ifOutsideTempOver:  *ioto,
		ifOutsideTempUnder: *iotu,
		ifSpeedAbove:       *isa,
		ifSpeedBelow:       *isb,
		ifBlockedDates:     *ibd,
		ifBatteryAbove:     *iba,
		ifBatteryBelow:     *ibb,
	}
}

func newTeslaClient(c config.Config) *tesla.Client {
	client, err := tesla.NewClient(context.Background(), tesla.WithTokenFile(c.OAuthTokenPath))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func vehicleByVin(client *tesla.Client, vin string) (*tesla.Vehicle, error) {
	vehicles, err := client.Vehicles()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range vehicles {
		if v.Vin == vin {
			return v, nil
		}
	}

	return nil, fmt.Errorf("could not find vehicle with vin %q", vin)
}

func wakeup(vehicle *tesla.Vehicle) {
	i := 1
	for {
		log.Printf("waking up car, attempt #%v", i)
		_, err := vehicle.Wakeup()
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
		i++
	}
}

func getData(vehicle *tesla.Vehicle) (ds *tesla.VehicleData) {
	i := 1
	for {
		log.Printf("getting drive state, attempt #%v, v:%v", i, vehicle)
		s, err := vehicle.Data()
		if err == nil {
			ds = s
			break
		} else {
			log.Errorf("error: %s", err)
		}
		time.Sleep(2 * time.Second)
		i++
	}

	return
}

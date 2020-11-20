package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/bogosj/go-tesla-go/config"
	"github.com/bogosj/tesla"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

type flags struct {
	configFilePath string
	configFileSet  bool

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

func isConfigFileFlagPassed() bool {
	if isFlagPassed("config_file") {
		w := color.New(color.FgRed, color.Bold).PrintFunc()
		w("Warning: ")
		m := "config_file is deprecated, please switch to use env based config"
		fmt.Println(m)
		log.Warn(m)
		return true
	}
	return false
}

func setupFlags() flags {
	cf := flag.String("config_file", "/gtg.config.json", "path to the config file (deprecated)")
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

	ibd := flag.String("if_blocked_dates", "", "dates (YYYY-MM-DD) comma separated to not execute on")

	flag.Parse()
	if *h1 || *h2 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	return flags{
		configFilePath:     *cf,
		configFileSet:      isConfigFileFlagPassed(),
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
	}
}

func newTeslaClient(c config.Config) *tesla.Client {
	client, err := tesla.NewClient(
		&tesla.Auth{
			ClientID:     c.ClientID,
			ClientSecret: c.ClientSecret,
			Email:        c.Email,
			Password:     c.Password,
		})
	if err != nil {
		panic(err)
	}
	return client
}

func vehicleByVin(client *tesla.Client, vin string) (*struct{ *tesla.Vehicle }, error) {
	vehicles, err := client.Vehicles()
	if err != nil {
		panic(err)
	}

	for _, v := range vehicles {
		if v.Vin == vin {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("could not find vehicle with vin %q", vin)
}

func wakeup(vehicle *struct{ *tesla.Vehicle }) {
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

func getData(vehicle *struct{ *tesla.Vehicle }) (ds *tesla.DriveState, chs *tesla.ChargeState, cls *tesla.ClimateState) {
	i := 1
	for {
		log.Printf("getting drive state, attempt #%v", i)
		s, err := vehicle.DriveState()
		if err == nil {
			ds = s
			break
		}
		time.Sleep(2 * time.Second)
		i++
	}

	i = 1
	for {
		log.Printf("getting charge state, attempt #%v", i)
		s, err := vehicle.ChargeState()
		if err == nil {
			chs = s
			break
		}
		time.Sleep(2 * time.Second)
		i++
	}

	i = 1
	for {
		log.Printf("getting climate state, attempt #%v", i)
		s, err := vehicle.ClimateState()
		if err == nil {
			cls = s
			break
		}
		time.Sleep(2 * time.Second)
		i++
	}

	return
}

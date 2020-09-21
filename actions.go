package main

import (
	"log"

	"github.com/bogosj/tesla"
)

func takeActions(flags flags, vehicle *struct{ *tesla.Vehicle }) {
	if flags.setChargeLimit != 0 {
		log.Printf("setting charge limit to %v\n", flags.setChargeLimit)
		vehicle.SetChargeLimit(flags.setChargeLimit)
	}

	if flags.startAC {
		log.Println("starting the A/C")
		vehicle.StartAirConditioning()
	}

	if flags.stopAC {
		log.Println("stopping the A/C")
		vehicle.StopAirConditioning()
	}

	if flags.setTemp != 0 {
		t := ftoc(flags.setTemp)
		log.Printf("setting the interior temp to %vC\n", t)
		vehicle.SetTemprature(t, t)
	}

}

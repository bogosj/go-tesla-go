package main

import (
	"github.com/bogosj/go-tesla-go/config"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	flags := setupFlags()
	c := config.New(flags.configFilePath)
	client := newTeslaClient(c)

	vehicle, err := vehicleByVin(client, c.VIN)
	if err != nil {
		panic(err)
	}

	wakeup(vehicle)
	ds, chs, cls := getData(vehicle)
	if flags.spew {
		spew.Dump(ds)
		spew.Dump(chs)
		spew.Dump(cls)
	}
	checkConditions(flags, chs, ds, cls)
	takeActions(flags, vehicle)
}

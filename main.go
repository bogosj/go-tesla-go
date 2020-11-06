package main

import (
	"github.com/bogosj/go-tesla-go/config"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	flags := setupFlags()
	var c config.Config
	if flags.configFileSet {
		c = config.New(flags.configFilePath)
	} else {
		c = config.FromEnv()
	}
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

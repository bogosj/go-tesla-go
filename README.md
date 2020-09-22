# go-tesla-go

A golang binary + docker image to make API calls to a Tesla vehicle.

## Usage

The binary can be run from the docker image [ghcr.io/bogosj/go-tesla-go](https://github.com/users/bogosj/packages/container/package/go-tesla-go).

```
docker run --rm \
  -v /home/bogosj/docker/go-tesla-go/gtg.config.json:/gtg.config.json \
  ghcr.io/bogosj/go-tesla-go --spew
```

The binary expects a configuration file to be mounted at /gtg.config.json, however this can be changed with `--config_file`. The structure of the configuration file can be found in [example.config.json](./example.config.json).

The values for ClientID and ClientSecret can be found in instructions at https://tesla-api.timdorr.com/api-basics/authentication. Email/Password/VIN should be self explanitory.

You can also install a local copy with `go get github.com/bogosj/go-tesla-go`.

## Flags
The flags that start "if" verify conditions before taking actions.

```
  -config_file string
        path to the config file (default "/gtg.config.json")
  -h    display this message
  -help
        display this message
  -if_in_geofence
        set to test car position. must set lat/lon/miles
  -if_inside_temp_over float
        set to test the inside temp, in F
  -if_inside_temp_under float
        set to test the inside temp, in F
  -if_outside_temp_over float
        set to test the outside temp, in F
  -if_outside_temp_under float
        set to test the outside temp, in F
  -if_plugged_in
        execute commands only if the car is plugged in (default true)
  -lat float
        latitude of the center of the geofence
  -lon float
        longitude of the center of the geofence
  -miles float
        distance in miles of the radius of the geofence
  -set_charge_limit int
        set to an integer for the percent SOC desired
  -set_temp float
        sets the interior temperature (driver and passenger), in F
  -spew
        emit the data retrieved from the API. emits regardless of conditions
  -start_ac
        turns the A/C on
  -stop_ac
        turns the A/C off
```

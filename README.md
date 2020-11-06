# go-tesla-go

A golang binary + docker image to make API calls to a Tesla vehicle.

## Usage

The binary can be run from the docker image [ghcr.io/bogosj/go-tesla-go](https://github.com/users/bogosj/packages/container/package/go-tesla-go).

```
docker run --rm ghcr.io/bogosj/go-tesla-go --spew
```

The binary expects five environment variables to be set:

```
GTG_CLIENTID
GTG_CLIENTSECRET
GTG_EMAIL
GTG_PASSWORD
GTG_VIN
```

These can be passed in the docker command with either the -e/--env flag or --env-file.

The values for GTG_CLIENTID and GTG_CLIENTSECRET can be found in instructions at https://tesla-api.timdorr.com/api-basics/authentication. Email/Password/VIN should be self explanitory.

### Deprecated
The binary used to expect a configuration file to be mounted at /gtg.config.json, however this can be changed with `--config_file`. The structure of the configuration file can be found in [config/example.config.json](./config/example.config.json). This flag will be removed in v2.0.

You can also install a local copy with `go get github.com/bogosj/go-tesla-go`.

## Flags
The flags that start "if" verify conditions before taking actions.

```
  -config_file string
        path to the config file (deprecated) (default "/gtg.config.json")
  -h    display this message
  -help
        display this message
  -if_blocked_dates string
        dates (YYYY-MM-DD) comma separated to not execute on
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
  -if_speed_above float
        set to test the speed of the car
  -if_speed_below float
        set to test the speed of the car
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

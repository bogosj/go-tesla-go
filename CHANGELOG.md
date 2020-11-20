# Changelog

## vNext

#### Bug Fixes:

* Removed \n from the end of log lines.

---

## v1.4.1

#### Bug Fixes:

* Fixed check --if_plugged_in not recognizing "NoPower" state.

---

## v1.4

#### Enhancements:

* Switch logs to the logrus library.

---

## v1.3.1

* Rebuilt with go v1.15.5.

---

## v1.3

#### Enhancements:

* Added --if_speed_above and --if_speed_below flags.

#### Bug Fixes:

* Allowed for testing temperature above or below 0.

---

## v1.2.1

* Rebuilt Docker image with go 1.15.4 and alpine 3.12.1.

---

## v1.2

#### Enhancements:

* Added the ability to configure via environment variables.

---

## v1.1

#### Enhancements:

* Added --if_blocked_dates flag.

---

## v1.0

Initial release.

# knozone
## Overview
The purpose of the knozone package is to be able to reliably use a known version of the IANA time zone database with Go.

In Go's time package, the LoadLocation method will search multiple possible locations to find a time zone database as described by this excerpt from the source code for the package.
```
// The time zone database needed by LoadLocation may not be
// present on all systems, especially non-Unix systems.
// LoadLocation looks in the directory or uncompressed zip file
// named by the ZONEINFO environment variable, if any, then looks in
// known installation locations on Unix systems,
// and finally looks in $GOROOT/lib/time/zoneinfo.zip.
```
For example, if Go fails to find a valid time zone file at the location specified by the ZONEINFO environment variable, then it will proceed to look in the OS on Unix systems and then look in the GOROOT. These backup locations may be using a different version of the time zone database. When Go fails to use the ZONEINFO location, it fails silently, so there is no way to tell if Go has defaulted to a backup location.

Fortunately, in Go 1.10 the [LoadLocationFromTZData function](https://golang.org/pkg/time/#LoadLocationFromTZData) was added that allows us to directly specify the data we'd like to load the location from. This is not the most user-friendly method though as we have to hand it the bytes that we want to use.

This is where knozone comes in. With knozone you can create a directory of timezone files that are structured in the exact same format as you would get from extracting Go's zoneinfo.zip file (found at $GOROOT/lib/time/zoneinfo.zip). You can then use knozone.LoadLocation to reliably load a time zone file from that directory or fail if it is not found in that specific directory. Knozone does the work of conjuring up the right bytes.

## Setting up the time zone database directory
I'm sure there are several ways to produce a directory of time zone database files, but the following method requires only a Go installation. **The instructions are for a Unix system.** If you are interested or have instructions for Windows, let me know and I'll see if I can figure out what's different there.

1. Find the following file <$GOROOT>/lib/time/update.bash
1. In the same directory will be a file called zoneinfo.zip, which is a zipped up form of the IANA time zone database, and is the product of running update.bash.
1. Read update.bash to see what values have been attributed to the CODE and DATA variables. These will likely be the same value and they refer to the version of the IANA time zone database that the script will generate. The format will be a year and a letter (e.g. 2018e). If you would like to change the version of the database, edit update.bash to change the value of these variables. If you are fine with the specified version then you can use the zoneinfo.zip file already present in the same directory and you can skip down to the extraction step.
1. (optional) update.bash will overwrite the zoneinfo.zip file, so if you'd like to make a backup, now is the time.
1. Run update.bash. This should generate a new zoneinfo.zip
1. Extract the files from zoneinfo.zip. Knozone cannot use the zip format of the database.
1. Move the extracted directory to whichever location you would like to put the files. By default knozone expects that a /zoneinfo directory will be located in the root of the main application, but the package allows you to specify the path to your time zone database directory.

## Functions
#### func GetZoneInfoPath
```func GetZoneInfoPath() string```
View the current value for the path to the time zone database directory.

#### func SetZoneInfoPath
```func SetZoneInfoPath(path string)```
Set the path to the time zone database directory

#### func LoadLocation
```func LoadLocation(name string) (*time.Location, error)```
LoadLocation has the same signature as the LoadLocation method from Go's time package. It looks in the specified time zone database directory and returns a \*time.Location if successful or an error if unsuccessful.

## Example
```
const MY_TIME_ZONE_DB_PATH string = "./data/zoneinfo"

//Set the path to the location of our time zone database directory
knozone.SetZoneInfoPath(MY_TIME_ZONE_DB_PATH)

//Load the Africa/Kigali location from our directory
loc, err := knozone.LoadLocation("Africa/Kigali")

//Set a time using our location - using the time package exclusively here
kigaliTime := time.Date(2020, time.January, 2, 9, 45, 0, 0, loc)
```

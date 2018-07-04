package knozone

import (
	"io/ioutil"
	"time"
)

//Default path to zoneinfo directory, can be changed with SetZoneInfoPath
var zoneInfoPath string = "./zoneinfo"

//GetZoneInfoPath allows the user to view the value of zoneInfoPath
func GetZoneInfoPath() string {
	return zoneInfoPath
}

//SetZoneInfoPath allows the user to reset the path to their zone info directory
func SetZoneInfoPath(path string) {
	zoneInfoPath = path
}

//getBytesFromZoneFile attempts to find the time zone file (name) in the
//directory and return the bytes for the file
func getBytesFromZoneFile(directory, name string) ([]byte, error) {
	if directory != "" {
		name = directory + "/" + name
	}
	return ioutil.ReadFile(name)
}

//LoadLocation searches the directory for the time zone files specified by name
//and returns a valid *time.Location if the corresponding file is found
func LoadLocation(name string) (*time.Location, error) {
	zoneBytes, err := getBytesFromZoneFile(zoneInfoPath, name)
	if err != nil {
		return nil, err
	}
	location, err := time.LoadLocationFromTZData(name, zoneBytes)
	if err != nil {
		return nil, err
	}
	return location, nil
}

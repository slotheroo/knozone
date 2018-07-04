package knozone

import "testing"

const TEST_PATH string = "./testdata"
const TEST_NAME_GOOD string = "Antarctica/Troll"
const TEST_NAME_BAD string = "Antarctica/Bad_Data"
const TEST_NAME_MISSING string = "Westeros/Kings_Landing"

func TestGetZoneInfoPath(t *testing.T) {
	path := GetZoneInfoPath()
	if path != zoneInfoPath {
		t.Error("Paths do not match. Expected " + path + " but got " + zoneInfoPath)
	}
}

func TestSetZoneInfoPath(t *testing.T) {
	newPath := "./data/tzdatabase"
	SetZoneInfoPath(newPath)
	path := GetZoneInfoPath()
	if path != newPath {
		t.Error("Setting path failed. Expected " + newPath + " but got " + path)
	}
}

func TestGetBytesFromZoneFile(t *testing.T) {
	_, err := getBytesFromZoneFile(TEST_PATH, TEST_NAME_GOOD)
	if err != nil {
		t.Error("Could not read test file " + TEST_NAME_GOOD + " in directory " + TEST_PATH)
	}
}

func TestLoadLocation(t *testing.T) {
	SetZoneInfoPath(TEST_PATH)
	location, err := LoadLocation(TEST_NAME_GOOD)
	if err != nil {
		t.Error("Could not load test location " + TEST_NAME_GOOD)
	}
	if location.String() != TEST_NAME_GOOD {
		t.Error("Location string does not match specified data")
	}

	location, err = LoadLocation(TEST_NAME_BAD)
	if err == nil {
		t.Error("Error was nil when a bad time zone file was provided")
	}

	location, err = LoadLocation(TEST_NAME_MISSING)
	if err == nil {
		t.Error("Error was nil when non-existent file name provided")
	}
}

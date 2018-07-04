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

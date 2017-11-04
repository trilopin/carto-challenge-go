package main

import (
	"fmt"
	"os"
)

// https://s3.amazonaws.com/carto-1000x/data/yellow_tripdata_2016-01.csv
func main() {
	switch os.Args[1] {
	case "stream_simple":
		StreamSimple(os.Args[2])
	default:
		fmt.Println("Usage: carto-challenge METHOD URL")
	}
}

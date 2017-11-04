package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func sumFloatField(r io.Reader, s byte, p int) (int, float64) {
	var (
		count int
		sum   float64
	)
	reader := bufio.NewReader(r)
	line, err := reader.ReadBytes(s) // avoid header
	for {
		line, err = reader.ReadBytes(s)
		if err != nil {
			break
		}
		strTip := strings.Split(string(line), ",")[p]
		partialTip, err := strconv.ParseFloat(strTip, 64)
		if err != nil {
			log.Fatalf("can not convert to float")
			return 0, 0.0
		}
		sum += partialTip
		count++
	}
	return count, sum
}

// StreamSimple takes an URL and process their content by streaming
// count all lines and get average tip_amount (15th field of csv)
func StreamSimple(URL string) {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	count, sum := sumFloatField(resp.Body, '\n', 15)

	fmt.Printf("\nTotal lines: %d", count)
	fmt.Printf("\nAverage Tip amount %0.4f\n", sum/float64(count))
	// Total lines: 10906858
	// Average Tip amount 1.7507
	// 40.78s user 31.88s system 4% cpu 26:18.64 total}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

var n = flag.Int("n", 3, "print strings longer than n x stddev")

func main() {
	flag.Parse()

	var stats Stats
	var ss []string

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		ss = append(ss, s.Text())
		l := len(s.Text())
		stats.cnt++
		stats.sum += l
		stats.sumsq += l * l
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}

	stddev := stats.stddev()

	for _, s := range ss {
		if len(s) > *n*int(stddev) {
			fmt.Printf("%d\t%s\n", len(s), s)
		}
	}
}

// Stats holds statistics about string lengths.
type Stats struct {
	cnt   int
	sum   int
	sumsq int
}

// https://rosettacode.org/wiki/Statistics/Basic#Go
func (s Stats) stddev() float64 {
	mean := s.sum / s.cnt
	return math.Sqrt(float64(s.sumsq/s.cnt - mean*mean))
}

package main

import (
	"fmt"
	"github.com/matchy109/ballclock/ballclock"
	//"github.com/davecheney/profile"
	"os"
	"strconv"
)

func main() {
	/*
			cfg := profile.Config{
				MemProfile:   true,
				BlockProfile: true,
				CPUProfile:   true,
				ProfilePath:  "./",
			}
		defer profile.Start(&cfg).Stop()
	*/

	var ball_cnt, iterations int

	switch len(os.Args) {
	case 2:
		ball_cnt, _ = strconv.Atoi(os.Args[1])
		iterations = -1
	case 3:
		ball_cnt, _ = strconv.Atoi(os.Args[1])
		iterations, _ = strconv.Atoi(os.Args[2])
	default:
		fmt.Printf("Number of Arg should be 2 or 3\n")
		os.Exit(1)
	}

	blck := ballclock.New(ball_cnt, iterations)
	fmt.Printf("%v balls cycle after %v days.\n", ball_cnt, blck.Run())
}

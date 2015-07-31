package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	MinCnt          = 27
	MaxCnt          = 127
	MinTraySize     = 4
	FiveMinTraySize = 11
	HourTraySize    = 11
)

type Tray []int8
type Trays struct {
	MinTray      Tray
	FiveMinTray  Tray
	HourTray     Tray
	MainTray     Tray
	OriginalTray Tray
}

func main() {

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
		os.Exit(0)
	}

	if ball_cnt >= MinCnt && ball_cnt <= MaxCnt {
		days := RanBallClock(int8(ball_cnt), iterations)
		fmt.Printf("%v balls cycle after %v days.\n", ball_cnt, days)
	} else {
		fmt.Printf("Number of ball should be in the range %d to %d\n", MinCnt, MaxCnt)
	}
}

func RanBallClock(ball_cnt int8, iterations int) int {
	var trays *Trays = InitializeTray(ball_cnt)

	for minutes := 1; ; minutes++ {
		ball := trays.MainTray[0]
		trays.MainTray = append(trays.MainTray[:0], trays.MainTray[1:]...)
		AddMinTray(&ball, trays)

		if iterations == minutes {
			fmt.Printf("{\"Min\":%s,\"FiveMin\":%s,\"Hour\":%s,\"Main\":%s}\n",
				trays.MinTray.JoinComma(), trays.FiveMinTray.JoinComma(), trays.HourTray.JoinComma(), trays.MainTray.JoinComma())
		}
		if minutes%(60*24) != 0 {
			continue
		}
		if reflect.DeepEqual(trays.OriginalTray, trays.MainTray[:]) {
			return (minutes / 60 / 24)
		}
	}
}

func InitializeTray(ball_cnt int8) *Trays {
	min_tray := make([]int8, 0, MinTraySize)
	five_min_tray := make([]int8, 0, FiveMinTraySize)
	hour_tray := make([]int8, 0, HourTraySize)
	main_tray := make([]int8, ball_cnt)
	original_tray := make([]int8, (cap(main_tray)))

	for i := int8(1); i <= ball_cnt; i++ {
		main_tray[i-1] = i
	}
	copy(original_tray, main_tray)
	trays := &Trays{MinTray: min_tray, FiveMinTray: five_min_tray, HourTray: hour_tray, MainTray: main_tray, OriginalTray: original_tray}

	return trays
}

func (tray *Tray) JoinComma() string {
	str := fmt.Sprintf("%v", *tray)
	return strings.Replace(str, " ", ",", -1)
}

func ReturnBall(tray *Tray, main_tray *Tray) {
	for i := len(*tray) - 1; i >= 0; i-- {
		*main_tray = append(*main_tray, (*tray)[i])
	}
}

func AddMinTray(ball *int8, trays *Trays) {
	if len(trays.MinTray) < MinTraySize {
		trays.MinTray = append(trays.MinTray, *ball)
	} else {
		ReturnBall(&trays.MinTray, &trays.MainTray)
		AddFiveMinTray(ball, trays)
		trays.MinTray = trays.MinTray[:0]
	}
}

func AddFiveMinTray(ball *int8, trays *Trays) {
	if len(trays.FiveMinTray) < FiveMinTraySize {
		trays.FiveMinTray = append(trays.FiveMinTray, *ball)
	} else {
		ReturnBall(&trays.FiveMinTray, &trays.MainTray)
		AddHourTray(ball, trays)
		trays.FiveMinTray = trays.FiveMinTray[:0]
	}
}

func AddHourTray(ball *int8, trays *Trays) {
	if len(trays.HourTray) < HourTraySize {
		trays.HourTray = append(trays.HourTray, *ball)
	} else {
		ReturnBall(&trays.HourTray, &trays.MainTray)
		trays.HourTray = trays.HourTray[:0]
		trays.MainTray = append(trays.MainTray, *ball)
	}
}

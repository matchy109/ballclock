package ballclock

import (
	"fmt"
	"os"
	"strings"
)

const (
	MinBallCnt      = 27
	MaxBallCnt      = 127
	MinTraySize     = 4
	FiveMinTraySize = 11
	HourTraySize    = 11
	MaxTraySize     = MinTraySize + FiveMinTraySize + HourTraySize + MaxBallCnt
)

type Trays struct {
	ClockTray      [MaxTraySize]int8
	MinTrayCnt     int64
	FiveMinTrayCnt int64
	HourTrayCnt    int64
	MainTraySCnt   int64
	MainTrayECnt   int64
	BallCnt        int8
	Iterations     int
}

func New(ball_cnt int, iterations int) *Trays {

	if ball_cnt < MinBallCnt || ball_cnt > MaxBallCnt {
		fmt.Printf("Number of ball should be in the range 27 to 127\n")
		os.Exit(1)
	}

	var trays Trays
	trays.BallCnt = int8(ball_cnt)
	trays.FiveMinTrayCnt = MinTraySize
	trays.HourTrayCnt = MinTraySize + FiveMinTraySize
	trays.MainTraySCnt = MinTraySize + FiveMinTraySize + HourTraySize
	trays.MainTrayECnt = trays.MainTraySCnt
	trays.Iterations = iterations

	for i := int8(1); i <= trays.BallCnt; i++ {
		trays.ClockTray[trays.MainTrayECnt] = i
		trays.MainTrayECnt++
	}
	return &trays
}

func (trays *Trays) ShowSituation() {
	str := fmt.Sprintf("{\"Min\":%v \"FiveMin\":%v \"Hour\":%v ",
		trays.ClockTray[:trays.MinTrayCnt], trays.ClockTray[4:trays.FiveMinTrayCnt], trays.ClockTray[15:trays.HourTrayCnt])

	switch {
	case trays.MainTraySCnt <= trays.MainTrayECnt:
		str += fmt.Sprintf("\"Main\":%v}", trays.ClockTray[trays.MainTraySCnt:trays.MainTrayECnt])
	case trays.MainTraySCnt > trays.MainTrayECnt:
		main_tray := append(trays.ClockTray[trays.MainTraySCnt:MaxTraySize], trays.ClockTray[26:trays.MainTrayECnt]...)
		str += fmt.Sprintf("\"Main\":%v}", main_tray)
	}

	fmt.Printf("%s\n", strings.Replace(str, " ", ",", -1))
	return
}

func (trays *Trays) Reset(start, end int8) {
	for i := start; i >= end; i-- {
		if trays.MainTrayECnt >= MaxTraySize {
			trays.MainTrayECnt = 26
		}
		trays.ClockTray[trays.MainTrayECnt] = trays.ClockTray[i]
		trays.MainTrayECnt++
	}
}

func (trays *Trays) IsEnd() bool {
	cnt := trays.MainTraySCnt

	for i := int8(1); i <= trays.BallCnt; i++ {
		if cnt >= MaxTraySize {
			cnt = 26
		}
		if trays.ClockTray[cnt] != i {
			return false
		}
		cnt++
	}
	return true
}

func (trays *Trays) Run() int {
	var ball int8

	for minutes := 1; ; minutes++ {

		if trays.Iterations == minutes-1 {
			trays.ShowSituation()
		}

		if trays.MainTraySCnt >= MaxTraySize {
			trays.MainTraySCnt = 26
		}
		ball = trays.ClockTray[trays.MainTraySCnt]
		trays.MainTraySCnt++

		if trays.MinTrayCnt < 4 {
			trays.ClockTray[trays.MinTrayCnt] = ball
			trays.MinTrayCnt++
			continue
		}
		trays.Reset(3, 0)
		trays.MinTrayCnt = 0

		if trays.FiveMinTrayCnt < 15 {
			trays.ClockTray[trays.FiveMinTrayCnt] = ball
			trays.FiveMinTrayCnt++
			continue
		}
		trays.Reset(14, 4)
		trays.FiveMinTrayCnt = 4

		if trays.HourTrayCnt < 26 {
			trays.ClockTray[trays.HourTrayCnt] = ball
			trays.HourTrayCnt++
			continue
		}
		trays.Reset(25, 15)
		if trays.MainTrayECnt >= MaxTraySize {
			trays.MainTrayECnt = 26
		}
		trays.ClockTray[trays.MainTrayECnt] = ball
		trays.MainTrayECnt++
		trays.HourTrayCnt = 15

		if minutes%1440 != 0 {
			continue
		}

		if trays.IsEnd() {
			return (minutes / 60 / 24)
		}
	}
}

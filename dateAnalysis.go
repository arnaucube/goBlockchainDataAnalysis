package main

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func map24hours() map[int]int {
	h := make(map[int]int)
	for i := 0; i < 24; i++ {
		h[i] = 0
	}
	return h
}
func decomposeDate(blockTime int64) (int, int, int, int) {
	/*i, err := strconv.ParseInt(blockTime, 10, 64)
	if err != nil {
		panic(err)
	}*/
	i := blockTime
	year := time.Unix(i, 0).Year()
	month := time.Unix(i, 0).Month()
	day := time.Unix(i, 0).Day()
	hour := time.Unix(i, 0).Hour()
	return year, int(month), day, hour
}
func unixTimeToTime(blockTime int64) time.Time {
	return time.Unix(blockTime, 0)
}
func timeToDate(blockTime int64) string {
	stringTime := strconv.FormatInt(blockTime, 10)
	i, err := strconv.ParseInt(stringTime, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	//fmt.Println(tm)
	stringDate := tm.String()
	//fmt.Println(stringDate)
	return stringDate
}
func hourAnalysis(e EdgeModel, blockTime int64) {
	//fmt.Println(blockTime)
	/*date := timeToDate(blockTime)
	dateHour := strings.Split(date, " ")[1]
	hourString := strings.Split(dateHour, ":")[0]*/
	_, _, _, hour := decomposeDate(blockTime)

	hourCount := ChartCountModel{}
	err := hourCountCollection.Find(bson.M{"hour": hour}).One(&hourCount)
	if err != nil {
		//date not yet in DB
		var hourCount ChartCountModel
		hourCount.Elem = hour
		hourCount.Count = 1
		err = hourCountCollection.Insert(hourCount)
		check(err)
	} else {
		hourCount.Count++
		err = hourCountCollection.Update(bson.M{"hour": hour}, &hourCount)
		check(err)
	}
}
func dateAnalysis(e EdgeModel, blockTime int64) {
	fmt.Println(blockTime)
	date := timeToDate(blockTime)

	dateCount := DateCountModel{}
	err := dateCountCollection.Find(bson.M{"date": date}).One(&dateCount)
	if err != nil {
		//date not yet in DB
		var dateCount DateCountModel
		dateCount.Date = date
		dateCount.Time = blockTime
		dateCount.Count = 1
		err = dateCountCollection.Insert(dateCount)
		check(err)
	} else {
		dateCount.Count++
		err = dateCountCollection.Update(bson.M{"date": date}, &dateCount)
		check(err)
	}
}

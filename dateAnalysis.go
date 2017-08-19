package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

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
	date := timeToDate(blockTime)
	dateHour := strings.Split(date, " ")[1]
	hour := strings.Split(dateHour, ":")[0]

	hourCount := HourCountModel{}
	err := hourCountCollection.Find(bson.M{"hour": hour}).One(&hourCount)
	if err != nil {
		//date not yet in DB
		var hourCount HourCountModel
		hourCount.Hour = hour
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

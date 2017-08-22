package main

import "gopkg.in/mgo.v2/bson"

func getStats() StatsModel {
	var stats StatsModel
	err := statsCollection.Find(bson.M{"title": "stats"}).One(&stats)
	check(err)
	stats.Title = "stats"
	return stats
}

func updateStats(stats StatsModel) {
	var oldStats StatsModel
	err := statsCollection.Find(bson.M{"title": "stats"}).One(&oldStats)
	if err != nil {
		err = statsCollection.Insert(stats)
		check(err)
	} else {
		err = statsCollection.Update(bson.M{"title": "stats"}, &stats)
		check(err)
	}
}

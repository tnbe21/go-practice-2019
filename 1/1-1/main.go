package main

import (
	"math/rand"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1).Format("2006/01/02")
	weathers := [...]string{"晴れ", "雨", "曇り", "雪"}
	rand.Seed(now.UnixNano())
	weather := weathers[rand.Intn(len(weathers))]
	rainyPercent := rand.Intn(101)
	println(tomorrow, weather, strconv.Itoa(rainyPercent)+"%")
}

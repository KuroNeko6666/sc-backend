package helper

import (
	"strconv"
	"time"
)

func TimeHandler(dateType string) []time.Time {
	var data []time.Time
	var reverse []time.Time
	now := time.Now()
	curYear, curMonth, curDay := now.Date()
	curHour, curMinute, curSecond := now.Clock()
	currentLocation := now.Location()

	for i := 0; i < 7; i++ {
		var currentDate time.Time
		switch dateType {
		case "second":
			currentDate = time.Date(curYear, curMonth, curDay, curHour, curMinute, (curSecond - i), 0, currentLocation)
		case "minute":
			currentDate = time.Date(curYear, curMonth, curDay, curHour, (curMinute - i), 0, 0, currentLocation)
		case "hour":
			currentDate = time.Date(curYear, curMonth, curDay, (curHour - i), 0, 0, 0, currentLocation)
		case "day":
			currentDate = time.Date(curYear, curMonth, (curDay - i), 0, 0, 0, 0, currentLocation)
		case "month":
			currentDate = time.Date(curYear, curMonth, 1, 0, 0, 0, 0, currentLocation).AddDate(0, (i * -1), 0)
		case "year":
			currentDate = time.Date((curYear - i), 1, 1, 0, 0, 0, 0, currentLocation)
		}
		data = append(data, currentDate)
	}

	for i := len(data); i > 0; i-- {
		reverse = append(reverse, data[i])
	}

	return reverse
}

func RangeTimeHandler(date []time.Time, dateType string) []time.Time {
	var rangeTime []time.Time
	for _, v := range date {
		switch dateType {
		case "second":
			rangeTime = append(rangeTime, v.Add(time.Second))
		case "minute":
			rangeTime = append(rangeTime, v.Add(time.Minute))
		case "hour":
			rangeTime = append(rangeTime, v.Add(time.Hour))
		case "day":
			rangeTime = append(rangeTime, v.Add(time.Hour*24))
		case "month":
			rangeTime = append(rangeTime, v.AddDate(0, 1, -1))
		case "year":
			rangeTime = append(rangeTime, v.AddDate(1, 0, -1))
		}
	}

	return rangeTime
}

func LabelHandler(date []time.Time, dateType string) []string {
	var label []string
	for _, v := range date {
		switch dateType {
		case "second":
			label = append(label, strconv.Itoa(v.Hour())+":"+strconv.Itoa(v.Minute())+":"+strconv.Itoa(v.Second()))
		case "minute":
			label = append(label, strconv.Itoa(v.Hour())+":"+strconv.Itoa(v.Minute()))
		case "hour":
			label = append(label, strconv.Itoa(v.Hour())+":"+strconv.Itoa(v.Minute()))
		case "day":
			label = append(label, v.Weekday().String())
		case "month":
			label = append(label, v.Month().String())
		case "year":
			label = append(label, strconv.Itoa(v.Year()))
		default:
			label = append(label, v.Weekday().String())
		}
	}
	return label

}

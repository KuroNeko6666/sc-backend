package helper

import (
	"errors"
	"strconv"
	"time"
)

func SingleTimeHandler(dateType string) []time.Time {
	var data []time.Time
	now := time.Now()
	curYear, curMonth, curDay := now.Date()
	curHour, curMinute, curSecond := now.Clock()
	currentLocation := now.Location()

	switch dateType {
	case "second":
		data = append(data, time.Date(curYear, curMonth, curDay, curHour, curMinute, (curSecond), 0, currentLocation))
		data = append(data, data[0].Add(time.Millisecond*999))
	case "minute":
		data = append(data, time.Date(curYear, curMonth, curDay, curHour, (curMinute), 0, 0, currentLocation))
		data = append(data, data[0].Add(time.Second*59))
	case "hour":
		data = append(data, time.Date(curYear, curMonth, curDay, (curHour), 0, 0, 0, currentLocation))
		data = append(data, data[0].Add(time.Minute*59))
	case "day":
		data = append(data, time.Date(curYear, curMonth, (curDay), 0, 0, 0, 0, currentLocation))
		data = append(data, data[0].Add(time.Hour*24))
	case "month":
		data = append(data, time.Date(curYear, curMonth, 1, 0, 0, 0, 0, currentLocation).AddDate(0, 0, 0))
		data = append(data, data[0].AddDate(0, 1, -1))
	case "year":
		data = append(data, time.Date((curYear), 1, 1, 0, 0, 0, 0, currentLocation))
		data = append(data, data[0].AddDate(1, 0, -1))
	}

	return data

}

func TimeHandler(dateType string) []time.Time {
	var data []time.Time
	// var reverse []time.Time
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

	return data
}

func RangeTimeHandler(date []time.Time, dateType string) []time.Time {
	var rangeTime []time.Time
	for _, v := range date {
		switch dateType {
		case "second":
			rangeTime = append(rangeTime, v.Add(time.Millisecond*999))
		case "minute":
			rangeTime = append(rangeTime, v.Add(time.Second*59))
		case "hour":
			rangeTime = append(rangeTime, v.Add(time.Minute*59))
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

func DateTypeValidate(dateType string) error {
	switch dateType {
	case "second":
		return nil
	case "minute":
		return nil
	case "hour":
		return nil
	case "day":
		return nil
	case "month":
		return nil
	case "year":
		return nil
	default:
		return errors.New("invalid date type")
	}
}

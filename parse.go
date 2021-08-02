package daygo

import (
	"regexp"
)

var templateRe = regexp.MustCompile(`^(\d{4})[-\/]?(\d{1,2})?[-\/]?(\d{0,2})[Tt\s]*(\d{1,2})?:?(\d{1,2})?:?(\d{1,2})?[.:]?(\d+)?$`)
var formatRe = regexp.MustCompile(`/\[([^\]]+)]|Y{1,4}|M{1,4}|D{1,2}|d{1,2}|H{1,2}|h{1,2}|m{1,2}|s{1,2}|Z{1,2}|SSS/g`)

// var duraFormatRe = regexp.MustCompile(`/\[([^\]]+)]|Y|M|D|H|m|s/g`)

func parseT(t string) []string {
	ret := templateRe.FindStringSubmatch(t)

	if len(ret) <= 0 {
		return nil
	}

	return ret
}

func parseList(list []int) (year, month, day, hour, minute, second int) {
	l := len(list)
	if l > 0 {
		year = list[0]
	}

	if l > 1 {
		month = list[1]
	} else {
		month = 1
	}

	if l > 2 {
		day = list[2]
	} else {
		day = 1
	}

	if l > 3 {
		hour = list[3]
	}
	if l > 4 {
		minute = list[4]
	}
	if l > 5 {
		second = list[5]
	}

	return year, month, day, hour, minute, second
}

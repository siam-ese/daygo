package day

import (
	"testing"
)

func TestTemplateRe(t *testing.T) {
	failStr1 := `2021-08-02`
	str2 := `2021-08-02T05:53:12`
	str3 := `20210802T055312`
	f := ErrorF("templateRe")

	if !templateRe.MatchString(failStr1) {
		t.Error(f(failStr1))
	}
	if !templateRe.MatchString(str2) {
		t.Error(f(str2))
	}
	if !templateRe.MatchString(str3) {
		t.Error(f(str3))
	}
}

func TestParseT(t *testing.T) {
	f := ErrorF("ParseT")

	p1 := "2021-07-28 12:26:36"
	r1 := parseT(p1)
	if r1 == nil {
		t.Error(f(p1))
	}
	p2 := "2004-05-03T17:30:08"
	r2 := parseT(p2)
	if r2 == nil {
		t.Error(f(p2))
	}
}

func TestParseList(t *testing.T) {
	list := []int{2021, 7, 30, 9, 42, 56}
	y, month, d, h, m, s := parseList(list)
	f := ErrorF("ParseList")

	if y != 2021 ||
		month != 7 ||
		d != 30 ||
		h != 9 ||
		m != 42 ||
		s != 56 {
		t.Error(f(list))
	}

	list2 := []int{2021}
	y, month, d, h, m, s = parseList(list2)

	if y != 2021 ||
		month != 1 ||
		d != 1 ||
		h != 0 ||
		m != 0 ||
		s != 0 {
		t.Error(f(list2))
	}
}

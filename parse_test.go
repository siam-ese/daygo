package day

import (
	"fmt"
	"testing"
)

func TestParseT(t *testing.T) {
	p1 := "2021-07-28 12:26:36"
	r1 := parseT(p1)
	if r1 == nil {
		t.Error(fmt.Printf("parseT exception error 1: Parameters (%s)", p1))
	}
	p2 := "2004-05-03T17:30:08"
	r2 := parseT(p2)
	if r2 == nil {
		t.Error(fmt.Printf("parseT exception error 2: Parameters (%s)", p2))
	}
}

func TestParseList(t *testing.T) {
	list := []int{2021, 7, 30, 9, 42, 56}
	y, month, d, h, m, s := parseList(list)
	if y != 2021 {
		t.Error("parseList Year error")
	}
	if month != 7 {
		t.Error("parseList month error")
	}
	if d != 30 {
		t.Error("parseList day error")
	}
	if h != 9 {
		t.Error("parseList hour error")
	}
	if m != 42 {
		t.Error("parseList minute error")
	}
	if s != 56 {
		t.Error("parseList second error")
	}
}

package locale

import (
	"testing"
)

func TestZHTranslator(t *testing.T) {
	trans := Translator{
		WeekMap:  ZH_CN.WeekMap,
		MonthMap: ZH_CN.MonthMap,
	}
	if trans.WT(1) != "ζζδΈ" {
		t.Error("test failed, ZH_CN translator")
	}
}

func TestENTranslator(t *testing.T) {
	trans := Translator{
		WeekMap:  EN.WeekMap,
		MonthMap: EN.MonthMap,
	}
	if trans.WT(1) != "Monday" {
		t.Error("test failed, EN translator")
	}
}

package locale

type Translator struct {
	WeekMap  [7]string
	MonthMap [12]string
}

func (t *Translator) WT(key int) string {
	key = key - 1
	if key > len(t.WeekMap) {
		return ""
	}
	return t.WeekMap[key]
}

func (t *Translator) MT(key int) string {
	key = key - 1
	if key > len(t.MonthMap) {
		return ""
	}
	return t.MonthMap[key]
}

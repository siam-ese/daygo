package locale

type Translator struct {
	WeekMap  [7]string
	MonthMap [12]string
}

func (t *Translator) WT(key int) string {
	return t.WeekMap[key]
}

func (t *Translator) MT(key int) string {
	return t.MonthMap[key]
}

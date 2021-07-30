package locale

const (
	Monday    = "Monday"
	Tuesday   = "Tuesday"
	Wednesday = "Wednesday"
	Thursday  = "Thursday"
	Friday    = "Friday"
	Saturday  = "Saturday"
	Sunday    = "Sunday"
	January   = "January"
	February  = "February"
	March     = "March"
	April     = "April"
	May       = "May"
	June      = "June"
	July      = "July"
	August    = "August"
	September = "September"
	October   = "October"
	November  = "November"
	December  = "December"
)

type Translator struct {
	Map map[string]string
}

func (t *Translator) T(key string) string {
	return t.Map[key]
}

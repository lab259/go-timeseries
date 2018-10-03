package timesrs

import "time"

type Granularity int

const (
	GranularitySecond Granularity = iota
	GranularityMinute
	GranularityHour
	GranularityDay
	GranularityWeek
	GranularityMonth
	GranularityYear
)

func (g Granularity) String() string {
	switch g {
	case GranularitySecond:
		return "second"
	case GranularityMinute:
		return "minute"
	case GranularityHour:
		return "hour"
	case GranularityDay:
		return "day"
	case GranularityWeek:
		return "week"
	case GranularityMonth:
		return "month"
	}
	return ""
}

func (g Granularity) RelativeTo(t, relativeTo time.Time) int {
	switch g {
	case GranularitySecond:
		return int(relativeTo.Sub(t.Truncate(time.Second)).Seconds())
	case GranularityMinute:
		return int(relativeTo.Sub(t.Truncate(time.Minute)).Minutes())
	case GranularityHour:
		return int(relativeTo.Sub(t.Truncate(time.Hour)).Hours())
	case GranularityDay:
		return int(relativeTo.Sub(t.Truncate(time.Hour * 24)).Hours() / 24)
	case GranularityWeek:
		return int(relativeTo.Sub(t.Truncate(time.Hour * 24)).Hours() / 24 / 7)
	case GranularityMonth:
		return int(t.Month())
	}
	return -1
}

func (g Granularity) Truncate(t time.Time) (time.Time) {
	switch g {
	case GranularitySecond:
		return t.Truncate(time.Second)
	case GranularityMinute:
		return t.Truncate(time.Minute)
	case GranularityHour:
		return t.Truncate(time.Hour)
	case GranularityDay:
		return t.Truncate(time.Hour * 24)
	case GranularityWeek:
		c := t.Truncate(time.Hour * 24)
		return c.Add(-time.Duration(int64(time.Hour) * 24 * int64(c.Weekday())))
	case GranularityMonth:
		return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	}
	return time.Time{}
}

// GranularityComposite is the configuration added to the pipeline to be passed
// to the `Storage`.
//
// The `Storage` create the collection and records according with this
// configuration.
type GranularityComposite struct {
	Record     Granularity
	Collection Granularity
}

package timesrs

import "time"

// Event is a happening that will be aggregated
type Event interface {
	Name() string
	Data() interface{}
	Time() time.Time
}

type evt struct {
	name   string
	data   interface{}
	moment time.Time
}

// NewEvent will return an instance of a event.
func NewEvent(name string, data interface{}, moment time.Time) *evt {
	return &evt{
		name:   name,
		data:   data,
		moment: moment,
	}
}

// Name returns the name of the event.
func (e *evt) Name() string {
	return e.name
}

// Data returns the payload stored into the event.
func (e *evt) Data() interface{} {
	return e.data
}

// Time returns the moment the event happened.
func (e *evt) Time() time.Time {
	return e.moment
}

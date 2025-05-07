package event

type Event struct {
	name    string
	payload string
}

func NewEvent(name string, payload string) Event {
	return Event{
		name:    name,
		payload: payload,
	}
}

func (e Event) Name() string {
	return e.name
}

func (e Event) Payload() string {
	return e.payload
}

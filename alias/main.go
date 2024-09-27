package alias

import "time"

// Type + kind
type MyCustomType string

func (s MyCustomType) IsDate() (time.Time, bool) {
	t, err := time.Parse(time.DateOnly, string(s))

	return t, err == nil
}

type Action func() error

func (a Action) DoPanic() {
	if err := a(); err != nil {
		panic(err)
	}
}
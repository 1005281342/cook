package cook

import "time"

type Food interface {
	Name() string
}

type Tool interface {
	Name() string
	Slice(Food)
	Waggle()
	Handle(Food, time.Duration)
}

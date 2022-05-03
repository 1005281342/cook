package cook

import (
	"time"

	"github.com/wuqinqiang/easyfsm"
)

type SliceParam struct {
	Tool
	Food
}

type MakeFoodParam struct {
	Tool
	Food
	Time time.Duration
}

type AddSeasoningParam struct {
	Seasoning
	Food
	Amount int // 单位ml
}

type WaggleParam struct {
	Tool
}

type EventData struct {
	Name easyfsm.EventName
	Data interface{}
}

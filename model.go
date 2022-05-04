package cook

import (
	"time"

	"github.com/wuqinqiang/easyfsm"
)

type FoodParam struct {
	Food
}

type ToolParam struct {
	Tool
	Time time.Duration
}

type SeasoningParam struct {
	Seasoning
	Amount int // 单位ml
}

type NotifyParam struct {
	Receiver string
	Message  string
}

type EventParam struct {
	Name easyfsm.EventName
	*easyfsm.FSM
}

type Param struct {
	FoodParam
	ToolParam
	SeasoningParam
	NotifyParam
	EventParam
}

type EventData struct {
	Name easyfsm.EventName
	Data interface{}
}

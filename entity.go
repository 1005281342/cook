package cook

import (
	"fmt"
	"log"
	"time"

	"github.com/wuqinqiang/easyfsm"
)

const (
	RinseEventName        easyfsm.EventName = "rinse"
	SliceEventName        easyfsm.EventName = "slice"
	MakeFoodEventName     easyfsm.EventName = "make-food"
	AddSeasoningEventName easyfsm.EventName = "add-seasoning"
	WaggleEventName       easyfsm.EventName = "waggle"
)

var rinseEntity = easyfsm.NewEventEntity(
	RinseEventName,
	func(opt *easyfsm.Param) (easyfsm.State, error) {
		param, ok := opt.Data.(Param)
		if !ok {
			return StateFailed, fmt.Errorf("%+v不是`Param`类型", opt.Data)
		}
		log.Printf("食物:%s 清洗中\n", param.FoodParam.Food.Name())
		time.Sleep(2 * time.Second) // TODO: 测试值，后续根据实际情况调整
		log.Printf("食物:%s 清洗完成\n", param.FoodParam.Food.Name())
		return StateRinsed, nil
	},
)

var sliceEntity = easyfsm.NewEventEntity(
	SliceEventName,
	func(opt *easyfsm.Param) (easyfsm.State, error) {
		param, ok := opt.Data.(Param)
		if !ok {
			return StateFailed, fmt.Errorf("%+v不是`Param`类型", opt.Data)
		}
		param.Tool.Slice(param.FoodParam.Food)
		return StateSliced, nil
	},
)

var makeFoodEntity = easyfsm.NewEventEntity(
	MakeFoodEventName,
	func(opt *easyfsm.Param) (easyfsm.State, error) {
		param, ok := opt.Data.(Param)
		if !ok {
			return StateFailed, fmt.Errorf("%+v不是`Param`类型", opt.Data)
		}
		if param.ToolParam.Time <= 0 {
			return StateFailed, fmt.Errorf("制作时长不符合要求")
		}
		param.ToolParam.Tool.Handle(param.FoodParam.Food, param.ToolParam.Time)
		time.Sleep(param.Time)
		return StateMade, nil
	}, easyfsm.WithObservers(Notify{}),
)

var addSeasoningEntity = easyfsm.NewEventEntity(
	AddSeasoningEventName,
	func(opt *easyfsm.Param) (easyfsm.State, error) {
		param, ok := opt.Data.(Param)
		if !ok {
			return StateFailed, fmt.Errorf("%+v不是`Param`类型", opt.Data)
		}
		log.Printf("给 %s 加 %dml %s", param.FoodParam.Food.Name(), param.SeasoningParam.Amount, param.SeasoningParam.Seasoning)
		return StateMade, nil
	},
)

var waggleEntity = easyfsm.NewEventEntity(
	WaggleEventName,
	func(opt *easyfsm.Param) (easyfsm.State, error) {
		param, ok := opt.Data.(Param)
		if !ok {
			return StateFailed, fmt.Errorf("%+v不是`Param`类型", opt.Data)
		}
		param.ToolParam.Tool.Waggle()
		return StateMade, nil
	},
)

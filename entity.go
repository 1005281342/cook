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
		param, ok := opt.Data.(Food)
		if !ok {
			return StateFailed, fmt.Errorf("%+v不是`Food`类型", opt.Data)
		}
		log.Printf("食物:%s 清洗中\n", param.Name())
		time.Sleep(2 * time.Second) // TODO: 测试值，后续根据实际情况调整
		log.Printf("食物:%s 清洗完成\n", param.Name())
		return StateRinsed, nil
	},
)

var sliceEntity = easyfsm.NewEventEntity(
	SliceEventName,
	func(opt *easyfsm.Param) (easyfsm.State, error) {
		param, ok := opt.Data.(SliceParam)
		if !ok {
			return StateFailed, fmt.Errorf("%+v不是`SliceParam`类型", opt.Data)
		}
		param.Tool.Slice(param.Food)
		time.Sleep(2 * time.Second) // TODO: 测试值，后续根据实际情况调整
		return StateSliced, nil
	},
)

var makeFoodEntity = easyfsm.NewEventEntity(
	MakeFoodEventName,
	func(opt *easyfsm.Param) (easyfsm.State, error) {
		param, ok := opt.Data.(MakeFoodParam)
		if !ok {
			return StateFailed, fmt.Errorf("%+v不是`MakeFoodParam`类型", opt.Data)
		}
		if param.Time <= 0 {
			return StateFailed, fmt.Errorf("制作时长不符合要求")
		}
		param.Tool.Handle(param.Food, param.Time)
		time.Sleep(param.Time)
		return StateMade, nil
	},
)

var addSeasoningEntity = easyfsm.NewEventEntity(
	AddSeasoningEventName,
	func(opt *easyfsm.Param) (easyfsm.State, error) {
		param, ok := opt.Data.(AddSeasoningParam)
		if !ok {
			return StateFailed, fmt.Errorf("%+v不是`AddSeasoningParam`类型", opt.Data)
		}
		log.Printf("给 %s 加 %dml %s", param.Food.Name(), param.Amount, param.Seasoning)
		return StateMade, nil
	},
)

var waggleEntity = easyfsm.NewEventEntity(
	WaggleEventName,
	func(opt *easyfsm.Param) (easyfsm.State, error) {
		param, ok := opt.Data.(WaggleParam)

		if !ok {
			return StateFailed, fmt.Errorf("%+v不是`AddSeasoningParam`类型", opt.Data)
		}
		param.Tool.Waggle()
		return StateMade, nil
	},
)

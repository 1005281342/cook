package cook

import (
	"fmt"
	"time"

	"github.com/wuqinqiang/easyfsm"
)

const (
	StateInit    easyfsm.State = 0
	StateCooking easyfsm.State = 1
	StateDone    easyfsm.State = 2
	StateFailed  easyfsm.State = 3
	StateRinsed  easyfsm.State = 4
	StateSliced  easyfsm.State = 5
	StateMade    easyfsm.State = 6
)

const (
	BizMakeFries easyfsm.BusinessName = "make-fries"
)

func init() {
	easyfsm.RegisterStateMachine(BizMakeFries, StateInit, rinseEntity)
	easyfsm.RegisterStateMachine(BizMakeFries, StateRinsed, sliceEntity)
	easyfsm.RegisterStateMachine(BizMakeFries, StateSliced, makeFoodEntity, addSeasoningEntity, waggleEntity)
	easyfsm.RegisterStateMachine(BizMakeFries, StateMade, makeFoodEntity, addSeasoningEntity, waggleEntity)
}

func MakeFries() error {

	const r = "jesonouyang"

	var (
		fsm   = easyfsm.NewFSM(BizMakeFries, StateInit)
		state easyfsm.State
		err   error
		k     = &AKnife{}
		af    = &AAirFryer{}
		p     = &APotato{}
	)

	for i, ed := range []EventData{
		// 1. 土豆清洗
		{RinseEventName, Param{
			FoodParam:  FoodParam{Food: p},
			EventParam: EventParam{Name: RinseEventName, FSM: fsm},
		}},
		// 2. 土豆切片
		{SliceEventName, Param{
			FoodParam:  FoodParam{Food: p},
			ToolParam:  ToolParam{Tool: k},
			EventParam: EventParam{Name: SliceEventName, FSM: fsm},
		}},
		// 3. 炸锅预热5分钟(TODO：众所周知，炸锅预热和1、2步是可以异步执行的)
		{MakeFoodEventName, Param{
			FoodParam:   FoodParam{Food: &AFoodNil{}},
			ToolParam:   ToolParam{Tool: af, Time: 5 * time.Minute},
			EventParam:  EventParam{Name: MakeFoodEventName, FSM: fsm},
			NotifyParam: NotifyParam{Receiver: r, Message: "炸锅预热完成"},
		}},
		// 4. 炸10min
		{MakeFoodEventName, Param{
			FoodParam:   FoodParam{Food: p},
			ToolParam:   ToolParam{Tool: af, Time: 10 * time.Minute},
			EventParam:  EventParam{Name: MakeFoodEventName, FSM: fsm},
			NotifyParam: NotifyParam{Receiver: r, Message: "1-薯条已炸10分钟"},
		}},
		// 5. 加油 20-30ml
		{AddSeasoningEventName, Param{
			FoodParam:      FoodParam{Food: p},
			SeasoningParam: SeasoningParam{Seasoning: Oil, Amount: 30},
			EventParam:     EventParam{Name: AddSeasoningEventName, FSM: fsm},
		}},
		// 6. 晃动锅体以使薯条受热均匀
		{WaggleEventName, Param{
			ToolParam:  ToolParam{Tool: af},
			EventParam: EventParam{Name: WaggleEventName, FSM: fsm},
		}},
		// 7. 继续炸10min
		{MakeFoodEventName, Param{
			FoodParam:   FoodParam{Food: p},
			ToolParam:   ToolParam{Tool: af, Time: 10 * time.Minute},
			EventParam:  EventParam{Name: MakeFoodEventName, FSM: fsm},
			NotifyParam: NotifyParam{Receiver: r, Message: "2-薯条已炸10分钟"},
		}},
	} {
		state, err = fsm.Call(ed.Name, easyfsm.WithData(ed.Data))
		if err != nil {
			return err
		}
		if state == StateFailed {
			return fmt.Errorf("执行到第%d步失败，事件：%s，参数: %+v", i, ed.Name, ed.Data)
		}
	}
	return nil
}

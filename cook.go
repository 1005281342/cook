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
		{RinseEventName, p},
		// 2. 土豆切片
		{SliceEventName, SliceParam{k, p}},
		// 3. 炸锅预热5分钟(TODO：众所周知，炸锅预热和1、2步是可以异步执行的)
		{MakeFoodEventName, MakeFoodParam{af, &AFoodNil{}, 5 * time.Minute}},
		// 4. 炸10min
		{MakeFoodEventName, MakeFoodParam{af, p, 10 * time.Minute}},
		// 5. 加油 20-30ml
		{AddSeasoningEventName, AddSeasoningParam{Oil, p, 30}},
		// 6. 晃动锅体以使薯条受热均匀
		{WaggleEventName, WaggleParam{af}},
		// 7. 继续炸10min
		{MakeFoodEventName, MakeFoodParam{af, p, 10 * time.Minute}},
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

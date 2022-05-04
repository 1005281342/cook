package cook

import (
	"log"

	"github.com/wuqinqiang/easyfsm"
)

type Notify struct{}

func (o Notify) Receive(opt *easyfsm.Param) {
	param, ok := opt.Data.(Param)
	if !ok {
		return
	}
	if param.NotifyParam.Message == "" || param.NotifyParam.Receiver == "" {
		return
	}
	log.Printf("发送消息给：%s, 内容是：%s\n", param.NotifyParam.Receiver, param.NotifyParam.Message)
}

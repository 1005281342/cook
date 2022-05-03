package main

import (
	"log"

	"github.com/1005281342/cook"
)

func main() {
	var err = cook.MakeFries()
	if err != nil {
		panic(err)
	}
	log.Println("薯条制作完成")
	//OUT:
	//2022/05/03 23:23:33 食物:potato 清洗中
	//2022/05/03 23:23:35 食物:potato 清洗完成
	//INFO eventName:=rinse beforeState:=0 afterState:=4
	//2022/05/03 23:23:35 用刀将potato切片
	//INFO eventName:=slice beforeState:=4 afterState:=5
	//2022/05/03 23:23:37 正在使用炸锅加工nil，大概持续300秒
	//INFO eventName:=make-food beforeState:=5 afterState:=6
	//2022/05/03 23:28:37 正在使用炸锅加工potato，大概持续600秒
	//INFO eventName:=make-food beforeState:=6 afterState:=6
	//2022/05/03 23:38:37 给 potato 加 30ml oil
	//INFO eventName:=add-seasoning beforeState:=6 afterState:=6
	//2022/05/03 23:38:37 晃动炸锅
	//INFO eventName:=waggle beforeState:=6 afterState:=6
	//2022/05/03 23:38:37 正在使用炸锅加工potato，大概持续600秒
	//INFO eventName:=make-food beforeState:=6 afterState:=6
	//2022/05/03 23:48:37 薯条制作完成
}

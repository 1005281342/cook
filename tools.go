package cook

import (
	"log"
	"time"
)

type ToolName string

const (
	AirFryer ToolName = "air-fryer" // 空气炸锅
	Knife    ToolName = "knife"     // 刀
)

type ABaseTool struct{}

func (a *ABaseTool) Name() string {
	return ""
}

func (a *ABaseTool) Slice(Food) {
}

func (a *ABaseTool) Waggle() {
}

func (a *ABaseTool) Handle(Food, time.Duration) {
}

type AKnife struct {
	ABaseTool
}

func (a *AKnife) Name() string {
	return string(Knife)
}

func (a *AKnife) Slice(food Food) {
	log.Printf("用刀将%s切片\n", food.Name())
}

type AAirFryer struct {
	ABaseTool
}

func (a *AAirFryer) Name() string {
	return string(AirFryer)
}

func (a *AAirFryer) Waggle() {
	log.Printf("晃动炸锅\n")
}

func (a *AAirFryer) Handle(food Food, t time.Duration) {
	log.Printf("正在使用炸锅加工%s，大概持续%d秒\n", food.Name(), t/time.Second)
}

package main

import (
	"fmt"
	"strconv"

	"github.com/jaeg/simple-ecs/component"
	"github.com/jaeg/simple-ecs/entity"
)

type PositionComponent struct {
	x, y int
}

func (pc PositionComponent) GetType() string {
	return "PositionComponent"
}

func (pc *PositionComponent) SetPosition(x int, y int) {
	pc.x = x
	pc.y = y
}

func AddPositionComponent(params []string) (component.Component, error) {
	pc := &PositionComponent{}
	x, err := strconv.Atoi(params[0])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(params[1])
	pc.x = x
	pc.y = y

	return pc, err
}

type AppearanceComponent struct {
	x, y int
}

func (pc AppearanceComponent) GetType() string {
	return "AppearanceComponent"
}

func AddAppearanceComponent(params []string) (component.Component, error) {
	ac := &AppearanceComponent{}
	x, err := strconv.Atoi(params[0])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(params[1])
	ac.x = x
	ac.y = y

	return ac, err
}

func main() {
	entity.RegisterComponentAddFunction("PositionComponent", AddPositionComponent)
	entity.RegisterComponentAddFunction("AppearanceComponent", AddAppearanceComponent)
	entity.FactoryLoad("entities.blueprints")
	e, err := entity.Create("tree")
	fmt.Println("Entity:", e, err)

	pc := e.GetComponent("PositionComponent").(*PositionComponent)
	fmt.Println("Position Component:", pc)

	ac := e.GetComponent("AppearanceComponent").(*AppearanceComponent)
	fmt.Println("Appearance Component:", ac)
}

package entity

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/jaeg/simple-ecs/component"
)

type ComponentAddFunction func(params []string) (component.Component, error)

var blueprints = make(map[string][]string)
var componentAddFunctions = make(map[string]ComponentAddFunction)

// FactoryLoad Loads the blueprints for the factory to construct entities
func FactoryLoad(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	entityName := ""
	for scanner.Scan() {
		value := scanner.Text()
		if value == "" {
			entityName = ""
			continue
		}
		if entityName == "" {
			entityName = value
			continue
		} else {
			blueprints[entityName] = append(blueprints[entityName], value)
		}
	}
}

func RegisterComponentAddFunction(name string, function ComponentAddFunction) {
	componentAddFunctions[name] = function
}

func Create(name string) (*Entity, error) {
	blueprint := blueprints[name]
	if blueprint != nil {
		entity := Entity{}
		entity.Blueprint = name

		for _, value := range blueprint {
			c := strings.Split(value, ":")
			params := strings.Split(c[1], ",")
			if componentAddFunctions[c[0]] != nil {
				newComp, err := componentAddFunctions[c[0]](params)
				if err != nil {
					return nil, err
				}
				entity.AddComponent(newComp)
			}

		}
		return &entity, nil
	}
	return nil, errors.New("no blueprint found")
}

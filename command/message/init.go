package messageCommand

import (
	"WRwolf_bot-Go/util"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type commandModule struct {
	Name                     string
	Alias                    []string
	Category                 string
	Description              string
	Usage                    string
	Example                  string
	IsRequireAdminPermission bool
	CommandFunc              interface{}
}

var CommandArr = make([]*commandModule, 0)

func init() {
	files, err := os.ReadDir("command/message")
	util.HandleError(err, "Error reading guild events directory")

	reflectType := reflect.TypeOf(&commandModule{})
	for _, file := range files {
		if file.Name() == "init.go" {
			continue
		}
		methodName := strings.Split(strings.Split(file.Name(), ".")[0], "_")[1]
		fmt.Println(methodName + " detected")
		method, isExist := reflectType.MethodByName(methodName)
		if isExist {
			commandStruct := &commandModule{}
			method.Func.Interface().(func(commandStruct *commandModule))(commandStruct)
			CommandArr = append(CommandArr, commandStruct)
			fmt.Println(commandStruct.Name + " is registering")
		}
	}
}

package commandMode

import (
	"fmt"
	"testing"
)

func TestCommand_Execute(t *testing.T) {
	electricCooker := new(ElectricCooker)
	invoker := new(ElectricCookerInvoker)

	steam := NewSteamCommand(electricCooker)
	invoker.SetCookCommand(steam)
	fmt.Println(invoker.ExecuteCookCommand())

	cook := NewCookCommand(electricCooker)
	invoker.SetCookCommand(cook)
	fmt.Println(invoker.ExecuteCookCommand())

	shutDown := NewShutDownCommand(electricCooker)
	invoker.SetCookCommand(shutDown)
	fmt.Println(invoker.ExecuteCookCommand())
}

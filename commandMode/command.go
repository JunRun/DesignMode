package commandMode

import "fmt"

// ElectricCooker 电饭煲
type ElectricCooker struct {
	fire     string // 火力
	pressure string // 压力
}

func (e *ElectricCooker) SetFire(fire string) {
	e.fire = fire
}

func (e *ElectricCooker) SetPressure(pressure string) {
	e.pressure = pressure
}

func (e *ElectricCooker) Run(duration string) string {
	return fmt.Sprintf("电饭煲火力设置为%s,压力为%s,持续时间为%s\n", e.fire, e.pressure, duration)
}
func (e *ElectricCooker) ShutDown() string {
	return fmt.Sprintf("电饭煲停止运行")
}

type Command interface {
	Execute() string
}
type steamCommand struct {
	electricCooker *ElectricCooker
}

func NewSteamCommand(cooker *ElectricCooker) *steamCommand {
	return &steamCommand{electricCooker: cooker}
}
func (s *steamCommand) Execute() string {
	s.electricCooker.SetFire("中")
	s.electricCooker.SetPressure("中")
	return fmt.Sprintf("蒸饭" + s.electricCooker.Run("30 分钟"))
}

type cookCommand struct {
	electricCooker *ElectricCooker
}

func NewCookCommand(cooker *ElectricCooker) *cookCommand {
	return &cookCommand{electricCooker: cooker}
}
func (c cookCommand) Execute() string {
	c.electricCooker.SetFire("低")
	c.electricCooker.SetPressure("中")
	return fmt.Sprintf(" 煮粥 " + c.electricCooker.Run("40 分钟"))
}

type shutDownCommand struct {
	electricCooker *ElectricCooker
}

func NewShutDownCommand(cooker *ElectricCooker) *shutDownCommand {
	return &shutDownCommand{electricCooker: cooker}
}

func (s *shutDownCommand) Execute() string {
	return s.electricCooker.ShutDown()
}

// ElectricCookerInvoker 电饭煲指令触发器
type ElectricCookerInvoker struct {
	cookCommand Command
}

// SetCookCommand 设置指令
func (e *ElectricCookerInvoker) SetCookCommand(cookCommand Command) {
	e.cookCommand = cookCommand
}

// ExecuteCookCommand 执行指令
func (e *ElectricCookerInvoker) ExecuteCookCommand() string {
	return e.cookCommand.Execute()
}

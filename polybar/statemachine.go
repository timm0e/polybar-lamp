package polybar

import (
	"polybar-lamp/lampcontroller"
)

type State interface {
	Increment(lampController lampcontroller.LampController)
	Decrement(lampController lampcontroller.LampController)
	OnOff(lampController lampcontroller.LampController)
	SwitchMode(lampController lampcontroller.LampController)

	EnterState(lampController lampcontroller.LampController)

	// maybe add an update hook
}

const (
	BRIGHTNESS = iota
	TEMPERATURE
	OFF
)

var stateToImpl = map[int]State{
	BRIGHTNESS:  &BrightnessState{},
	TEMPERATURE: &TemperatureState{},
	OFF:         OffState{},
}

var CurrentState int

func getCurrentState() State {
	return stateToImpl[CurrentState]
}

func SwitchState(state int, lampController lampcontroller.LampController) {
	CurrentState = state
	EnterState(lampController)
}

func Increment(lampController lampcontroller.LampController) {
	getCurrentState().Increment(lampController)
}
func Decrement(lampController lampcontroller.LampController) {
	getCurrentState().Decrement(lampController)
}
func OnOff(lampController lampcontroller.LampController) {
	getCurrentState().OnOff(lampController)
}
func SwitchMode(lampController lampcontroller.LampController) {
	getCurrentState().SwitchMode(lampController)
}
func EnterState(lampController lampcontroller.LampController) {
	getCurrentState().EnterState(lampController)
}

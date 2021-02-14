package polybar

import (
	"polybar-lamp/lampcontroller"
)

type OffState struct{}

func (o OffState) Increment(lampController lampcontroller.LampController) {
}

func (o OffState) Decrement(lampController lampcontroller.LampController) {
}

func (o OffState) OnOff(lampController lampcontroller.LampController) {
	err := lampController.SetOnState(true)
	if err != nil {
		println(err)
	}
	SwitchState(BRIGHTNESS, lampController)
}

func (o OffState) SwitchMode(lampController lampcontroller.LampController) {
}

func (o OffState) EnterState(lampController lampcontroller.LampController) {
	Output("\uF011 Off")
}

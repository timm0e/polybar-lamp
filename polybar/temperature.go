package polybar

import (
	"fmt"
	"polybar-lamp/lampcontroller"
)

type TemperatureState struct {
	currentTemperature int
}

func outputTemperature(temperature int) {
	Output(fmt.Sprintf("\uF2C7 %d%%", temperature))
}

func (t *TemperatureState) Increment(lampController lampcontroller.LampController) {
	t.currentTemperature += 5 - t.currentTemperature % 5
	if t.currentTemperature > 100 {
		t.currentTemperature = 100
	}
	err := lampController.SetColorTemp(t.currentTemperature)
	if err != nil {
		println(err)
	}
	outputTemperature(t.currentTemperature)
}

func (t *TemperatureState) Decrement(lampController lampcontroller.LampController) {
	t.currentTemperature -= 5 + t.currentTemperature % 5
	if t.currentTemperature < 0 {
		t.currentTemperature = 0
	}
	err := lampController.SetColorTemp(t.currentTemperature)
	if err != nil {
		println(err)
	}
	outputTemperature(t.currentTemperature)
}

func (t TemperatureState) OnOff(lampController lampcontroller.LampController) {
	err := lampController.SetOnState(false)
	if err != nil {
		println(err)
	}
	SwitchState(OFF, lampController)
}

func (t TemperatureState) SwitchMode(lampController lampcontroller.LampController) {
	SwitchState(BRIGHTNESS, lampController)
}

func (t *TemperatureState) EnterState(lampController lampcontroller.LampController) {
	t.currentTemperature, _ = lampController.GetColorTemp()
	outputTemperature(t.currentTemperature)
}

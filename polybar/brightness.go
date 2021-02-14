package polybar

import (
	"fmt"
	"polybar-lamp/lampcontroller"
)

type BrightnessState struct {
	currentBrightness int
}

func outputBrightness(brightness int) {
	Output(fmt.Sprintf("\uF185 %d%%", brightness))
}

func (b *BrightnessState) Increment(lampController lampcontroller.LampController) {
	b.currentBrightness += 5 - b.currentBrightness % 5
	if b.currentBrightness > 100 {
		b.currentBrightness = 100
	}
	err := lampController.SetBrightness(b.currentBrightness)
	if err != nil {
		println(err)
	}
	outputBrightness(b.currentBrightness)
}

func (b *BrightnessState) Decrement(lampController lampcontroller.LampController) {
	b.currentBrightness -= 5 + b.currentBrightness % 5
	if b.currentBrightness < 5 {
		b.currentBrightness = 5
	}
	err := lampController.SetBrightness(b.currentBrightness)
	if err != nil {
		println(err)
	}
	outputBrightness(b.currentBrightness)
}

func (b *BrightnessState) OnOff(lampController lampcontroller.LampController) {
	err := lampController.SetOnState(false)
	if err != nil {
		println(err)
	}
	SwitchState(OFF, lampController)
}

func (b BrightnessState) SwitchMode(lampController lampcontroller.LampController) {
	SwitchState(TEMPERATURE, lampController)
}

func (b *BrightnessState) EnterState(lampController lampcontroller.LampController) {
	b.currentBrightness, _ = lampController.GetBrightness()

	outputBrightness(b.currentBrightness)
}

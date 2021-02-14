package lampcontroller

import (
	"errors"
	"polybar-lamp/config"
	"polybar-lamp/hass"
)

type LampController struct {
	EntityID string

	apiClient *hass.APIClient

	minColorTemp int
	maxColorTemp int
}

func NewLampController(lampEntityID string, apiClient *hass.APIClient, colorTempLimitOverride *config.ColorTempLimitOverride) (LampController, error) {
	lampController := LampController{EntityID: lampEntityID, apiClient: apiClient}

	var err error = nil

	if colorTempLimitOverride != nil {
		lampController.minColorTemp = colorTempLimitOverride.Min
		lampController.maxColorTemp = colorTempLimitOverride.Max
	} else {
		err = lampController.updateColorTempLimitsFromAPI()
	}

	return lampController, err
}

func (lc LampController) getState() (hass.LampState, error) {
	return lc.apiClient.GetState(lc.EntityID)
}

func (lc LampController) getStateWithAttributes() (hass.LampState, error) {
	lampState, err := lc.getState()

	if err != nil {
		return lampState, err
	}

	if lampState.Attributes == nil {
		return lampState, errors.New("response did not contain `attributes`")
	}

	return lampState, err
}

func (lc *LampController) updateColorTempLimitsFromAPI() error {
	lampState, err := lc.getStateWithAttributes()

	if err != nil {
		return err
	}

	lc.minColorTemp = lampState.Attributes.MinColorTemp
	lc.maxColorTemp = lampState.Attributes.MaxColorTemp

	return nil
}

func (lc LampController) GetOnState() (bool, error) {
	lampState, err := lc.getState()
	if err != nil {
		return false, err
	}

	return lampState.State == "on", nil
}

func (lc LampController) SetOnState(onState bool) error {
	if !onState {
		return lc.apiClient.TurnOff(lc.EntityID)
	}
	return lc.apiClient.TurnOn(lc.EntityID)
}

func (lc LampController) GetBrightness() (int, error) {
	lampState, err := lc.getStateWithAttributes()

	if err != nil {
		return -1, err
	}

	return lampState.Attributes.Brightness * 100 / 255, nil
}

func (lc LampController) SetBrightness(brightness int) error {
	if brightness < 0 || brightness > 100 {
		return errors.New("brightness needs to be in range [0, 100]")
	}
	return lc.apiClient.SetBrightness(lc.EntityID, brightness)
}

func (lc LampController) GetColorTemp() (int, error) {
	lampState, err := lc.getStateWithAttributes()

	if err != nil {
		return -1, err
	}

	colorTempPercent := (lampState.Attributes.ColorTemp - lc.minColorTemp) / ((lc.maxColorTemp - lc.minColorTemp) / 100)

	return colorTempPercent, nil
}

func (lc LampController) SetColorTemp(colorTempPercent int) error {
	if colorTempPercent < 0 || colorTempPercent > 100 {
		return errors.New("colorTempPercent needs to be in range [0, 100]")
	}

	colorTempAbs := lc.minColorTemp

	if colorTempPercent != 0 {
		colorTempAbs = ((lc.maxColorTemp - lc.minColorTemp) / 100) * colorTempPercent + lc.minColorTemp
	}

	return lc.apiClient.SetColorTemp(lc.EntityID, colorTempAbs)
}

package hass

type TurnOnData struct {
	EntityID          string `json:"entity_id"`
	ColorTemp         *int   `json:"color_temp,omitempty"`
	BrightnessPercent *int   `json:"brightness_pct,omitempty"`
}

type TurnOffData struct {
	EntityID string `json:"entity_id"`
}

type LampAttributes struct {
	Brightness   int `json:"brightness"`
	ColorTemp    int `json:"color_temp"`
	MinColorTemp int `json:"min_mireds"`
	MaxColorTemp int `json:"max_mireds"`
}

type LampState struct {
	EntityID   string          `json:"entity_id"`
	State      string          `json:"state"`
	Attributes *LampAttributes `json:"attributes,omitempty"`
}

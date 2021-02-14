# polybar-lamp
Control a light controlled by HomeAssistant directly from your polybar

> Note: This piece of software was thrown together in a day and is currently in a "works for me" state. See "Known Issues" to get an overview if it has any use for you in its current state.
> 
> Future improvements may follow though 

## Usage
![](https://imgur.com/Dsmi31g.gif)

**Left-Click**        to toggle the lamp **on/off**  
**Right-Click**       to switch between controlling **brightness or color-temperature**  
**Mouse-Wheel-Up**    to **increase** the selected value  
**Mouse-Wheel-Down**  to **decrease** the selected value

_if you use the default polybar configuration provided below_

## Setup
### Installation
Grab the latest release or build it using `go build`, and store it permanently in a place you like.

### Configuration of polybar-lamp
Run `./polybar-lamp makeconfig` to generate a blank config.yaml in the same folder as the executable.

Set `lampEntity` to the entity id of your lamp.

See https://developers.home-assistant.io/docs/api/rest/ on how to determine the `apiBaseUrl` and `apiKey` values.

#### Fixing a wrong color temperature range
For me the color temperature of my lamp can only be set from 28% - 87% in the HomeAssistant UI. If this also happens to you do the following:

1. Determine the minimum and maximum `color_temp` when dragging the slider to the possible minimum or maximum. You can find these values in the HomeAssistant Developer Tools
2. Add the following block to the config.yaml and fill in your respective values:
```yaml
colorTempLimitOverride:
  min: 250
  max: 454
```

### Configuration of polybar
Add a new module to polybar:

```
[module/lamp]
type = custom/script

exec = <path-to-polybar-lamp>/polybar-lamp server
tail = true

click-left = <path-to-polybar-lamp>/polybar-lamp sendcommand onoff
click-right = <path-to-polybar-lamp>/polybar-lamp sendcommand switchmode
scroll-up = <path-to-polybar-lamp>/polybar-lamp sendcommand increment
scroll-down = <path-to-polybar-lamp>/polybar-lamp sendcommand decrement
```

You can now add the `lamp` module to *one* of your bars.
_If you wish to rename the module just change `[module/lamp]` to `[module/whateveryoulike]` 

**Make sure you include the Fontawesome 5 icon font in your polybar config for the icons to be displayed correctly**

## Known Issues

- updates to the lamp state (e.g. turning off the lamp) outside the application do not trigger an update of the ui yet
- multiple instances are not supported at the moment 
    - path of the server socket is not configurable, which only allows one server instance at a time
    - multiple instances controlling the same light would not be able to talk to each other - which (in case the socket path was configurable) would not sync state changes across the instances
    - the config.yaml path is not configurable either, which (in case the socket path was configurable) would require multiple copies of the polybar-lamp executable to support multiple configurations
- there is no log output of errors yet

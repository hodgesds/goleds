// +build linux

package leds

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// LEDs returns the set of LEDs configured for the system.
func LEDs() ([]*LED, error) {
	files, err := ioutil.ReadDir(SysFSBase)
	if err != nil {
		return nil, err
	}

	leds := []*LED{}

	for _, f := range files {
		switch mode := f.Mode(); {
		case mode.IsDir(), mode&os.ModeSymlink != 0:
			leds = append(leds, &LED{
				basePath: SysFSBase + "/" + f.Name(),
				name:     f.Name(),
			})
		}
	}

	return leds, nil
}

// Trigger is an LED trigger.
type Trigger struct {
	Name   string
	Active bool
}

// LED represents an LED.
type LED struct {
	basePath string
	name     string
}

// Name returns the name of the LED.
func (l *LED) Name() string {
	return l.name
}

// SetBrightness is used to set the brightness of a LED device.
func (l *LED) SetBrightness(b int) error {
	max, err := l.MaxBrightness()
	if err != nil {
		return err
	}
	if b > max {
		return fmt.Errorf("can't set beyond max brightness: %d", max)
	}

	return ioutil.WriteFile(l.basePath+"/brightness", []byte(string(b)), 0644)
}

// SetTrigger is used to set the trigger of a LED device.
func (l *LED) SetTrigger(trigger string) error {
	return ioutil.WriteFile(l.basePath+"/trigger", []byte(trigger), 0644)
}

// Brightness returns the brightness setting.
func (l *LED) Brightness() (int, error) {
	b, err := ioutil.ReadFile(l.basePath + "/brightness")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.Replace(string(b), "\n", "", -1))
}

// MaxBrightness returns the max brightness.
func (l *LED) MaxBrightness() (int, error) {
	b, err := ioutil.ReadFile(l.basePath + "/max_brightness")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.Replace(string(b), "\n", "", -1))
}

// Triggers returns the set of triggers for the LED.
func (l *LED) Triggers() ([]*Trigger, error) {
	b, err := ioutil.ReadFile(l.basePath + "/trigger")
	if err != nil {
		return nil, err
	}

	triggerStrs := strings.Split(strings.Replace(string(b), "\n", "", -1), " ")
	triggers := make([]*Trigger, len(triggerStrs))

	for i, triggerStr := range triggerStrs {
		var trigger Trigger
		if strings.ContainsAny(triggerStr, "[]") {
			trigger.Active = true
		}
		trigger.Name = strings.Replace(strings.Replace(triggerStr, "[", "", -1), "]", "", -1)
		triggers[i] = &trigger
	}

	return triggers, nil
}

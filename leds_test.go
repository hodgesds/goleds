// +build linux

package leds

import "testing"

func TestLEDs(t *testing.T) {
	leds, err := LEDs()
	if err != nil {
		t.Fatal(err)
	}
	if len(leds) == 0 {
		t.Errorf("Expected leds to be >0")
	}

	led := leds[0]
	if _, err := led.Brightness(); err != nil {
		t.Fatal(err)
	}
	if _, err := led.MaxBrightness(); err != nil {
		t.Fatal(err)
	}
	if _, err := led.Triggers(); err != nil {
		t.Fatal(err)
	}
}

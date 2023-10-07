package main

import (
	"fmt"
	"testing"
)

func TestSetBG(t *testing.T) {
	type testContainer struct {
		f func(...interface{}) string
		v string
	}

	BGColors := map[string]testContainer{
		"black":   {f: BlackBG, v: blackbg},
		"red":     {f: RedBG, v: redbg},
		"green":   {f: GreenBG, v: greenbg},
		"yellow":  {f: YellowBG, v: yellowbg},
		"blue":    {f: BlueBG, v: bluebg},
		"magenta": {f: MagentaBG, v: magentabg},
		"cyan":    {f: CyanBG, v: cyanbg},
		"white":   {f: WhiteBG, v: whitebg},
		"256": {f: func(v ...interface{}) string {
			return Color256BG(42, v...)
		}, v: fmt.Sprintf(bg256, 42)},
		"true": {f: func(v ...interface{}) string {
			return ColorTrueBG(42, 420, 69, v...)
		}, v: fmt.Sprintf(bgtrue, 42, 420, 69)},
	}

	for name, data := range BGColors {
		t.Run(name, func(t *testing.T) {
			if data.f("teststring") != escape+data.v+set+"teststring"+escape+resetbg+set {
				t.Log("did not correctly build for BG color: ", name)
				t.Fail()
			}
			fmt.Println(data.f(" ", name, " "))

			nocolorIsSet = true
			t.Cleanup(func() { nocolorIsSet = false })

			if data.f("teststring") != "teststring" {
				t.Log("did not honor NO_COLOR for BG color: ", name)
				t.Fail()
			}
		})
	}
}

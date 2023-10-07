package main

import (
	"fmt"
	"testing"
)

func TestSetColor(t *testing.T) {
	type testContainer struct {
		f func(...interface{}) string
		v string
	}

	BGColors := map[string]testContainer{
		"black":   {f: Black, v: blackfg},
		"red":     {f: Red, v: redfg},
		"green":   {f: Green, v: greenfg},
		"yellow":  {f: Yellow, v: yellowfg},
		"blue":    {f: Blue, v: bluefg},
		"magenta": {f: Magenta, v: magentafg},
		"cyan":    {f: Cyan, v: cyanfg},
		"white":   {f: White, v: whitefg},
		"256": {f: func(v ...interface{}) string {
			return Color256(42, v...)
		}, v: fmt.Sprintf(fg256, 42)},
		"true": {f: func(v ...interface{}) string {
			return ColorTrue(42, 420, 69, v...)
		}, v: fmt.Sprintf(fgtrue, 42, 420, 69)},
	}

	for name, data := range BGColors {
		t.Run(name, func(t *testing.T) {
			if data.f("teststring") != escape+data.v+set+"teststring"+escape+resetfg+set {
				t.Log("did not correctly build for FG color: ", name)
				t.Fail()
			}
			fmt.Println(data.f(" ", name, " "))

			nocolorIsSet = true
			t.Cleanup(func() { nocolorIsSet = false })

			if data.f("teststring") != "teststring" {
				t.Log("did not honor NO_COLOR for FG color: ", name)
				t.Fail()
			}
		})
	}
}

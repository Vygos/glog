package glog

import "fmt"

func ApplyColor(input string, color Color) string {
	return fmt.Sprintf("%s%s%s", color, input, string(Reset))
}

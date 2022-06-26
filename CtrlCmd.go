package ctrl_cmd

import (
	"fmt"
	"os/exec"
)

// alternative fmt.Sprintf
var F = fmt.Sprintf

// custom symbols
var (
	Block     = "█"
	TriangleL = "▶"
	TriangleR = "◀"
	TriangleU = "▲"
	TriangleD = "▼"
)

// convert  user params in eq cmd args
func ColorMode(c string, mode string) (int, string) {
	rcolor := -1
	rmode := "ab"
	switch c {
	case "black":
		rcolor = 0
	case "red":
		rcolor = 1
	case "green":
		rcolor = 2
	case "yellow":
		rcolor = 3
	case "blue":
		rcolor = 4
	case "violet":
		rcolor = 5
	case "cyan":
		rcolor = 6
	case "white":
		rcolor = 7
	}

	switch mode {
	case "bg":
		rmode = "ab"
	case "fg":
		rmode = "af"
	}
	return rcolor, rmode
}

// print color command line
func LineColor(line string, color string, mode string) {
	c, m := ColorMode(color, mode)
	re := exec.Command("bash", "-c", F("echo `tput  set%v %v`\"%+v\"`tput sgr0`", m, c, line))
	res, _ := re.Output()
	fmt.Print(string(res))
}

// print format command line
func LineF(line string) {
	re := exec.Command("bash", "-c", F("%v", line))
	res, _ := re.Output()
	fmt.Print(string(res))
}

// print easy command line
func Line(line string) {
	re := exec.Command("bash", "-c", line)
	res, _ := re.Output()
	fmt.Print(string(res))
}

// get result command line
func LineGet(line string) string {
	re := exec.Command("bash", "-c", line)
	res, _ := re.Output()
	return string(res)
}

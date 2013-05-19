package term

func ColorPrintf(c string, f string, a ...interface{}) {
	// start escape \033[
	// end escape   m
	// Black       0;30     Dark Gray     1;30
	// Blue        0;34     Light Blue    1;34
	// Green       0;32     Light Green   1;32
	// Cyan        0;36     Light Cyan    1;36
	// Red         0;31     Light Red     1;31
	// Purple      0;35     Light Purple  1;35
	// Brown       0;33     Yellow        1;33
	// Light Gray  0;37     White         1;37
	var buf bytes.Buffer

	color := map[string]string{
		"black":  "0;30",
		"blue":   "0;34",
		"green":  "0;32",
		"cyan":   "0;36",
		"red":    "0;31",
		"purple": "0;35",
		"brown":  "0;33",
	}

	start := "\033["
	end := "m"

	// create color string
	buf.WriteString(start)
	buf.WriteString(color[c])
	buf.WriteString(end)
	buf.WriteString(f)

	fmt.Printf(buf.String(), a...)

	// reset color
	fmt.Printf("\033[0m")
}

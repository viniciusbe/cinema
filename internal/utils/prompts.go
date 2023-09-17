package utils

import (
	"bufio"
	"os"
	"strings"
)

func StringPrompt() string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

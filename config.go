package main

import (
	"fmt"
	"os"
	"strings"
)

func findConfigPath() string {
	return "trrr.ini"
}

func readConfig(configPath string, c *config) error {
	bytes, err := os.ReadFile(configPath)
	if err != nil { return err }
	for line := range strings.SplitSeq(string(bytes), "\n") {
		if strings.TrimSpace(line) == "" { continue }
		pair := strings.Split(line, "=")
		k, v := pair[0], pair[1]
		k = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		fmt.Printf("conf: %s=%s\n", k, v)

		switch k {
		case "source": c.source = v
		case "target": c.target = v
		case "backend": c.backend = v
		case "lingvaDomain": c.lingvaDomain = v
		default: return fmt.Errorf("unknown key: %s", k)
		}
	}
	// .......
	return nil
}

package lib

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// LoadEnv loads env values from file.
func LoadEnv(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	values := ParseEnv(f)
	return setEnv(values)
}

type env struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

// ParseEnv knows how to parse env values from reader.
func ParseEnv(reader io.Reader) []env {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var values []env

	for scanner.Scan() {
		fullLine := scanner.Text()
		fullLine = strings.TrimSpace(fullLine)

		if len(fullLine) == 0 {
			continue
		}
		if strings.HasPrefix(fullLine, "#") || strings.HasPrefix(fullLine, "=") {
			continue
		}

		parts := strings.SplitN(fullLine, "=", 2)
		kv := env{
			Key: strings.TrimSpace(parts[0]),
			Val: strings.TrimSpace(parts[1]),
		}

		values = append(values, kv)
	}

	return values
}

func setEnv(values []env) error {
	for _, v := range values {
		if err := os.Setenv(v.Key, v.Val); err != nil {
			return err
		}
	}

	return nil
}

package godotenvgenerator

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// Write serializes the given environment and writes it to a file
func Write(envMap map[string]string, filename string) error {
	content, err := Marshal(envMap)
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err2 := file.WriteString(content)
	return err2
}

// Marshal outputs the given environment as a dotenv-formatted environment file.
// Each line is in the format: KEY="VALUE" where VALUE is backslash-escaped.
func Marshal(envMap map[string]string) (string, error) {
	lines := make([]string, 0, len(envMap))
	for k, v := range envMap {
		lines = append(lines, fmt.Sprintf(`%s=%s`, k, v))
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n"), nil
}

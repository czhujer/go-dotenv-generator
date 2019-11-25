package godotenvgenerator

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

const doubleQuoteSpecialChars = "\\\n\r\"!$`"

// Write serializes the given environment and writes it to a file
func Write(envMap map[string]string, filename string) error {
	content, error := Marshal(envMap)
	if error != nil {
		return error
	}
	file, error := os.Create(filename)
	if error != nil {
		return error
	}
	_, err := file.WriteString(content)
	return err
}

// Marshal outputs the given environment as a dotenv-formatted environment file.
// Each line is in the format: KEY="VALUE" where VALUE is backslash-escaped.
func Marshal(envMap map[string]string) (string, error) {
	lines := make([]string, 0, len(envMap))
	for k, v := range envMap {
		lines = append(lines, fmt.Sprintf(`%s="%s"`, k, doubleQuoteEscape(v)))
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n"), nil
}

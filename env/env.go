package env

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Read this function receives a path to a file and retrieves the variables defined in that file as environment variables.
// There must be one and only one variable per line and its format must be name=value
// any leading and trailing whitespaces in both name and value are trimmed.
// Any incorrect formatted line or with a leading # is ignored.
func Read(path string) {
	var file *os.File
	var err error
	if file, err = os.Open(path); err != nil {
		log.Fatalf("Error opening environment variable file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if validateLine(scanner.Text()) {
			pairKeyValue := strings.SplitN(scanner.Text(), "=", 2)
			if len(pairKeyValue) == 2 {
				err = os.Setenv(strings.Trim(pairKeyValue[0], " "),
					strings.Trim(pairKeyValue[1], " "))
				if err != nil {
					log.Fatalf("Error reading environment variable file: %s", err)
				}
			}
		}
	}

	err = file.Close()
	if err != nil {
		log.Fatalf("Error closing environment variable file: %s", err)
	}

}

func validateLine(text string) bool {
	if len(text) > 0 && !strings.HasPrefix(text, "#") {
		return true
	} else {
		return false
	}
}

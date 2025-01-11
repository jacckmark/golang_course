// new packages need to be placed in go into separate folders named like the packages
// (here fileops). The files that are part of this package does not have to be named
// like the package (they only need to have the 'package nameOfThePackage' in them)
package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// when exporting the functions/variables from the package the casing indicates
// if the function/variable is exported (no 'export' keyword). So here we had to
// make the function names uppercase to make them usable in other parts of the app
func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

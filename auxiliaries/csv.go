package auxiliaries

import (
	"encoding/csv"
	"os"
	"fmt"
	"strings"
)

func GetCSV(path string) (map[string] []string, error) {
	// try to get the data from the file at path
	data, errReadingFile := os.ReadFile(path)
	if errReadingFile != nil {
		return map[string] []string{}, errReadingFile
	}

	// read the data string using csv library
	r := csv.NewReader(strings.NewReader( fmt.Sprintf("%s", data) ))
	r.Comma = ','

	// format using csv library
	records, err := r.ReadAll()
	if err != nil {
		return map[string] []string{}, err
	}
	/* records have the format:
		[[column1, column2, column3], 
		[data1.1, data1.2, data1.3],
		[data2.1, data2.2, data2.3]] */

	// Format to the way I want
	toReturn := make(map[string] []string)
	for columnPosition, columnName := range records[0] {
		dataSlice := make([]string, len(records[1:]), len(records[1:]))

		for i := range (len(records) - 1) {
			dataSlice[i] = records[1+i][columnPosition]
		}
		
		toReturn[columnName] = dataSlice
	}
	/* 	Now to return have the format 
		column1: [data1.1, data1.2]
		column2: [data2.1, data2.2]
		column3: [data3.1, data3.2] */

	fmt.Println(toReturn)
	return toReturn, nil
}


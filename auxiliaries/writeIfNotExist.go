package auxiliaries

import "os"

func WriteIfNotExist(filePath string, toWrite string) (error) {
	// try to open the file
	_, openFileErr := os.OpenFile(filePath, 0, 0644)

	// if the file dosen't exist
	if openFileErr != nil && !os.IsExist(openFileErr) {
		writeErr := os.WriteFile(filePath, []byte(toWrite), 0666)
		if writeErr != nil {
			return writeErr
		}
	// if the file exist and have some kind of error
	} else if os.IsExist(openFileErr) {
		return openFileErr
	}

	// if the file exist and don't have any kind of error
	return nil
}


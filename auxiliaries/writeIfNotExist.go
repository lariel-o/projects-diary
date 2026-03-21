package auxiliaries

import "os"

func WriteIfNotExist(filePath string, toWrite string) (error) {
	_, openFileErr := os.OpenFile(filePath, 0, 0644)

	if openFileErr != nil && !os.IsExist(openFileErr) {
		writeErr := os.WriteFile(filePath, []byte(toWrite), 0666)
		if writeErr != nil {
			return writeErr
		}
	} else if os.IsExist(openFileErr) {
		return openFileErr
	}

	return nil
}


package auxiliaries

import "os"

func CreateProjectDir(projectPath string) (error) {
	err := os.Mkdir(projectPath, 0750)	
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}


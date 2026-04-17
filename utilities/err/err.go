package err

import (
	"os"
	"log"
)

func LogErr(e error) {
	log.Fatal(e)
	os.Exit(1)
}


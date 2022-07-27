package args

import (
	"os"
	"strings"
)

func Get() string {
	var args = strings.Join(os.Args[1:len(os.Args)]," ")
	return args
}

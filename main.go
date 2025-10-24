package main

import (
	"gbase/src/migrate"
	"gbase/src/route"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			migrate.Run()
			return
		}
	}

	route.StartService()
}

package main

import (
	"hostgator/log"
	"hostgator/router"
)

func main() {
	// mocks.Populate()
	// log.InicializaLog()
	var r = router.StartRouter()
	var err = r.Run(":4000")
	if err != nil {
		log.Error(err)
	}
}

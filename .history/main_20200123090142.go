package main

import (
	"agplanservoperadoraapi/router"
	"hostgator/log"
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

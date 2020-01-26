package main

import (
	"hostgatorBackend/log"
	"hostgatorBackend/router"
)

func main() {
	var r = router.StartRouter()
	var err = r.Run(":4000")
	if err != nil {
		log.Error(err)
	}
}

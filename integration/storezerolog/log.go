package storezerolog

import (
	"github.com/gubtos/gosubmoduleexample"
	"github.com/rs/zerolog/log"
)

func Log(data gosubmoduleexample.Data) {
	log.Printf("data A: %s\ndata B: %s", data.ValA, data.ValB)
}

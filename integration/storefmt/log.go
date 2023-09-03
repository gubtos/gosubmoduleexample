package fmtlog

import (
	"fmt"

	"github.com/gubtos/gosubmoduleexample"
)

func Log(data gosubmoduleexample.Data) {
	fmt.Printf("data A: %s\ndata B: %s", data.ValA, data.ValB)
}

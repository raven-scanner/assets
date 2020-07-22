package main

import (
	"github.com/raven-scanner/assets/model"

	utils "github.com/circuit-platform/models-utils"
)

func main() {
	utils.Run(model.CreateAssetsIndex, SyncNamespaces)
}
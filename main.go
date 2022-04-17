package main

import (
	"github.com/Vause/pet-information/internal"
	"github.com/Vause/pet-information/internal/services"
	"github.com/Vause/pet-information/internal/utils"
)

func main() {
	configuration := utils.GetConfig()
	services.SetUp(configuration)
	internal.GetBreeds()
}

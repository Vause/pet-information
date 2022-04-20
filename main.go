package main

import (
	"github.com/Vause/pet-information/internal/petInformation"
	"github.com/Vause/pet-information/internal/utils"
)

func main() {
	utils.SetUpConfig()
	petInformation.GetCatBreeds()
	petInformation.GetDogBreeds()
}

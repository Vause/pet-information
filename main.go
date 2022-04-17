package main

import (
	"github.com/Vause/pet-information/internal/petInformation"
	"github.com/Vause/pet-information/internal/utils"
)

func main() {
	petInformation.GetCatBreeds(utils.GetConfig("cat"))
	petInformation.GetDogBreeds(utils.GetConfig("dog"))
}

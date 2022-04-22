package main

import (
	"fmt"

	"github.com/Vause/pet-information/internal/petInformation"
	"github.com/Vause/pet-information/internal/petInformation/shared"
	"github.com/Vause/pet-information/internal/utils"
)

func main() {
	utils.SetUpConfig()
	cats := petInformation.GetCats()
	dogs := petInformation.GetDogs()
	fishes := petInformation.GetFishes()

	for _, cat := range cats {
		fmt.Println(GetPetName(&cat))
	}

	for _, dog := range dogs {
		fmt.Println(GetPetName(&dog))
	}

	for _, fish := range fishes {
		fmt.Println(GetPetName(&fish))
	}
}

func GetPetName(p shared.Pet) string {
	return p.GetName()
}

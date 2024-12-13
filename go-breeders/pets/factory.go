package pets

import "github.com/rishabhsvats/go-patterns/go-breeders/models"

func NewPet(species string) *models.Pet {
	pet := models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "no description entered yet",
		Lifespan:    0,
	}

	return &pet
}

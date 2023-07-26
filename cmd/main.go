package main

import (
	"fmt"
	"github.com/ferminhg/learning-go/internal/application"
	"github.com/ferminhg/learning-go/internal/infra/storage"
)

func main() {
	fmt.Println("Marketplace: wop wop 🌍")

	repository := storage.NewInMemoryAdRepository()
	service := application.AdService{Repository: repository}

	fmt.Println("💾 Posting Adds")
	ad1, _ := service.Post("t1", "d1", 1)
	service.Post("t2", "d2", 2)
	service.Post("t3", "d3", 3)
	service.Post("t4", "d4", 4)
	service.Post("t5", "d5", 5)
	service.Post("t6", "d6", 6)
	service.Post("t7", "d7", 7)
	service.Post("t8", "d8", 8)

	fmt.Println("🔎 Finding Add")
	ad, _ := service.Find(ad1.Id.String())

	fmt.Println(ad)

	fmt.Println("🎟️ Finding Random Adds")

	ads, _ := service.FindRandom()
	fmt.Println(ads)
}

package domain

import (
	"fmt"
	"github.com/go-faker/faker/v4"
)

func RandomAdFactory() Ad {
	a := Ad{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Random Ad", a)
	return a
}

package main

import (
	"fmt"
	"github.com/soma115/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	o := drukuj("aaa", "bbb")
	fmt.Println(o)

	value, exist := 2, 1
	fmt.Println(value, exist)

	obiekt := Klasa {
		Name: "Jan",
		Power: 9002
	}				//deklaracja :=
	obiekt = Klasa{"Tom", 9004}	//reasignement =

	fmt.Println(obiekt.Name, obiekt.Power)

	goku := Klasa{"Goku", 9000}
	Super(&goku)
	fmt.Println(goku.Power)
}

func Super(s *Klasa) {
	s.Power += 10000
}

func drukuj(x, y string) int {
	fmt.Println(x, y)
	return 2
}

type Klasa struct {
	Name  string
	Power int
}

// zmienna typ
// & = wartość -> adres. &pokaz_adres_pod_którym_jest_wartość_tej_zmiennej
// * = adres -> wartość. *wartość_typu_adres. Do przyjmowania adresu.

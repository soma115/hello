package main

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/golang/example/stringutil"
	"github.com/tylerwince/godbg"
)

// Zmienna globalna musi mieć obowiązkowo komentarz
var Zmienna int

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	o := drukuj("aaa", "bbb")
	fmt.Println(o)

	value, exist := 2, 1
	fmt.Println(value, exist)

	obiekt := Klasa{ //deklaracja :=
		Name:  "Jan",
		Power: 9002,
	}
	obiekt = Klasa{ //reasignement =
		"Tom",
		9004,
	}

	fmt.Println("obiekt.Name, obiekt.Power", obiekt.Name, obiekt.Power)

	goku := Klasa{"Goku", 9000}
	Super(&goku)
	fmt.Println(goku.Power)

	zmienna := 1
	zmienna += 2

	scores := [3]int{99, 200, 300}
	scores[0] = 100
	fmt.Print("scores len: ", len(scores), "\n")
	for i, n := range scores {
		fmt.Println(i, n)
	}

	/////////////////////////////////////////////////

	scores1 := []int{1, 2, 3, 4, 5}
	slice := scores1[2:4]
	fmt.Print("slice: ", slice, "\n")
	slice[0] = 999
	fmt.Println(scores1)

	/////////////////////////////////////////////////

	scores2 := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores2[i] = int(rand.Int31n(999))
	}
	sort.Ints(scores2)
	worst := make([]int, 5)
	copy(worst, scores2[:5])
	fmt.Println(worst)

	/////////////////////////////////////////////////

	// Array
	scores3 := make([]int, 100)
	scores4 := []int{1, 2, 3, 4, 5} //wraz z wypełnieniem wartościami
	fmt.Print(scores3, scores4)

	// Slice
	slice1 := scores3[2:4]
	godbg.Dbg(slice1)

	// Map (dictionary)
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	// power, exists := lookup["vegeta"]
}

// Super obowiązkowy komantarz
func Super(s *Klasa) {
	s.Power += 10000
}

func drukuj(x, y string) int {
	fmt.Println(x, y)
	return 2
}

// Klasa obowiązkowy komentarz
type Klasa struct {
	Name  string
	Power int
}

// Metoda Klasy
func (p *Klasa) Metoda() {
	fmt.Printf("HiHiHi %s\n", p.Name)
}

// zmienna typ
// & = wartość -> adres. &pokaz_adres_pod_którym_jest_wartość_tej_zmiennej
// * = adres -> wartość. *wartość_typu_adres. Do przyjmowania adresu.

package main

import "fmt"

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (t *Trade) Value() float64 {

	v := t.Price * float64(t.Volume)

	if t.Buy {
		v = -v
	}

	return v
}

func main() {

	t1 := Trade{"MSFT", 10, 1000, false}

	fmt.Println(t1)

	fmt.Printf("%+v\n", t1)

	t2 := Trade{

		Symbol: "GOOG",
		Volume: 1000,
		Price:  10,
		Buy:    true,
	}

	t3 := t2

	fmt.Println(t2, t3)
	fmt.Println(t2.Value())
}

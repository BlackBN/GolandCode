package point

import "fmt"

type Action interface {
	Fight()
}

type Person struct {
}

func (p Person) Fight() {
	fmt.Println("Fight")
}

type Point struct {
	x, y int
}

func New(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p Point) Change(x, y int) {
	p.x = x + y
	p.y = y * y

}

func (p *Point) Change2(x, y int) {
	p.x = x + y
	p.y = y * y
}

type DefInt int

func (d DefInt) Print() {}

type DefFloat float64

func (d DefFloat) Print() {}

type DefString string

func (d DefString) Print() {}

type DefArray [30]DefMap

func (d DefArray) Print() {}

type DefChan chan struct{}

func (d DefChan) Send() {}

type DefSlice []DefChan

func (d DefSlice) Print() {}

type DefMap map[string]DefSlice

func (d DefMap) Print() {}

type DefStruct struct {
	DefArray
}

func (d DefStruct) Print() {}

// compile error
// type DefInt *int
// func (d DefInt) Print() {}
// type DefInterface interface{}
// func (d DefInterface) Print() {}

package makeCars

type Car interface {
	DescribeCar() (string, int)
}

type sedan struct {
}

func MakeSedan() Car {
	return &sedan{}
}

func (s sedan) DescribeCar() (string, int) {
	return "x", 1
}

type suv struct {
}

func MakeSUV() Car {
	return &suv{}
}

func (u suv) DescribeCar() (string, int) {
	return "x", 1
}

type coupe struct {
}

func MakeCoupe() Car {
	return &coupe{}
}

func (c coupe) DescribeCar() (string, int) {
	return "x", 1
}

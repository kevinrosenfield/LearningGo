package pymate

type Primate struct {
	age			 int
	currX		 int
	currY		 int
	currCalories int
	sex 		 string
}

type Group struct {
	primates    []Primate
	numPrimates int
}

func (g *Group) CreatePrimates(numPrimates int) {
	for i := 0; i < numPrimates; i++ {
		p := Primate{
			age: 0,
			currX: 0,
			currY: 0,
			sex: "",
		}
		g.primates = append(g.primates, p)
	}
}
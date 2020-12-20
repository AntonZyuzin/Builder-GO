package main

import "fmt"

type Pastry string
type Sauce string
type Filling string

const (
	YeastfreePastry Pastry = "бездрожжевое тесто,"
	YeastPastry            = "дрожжевое тесто,"
)

const (
	TomatoSauce Sauce = "томатный соус,"
)

const (
	MargaritaFilling Filling = "сыр Моцарелла, базилик, помидоры."
	DiabloFilling            = "сыр Моцарелла, шампиньоны, чили, паприка."
)

type Pizza interface {
	Pastry(Pastry) Pizza
	Sauce(Sauce) Pizza
	Fill(Filling) Pizza
	TakePizza() error
}

type newPizza struct {
	pastry  Pastry
	sauce   Sauce
	filling Filling
}

func (np *newPizza) Pastry(p Pastry) Pizza {
	np.pastry = p
	return np
}
func (np *newPizza) Sauce(s Sauce) Pizza {
	np.sauce = s
	return np
}
func (np *newPizza) Fill(f Filling) Pizza {
	np.filling = f
	return np
}
func (np *newPizza) TakePizza() error {
	return nil
}

func New() Pizza {
	return &newPizza{}
}

func main() {
	builder := New()
	margarita := builder.Pastry(YeastfreePastry).Sauce(TomatoSauce).Fill(MargaritaFilling)
	margarita.TakePizza()
	fmt.Println(margarita)
	diablo := builder.Pastry(YeastPastry).Sauce(TomatoSauce).Fill(DiabloFilling)
	diablo.TakePizza()
	fmt.Println(diablo)
}

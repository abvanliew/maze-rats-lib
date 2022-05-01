package components

import "fmt"

var Arcane magic

type magic struct {
	spellCompositions [3]spellComposition
}

func (m *magic) PickSpell() string {
	comp := dice.d3()
	fistSub := dice.d2()
	secondSub := dice.d2()
	return m.getSpell(comp, fistSub, secondSub)
}

func (m *magic) getSpell(comp, firstSub, secondSub int) string {
	return m.spellCompositions[comp].get(firstSub, secondSub)
}

type spellComposition struct {
	first  spellHalves
	second spellHalves
}

func (composistion *spellComposition) get(firstSub, secondSub int) string {
	return fmt.Sprintf(
		"%s %s",
		composistion.first.substance[firstSub].pick(),
		composistion.second.substance[secondSub].pick(),
	)
}

type spellHalves struct {
	substance [2]*randTable
}

func shapeMagic() {
	effect := spellHalves{[2]*randTable{
		tables["PhysicalEffect"],
		tables["EtherealEffect"],
	}}
	element := spellHalves{[2]*randTable{
		tables["PhysicalElement"],
		tables["EtherealElement"],
	}}
	form := spellHalves{[2]*randTable{
		tables["PhysicalForm"],
		tables["EtherealForm"],
	}}
	Arcane = magic{
		[3]spellComposition{
			{effect, form},
			{element, form},
			{effect, element},
		},
	}
}

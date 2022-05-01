package components

import (
	"math/rand"
	"time"
)

type diceBag struct {
	randomizer *rand.Rand
}

func (d *diceBag) roll() int {
	return d.randomizer.Intn(36)
}

func (d *diceBag) d2() int {
	return d.randomizer.Intn(2)
}

func (d *diceBag) d3() int {
	return d.randomizer.Intn(3)
}

func (d *diceBag) d6() int {
	return d.randomizer.Intn(6)
}

var dice diceBag

func openDiceBag() {
	dice.randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))
}

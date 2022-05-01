package components

import (
	"fmt"
)

var tables map[string]*randTable = make(map[string]*randTable)

type randTable struct {
	entries [36]entry
}

func newTable(values [36]string) *randTable {
	var t [36]entry
	for i, v := range values {
		t[i] = *newEntry(v)
	}

	return &randTable{entries: t}
}

func (r *randTable) get(i int) string {
	return r.entries[i].get()
}

func (r *randTable) pick() string {
	return r.get(dice.roll())
}

func (r *randTable) addReference(refTable string, i int) {
	r.entries[i].refTable = refTable
}

type entry struct {
	text     string
	refTable string
}

func newEntry(text string) *entry {
	return &entry{text: text}
}

func (e *entry) get() string {
	res := e.text
	if ref, ok := tables[e.refTable]; ok {
		res = fmt.Sprintf(res, ref.pick())
	}
	return res
}

func generateTables() {
	tables["PhysicalEffect"] = newTable([36]string{
		"Animating", "Attracting", "Binding", "Blossoming", "Consuming", "Creeping",
		"Crushing", "Diminishing", "Dividing", "Duplicating", "Eveloping", "Expanding",
		"Fusing", "Grasping", "Hastening", "Hindering", "Illuminating", "Imprisoning",
		"Levitating", "Opening", "Petrifying", "Phasing", "Piercing", "Pursing",
		"Reflecting", "Regenerating", "Rending", "Repeling", "Resurrecting", "Screaming",
		"Sealing", "Shapeshifting", "Shielding", "Spawning", "Transmuting", "Transporting"},
	)
	tables["EtherealEffect"] = newTable([36]string{
		"Avenging", "Bashing", "Bewildering", "Blinding", "Charming", "Communication",
		"Compelling", "Concealing", "Deafening", "Deceiving", "Deciphering", "Disgusing",
		"Dispelling", "Emboldening", "Encoding", "Energizing", "Enlightening", "Enraging",
		"Excrutiating", "Foreseeing", "Intoxicating", "Maddening", "Mesmerizing", "Mindreading",
		"Nullifying", "Paralyzing", "Revealing", "Revolting", "Scrying", "Silencing",
		"Smoothing", "Summoning", "Terrifying", "Warding", "Wearying", "Withering"},
	)
	tables["PhysicalElement"] = newTable([36]string{
		"Acid", "Amber", "Bark", "Blood", "Bone", "Brine",
		"Clay", "Crow", "Crystal", "Ember", "Flesh", "Fungus",
		"Glass", "Honey", "Ice", "Insect", "Wood", "Lava",
		"Moss", "Obsidian", "Oil", "Poison", "Rat", "Salt",
		"Sand", "Sap", "Serpent", "Slime", "Stone", "Tar",
		"Thorn", "Vine", "Water", "Wine", "Wood", "Worm"},
	)
	tables["EtherealElement"] = newTable([36]string{
		"Ash", "Chaos", "Distoriation", "Dream", "Dust", "Echo",
		"Ectoplasm", "Fire", "Fog", "Ghost", "Harmony", "Heat",
		"Light", "Lightning", "Memory", "Mind", "Mutation", "Negation",
		"Plague", "Plasma", "Probability", "Rain", "Rot", "Shadow",
		"Smoke", "Snow", "Soul", "Star", "Stasis", "Steam",
		"Thunder", "time", "Void", "Warp", "Whisper", "Wind"},
	)
	tables["PhysicalForm"] = newTable([36]string{
		"Alter", "Armor", "Arrow", "Beast", "Blade", "Cauldron",
		"Chain", "Chariot", "Claw", "Cloak", "Colossus", "Crown",
		"Elemental", "Eye", "Fountain", "Gate", "Golem", "Hammer",
		"Horn", "Key", "Mask", "Monolith", "Pit", "Prison",
		"Sentinel", "Servant", "Shield", "Spear", "Steed", "Swarm",
		"Tentacle", "Throne", "Torch", "Trap", "Wall", "Web"},
	)
	tables["EtherealForm"] = newTable([36]string{
		"Aura", "Beacon", "Beam", "Blast", "Blob", "Bolt",
		"Bubble", "Call", "Cascade", "Circle", "Cloud", "Coil",
		"Cone", "Cube", "Dance", "Disk", "Field", "Form",
		"Gaze", "Loop", "Moment", "Nexus", "Portal", "Pulse",
		"Pyramid", "Ray", "Shard", "Sphere", "Spray", "Storm",
		"Swarm", "Torrent", "Touch", "Vortex", "Wave", "Word"},
	)
	shapeMagic()

	tables["ReferenceTest"] = newTable([36]string{
		"%v-skin", "%v-skin", "%v-skin", "%v-skin", "%v-skin", "%v-skin",
		"%v-feet", "%v-feet", "%v-feet", "%v-feet", "%v-feet", "%v-feet",
		"%v-arms", "%v-arms", "%v-arms", "%v-arms", "%v-arms", "%v-arms",
		"%v-legs", "%v-legs", "%v-legs", "%v-legs", "%v-legs", "%v-legs",
		"%v-head", "%v-head", "%v-head", "%v-head", "%v-head", "%v-head",
		"%v-hands", "%v-hands", "%v-hands", "%v-hands", "%v-hands", "%v-hands"},
	)

	for i := 0; i < 36; i++ {
		tables["ReferenceTest"].addReference("PhysicalEffect", i)
	}
}

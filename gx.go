package genetix

import (
	"math/rand"
	"sort"
)

// Entity represents something to be evolved.
// You should update the scores after each evolution.
type Entity interface {
	Score() float64
	Mutate()
	CrossOver(Entity)
	Clone() Entity
	String() string
}

// Population is a collection of entities.
// The only purpose of defining this type, is to make it sort-able.
type Population []Entity

func (p Population) Len() int           { return len(p) }
func (p Population) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Population) Less(i, j int) bool { return p[i].Score() < p[j].Score() }

// Evolve the population according to the scores.
// The top `elites` entities will always be perserved.
func Evolve(pop Population, elites int, mutate, cross float64) {
	n := pop.Len()
	// prepare
	sort.Sort(pop)
	// perserve
	top := make([]Entity, elites)
	for i := range top {
		top[i] = pop[i].Clone()
	}
	// mutate
	for loop := int(float64(n) * mutate); loop >= 0; loop-- {
		i := rand.Intn(n)
		pop[i].Mutate()
	}
	// crossover
	for loop := int(float64(n*(n-1)/2) * cross); loop >= 0; loop-- {
		i := rand.Intn(n)
		j := rand.Intn(n)
		pop[i].CrossOver(pop[j])
	}
	// restore elites
	for i, e := range top {
		pop[i] = e
	}
}

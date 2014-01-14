package genetix

import (
	"math/rand"
	"sort"
)

// Population represents the entities to be evolved.
type Population interface {
	sort.Interface
	// Mutate the entity at i.
	Mutate(i int)
	// CrossOver the two entities at i and j. Both will be modified.
	CrossOver(i, j int)
	// Copy the entity from src to dst.
	Copy(dst, src int)
	// Dedup removes all duplicates, and fills in new entities if needed.
	Dedup()
}

// Evolve the population according to the scores.
//
// The top entities will be preserved as elites, and they will
// always be sorted and placed at [0..elites-1].
// The last elements of pop will be used for temporary storage for elites.
func Evolve(pop Population, elites int, mutate, cross float64) {
	n := pop.Len()
	reduced := n - elites
	// prepare
	sort.Sort(pop)
	// preserve
	for i := 0; i < elites; i++ {
		pop.Copy(n-i-1, i)
	}
	// mutate
	for loop := int(float64(n) * mutate); loop >= 0; loop-- {
		i := rand.Intn(reduced)
		pop.Mutate(i)
	}
	// crossover
	for loop := int(float64(n*(n-1)/2) * cross); loop >= 0; loop-- {
		pop.CrossOver(rand.Intn(reduced), rand.Intn(reduced))
	}
	// restore
	for i := 0; i < elites; i++ {
		pop.Swap(i, n-i-1)
	}
	// clean up
	pop.Dedup()
}

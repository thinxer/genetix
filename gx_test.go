package genetix

import (
	"math/rand"
	"testing"
)

var panda = []byte("giantpanda")

type findPanda []byte

func (f findPanda) Score() (s float64) {
	for i, c := range f {
		if c != panda[i] {
			s += 1.0
		}
	}
	return
}

func (f findPanda) Mutate() {
	i := rand.Intn(len(f))
	f[i] = byte('a') + byte(rand.Intn(26))
}

func (f findPanda) CrossOver(oe Entity) {
	o := oe.(findPanda)
	index := rand.Intn(len(o))
	for i := 0; i <= index; i++ {
		f[i], o[i] = o[i], f[i]
	}
}

func (f findPanda) Clone() Entity {
	o := findPanda(make([]byte, len(f)))
	copy(o, f)
	return o
}

func (f findPanda) String() string {
	return string(f)
}

func TestEvolution(t *testing.T) {
	var pop = Population(make([]Entity, 100))
	for i := range pop {
		pop[i] = findPanda(make([]byte, len(panda)))
	}
	for i := 0; i < 1000; i++ {
		Evolve(pop, 10, 0.2, 0.1)
		t.Logf("Iter: %v, top: %v %v", i, pop[0].Score(), pop[0])
		if pop[0].Score() <= 0.0001 {
			break
		}
	}
}

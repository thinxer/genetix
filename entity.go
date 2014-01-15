package genetix

// Entity represents something to be evolved.
type Entity interface {
	Score() float64

	// Reset sets the entity to its zero value.
	Reset()
	// Mutate the entity randomly.
	Mutate()
	// CrossOver with another entity. Both will be modified.
	CrossOver(Entity)
	// Clone returns a deep copy of this entity.
	Clone() Entity
}

// EntityPopulation implements Population for easier usage.
type EntityPopulation []Entity

func (p EntityPopulation) Len() int           { return len(p) }
func (p EntityPopulation) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p EntityPopulation) Less(i, j int) bool { return p[i].Score() < p[j].Score() }
func (p EntityPopulation) Mutate(i int)       { p[i].Mutate() }
func (p EntityPopulation) CrossOver(i, j int) { p[i].CrossOver(p[j]) }
func (p EntityPopulation) Copy(i, j int)      { p[i] = p[j].Clone() }

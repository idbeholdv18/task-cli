package ids

type Id uint

// Returns the next in order, after the given ID.
//
// In other words, it increments the given ID.
func (id Id) Next() Id {
	return id + 1
}

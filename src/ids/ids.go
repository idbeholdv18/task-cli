package ids

type Id int

type CompareResult int

const Less, Equal, Greater = -1, 0, 1

func (id Id) Next() Id {
	return Id(int(id) + 1)
}

// Returns the result of comparing a and b:
//
//   - a < b: 	returns ids.Less
//   - a == b: 	returns ids.Equal
//   - a > b: 	returns ids.Greater
func (a Id) Compare(b Id) CompareResult {
	if a < b {
		return Less
	}
	if a > b {
		return Greater
	}
	return Equal
}

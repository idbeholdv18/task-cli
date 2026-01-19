package shared

type Id int

type CompareResult int

const (
	Less    = -1
	Equal   = 0
	Greater = 1
)

func (id Id) Next() Id {
	return Id(int(id) + 1)
}

func (a Id) Max(b Id) Id {
	if int(a) >= int(b) {
		return a
	}
	return b
}

func (a Id) Compare(b Id) CompareResult {
	if a < b {
		return Less
	}
	if a > b {
		return Greater
	}
	return Equal
}

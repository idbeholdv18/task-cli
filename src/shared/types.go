package shared

type Id int

func CreateId(payload int) Id {
	return Id(payload)
}

func (id Id) Next() Id {
	return Id(int(id) + 1)
}

func (a Id) Compare(b Id) Id {
	if int(a) >= int(b) {
		return a
	}
	return b
}

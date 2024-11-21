package profiles

type ProfileVariable struct {
	Name  string
	Value string
}

type Profile struct {
	Name          string
	InheritFrom   string
	DefaultSwitch bool
	Variables     []ProfileVariable
}

func (p Profile) Equal(q Profile) bool {
	if p.Name != q.Name {
		return false
	}

	if len(p.Variables) != len(q.Variables) {
		return false
	}

	for i, item := range p.Variables {
		if item.Name != q.Variables[i].Name {
			return false
		}

		if item.Value != q.Variables[i].Value {
			return false
		}
	}

	return true
}

func (p Profile) NotEqual(q Profile) bool {
	return !p.Equal(q)
}

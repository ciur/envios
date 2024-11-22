package profiles

import "fmt"

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

type ProfilesList []Profile

func (p Profile) Equal(q Profile) bool {
	if p.Name != q.Name {
		return false
	}

	if p.InheritFrom != q.InheritFrom {
		return false
	}

	if p.DefaultSwitch != q.DefaultSwitch {
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

func (arr ProfilesList) FindProfile(name string) *Profile {
	for _, profile := range arr {
		if profile.Name == name {
			return &profile
		}
	}

	return nil
}

func (arr ProfilesList) PrintExports(name string) {
	found := arr.FindProfile(name)

	if found != nil {
		found.PrintExports()
		if found.InheritFrom != "" {
			inheritFrom := arr.FindProfile(found.InheritFrom)
			if inheritFrom != nil {
				inheritFrom.PrintExports()
			}
		}
	} else {
		fmt.Printf("Profile '%s' not found\n", name)
	}

}

func (p Profile) PrintExports() {
	for _, v := range p.Variables {
		fmt.Printf("export %s=%s\n", v.Name, v.Value)
	}
}

func (arr ProfilesList) List() {
	for i, p := range arr {
		if p.InheritFrom != "" {
			fmt.Printf("%d. %s inherits from %s\n", i+1, p.Name, p.InheritFrom)
		} else {
			fmt.Printf("%d. %s\n", i+1, p.Name)
		}
	}
}

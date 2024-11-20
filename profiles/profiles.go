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

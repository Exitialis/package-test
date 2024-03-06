package names

type Name struct {
	name string
}

func New(name string) *Name {
	return &Name{name: name}
}

func (name *Name) ToString() string {
	return name.name
}

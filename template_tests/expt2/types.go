package expt2
type Data struct {
	Name   string
	Desc   string
	Fields []Field
}

type Field struct {
	Name     string
	TypeName string
}
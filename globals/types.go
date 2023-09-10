package globals

type ScaffoldInfoT struct {
	Filepath string
	Content  string
}

type ScaffoldInfoTListT []ScaffoldInfoT

type MoveItemDetailsT struct {
	Key   string
	Index int
}

type SubstitionTemplateT struct {
	FullQueryId string
	MoveItemsInfo map[string]MoveItemDetailsT
}


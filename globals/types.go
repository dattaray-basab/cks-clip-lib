package globals

type ScaffoldInfoT struct {
	Filepath string
	Content  string
}

type ScaffoldInfoTListT []ScaffoldInfoT

type MoveItemDetailsT struct {
	Key   string
	Index int
	IsLastItem bool 
}

type SubstitionTemplateT struct {
	FullQueryId string
	ShortQueryId string
	MoveItemsInfo map[string]MoveItemDetailsT
}


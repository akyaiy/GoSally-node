package config

type _config map[string]ConfValue

type TypeConfig _config

const (
	SrcNONE Source = iota
	SrcDEF
	SrcCMD
	SrcCFG
	SrcENV
	SrcCST
)

const (
	True = "true"
)

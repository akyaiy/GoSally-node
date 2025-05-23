package config

//type _configParser interface {
//	ParseConfigFile(path string)
//}

type _paramParser interface {
	ProgramConfig() map[string]string
	ParseCMDlineArgs() error
	ParseArgs(args []string) error
}

type Parser interface {
	//	_configParser
	_paramParser
}

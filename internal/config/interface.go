package config

//type _configParser interface {
//	ParseConfigFile(path string)
//}

type _CMDLineParser interface {
	ProgramConfig() TypeConfig
	ParseCMDlineArgs() error
	ParseArgs(args []string) error
}

type Parser interface {
	//	_configParser
	_CMDLineParser
}

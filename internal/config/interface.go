package config

type _configParser interface {
	ParseConfigFile(path string)
}

type _paramParser interface {
	ProgramConfig() map[string]string
	ParseCMDlineArgs()
	ParseArgs(args []string)
}

type Parser interface {
	//	_configParser
	_paramParser
}

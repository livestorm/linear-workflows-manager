package linear

var boardConfigurationMapping = map[string]BoardConfig{
	"LIV": techBoardConfig,
}

func GetBoardConfig(team string) BoardConfig {
	return boardConfigurationMapping[team]
}

package data

//Levels contains an array of the different available levels
type Levels []string

//LoggerLevel encapsulate the levels of a logger
type LoggerLevel struct {
	ConfiguredLevel string `json:"configuredLevel"`
	EffectiveLevel  string `json:"effectiveLevel"`
}

//Loggers contains a map between loggers and their levels
type Loggers map[string]LoggerLevel

//ActuatorData is the top structure to manage actuators and their levels
type ActuatorData struct {
	Levels  Levels  `json:"levels"`
	Loggers Loggers `json:"loggers"`
}

//PredefinedLevels contains an array with the possible logger levels
var PredefinedLevels = []string{
	"OFF",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
	"TRACE",
}

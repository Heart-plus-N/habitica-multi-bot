package observer_pattern

type SharedConfig struct {
	HabiticaUsername string `mapstructure:"HABITICA_USERNAME"`
	HabiticaPassword string `mapstructure:"HABITICA_PASSWORD"`
}

package config

type GlobalConfig struct {
	MODE        string
	ProgramName string
	AUTHOR      string
	VERSION     string
	REDIS       struct {
		Use    bool
		Config struct {
			IP       string
			PORT     string
			PASSWORD string
			DB       int
		}
	}
}

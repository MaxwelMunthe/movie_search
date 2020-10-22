package infrastructure

type app struct {
	Appname string `yaml:"name"`
	Debug   bool   `yaml:"debug"`
	Port    string `yaml:"port"`
	Service string `yaml:"service"`
	Host    string `yaml:"host"`
}

type OmdbApi struct {
	Url string `yaml:"url"`
	Key string `yaml:"key"`
}

type Environment struct {
	App     app     `yaml:"app"`
	OmDbApi OmdbApi `yaml:"omDbApi"`
	path    string
}

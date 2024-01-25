package option

type Options struct {
	Provider Provider `yaml:"provider"`
}

type Provider struct {
	HTTP HTTPProvider `yaml:"http"`
}

type HTTPProvider struct {
	Addr string `yaml:"addr"`
}

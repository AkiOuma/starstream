package definition

type Definition struct {
	Destination     string    `yaml:"destination"`
	ServiceName     string    `yaml:"serviceName"`
	Type            string    `yaml:"type"`
	ApiVersion      string    `yaml:"apiVersion"`
	ThirdPartyProto []string  `yaml:"thirdPartyProto"`
	Entity          []*Entity `yaml:"entity"`
	Config          Config    `yaml:"config"`
}

type Entity struct {
	Name  string   `yaml:"name"`
	Field []*Field `yaml:"field"`
}

type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Root bool   `yaml:"root"`
}

type Config struct {
	Transport  *Transport  `yaml:"transport"`
	Datasource *Datasource `yaml:"datasource"`
}

type Transport struct {
	Grpc *Grpc `yaml:"grpc"`
}

type Datasource struct {
	Sqldb *Sqldb `yaml:"sqldb"`
}

type Grpc struct {
	Addr string `yaml:"addr"`
}

type Sqldb struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Database string `yaml:"database"`
}

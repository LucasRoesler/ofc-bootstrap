package types

type Plan struct {
	Orchestration string                   `yaml:"orchestration"`
	Secrets       []KeyValueNamespaceTuple `yaml:"secrets"`
	RootDomain    string                   `yaml:"root_domain"`
	Registry      string                   `yaml:"registry"`
	CustomersURL  string                   `yaml:"customers_url"`
	Github        Github                   `yaml:"github"`
}

type KeyValueTuple struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type FileSecret struct {
	Name      string `yaml:"name"`
	ValueFrom string `yaml:"value_from"`
}

type KeyValueNamespaceTuple struct {
	Name      string          `yaml:"name"`
	Literals  []KeyValueTuple `yaml:"literals"`
	Namespace string          `yaml:"namespace"`
	Files     []FileSecret    `yaml:"files"`
}

type Github struct {
	AppID          string `yaml:"app_id"`
	PrivateKeyFile string `yaml:"private_key_filename"`
}

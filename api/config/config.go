package config

type config struct {
	Mode   string `json:"mode"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	DB struct {
		Postgres struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		} `yaml:"postgres"`
	} `yaml:"db"`
}

func New() *config {
	return &config{}
}

// func readConfig(path string) (*config, error) {
// 	config := &config{}

// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	d := yaml.NewDecoder(file)

// 	if err := d.Decode(&config); err != nil {
// 		return nil, err
// 	}

// 	return config, nil
// }

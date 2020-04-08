package user_config

type Config struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Editor string `json:"editor"`
}

func Get() (*Config, error) {

	newConfig := &Config{
		Name:   "test",
		Email:  "email@example.com",
		Editor: "subl",
	}

	return newConfig, nil
}

func Set(key string, value string) error {

}

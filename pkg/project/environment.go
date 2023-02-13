package project

type Environment string

const (
	EnvironmentDevelopment Environment = "development"
	EnvironmentProduction  Environment = "production"
)

// ValidEnvironment method
func ValidEnvironment(e Environment) bool {
	environments := []Environment{
		EnvironmentDevelopment,
		EnvironmentProduction,
	}

	for _, environment := range environments {
		if environment == e {
			return true
		}
	}

	return false
}

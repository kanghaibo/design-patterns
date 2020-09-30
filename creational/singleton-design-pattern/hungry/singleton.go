package hungry

type singleton struct{}

var instance *singleton

func init() {
	instance = &singleton{}
}

func GetInstance() *singleton {
	return instance
}

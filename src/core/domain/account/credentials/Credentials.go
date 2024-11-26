package credentials

type Credentials struct {
	email    string
	password string
}

func (instance *Credentials) Email() string {
	return instance.email
}

func (instance *Credentials) Password() string {
	return instance.password
}

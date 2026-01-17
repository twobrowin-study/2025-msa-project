package config

type secretString string

// Отвечает интерфейсу cleanenv.Setter
func (sec *secretString) SetValue(s string) error {
	*sec = secretString(s)
	return nil
}

// Отвечает интерфейсу json.Marshaler
func (sec *secretString) MarshalJSON() ([]byte, error) {
	return []byte("\"********\""), nil
}

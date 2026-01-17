package user

// Модель пользователя
type User struct {
	ID int64 `bun:",pk,autoincrement" json:"id"`

	Username string `bun:",notnull" json:"username"`

	FirstName string `bun:",notnull" json:"firstName"`
	LastName  string `bun:",notnull" json:"lastName"`

	Email string `bun:",notnull" json:"email"`
	Phone string `bun:",notnull" json:"phone"`
}

package DataModel

type User struct {
	Id             int
	Reliable_grade int
	Name           string
	Email          string
	Password       string
	Location       string
	Rooms          int
	Price          int
}

//constr
func NewUser(id int) *User {
	return &User{Id: id, Reliable_grade: 100}
}

//the more negative comments the user make => his reliable grade getting smaller
func AddedNagetiveComment(user User) {
	user.Reliable_grade = int(float32(user.Reliable_grade) * 0.95)
}

//the more Positive comments the user make => his reliable grade getting bigger
func AddedPositiveComment(user User) {
	user.Reliable_grade = int(float32(user.Reliable_grade) * 1.01)
}

// set user email
func SetEmail(user User, email string) {
	user.Email = email
}

// set user password
func SetPassword(user User, password string) {
	user.Password = password
}

// set user location
func SetLocation(user User, location string) {
	user.Location = location
}

func SetPrice(user User, price int) {
	user.Price = price
}

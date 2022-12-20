package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextId
	nextId++

	/* address of operator to grab the reference to the user object that comes in.
	 * Because the users collection is storing pointers
	 */
	users = append(users, &u)

	return u, nil

}

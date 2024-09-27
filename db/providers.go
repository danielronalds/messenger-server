package db;

type UserProvider interface  {
    GetUsers() ([]User, error);
    GetUserWithPass(username string, password string) (User, error);
	CreateUser(username, displayName string, hashedPassword, salt []byte) (User, error)
}

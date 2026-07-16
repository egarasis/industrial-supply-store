package entity

const (
	RoleAdmin    = "admin"
	RoleCustomer = "customer"
)

type User struct {
	ID       int
	Email    string
	Password string
	Role     string
}

type UserProfile struct {
	ID          int
	UserID      int
	Email       string
	Password    string
	Role        string
	CompanyName string
	ContactName string
	Phone       string
	Address     string
	FullName    string
}

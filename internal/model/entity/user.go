package entity

const (
	RoleAdmin    = "admin"
	RoleCustomer = "customer"
)

type User struct {
	ID       int
	Email    string
	Password string
	Role     string // ADMIN, CUSTOMER
}

type UserProfile struct {
	ID          int
	UserID      int
	Email       string
	Password    string
	Role        string // ADMIN, CUSTOMER
	CompanyName string
	ContactName string
	Phone       string
	Address     string
}

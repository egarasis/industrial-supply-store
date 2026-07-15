package handlers

import (
	"bufio"
	"fmt"
	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
	"os"
	"strings"
)

type userHandler struct {
	uc           domain.UserUsecase
	adminHandler domain.AdminHandler
}

func NewUserHandler(uc domain.UserUsecase) domain.UserHandler {
	return &userHandler{
		uc: uc,
	}
}

func (h *userHandler) Run() {
	for {
		fmt.Println("\n===== Welcome =====")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")

		fmt.Print("Choose: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			h.login()
		case 2:
			h.create()
		case 0:
			fmt.Println("Good Bye!")
			return
		default:
			fmt.Println("Invalid Menu")
		}
	}
}

func (h *userHandler) create() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Email  : ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Password : ")
	password, _ := reader.ReadString('\n')

	fmt.Print("Role (ADMIN, CUSTOMER): ")
	role, _ := reader.ReadString('\n')

	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	role = strings.TrimSpace(role)

	err := h.uc.Register(email, password, role)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	fmt.Println("User created successfully.")
}
func (h *userHandler) login() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Email  : ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Password : ")
	password, _ := reader.ReadString('\n')

	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	user, err := h.uc.Login(email, password)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	fmt.Println("Login Success")
	fmt.Println("Welcome", email)

	if user.Role == entity.RoleAdmin {
		h.adminHandler.Run()
	}
}

// func (h *userHandler) list() {

// 	users, err := h.uc.GetAll()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println()

// 	for _, user := range users {
// 		fmt.Printf("%d | %s | %s\n",
// 			user.ID,
// 			user.Name,
// 			user.Email,
// 		)
// 	}
// }

// func (h *userHandler) detail() {

// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("User ID : ")

// 	text, _ := reader.ReadString('\n')

// 	id, _ := strconv.Atoi(strings.TrimSpace(text))

// 	user, err := h.uc.GetByID(id)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println()
// 	fmt.Println("ID    :", user.ID)
// 	fmt.Println("Name  :", user.Name)
// 	fmt.Println("Email :", user.Email)
// }

// func (h *userHandler) update() {

// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("ID : ")
// 	idText, _ := reader.ReadString('\n')

// 	id, _ := strconv.Atoi(strings.TrimSpace(idText))

// 	fmt.Print("New Name : ")
// 	name, _ := reader.ReadString('\n')

// 	fmt.Print("New Email : ")
// 	email, _ := reader.ReadString('\n')

// 	name = strings.TrimSpace(name)
// 	email = strings.TrimSpace(email)

// 	err := h.uc.Update(id, name, email)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println("User updated successfully.")
// }

// func (h *userHandler) delete() {

// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("User ID : ")

// 	text, _ := reader.ReadString('\n')

// 	id, _ := strconv.Atoi(strings.TrimSpace(text))

// 	err := h.uc.Delete(id)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println("User deleted successfully.")
// }

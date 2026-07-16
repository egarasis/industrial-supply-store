package handlers

import (
	"bufio"
	"fmt"
	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
	"os"
	"strings"

	"golang.org/x/term"
)

type userHandler struct {
	uc              domain.UserUsecase
	adminHandler    domain.AdminHandler
	customerHandler domain.CustomerHandler
}

func NewUserHandler(
	uc domain.UserUsecase,
	adminHandler domain.AdminHandler,
	customerHandler domain.CustomerHandler,
) domain.UserHandler {
	return &userHandler{
		uc:              uc,
		adminHandler:    adminHandler,
		customerHandler: customerHandler,
	}
}

func (h *userHandler) Run() {
	for {
		fmt.Println("\n===== Welcome To Industrial Tool Store =====")
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
	// To make it invisible when input
	passwordBytes, passErr := term.ReadPassword(int(os.Stdin.Fd()))
	if passErr != nil {
		fmt.Println("Error:", passErr)
		return
	}
	fmt.Println()

	fmt.Print("Role (ADMIN, CUSTOMER): ")
	role, _ := reader.ReadString('\n')

	email = strings.TrimSpace(email)
	password := strings.TrimSpace(string(passwordBytes))
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
	passwordBytes, passErr := term.ReadPassword(int(os.Stdin.Fd()))
	if passErr != nil {
		fmt.Println("Error:", passErr)
		return
	}
	fmt.Println()

	email = strings.TrimSpace(email)
	password := strings.TrimSpace(string(passwordBytes))

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

	if user.Role == entity.RoleCustomer {
		h.customerHandler.Run(user.ID)
	}
}

package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Register() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, _ := scanner.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := scanner.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Print("Enter role [bank/customer]: ")
	roleStr, _ := scanner.ReadString('\n')
	roleStr = strings.ToLower(strings.TrimSpace(roleStr))

	var role Role
	if roleStr == "bank" {
		role = RoleBank
	} else {
		role = RoleCustomer
	}

	user, err := h.service.Register(username, password, role)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User registered:", user.Username, "with role:", user.Role)
}

func (h *Handler) Login() *User {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, _ := scanner.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := scanner.ReadString('\n')
	password = strings.TrimSpace(password)

	user, err := h.service.Login(username, password)
	if err != nil {
		return nil
	}

	fmt.Println("Welcome", user.Username)
	return &user
}

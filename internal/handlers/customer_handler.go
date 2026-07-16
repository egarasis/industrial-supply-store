package handlers

import (
	"bufio"
	"context"
	"fmt"
	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
	"os"
	"strings"
)

type CustomerHandler struct {
	orderUC domain.OrderUsecase
	userUC  domain.UserUsecase // 1. Kita tambahkan dependency UserUsecase di sini
}

// 2. Constructor sekarang menerima userUC agar bisa update profile dengan benar
func NewCustomerHandler(orderUC domain.OrderUsecase, userUC domain.UserUsecase) domain.CustomerHandler {
	return &CustomerHandler{
		orderUC: orderUC,
		userUC:  userUC,
	}
}

func (h *CustomerHandler) Run(userID int) {
	for {
		fmt.Println("\n===== CUSTOMER MENU =====")
		fmt.Println("1. View Products")
		fmt.Println("2. Checkout")
		fmt.Println("3. My Orders")
		fmt.Println("4. Order Detail")
		fmt.Println("5. Update Profile")
		fmt.Println("0. Logout")

		fmt.Print("Choose : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			h.viewProducts()

		case 2:
			// 3. Diperbaiki: memanggil Checkout (C besar) untuk belanja asli
			h.Checkout(userID)

		case 3:
			h.getMyOrders(userID)

		case 4:
			h.getOrderDetail(userID)

		case 5:
			// 4. Diperbaiki: memanggil updateProfile (u kecil) untuk ganti nama/perusahaan
			h.updateProfile(userID)

		case 0:
			fmt.Println("Logout...")
			return

		default:
			fmt.Println("Invalid menu")
		}
	}
}

// 5. Diperbaiki: Nama diubah dari checkout menjadi updateProfile agar sesuai fungsinya
func (h *CustomerHandler) updateProfile(userID int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n===== UPDATE PROFILE =====")

	fmt.Print("Enter Full Name   : ")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)

	fmt.Print("Enter Company Name: ")
	companyName, _ := reader.ReadString('\n')
	companyName = strings.TrimSpace(companyName)

	// --- TAMBAHKAN 3 BARIS INPUT BARU INI ---
	fmt.Print("Enter Contact Name: ")
	contactName, _ := reader.ReadString('\n')
	contactName = strings.TrimSpace(contactName)

	fmt.Print("Enter Phone Number: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	fmt.Print("Enter Address     : ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)
	// ----------------------------------------

	ctx := context.Background()

	// Pastikan semua variabel baru dikirim ke usecase
	err := h.userUC.UpdateProfile(ctx, entity.UserProfile{
		UserID:      userID,
		FullName:    fullName,
		CompanyName: companyName,
		ContactName: contactName, // <-- masukkan ini
		Phone:       phone,       // <-- masukkan ini
		Address:     address,     // <-- masukkan ini
	})

	if err != nil {
		fmt.Println("Error updating profile:", err)
		return
	}

	fmt.Println("Profile updated successfully!")
}

func (h *CustomerHandler) Checkout(userID int) {
	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	var cart []entity.CartItem

	for {
		var productID int
		var qty int
		var choice string

		fmt.Print("Product ID : ")
		fmt.Fscanln(reader, &productID)

		fmt.Print("Quantity : ")
		fmt.Fscanln(reader, &qty)

		cart = append(cart, entity.CartItem{
			ProductID: productID,
			Quantity:  qty,
		})

		fmt.Print("Add another product? (y/n): ")
		fmt.Fscanln(reader, &choice)

		if choice == "n" || choice == "N" {
			break
		}
	}

	err := h.orderUC.Checkout(
		ctx,
		userID,
		cart,
	)

	if err != nil {
		fmt.Println("Checkout Failed :", err)
		return
	}

	fmt.Println("Checkout Success")
}

func (h *CustomerHandler) getMyOrders(userID int) {
	ctx := context.Background()

	orders, err := h.orderUC.GetMyOrders(
		ctx,
		userID,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("========== MY ORDERS ==========")

	for _, order := range orders {
		fmt.Printf(
			"Order ID : %d\n",
			order.ID,
		)
		fmt.Printf(
			"Total : %.2f\n",
			order.TotalPrice,
		)
		fmt.Printf(
			"Status : %s\n",
			order.Status,
		)
		fmt.Println("----------------------------")
	}
}

func (h *CustomerHandler) getOrderDetail(userID int) {
	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	var orderID int

	fmt.Print("Order ID : ")
	fmt.Fscanln(reader, &orderID)

	items, err := h.orderUC.GetOrderDetail(
		ctx,
		orderID,
		userID,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("====== ORDER DETAIL ======")

	var total float64

	for _, item := range items {
		fmt.Printf(
			"%s\n",
			item.ProductName,
		)
		fmt.Printf(
			"Price : %.2f\n",
			item.Price,
		)
		fmt.Printf(
			"Qty : %d\n",
			item.Quantity,
		)
		fmt.Printf(
			"Subtotal : %.2f\n",
			item.Subtotal,
		)
		total += item.Subtotal
		fmt.Println("----------------------")
	}
	fmt.Printf("TOTAL : %.2f\n", total)
}

func (h *CustomerHandler) viewProducts() {
	ctx := context.Background()

	products, err := h.orderUC.GetAllProducts(ctx)
	if err != nil {
		fmt.Println("Failed to get products:", err)
		return
	}

	if len(products) == 0 {
		fmt.Println("No products available.")
		return
	}

	fmt.Println("\n============== PRODUCT LIST ==============")
	fmt.Printf("%-5s %-25s %-15s %-10s %-8s\n",
		"ID", "Product", "Supplier", "Price", "Stock")
	fmt.Println("---------------------------------------------------------------")

	for _, product := range products {
		fmt.Printf(
			"%-5d %-25s %-15s Rp%-8.0f %-8d\n",
			product.ID,
			product.ProductName,
			product.SupplierName,
			product.Price,
			product.Stock,
		)
	}
	fmt.Println("---------------------------------------------------------------")
}

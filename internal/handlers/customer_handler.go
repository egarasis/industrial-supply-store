package handlers

import (
	"bufio"
	"context"
	"fmt"
	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
	"os"
)

type CustomerHandler struct {
	orderUC domain.OrderUsecase
}

func NewCustomerHandler(orderUC domain.OrderUsecase) domain.CustomerHandler {
	return &CustomerHandler{
		orderUC: orderUC,
	}
}

func (h *CustomerHandler) Run(userID int) {
	for {
		fmt.Println("\n===== CUSTOMER MENU =====")
		fmt.Println("1. View Products")
		fmt.Println("2. Checkout")
		fmt.Println("3. My Orders")
		fmt.Println("4. Order Detail")
		fmt.Println("0. Logout")

		fmt.Print("Choose : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			h.viewProducts()

		case 2:
			h.Checkout(userID)

		case 3:
			h.GetMyOrders(userID)

		case 4:
			h.GetOrderDetail()

		case 0:
			fmt.Println("Logout...")
			return

		default:
			fmt.Println("Invalid menu")
		}
	}
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

func (h *CustomerHandler) GetMyOrders(userID int) {

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

func (h *CustomerHandler) GetOrderDetail() {

	ctx := context.Background()

	reader := bufio.NewReader(os.Stdin)

	var orderID int

	fmt.Print("Order ID : ")
	fmt.Fscanln(reader, &orderID)

	items, err := h.orderUC.GetOrderDetail(
		ctx,
		orderID,
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

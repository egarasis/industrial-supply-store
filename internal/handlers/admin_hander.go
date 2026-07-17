package handlers

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type adminHandler struct {
	uc domain.AdminUsecase
}

func NewAdminHandler(uc domain.AdminUsecase) domain.AdminHandler {
	return &adminHandler{
		uc: uc,
	}
}

func (h *adminHandler) Run() {
	for {
		fmt.Println("\n===== ADMIN MENU =====")
		fmt.Println("1. List Products")
		fmt.Println("2. Add Product")
		fmt.Println("3. Update Product")
		fmt.Println("4. Delete Product")
		fmt.Println("\n======================")
		fmt.Println("5. List Orders")
		fmt.Println("6. Update Status Orders to Completed")
		fmt.Println("\n======================")
		fmt.Println("7. List Category")
		fmt.Println("8. Add Category")
		fmt.Println("9. Assign Category to Product")
		fmt.Println("\n======================")
		fmt.Println("10. Report Completed Orders")
		fmt.Println("11. Report Top Purcase Users")
		fmt.Println("12. Report Out of Stock Products")
		fmt.Println("0. Logout")

		fmt.Print("Choose : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			h.newListProducts()
		case 2:
			h.newCreateProduct()
		case 3:
			h.newUpdateProduct()
		case 4:
			h.newDeleteProduct()
		case 5:
			h.listOrders()
		case 6:
			h.updateOrderStatusToCompleted()
		case 7:
			h.listCategories()
		case 8:
			h.addCategory()
		case 9:
			h.assignCategory()
		case 10:
			h.reportCompletedOrder()
		case 11:
			h.topUserReport()
		case 12:
			h.outOfStockReport()
		case 0:
			fmt.Println("Logout...")
			return
		default:
			fmt.Println("Invalid menu")
		}
	}
}

func (h *adminHandler) listProducts() {
	fmt.Println("\n--- List Products ---")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	products, err := h.uc.ListProducts(ctx)
	if err != nil {
		fmt.Printf("Error get products: %v\n", err)
		return
	}

	if len(products) == 0 {
		fmt.Println("No products available.")
		return
	}

	fmt.Printf("%-5s | %-20s | %-15s | %-8s\n", "ID", "Name", "Price", "Stock")
	fmt.Println("-----------------------------------------------------------------")

	for _, p := range products {
		priceFormatted := formatRupiah(p.Price)
		fmt.Printf("%-5d | %-20s | %-15s | %-8d\n", p.ID, p.Name, priceFormatted, p.Stock)
	}
}

func (h *adminHandler) addProduct() {
	fmt.Println("\n--- Add Product ---")
	var name string
	var price float64
	var stock int

	fmt.Print("Product Name  : ")
	fmt.Scanln(&name)
	fmt.Print("Product Price : ")
	fmt.Scanln(&price)
	fmt.Print("Product Stock : ")
	fmt.Scanln(&stock)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	product := entity.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}

	err := h.uc.AddProduct(ctx, &product)
	if err != nil {
		fmt.Printf("Error add product: %v\n", err)
		return
	}

	fmt.Println("Product added successfully!")
}

func (h *adminHandler) updateProduct() {
	fmt.Println("\n--- Update Product ---")
	var id int
	var name string
	var price float64
	var stock int

	fmt.Print("Product ID to Update : ")
	fmt.Scanln(&id)
	fmt.Print("New Name             : ")
	fmt.Scanln(&name)
	fmt.Print("New Price            : ")
	fmt.Scanln(&price)
	fmt.Print("New Stock            : ")
	fmt.Scanln(&stock)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	product := entity.Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}

	err := h.uc.UpdateProduct(ctx, &product)
	if err != nil {
		fmt.Printf("Error update product: %v\n", err)
		return
	}

	fmt.Println("Product updated successfully!")
}

func (h *adminHandler) deleteProduct() {
	fmt.Println("\n--- Delete Product ---")
	var id int
	fmt.Print("Product ID to Delete : ")
	fmt.Scanln(&id)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.uc.DeleteProduct(ctx, id)
	if err != nil {
		fmt.Printf("Error delete product: %v\n", err)
		return
	}

	fmt.Println("Product deleted successfully!")
}

// Fungsi pembantu untuk format Rupiah
func formatRupiah(amount float64) string {
	s := fmt.Sprintf("%.0f", amount)
	var result []string
	n := len(s)
	for i := n; i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		result = append([]string{s[start:i]}, result...)
	}
	return "Rp " + strings.Join(result, ".")
}

func (h *adminHandler) listOrders() {

	ctx := context.Background()

	orders, err := h.uc.ListOrders(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(orders) == 0 {
		fmt.Println("No orders found.")
		return
	}

	fmt.Println("\n========================================== ALL ORDERS ==========================================")
	fmt.Printf("%-5s %-30s %-12s %-12s %-20s\n",
		"ID",
		"Email",
		"Total",
		"Status",
		"Order Date",
	)

	fmt.Println("-----------------------------------------------------------------------------------------------")

	for _, order := range orders {

		fmt.Printf("%-5d %-30s Rp%-10.0f %-12s %-20s\n",
			order.ID,
			order.Email,
			order.TotalPrice,
			order.Status,
			order.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}

	fmt.Println("-----------------------------------------------------------------------------------------------")
}

func (h *adminHandler) updateOrderStatusToCompleted() {

	ctx := context.Background()

	orders, err := h.uc.GetOrdersByStatus(ctx, entity.StatusPending)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(orders) == 0 {
		fmt.Println("No pending orders.")
		return
	}

	fmt.Println("\n==================== PENDING ORDERS ====================")
	fmt.Printf("%-8s %-30s %-12s %-12s\n",
		"ID",
		"Email",
		"Total",
		"Status",
	)

	for _, order := range orders {

		fmt.Printf("%-8d %-30s Rp%-10.0f %-12s\n",
			order.ID,
			order.Email,
			order.TotalPrice,
			order.Status,
		)
	}

	var orderID int

	fmt.Print("\nEnter Order ID to Complete : ")
	fmt.Scanln(&orderID)

	// Change to completed
	err = h.uc.UpdateOrderStatus(ctx, orderID, entity.StatusCompleted)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Order updated successfully.")
}

func (h *adminHandler) reportCompletedOrder() {

	ctx := context.Background()

	orders, err := h.uc.GetOrdersByStatus(ctx, entity.StatusCompleted)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(orders) == 0 {
		fmt.Println("No pending orders.")
		return
	}

	fmt.Println("\n==================== COMPLETED ORDERS ====================")
	fmt.Printf("%-8s %-30s %-12s %-12s\n",
		"ID",
		"Email",
		"Total",
		"Status",
	)

	for _, order := range orders {

		fmt.Printf("%-8d %-30s Rp%-10.0f %-12s\n",
			order.ID,
			order.Email,
			order.TotalPrice,
			order.Status,
		)
	}
}

func (h *adminHandler) newCreateProduct() {

	reader := bufio.NewReader(os.Stdin)
	var product entity.ProductWithSupplier

	fmt.Print("Supplier ID      : ")
	input, _ := reader.ReadString('\n')
	product.SupplierID, _ = strconv.Atoi(strings.TrimSpace(input))

	fmt.Print("Product Name     : ")
	product.ProductName, _ = reader.ReadString('\n')
	product.ProductName = strings.TrimSpace(product.ProductName)

	fmt.Print("Description      : ")
	product.Description, _ = reader.ReadString('\n')
	product.Description = strings.TrimSpace(product.Description)

	fmt.Print("Price            : ")
	input, _ = reader.ReadString('\n')
	product.Price, _ = strconv.ParseFloat(strings.TrimSpace(input), 64)

	fmt.Print("Stock            : ")
	input, _ = reader.ReadString('\n')
	product.Stock, _ = strconv.Atoi(strings.TrimSpace(input))

	err := h.uc.NewCreateProduct(context.Background(), product)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Product created successfully.")
}

func (h *adminHandler) newUpdateProduct() {
	reader := bufio.NewReader(os.Stdin)
	var product entity.ProductWithSupplier

	fmt.Print("Product ID       : ")
	input, _ := reader.ReadString('\n')
	product.ID, _ = strconv.Atoi(strings.TrimSpace(input))

	fmt.Print("Supplier ID      : ")
	input, _ = reader.ReadString('\n')
	product.SupplierID, _ = strconv.Atoi(strings.TrimSpace(input))

	fmt.Print("Product Name     : ")
	product.ProductName, _ = reader.ReadString('\n')
	product.ProductName = strings.TrimSpace(product.ProductName)

	fmt.Print("Description      : ")
	product.Description, _ = reader.ReadString('\n')
	product.Description = strings.TrimSpace(product.Description)

	fmt.Print("Price            : ")
	input, _ = reader.ReadString('\n')
	product.Price, _ = strconv.ParseFloat(strings.TrimSpace(input), 64)

	fmt.Print("Stock            : ")
	input, _ = reader.ReadString('\n')
	product.Stock, _ = strconv.Atoi(strings.TrimSpace(input))

	err := h.uc.NewUpdateProduct(context.Background(), product)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Product updated successfully.")
}

func (h *adminHandler) newDeleteProduct() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Product ID : ")

	input, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(input))

	err := h.uc.NewDeleteProduct(context.Background(), id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Product deleted successfully.")
}

func (h *adminHandler) newListProducts() {

	products, err := h.uc.NewListProducts(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(products) == 0 {
		fmt.Println("No products found.")
		return
	}

	fmt.Println("\n==================== PRODUCT LIST ====================")

	fmt.Printf("%-5s %-20s %-20s %-12s %-8s\n",
		"ID",
		"Product",
		"Supplier",
		"Price",
		"Stock",
	)

	fmt.Println("--------------------------------------------------------------")

	for _, p := range products {
		fmt.Printf(
			"%-5d %-20s %-20s %-12.2f %-8d\n",
			p.ID,
			p.ProductName,
			p.SupplierName,
			p.Price,
			p.Stock,
		)
	}
}

func (h *adminHandler) listCategories() {

	categories, err := h.uc.GetAllCategories(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n==================== CATEGORY LIST ====================")
	for _, c := range categories {
		fmt.Printf("%d. %s\n", c.ID, c.CategoryName)
	}
}

func (h *adminHandler) addCategory() {
	reader := bufio.NewReader(os.Stdin)
	var c entity.Category

	fmt.Print("Category Name : ")
	c.CategoryName, _ = reader.ReadString('\n')
	c.CategoryName = strings.TrimSpace(c.CategoryName)

	err := h.uc.CreateCategory(context.Background(), c)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Category created successfully.")
}

func (h *adminHandler) assignCategory() {
	reader := bufio.NewReader(os.Stdin)
	var pc entity.ProductCategory

	fmt.Print("Product ID : ")
	input, _ := reader.ReadString('\n')
	pc.ProductID, _ = strconv.Atoi(strings.TrimSpace(input))

	fmt.Print("Category ID : ")
	input, _ = reader.ReadString('\n')
	pc.CategoryID, _ = strconv.Atoi(strings.TrimSpace(input))

	err := h.uc.AssignCategory(context.Background(), pc)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Category assigned successfully.")
}

func (h *adminHandler) topUserReport() {

	users, err := h.uc.GetTopUser(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n========== USER REPORT ==========")

	fmt.Printf("%-5s %-20s %-15s %-15s\n",
		"ID",
		"Username",
		"Orders",
		"Total Spent")

	for _, u := range users {

		fmt.Printf("%-5d %-20s %-15d Rp %.2f\n",
			u.UserID,
			u.Email,
			u.TotalOrder,
			u.TotalSpent)
	}
}

func (h *adminHandler) outOfStockReport() {

	products, err := h.uc.GetOutOfStock(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n======= OUT OF STOCK =======")

	fmt.Printf("%-5s %-30s %-10s\n",
		"ID",
		"Product",
		"Stock")

	for _, p := range products {

		fmt.Printf("%-5d %-30s %-10d\n",
			p.ProductID,
			p.ProductName,
			p.Stock)
	}
}

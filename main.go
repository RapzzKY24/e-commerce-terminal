package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	admin "tubes_alpro/Admin"
	algorithmn "tubes_alpro/Algorithmn"
	cart "tubes_alpro/Cart"
	menu "tubes_alpro/Menu"
	order "tubes_alpro/Order"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n===== Selamat Datang di Aplikasi Restoran =====")
		fmt.Println("1. Masuk sebagai Customer")
		fmt.Println("2. Masuk sebagai Admin")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih mode: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			customerMode(scanner)
		case "2":
			adminMode(scanner)
		case "3":
			fmt.Println("Terima kasih telah menggunakan aplikasi kami!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func customerMode(scanner *bufio.Scanner) {
	myCart := cart.Cart{}

	for {
		fmt.Println("\n===== Mode Customer - Aplikasi Keranjang Belanja =====")
		fmt.Println("1. Lihat menu")
		fmt.Println("2. Tambah barang ke keranjang")
		fmt.Println("3. Hapus barang dari keranjang")
		fmt.Println("4. Perbarui jumlah barang")
		fmt.Println("5. Lihat keranjang")
		fmt.Println("6. Kosongkan keranjang")
		fmt.Println("7. Urutkan barang di keranjang")
		fmt.Println("8. Cari barang")
		fmt.Println("9. Checkout")
		fmt.Println("10. Kembali ke menu utama")
		fmt.Print("Pilih opsi: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			menu.DisplayMenu()
		case "2":
			addItemFromMenu(scanner, &myCart)
		case "3":
			removeItem(scanner, &myCart)
		case "4":
			updateItem(scanner, &myCart)
		case "5":
			viewCart(&myCart)
		case "6":
			myCart.ClearCart()
			fmt.Println("Keranjang berhasil dikosongkan!")
		case "7":
			sortCartItems(scanner, &myCart)
		case "8":
			searchItem(scanner, &myCart)
		case "9":
			checkout(scanner, &myCart)
		case "10":
			return
		default:
			fmt.Println("Opsi tidak valid. Silakan coba lagi.")
		}
	}
}

func adminMode(scanner *bufio.Scanner) {
	admin.AdminMenu(scanner)
}

func addItemFromMenu(scanner *bufio.Scanner, c *cart.Cart) {
	menu.DisplayMenu()

	fmt.Print("Masukkan ID menu yang ingin dipesan: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	menuItem, exists := menu.GetMenuByID(id)
	if !exists {
		fmt.Println("Menu tidak ditemukan.")
		return
	}

	if menuItem.Stok == 0 {
		fmt.Println("Maaf, menu ini sedang habis.")
		return
	}

	fmt.Printf("Menu: %s - Rp%d (Stok: %d)\n", menuItem.Nama, menuItem.Harga, menuItem.Stok)
	fmt.Print("Masukkan jumlah yang ingin dipesan: ")
	scanner.Scan()
	quantityStr := scanner.Text()
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity <= 0 {
		fmt.Println("Jumlah tidak valid.")
		return
	}

	if quantity > menuItem.Stok {
		fmt.Printf("Maaf, stok tidak mencukupi. Stok tersedia: %d\n", menuItem.Stok)
		return
	}

	// Konversi ke cart item dan tambahkan ke keranjang
	cartItem := menu.ConvertToCartItem(menuItem, quantity)
	c.AddItem(cartItem)

	// Update stok menu (kurangi stok)
	menu.UpdateStok(id, menuItem.Stok-quantity)

	fmt.Printf("%s x%d berhasil ditambahkan ke keranjang!\n", menuItem.Nama, quantity)
}

func addItem(scanner *bufio.Scanner, c *cart.Cart) {
	var name string
	var quantity, price int

	fmt.Print("Masukkan nama barang: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Print("Masukkan jumlah: ")
	scanner.Scan()
	quantity, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Masukkan harga per barang: ")
	scanner.Scan()
	price, _ = strconv.Atoi(scanner.Text())

	item := cart.Item{
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}

	c.AddItem(item)
	fmt.Println("Barang berhasil ditambahkan ke keranjang!")
}

func removeItem(scanner *bufio.Scanner, c *cart.Cart) {
	fmt.Print("Masukkan nama barang yang akan dihapus: ")
	scanner.Scan()
	name := scanner.Text()

	c.RemoveItem(name)
	fmt.Println("Barang berhasil dihapus dari keranjang!")
}

func updateItem(scanner *bufio.Scanner, c *cart.Cart) {
	fmt.Print("Masukkan nama barang yang akan diperbarui: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Masukkan jumlah baru: ")
	scanner.Scan()
	quantity, _ := strconv.Atoi(scanner.Text())

	c.UpdateItem(name, quantity)
	fmt.Println("Jumlah barang berhasil diperbarui!")
}

func viewCart(c *cart.Cart) {
	if len(c.Items) == 0 {
		fmt.Println("Keranjang Anda kosong.")
		return
	}

	fmt.Println("\n===== Keranjang Anda =====")
	totalPrice := 0

	for i, item := range c.Items {
		itemTotal := item.Price * item.Quantity
		totalPrice += itemTotal
		fmt.Printf("%d. %s - Jumlah: %d - Harga: %d - Total: %d\n",
			i+1, item.Name, item.Quantity, item.Price, itemTotal)
	}

	fmt.Printf("\nTotal Nilai Keranjang: %d\n", totalPrice)
}

func checkout(scanner *bufio.Scanner, c *cart.Cart) {
	if len(c.Items) == 0 {
		fmt.Println("Keranjang Anda kosong. Tidak dapat checkout.")
		return
	}

	fmt.Print("Masukkan nama Anda: ")
	scanner.Scan()
	customerName := scanner.Text()

	orderID := fmt.Sprintf("ORD-%d", time.Now().Unix())

	// Tambahkan transaksi ke log menu
	for _, item := range c.Items {
		// Cari menu berdasarkan nama untuk mendapatkan ID
		menuItems := menu.SearchMenuByName(item.Name)
		if len(menuItems) > 0 {
			// Ambil menu pertama yang cocok
			menuItem := menuItems[0]
			menu.TransaksiLog = append(menu.TransaksiLog, menu.Transaksi{
				IDMenu: menuItem.ID,
				Jumlah: item.Quantity,
			})
		}
	}

	newOrder := order.CreateOrder(orderID, customerName, *c)

	fmt.Println("\n===== Konfirmasi Pesanan =====")
	fmt.Printf("ID Pesanan: %s\n", newOrder.ID)
	fmt.Printf("Pelanggan: %s\n", newOrder.CustomerName)
	fmt.Printf("Tanggal: %s\n", newOrder.OrderDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("Status: %s\n", newOrder.Status)

	fmt.Println("\nBarang:")
	for i, item := range newOrder.Cart.Items {
		fmt.Printf("%d. %s - Jumlah: %d - Harga: Rp%d - Total: Rp%d\n",
			i+1, item.Name, item.Quantity, item.Price, item.Price*item.Quantity)
	}

	fmt.Printf("\nTotal Nilai Pesanan: Rp%d\n", newOrder.TotalPrice)
	fmt.Println("\nTerima kasih atas pesanan Anda!")
	fmt.Println("Pesanan Anda sedang diproses...")

	c.ClearCart()
}

func sortCartItems(scanner *bufio.Scanner, c *cart.Cart) {
	if len(c.Items) <= 1 {
		fmt.Println("Keranjang memiliki 1 atau kurang barang. Tidak perlu diurutkan.")
		return
	}

	fmt.Println("\n===== Opsi Pengurutan =====")
	fmt.Println("1. Urutkan berdasarkan harga (rendah ke tinggi)")
	fmt.Println("2. Urutkan berdasarkan harga (tinggi ke rendah)")
	fmt.Println("3. Urutkan berdasarkan nama (A-Z)")
	fmt.Print("Pilih opsi pengurutan: ")

	scanner.Scan()
	sortOption := scanner.Text()

	fmt.Println("\n===== Algoritma Pengurutan =====")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilih algoritma pengurutan: ")

	scanner.Scan()
	algorithmOption := scanner.Text()

	if sortOption == "1" || sortOption == "2" {
		prices := make([]int, len(c.Items))
		for i, item := range c.Items {
			prices[i] = item.Price
		}

		if algorithmOption == "1" {
			prices = algorithmn.SelectionSort(prices)
		} else {
			prices = algorithmn.InsertionSort(prices)
		}

		if sortOption == "2" {
			for i, j := 0, len(prices)-1; i < j; i, j = i+1, j-1 {
				prices[i], prices[j] = prices[j], prices[i]
			}
		}

		sortedItems := make([]cart.Item, 0)
		for _, price := range prices {
			for _, item := range c.Items {
				if item.Price == price {
					found := false
					for _, sortedItem := range sortedItems {
						if sortedItem.Name == item.Name && sortedItem.Price == item.Price {
							found = true
							break
						}
					}
					if !found {
						sortedItems = append(sortedItems, item)
					}
				}
			}
		}
		c.Items = sortedItems

	} else if sortOption == "3" {
		sort.Slice(c.Items, func(i, j int) bool {
			return c.Items[i].Name < c.Items[j].Name
		})
	}

	fmt.Println("Barang di keranjang berhasil diurutkan!")
	viewCart(c)
}

func searchItem(scanner *bufio.Scanner, c *cart.Cart) {
	if len(c.Items) == 0 {
		fmt.Println("Keranjang Anda kosong.")
		return
	}

	fmt.Println("\n===== Opsi Pencarian =====")
	fmt.Println("1. Cari berdasarkan nama")
	fmt.Println("2. Cari berdasarkan harga")
	fmt.Print("Pilih opsi pencarian: ")

	scanner.Scan()
	searchOption := scanner.Text()

	fmt.Println("\n===== Algoritma Pencarian =====")
	fmt.Println("1. Linear Search")
	fmt.Println("2. Binary Search (membutuhkan data terurut)")
	fmt.Print("Pilih algoritma pencarian: ")

	scanner.Scan()
	algorithmOption := scanner.Text()

	if searchOption == "1" {
		fmt.Print("Masukkan nama barang yang dicari: ")
		scanner.Scan()
		name := scanner.Text()

		if algorithmOption == "1" {
			found := false
			for i, item := range c.Items {
				if strings.EqualFold(item.Name, name) {
					fmt.Printf("\nBarang ditemukan di posisi %d:\n", i+1)
					fmt.Printf("Nama: %s, Jumlah: %d, Harga: %d\n",
						item.Name, item.Quantity, item.Price)
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Barang tidak ditemukan di keranjang.")
			}
		} else {
			fmt.Println("Binary search berdasarkan nama tidak diimplementasikan.")
		}
	} else if searchOption == "2" {
		fmt.Print("Masukkan harga yang dicari: ")
		scanner.Scan()
		price, _ := strconv.Atoi(scanner.Text())

		if algorithmOption == "1" {
			prices := make([]int, len(c.Items))
			for i, item := range c.Items {
				prices[i] = item.Price
			}

			index := algorithmn.LinearSearch(prices, price)
			if index != -1 {
				fmt.Printf("\nBarang ditemukan di posisi %d:\n", index+1)
				fmt.Printf("Nama: %s, Jumlah: %d, Harga: %d\n",
					c.Items[index].Name, c.Items[index].Quantity, c.Items[index].Price)
			} else {
				fmt.Println("Barang dengan harga tersebut tidak ditemukan di keranjang.")
			}
		} else if algorithmOption == "2" {
			prices := make([]int, len(c.Items))
			for i, item := range c.Items {
				prices[i] = item.Price
			}

			sort.Ints(prices)

			index := algorithmn.BinarySearch(prices, price)
			if index != -1 {
				fmt.Println("\nBarang dengan harga tersebut ada di keranjang.")
				for i, item := range c.Items {
					if item.Price == price {
						fmt.Printf("Posisi %d: %s, Jumlah: %d, Harga: %d\n",
							i+1, item.Name, item.Quantity, item.Price)
					}
				}
			} else {
				fmt.Println("Barang dengan harga tersebut tidak ditemukan di keranjang.")
			}
		}
	}
}

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
		fmt.Println("2. Urutkan menu")
		fmt.Println("3. Cari menu")
		fmt.Println("4. Tambah barang ke keranjang")
		fmt.Println("5. Hapus barang dari keranjang")
		fmt.Println("6. Perbarui jumlah barang")
		fmt.Println("7. Lihat keranjang")
		fmt.Println("8. Kosongkan keranjang")
		fmt.Println("9. Urutkan barang di keranjang")
		fmt.Println("10. Cari barang di keranjang")
		fmt.Println("11. Checkout")
		fmt.Println("12. Kembali ke menu utama")
		fmt.Print("Pilih opsi: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			menu.DisplayMenu()
		case "2":
			sortMenuItems(scanner)
		case "3":
			searchMenuItems(scanner)
		case "4":
			addItemFromMenu(scanner, &myCart)
		case "5":
			removeItem(scanner, &myCart)
		case "6":
			updateItem(scanner, &myCart)
		case "7":
			viewCart(&myCart)
		case "8":
			myCart.ClearCart()
			fmt.Println("Keranjang berhasil dikosongkan!")
		case "9":
			sortCartItems(scanner, &myCart)
		case "10":
			searchItem(scanner, &myCart)
		case "11":
			checkout(scanner, &myCart)
		case "12":
			return
		default:
			fmt.Println("Opsi tidak valid. Silakan coba lagi.")
		}
	}
}

func adminMode(scanner *bufio.Scanner) {
	admin.AdminMenu(scanner)
}

func sortMenuItems(scanner *bufio.Scanner) {
	menuItems := menu.GetAllMenuItems()
	if len(menuItems) <= 1 {
		fmt.Println("Menu memiliki 1 atau kurang item. Tidak perlu diurutkan.")
		return
	}

	fmt.Println("\n===== Opsi Pengurutan Menu =====")
	fmt.Println("1. Urutkan berdasarkan harga (rendah ke tinggi)")
	fmt.Println("2. Urutkan berdasarkan harga (tinggi ke rendah)")
	fmt.Println("3. Urutkan berdasarkan nama (A-Z)")
	fmt.Println("4. Urutkan berdasarkan stok (rendah ke tinggi)")
	fmt.Print("Pilih opsi pengurutan: ")

	scanner.Scan()
	sortOption := scanner.Text()

	fmt.Println("\n===== Algoritma Pengurutan =====")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilih algoritma pengurutan: ")

	scanner.Scan()
	algorithmOption := scanner.Text()

	var sortedItems []menu.MenuItem

	switch sortOption {
	case "1": // Harga rendah ke tinggi
		prices := make([]int, len(menuItems))
		for i, item := range menuItems {
			prices[i] = item.Harga
		}

		if algorithmOption == "1" {
			prices = algorithmn.SelectionSort(prices)
		} else {
			prices = algorithmn.InsertionSort(prices)
		}

		// Buat slice terurut berdasarkan harga
		for _, price := range prices {
			for _, item := range menuItems {
				if item.Harga == price {
					// Cek apakah sudah ada di sortedItems
					found := false
					for _, sortedItem := range sortedItems {
						if sortedItem.ID == item.ID {
							found = true
							break
						}
					}
					if !found {
						sortedItems = append(sortedItems, item)
						break
					}
				}
			}
		}

	case "2": // Harga tinggi ke rendah
		prices := make([]int, len(menuItems))
		for i, item := range menuItems {
			prices[i] = item.Harga
		}

		if algorithmOption == "1" {
			prices = algorithmn.SelectionSort(prices)
		} else {
			prices = algorithmn.InsertionSort(prices)
		}

		// Reverse untuk tinggi ke rendah
		for i, j := 0, len(prices)-1; i < j; i, j = i+1, j-1 {
			prices[i], prices[j] = prices[j], prices[i]
		}

		// Buat slice terurut berdasarkan harga
		for _, price := range prices {
			for _, item := range menuItems {
				if item.Harga == price {
					// Cek apakah sudah ada di sortedItems
					found := false
					for _, sortedItem := range sortedItems {
						if sortedItem.ID == item.ID {
							found = true
							break
						}
					}
					if !found {
						sortedItems = append(sortedItems, item)
						break
					}
				}
			}
		}

	case "3": // Nama A-Z
		sortedItems = make([]menu.MenuItem, len(menuItems))
		copy(sortedItems, menuItems)

		// Simple bubble sort untuk nama
		for i := 0; i < len(sortedItems)-1; i++ {
			for j := 0; j < len(sortedItems)-i-1; j++ {
				if sortedItems[j].Nama > sortedItems[j+1].Nama {
					sortedItems[j], sortedItems[j+1] = sortedItems[j+1], sortedItems[j]
				}
			}
		}

	case "4": // Stok rendah ke tinggi
		stoks := make([]int, len(menuItems))
		for i, item := range menuItems {
			stoks[i] = item.Stok
		}

		if algorithmOption == "1" {
			stoks = algorithmn.SelectionSort(stoks)
		} else {
			stoks = algorithmn.InsertionSort(stoks)
		}

		// Buat slice terurut berdasarkan stok
		for _, stok := range stoks {
			for _, item := range menuItems {
				if item.Stok == stok {
					// Cek apakah sudah ada di sortedItems
					found := false
					for _, sortedItem := range sortedItems {
						if sortedItem.ID == item.ID {
							found = true
							break
						}
					}
					if !found {
						sortedItems = append(sortedItems, item)
						break
					}
				}
			}
		}

	default:
		fmt.Println("Opsi pengurutan tidak valid.")
		return
	}

	fmt.Println("\n===== MENU TERURUT =====")
	for _, item := range sortedItems {
		status := "Tersedia"
		if item.Stok == 0 {
			status = "Habis"
		}
		fmt.Printf("ID: %d | %s | Rp%d | Stok: %d (%s)\n",
			item.ID, item.Nama, item.Harga, item.Stok, status)
	}
}

func searchMenuItems(scanner *bufio.Scanner) {
	menuItems := menu.GetAllMenuItems()
	if len(menuItems) == 0 {
		fmt.Println("Belum ada menu tersedia.")
		return
	}

	fmt.Println("\n===== Opsi Pencarian Menu =====")
	fmt.Println("1. Cari berdasarkan nama")
	fmt.Println("2. Cari berdasarkan harga")
	fmt.Println("3. Cari berdasarkan rentang harga")
	fmt.Print("Pilih opsi pencarian: ")

	scanner.Scan()
	searchOption := scanner.Text()

	switch searchOption {
	case "1": // Cari berdasarkan nama
		fmt.Print("Masukkan nama menu yang dicari: ")
		scanner.Scan()
		nama := scanner.Text()

		hasil := menu.SearchMenuByName(nama)
		if len(hasil) == 0 {
			fmt.Println("Menu tidak ditemukan.")
		} else {
			fmt.Println("\n===== HASIL PENCARIAN =====")
			for _, item := range hasil {
				status := "Tersedia"
				if item.Stok == 0 {
					status = "Habis"
				}
				fmt.Printf("ID: %d | %s | Rp%d | Stok: %d (%s)\n",
					item.ID, item.Nama, item.Harga, item.Stok, status)
			}
		}

	case "2": // Cari berdasarkan harga
		fmt.Print("Masukkan harga yang dicari: ")
		scanner.Scan()
		hargaStr := scanner.Text()
		harga, err := strconv.Atoi(hargaStr)
		if err != nil {
			fmt.Println("Format harga tidak valid.")
			return
		}

		fmt.Println("\n===== Algoritma Pencarian =====")
		fmt.Println("1. Linear Search")
		fmt.Println("2. Binary Search (membutuhkan data terurut)")
		fmt.Print("Pilih algoritma pencarian: ")

		scanner.Scan()
		algorithmOption := scanner.Text()

		if algorithmOption == "1" {
			// Linear search
			found := false
			fmt.Println("\n===== HASIL PENCARIAN =====")
			for _, item := range menuItems {
				if item.Harga == harga {
					status := "Tersedia"
					if item.Stok == 0 {
						status = "Habis"
					}
					fmt.Printf("ID: %d | %s | Rp%d | Stok: %d (%s)\n",
						item.ID, item.Nama, item.Harga, item.Stok, status)
					found = true
				}
			}
			if !found {
				fmt.Println("Menu dengan harga tersebut tidak ditemukan.")
			}
		} else if algorithmOption == "2" {
			// Binary search
			prices := make([]int, len(menuItems))
			for i, item := range menuItems {
				prices[i] = item.Harga
			}

			// Sort prices untuk binary search
			sort.Ints(prices)

			index := algorithmn.BinarySearch(prices, harga)
			if index != -1 {
				fmt.Println("\n===== HASIL PENCARIAN =====")
				fmt.Println("Menu dengan harga tersebut ditemukan:")
				for _, item := range menuItems {
					if item.Harga == harga {
						status := "Tersedia"
						if item.Stok == 0 {
							status = "Habis"
						}
						fmt.Printf("ID: %d | %s | Rp%d | Stok: %d (%s)\n",
							item.ID, item.Nama, item.Harga, item.Stok, status)
					}
				}
			} else {
				fmt.Println("Menu dengan harga tersebut tidak ditemukan.")
			}
		}

	case "3": // Cari berdasarkan rentang harga
		fmt.Print("Masukkan harga minimum: ")
		scanner.Scan()
		minStr := scanner.Text()
		minHarga, err := strconv.Atoi(minStr)
		if err != nil {
			fmt.Println("Format harga minimum tidak valid.")
			return
		}

		fmt.Print("Masukkan harga maksimum: ")
		scanner.Scan()
		maxStr := scanner.Text()
		maxHarga, err := strconv.Atoi(maxStr)
		if err != nil {
			fmt.Println("Format harga maksimum tidak valid.")
			return
		}

		if minHarga > maxHarga {
			fmt.Println("Harga minimum tidak boleh lebih besar dari harga maksimum.")
			return
		}

		found := false
		fmt.Println("\n===== HASIL PENCARIAN =====")
		for _, item := range menuItems {
			if item.Harga >= minHarga && item.Harga <= maxHarga {
				status := "Tersedia"
				if item.Stok == 0 {
					status = "Habis"
				}
				fmt.Printf("ID: %d | %s | Rp%d | Stok: %d (%s)\n",
					item.ID, item.Nama, item.Harga, item.Stok, status)
				found = true
			}
		}
		if !found {
			fmt.Printf("Menu dengan rentang harga Rp%d - Rp%d tidak ditemukan.\n", minHarga, maxHarga)
		}

	default:
		fmt.Println("Opsi pencarian tidak valid.")
	}
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

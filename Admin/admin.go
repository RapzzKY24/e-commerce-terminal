package admin

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	menu "tubes_alpro/Menu"
	order "tubes_alpro/Order"
)

// AdminMenu menampilkan menu utama admin
func AdminMenu(scanner *bufio.Scanner) {
	for {
		fmt.Println("\n===== MENU UTAMA ADMIN =====")
		fmt.Println("1. Kelola Menu Makanan/Minuman")
		fmt.Println("2. Update Stok Menu")
		fmt.Println("3. Lihat Riwayat Transaksi")
		fmt.Println("4. Lihat Proses Pesanan")
		fmt.Println("5. Kembali ke menu utama")
		fmt.Print("Pilih menu: ")

		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			kelolaMenu(scanner)
		case "2":
			updateStok(scanner)
		case "3":
			tampilkanTransaksi()
		case "4":
			tampilkanPesanan()
		case "5":
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

// kelolaMenu mengelola menu: tambah, lihat, edit, hapus
func kelolaMenu(scanner *bufio.Scanner) {
	for {
		fmt.Println("\n--- Kelola Menu ---")
		fmt.Println("1. Tambah Menu")
		fmt.Println("2. Lihat Menu")
		fmt.Println("3. Edit Menu")
		fmt.Println("4. Hapus Menu")
		fmt.Println("5. Kembali")
		fmt.Print("Pilih aksi: ")

		scanner.Scan()
		opsi := scanner.Text()

		switch opsi {
		case "1":
			tambahMenu(scanner)
		case "2":
			menu.DisplayMenu()
		case "3":
			editMenu(scanner)
		case "4":
			hapusMenu(scanner)
		case "5":
			return
		default:
			fmt.Println("Pilihan tidak dikenal.")
		}
	}
}

func tambahMenu(scanner *bufio.Scanner) {
	fmt.Print("Nama menu: ")
	scanner.Scan()
	nama := strings.TrimSpace(scanner.Text())
	
	if nama == "" {
		fmt.Println("Nama menu tidak boleh kosong.")
		return
	}

	fmt.Print("Harga: ")
	scanner.Scan()
	hargaStr := scanner.Text()

	fmt.Print("Stok awal: ")
	scanner.Scan()
	stokStr := scanner.Text()

	harga, stok, err := menu.ValidateMenuInput(hargaStr, stokStr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	menu.TambahMenu(nama, harga, stok)
	fmt.Println("Menu berhasil ditambahkan.")
}

func editMenu(scanner *bufio.Scanner) {
	fmt.Print("Masukkan ID menu yang ingin diedit: ")
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

	fmt.Printf("Menu saat ini: %s - Rp%d\n", menuItem.Nama, menuItem.Harga)
	
	fmt.Print("Nama baru: ")
	scanner.Scan()
	nama := strings.TrimSpace(scanner.Text())
	
	if nama == "" {
		fmt.Println("Nama menu tidak boleh kosong.")
		return
	}

	fmt.Print("Harga baru: ")
	scanner.Scan()
	hargaStr := scanner.Text()
	
	harga, err := strconv.Atoi(hargaStr)
	if err != nil || harga <= 0 {
		fmt.Println("Harga tidak valid.")
		return
	}

	if menu.EditMenu(id, nama, harga) {
		fmt.Println("Menu berhasil diperbarui.")
	} else {
		fmt.Println("Gagal memperbarui menu.")
	}
}

func hapusMenu(scanner *bufio.Scanner) {
	fmt.Print("Masukkan ID menu yang ingin dihapus: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	if menu.HapusMenu(id) {
		fmt.Println("Menu berhasil dihapus.")
	} else {
		fmt.Println("Menu tidak ditemukan.")
	}
}

func updateStok(scanner *bufio.Scanner) {
	fmt.Print("ID menu untuk update stok: ")
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

	fmt.Printf("Menu: %s - Stok saat ini: %d\n", menuItem.Nama, menuItem.Stok)
	fmt.Print("Masukkan stok baru: ")
	scanner.Scan()
	stokStr := scanner.Text()
	stok, err := strconv.Atoi(stokStr)
	if err != nil || stok < 0 {
		fmt.Println("Stok tidak valid.")
		return
	}

	if menu.UpdateStok(id, stok) {
		fmt.Println("Stok berhasil diperbarui.")
	} else {
		fmt.Println("Gagal memperbarui stok.")
	}
}

// tampilkanTransaksi menampilkan riwayat transaksi
func tampilkanTransaksi() {
	fmt.Println("\n===== RIWAYAT TRANSAKSI =====")
	if len(menu.TransaksiLog) == 0 {
		fmt.Println("Belum ada transaksi.")
		return
	}

	totalPendapatan := 0
	for i, trx := range menu.TransaksiLog {
		menuItem, exists := menu.GetMenuByID(trx.IDMenu)
		if exists {
			subtotal := menuItem.Harga * trx.Jumlah
			totalPendapatan += subtotal
			fmt.Printf("%d. %s x%d = Rp%d\n", i+1, menuItem.Nama, trx.Jumlah, subtotal)
		}
	}
	fmt.Printf("\nTotal Pendapatan: Rp%d\n", totalPendapatan)
}

// tampilkanPesanan menampilkan pesanan aktif
func tampilkanPesanan() {
	fmt.Println("\n===== PESANAN AKTIF =====")
	
	// Menampilkan pesanan dari order history
	orders := order.GetAllOrders()
	if len(orders) == 0 {
		fmt.Println("Belum ada pesanan.")
		return
	}

	for _, ord := range orders {
		fmt.Printf("\nPesanan ID: %s\n", ord.ID)
		fmt.Printf("Customer: %s\n", ord.CustomerName)
		fmt.Printf("Status: %s\n", ord.Status)
		fmt.Printf("Total: Rp%d\n", ord.TotalPrice)
		fmt.Println("Items:")
		for _, item := range ord.Cart.Items {
			fmt.Printf("  - %s x%d @ Rp%d\n", item.Name, item.Quantity, item.Price)
		}
	}
}

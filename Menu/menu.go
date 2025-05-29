package menu

import (
	"fmt"
	"strconv"
	"strings"
	cart "tubes_alpro/Cart"
)

// MenuItem untuk merepresentasikan item makanan atau minuman
type MenuItem struct {
	ID    int
	Nama  string
	Harga int
	Stok  int
}

// Transaksi untuk menyimpan data transaksi yang pernah dilakukan
type Transaksi struct {
	IDMenu int
	Jumlah int
}

var (
	MenuList     = make(map[int]MenuItem) // daftar menu (dengan ID sebagai key)
	TransaksiLog = []Transaksi{}
	NextMenuID   = 1
)

// Inisialisasi menu default
func init() {
	// Menambahkan beberapa menu default
	MenuList[1] = MenuItem{ID: 1, Nama: "Nasi Goreng", Harga: 15000, Stok: 20}
	MenuList[2] = MenuItem{ID: 2, Nama: "Mie Ayam", Harga: 12000, Stok: 15}
	MenuList[3] = MenuItem{ID: 3, Nama: "Es Teh", Harga: 5000, Stok: 30}
	MenuList[4] = MenuItem{ID: 4, Nama: "Ayam Bakar", Harga: 25000, Stok: 10}
	NextMenuID = 5
}

// TambahMenu menambahkan menu baru
func TambahMenu(nama string, harga int, stok int) {
	MenuList[NextMenuID] = MenuItem{
		ID:    NextMenuID,
		Nama:  nama,
		Harga: harga,
		Stok:  stok,
	}
	NextMenuID++
}

// DisplayMenu menampilkan semua menu yang tersedia
func DisplayMenu() {
	fmt.Println("\n===== DAFTAR MENU =====")
	if len(MenuList) == 0 {
		fmt.Println("Belum ada menu tersedia.")
		return
	}

	for _, item := range MenuList {
		status := "Tersedia"
		if item.Stok == 0 {
			status = "Habis"
		}
		fmt.Printf("ID: %d | %s | Rp%d | Stok: %d (%s)\n", 
			item.ID, item.Nama, item.Harga, item.Stok, status)
	}
}

// GetMenuByID mengambil menu berdasarkan ID
func GetMenuByID(id int) (MenuItem, bool) {
	menu, exists := MenuList[id]
	return menu, exists
}

// UpdateStok memperbarui stok menu
func UpdateStok(id int, stok int) bool {
	if menu, exists := MenuList[id]; exists {
		menu.Stok = stok
		MenuList[id] = menu
		return true
	}
	return false
}

// EditMenu mengedit menu yang sudah ada
func EditMenu(id int, nama string, harga int) bool {
	if menu, exists := MenuList[id]; exists {
		menu.Nama = nama
		menu.Harga = harga
		MenuList[id] = menu
		return true
	}
	return false
}

// HapusMenu menghapus menu dari daftar
func HapusMenu(id int) bool {
	if _, exists := MenuList[id]; exists {
		delete(MenuList, id)
		return true
	}
	return false
}

// PesanMenu memproses pesanan dari customer
func PesanMenu(id int, jumlah int) bool {
	if menu, exists := MenuList[id]; exists {
		if menu.Stok >= jumlah {
			// Kurangi stok
			menu.Stok -= jumlah
			MenuList[id] = menu
			
			// Tambahkan ke log transaksi
			TransaksiLog = append(TransaksiLog, Transaksi{
				IDMenu: id,
				Jumlah: jumlah,
			})
			return true
		}
	}
	return false
}

// ConvertToCartItem mengkonversi MenuItem ke cart.Item
func ConvertToCartItem(menuItem MenuItem, quantity int) cart.Item {
	return cart.Item{
		Name:     menuItem.Nama,
		Quantity: quantity,
		Price:    menuItem.Harga,
	}
}

// SearchMenuByName mencari menu berdasarkan nama
func SearchMenuByName(nama string) []MenuItem {
	var hasil []MenuItem
	for _, menu := range MenuList {
		if strings.Contains(strings.ToLower(menu.Nama), strings.ToLower(nama)) {
			hasil = append(hasil, menu)
		}
	}
	return hasil
}

// GetAllMenuItems mengembalikan semua menu dalam bentuk slice
func GetAllMenuItems() []MenuItem {
	var items []MenuItem
	for _, menu := range MenuList {
		items = append(items, menu)
	}
	return items
}

// ValidateMenuInput memvalidasi input menu
func ValidateMenuInput(hargaStr, stokStr string) (int, int, error) {
	harga, err := strconv.Atoi(hargaStr)
	if err != nil {
		return 0, 0, fmt.Errorf("format harga tidak valid")
	}
	
	stok, err := strconv.Atoi(stokStr)
	if err != nil {
		return 0, 0, fmt.Errorf("format stok tidak valid")
	}
	
	if harga <= 0 {
		return 0, 0, fmt.Errorf("harga harus lebih dari 0")
	}
	
	if stok < 0 {
		return 0, 0, fmt.Errorf("stok tidak boleh negatif")
	}
	
	return harga, stok, nil
}

package main

import (
	"bufio" // untuk input/output buffer, digunakan untuk membaca input baris per baris.
	"fmt"   // untuk format I/O, seperti mencetak ke konsol.
	"os"    // untuk interaksi dengan sistem operasi, contohnya membaca dari standard input.
	"strconv" // untuk konversi string ke tipe data numerik atau sebaliknya.
	"strings" // untuk manipulasi string.
)

const NMAX = 100

type User struct {
	username, password, nama string
}

type Buku struct {
	judul, penulis string
	halaman        int
}

type tabUser [NMAX]User
type tabBuku [NMAX]Buku

// scantext membaca satu baris input teks dari pengguna dan membersihkan spasi di awal/akhir.
func scantext() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// periksaAngka meminta pengguna untuk memasukkan angka melalui pesan yang diberikan (parameter 'inputPesan'),
// memastikan input adalah integer yang valid, dan mengembalikan nilai integer tersebut.
// Fungsi ini akan terus meminta input hingga valid dan juga mengatur warna teks konsol.
func periksaAngka(inputPesan string) int {
	var hasil int
	var input string
	var inputValid bool = false

	for !inputValid {
		fmt.Print("\033[0m")
		fmt.Print(inputPesan)
		input = scantext()
		fmt.Print("\033[96m")

		integerHasilParsing, err := strconv.Atoi(input)
		if err == nil {
			hasil = integerHasilParsing
			inputValid = true
		} else {
			fmt.Println("\n+------------------------------+")
			fmt.Println("|   Input tidak valid (angka)  |")
			fmt.Println("+------------------------------+")
		}
	}
	return hasil
}

// main adalah fungsi utama yang menginisialisasi data dan menjalankan loop menu utama aplikasi perpustakaan.
func main() {
	var daftarUser tabUser
	var daftarBuku tabBuku 
	var jumlahUser int
	var jumlahBuku int 
	var menu int
	var lanjut bool 

	fmt.Print("\033[96m")
	defer fmt.Print("\033[0m")

	jumlahUser = 1
	daftarUser[0] = User{"admin", "admin123", "Administrator"}

	jumlahBuku = 27
	daftarBuku[0] = Buku{"Harry Potter and the Sorcerer's Stone", "J.K. Rowling", 350}
	daftarBuku[1] = Buku{"The Hobbit", "J.R.R. Tolkien", 295}
	daftarBuku[2] = Buku{"Clean Code", "Robert C. Martin", 464}
	daftarBuku[3] = Buku{"To Kill a Mockingbird", "Harper Lee", 324}
	daftarBuku[4] = Buku{"1984", "George Orwell", 328}
	daftarBuku[5] = Buku{"Pride and Prejudice", "Jane Austen", 432}
	daftarBuku[6] = Buku{"The Great Gatsby", "F. Scott Fitzgerald", 180}
	daftarBuku[7] = Buku{"Moby Dick", "Herman Melville", 635}
	daftarBuku[8] = Buku{"War and Peace", "Leo Tolstoy", 1225}
	daftarBuku[9] = Buku{"The Catcher in the Rye", "J.D. Salinger", 224}
	daftarBuku[10] = Buku{"The Lord of the Rings", "J.R.R. Tolkien", 1178}
	daftarBuku[11] = Buku{"Jane Eyre", "Charlotte Brontë", 500}
	daftarBuku[12] = Buku{"Brave New World", "Aldous Huxley", 311}
	daftarBuku[13] = Buku{"The Chronicles of Narnia", "C.S. Lewis", 208}
	daftarBuku[14] = Buku{"Animal Farm", "George Orwell", 112}
	daftarBuku[15] = Buku{"Fahrenheit 451", "Ray Bradbury", 194}
	daftarBuku[16] = Buku{"Gone with the Wind", "Margaret Mitchell", 1037}
	daftarBuku[17] = Buku{"The Odyssey", "Homer", 541}
	daftarBuku[18] = Buku{"The Divine Comedy", "Dante Alighieri", 798}
	daftarBuku[19] = Buku{"Don Quixote", "Miguel de Cervantes", 863}
	daftarBuku[20] = Buku{"Frankenstein", "Mary Shelley", 280}
	daftarBuku[21] = Buku{"Alice's Adventures in Wonderland", "Lewis Carroll", 100}
	daftarBuku[22] = Buku{"One Hundred Years of Solitude", "Gabriel García Márquez", 417}
	daftarBuku[23] = Buku{"The Alchemist", "Paulo Coelho", 197}
	daftarBuku[24] = Buku{"Sapiens: A Brief History of Humankind", "Yuval Noah Harari", 464}
	daftarBuku[25] = Buku{"Atomic Habits", "James Clear", 320}
	daftarBuku[26] = Buku{"Omniscient Reader's Viewpoint", "King Dokja", 4951}

	lanjut = true
	for lanjut {
		fmt.Println("\n+------------------------------+")
		fmt.Println("|      SISTEM PERPUSTAKAAN     |")
		fmt.Println("+----+-------------------------+")
		fmt.Println("| No | Pilihan                 |")
		fmt.Println("+----+-------------------------+")
		fmt.Println("| 1  | Register                |")
		fmt.Println("| 2  | Login                   |")
		fmt.Println("| 3  | Keluar                  |")
		fmt.Println("+----+-------------------------+")

		menu = periksaAngka("Pilih menu: ")

		switch menu {
		case 1:
			jumlahUser = register(&daftarUser, jumlahUser)
		case 2:
			login(&daftarUser, jumlahUser, &daftarBuku, &jumlahBuku)
		case 3:
			fmt.Println("\n+------------------------------+")
			fmt.Println("|       Keluar dari program    |")
			fmt.Println("+------------------------------+")
			lanjut = false
		default:
			fmt.Println("\n+------------------------------+")
			fmt.Println("|     Pilihan tidak valid      |")
			fmt.Println("+------------------------------+")
		}
	}
}

// register sebagai proses pendaftaran pengguna baru dan mengembalikan jumlah pengguna
func register(A *tabUser, jumlahUser int) int {
	var username string
	var password string
	var nama string
	var userAda bool = false
	var i int

	fmt.Print("\033[0m")
	fmt.Print("Username: ")
	username = scantext()
	fmt.Print("\033[96m")

	i = 0
	for i < jumlahUser && !userAda {
		if A[i].username == username {
			userAda = true
		}
		i++
	}

	if userAda {
		fmt.Println("\n+---------------------------------+")
		fmt.Println("|        REGISTRASI GAGAL         |")
		fmt.Println("+---------------------------------+")
		fmt.Println("| Username sudah digunakan        |")
		fmt.Println("+---------------------------------+\n")
		return jumlahUser
	}

	fmt.Print("\033[0m")
	fmt.Print("Password: ")
	password = scantext()
	fmt.Print("\033[96m")

	fmt.Print("\033[0m")
	fmt.Print("Nama Lengkap: ")
	nama = scantext()
	fmt.Print("\033[96m")

	if jumlahUser >= NMAX {
		fmt.Println("\n+---------------------------------+")
		fmt.Println("|        REGISTRASI GAGAL         |")
		fmt.Println("+---------------------------------+")
		fmt.Println("| Kapasitas pengguna penuh        |")
		fmt.Println("+---------------------------------+\n")
		return jumlahUser
	}

	A[jumlahUser].username = username
	A[jumlahUser].password = password
	A[jumlahUser].nama = nama
	jumlahUser++

	fmt.Println("\n+---------------------------------+")
	fmt.Println("|       REGISTRASI BERHASIL       |")
	fmt.Println("+---------------------------------+")
	fmt.Printf("| Selamat datang, %-15s |\n", nama)
	fmt.Println("+---------------------------------+\n")

	return jumlahUser
}

// login digunakan untuk beralih ke menu perpustakaan jika berhasil.
func login(A *tabUser, jumlahUser int, daftarBuku *tabBuku, jumlahBuku *int) {
	var username string 
	var password string 
	var loginBerhasil bool = false
	var namaLogin string
	var userDitemukan bool = false
	var i int

	fmt.Print("\033[0m")
	fmt.Print("Username: ")
	username = scantext()
	fmt.Print("\033[96m")

	fmt.Print("\033[0m")
	fmt.Print("Password: ")
	password = scantext()
	fmt.Print("\033[96m")

	i = 0
	for i < jumlahUser && !userDitemukan {
		if A[i].username == username && A[i].password == password {
			namaLogin = A[i].nama
			fmt.Println("\n+---------------------------------+")
			fmt.Println("|       LOGIN BERHASIL            |")
			fmt.Println("+---------------------------------+")
			fmt.Printf("| Selamat datang, %-15s |\n", namaLogin)
			fmt.Println("+---------------------------------+\n")

			menuPerpustakaan(daftarBuku, jumlahBuku, namaLogin)
			loginBerhasil = true
			userDitemukan = true
		}
		i++
	}

	if !loginBerhasil {
		fmt.Println("\n+------------------------------+")
		fmt.Println("|          LOGIN GAGAL         |")
		fmt.Println("+------------------------------+")
		fmt.Println("| Username atau password salah |")
		fmt.Println("+------------------------------+")
	}
}

// menuPerpustakaan menampilkan menu-menu terkait data buku.
func menuPerpustakaan(daftarBuku *tabBuku, jumlahBuku *int, namaUser string) {
	var lanjutMenu bool = true
	var pilih int

	for lanjutMenu {
		fmt.Println("\n+-----------------------------------------------+")
		fmt.Println("|              MENU PERPUSTAKAAN                |")
		fmt.Println("+----+------------------------------------------+")
		fmt.Println("| No | Pilihan                                  |")
		fmt.Println("+----+------------------------------------------+")
		fmt.Println("| 1  | Tambah Buku                              |")
		fmt.Println("| 2  | Tampilkan Semua Buku                     |")
		fmt.Println("| 3  | Edit Data Buku                           |")
		fmt.Println("| 4  | Hapus Buku                               |")
		fmt.Println("| 5  | Cari Buku by Penulis                     |")
		fmt.Println("| 6  | Cari Buku by Judul                       |")
		fmt.Println("| 7  | Cari Buku by Halaman                     |")
		fmt.Println("| 8  | Urutkan Buku (Halaman Terbanyak)         |")
		fmt.Println("| 9  | Urutkan Buku (Halaman Tersedikit)        |")
		fmt.Println("| 10 | Urutkan Buku by Judul (A-Z)              |")
		fmt.Println("| 11 | Urutkan Buku by Judul (Z-A)              |")
		fmt.Println("| 12 | Tampilkan Statistik Halaman Buku         |")
		fmt.Println("| 13 | Logout                                   |")
		fmt.Println("+----+------------------------------------------+")

		pilih = periksaAngka("Pilih menu: ")

		switch pilih {
		case 1:
			tambahBuku(daftarBuku, jumlahBuku)
		case 2:
			tampilkanBuku(*daftarBuku, *jumlahBuku)
		case 3:
			editBuku(daftarBuku, *jumlahBuku)
		case 4:
			hapusBuku(daftarBuku, jumlahBuku)
		case 5:
			cariBukuByPenulis(*daftarBuku, *jumlahBuku) 
		case 6:
			cariBukuByJudul(*daftarBuku, *jumlahBuku)
		case 7:
			cariBukuByHalaman(daftarBuku, *jumlahBuku)
		case 8:
			urutanBukuTerbesar(daftarBuku, *jumlahBuku) 
		case 9:
			urutanBukuTerkecil(daftarBuku, *jumlahBuku) 
		case 10:
			urutkanBukuByJudulAsc(daftarBuku, *jumlahBuku) 
		case 11:
			urutkanBukuByJudulDesc(daftarBuku, *jumlahBuku)
		case 12:
			tampilkanStatistikBuku(*daftarBuku, *jumlahBuku) 
		case 13:
			fmt.Println("\n+----------------------------------------+")
			fmt.Println("| Logout berhasil. Kembali ke menu utama.|")
			fmt.Println("+----------------------------------------+")
			lanjutMenu = false
		default:
			fmt.Println("\n+--------------------------------------+")
			fmt.Println("|          Pilihan tidak valid.        |")
			fmt.Println("+--------------------------------------+")
		}
	}
}

// tambahBuku menambahkan data buku baru ke dalam array B.
func tambahBuku(B *tabBuku, jumlahBuku *int) {
	var judul string
	var penulis string
	var halaman int 
	var halamanValid bool = false
	
	if *jumlahBuku >= NMAX {
		fmt.Println("\n+------------------------------+")
		fmt.Println("|   Kapasitas buku penuh.    |")
		fmt.Println("+------------------------------+")
		return
	}

	fmt.Print("\033[0m")
	fmt.Print("Judul Buku: ")
	judul = scantext()
	fmt.Print("\033[96m")

	fmt.Print("\033[0m")
	fmt.Print("Penulis: ")
	penulis = scantext()
	fmt.Print("\033[96m")

	for !halamanValid {
		halaman = periksaAngka("Jumlah Halaman: ")
		if halaman > 0 {
			halamanValid = true
		} else {
			fmt.Println("\n+------------------------------+")
			fmt.Println("| Halaman harus angka positif  |")
			fmt.Println("+------------------------------+")
		}
	}

	B[*jumlahBuku].judul = judul
	B[*jumlahBuku].penulis = penulis
	B[*jumlahBuku].halaman = halaman
	*jumlahBuku++

	fmt.Println("\n+------------------------------+")
	fmt.Println("| Buku berhasil ditambahkan.   |")
	fmt.Println("+------------------------------+")
}

// tampilkanBuku mencetak semua data buku yang tersimpan dalam array B.
func tampilkanBuku(B tabBuku, jumlahBuku int) {
	if jumlahBuku == 0 {
		fmt.Println("\n+------------------------------+")
		fmt.Println("| Tidak ada buku dalam daftar. |")
		fmt.Println("+------------------------------+")
		return
	}

	fmt.Println("\n+----+------------------------------------------+-------------------------+--------+")
	fmt.Println("| No | Judul                                    | Penulis                 | Halaman|")
	fmt.Println("+----+------------------------------------------+-------------------------+--------+")
	for i := 0; i < jumlahBuku; i++ {
		fmt.Printf("| %-2d | %-40s | %-23s | %-6d |\n", i+1, B[i].judul, B[i].penulis, B[i].halaman)
	}
	fmt.Println("+----+------------------------------------------+-------------------------+--------+")
}

// editBuku digunakan pengguna untuk mengubah data buku yang sudah ada berdasarkan nomor urut.
func editBuku(B *tabBuku, jumlahBuku int) {
	var idx int
	var newJudul string
	var newPenulis string
	var halamanStr string
	var newHalaman int
	
	if jumlahBuku == 0 {
		fmt.Println("\n+-------------------------------+")
		fmt.Println("|  Tidak ada buku untuk diedit. |")
		fmt.Println("+-------------------------------+")
		return
	}

	tampilkanBuku(*B, jumlahBuku)
	fmt.Println("Tips: Masukkan nomor urut buku yang ingin diedit.")

	idx = periksaAngka("Pilih nomor buku: ")
	idx--

	if idx < 0 || idx >= jumlahBuku {
		fmt.Println("\n+------------------------------+")
		fmt.Println("|    Nomor buku tidak valid.   |")
		fmt.Println("+------------------------------+")
		return
	}

	fmt.Println("\nMengedit buku:", B[idx].judul)

	fmt.Print("\033[0m")
	fmt.Print("Judul baru ('" + B[idx].judul + "', kosongkan jika tidak diubah): ")
	newJudul = scantext()
	fmt.Print("\033[96m")

	fmt.Print("\033[0m")
	fmt.Print("Penulis baru ('" + B[idx].penulis + "', kosongkan jika tidak diubah): ")
	newPenulis = scantext()
	fmt.Print("\033[96m")

	fmt.Print("\033[0m")
	fmt.Print("Jumlah halaman baru ('" + strconv.Itoa(B[idx].halaman) + "', 0 atau kosong jika tidak diubah): ")
	halamanStr = scantext()
	fmt.Print("\033[96m")

	if newJudul != "" {
		B[idx].judul = newJudul
	}
	if newPenulis != "" {
		B[idx].penulis = newPenulis
	}

	if halamanStr != "" {
		if halamanStr == "0" {
		} else {
			nilaiHalamanBaru, err := strconv.Atoi(halamanStr)
			if err == nil {
				if nilaiHalamanBaru > 0 {
					newHalaman = nilaiHalamanBaru
					B[idx].halaman = newHalaman
				} else {
					fmt.Println("Input halaman tidak valid (harus > 0), halaman tidak diubah.")
				}
			} else {
				fmt.Println("Input halaman tidak valid (bukan angka), halaman tidak diubah.")
			}
		}
	}
	fmt.Println("\n+------------------------------+")
	fmt.Println("|Data buku berhasil diperbarui.|")
	fmt.Println("+------------------------------+")
}

// hapusBuku menghapus data buku dari array berdasarkan nomor urut yang dipilih pengguna.
func hapusBuku(B *tabBuku, jumlahBuku *int) {
	var idx int

	if *jumlahBuku == 0 {
		fmt.Println("\n+------------------------------+")
		fmt.Println("| Tidak ada buku untuk dihapus.|")
		fmt.Println("+------------------------------+")
		return
	}
	tampilkanBuku(*B, *jumlahBuku)
	fmt.Println("Tips: Masukkan nomor urut buku yang ingin dihapus.")

	idx = periksaAngka("Pilih nomor buku: ")
	idx--

	if idx < 0 || idx >= *jumlahBuku {
		fmt.Println("\n+------------------------------+")
		fmt.Println("|    Nomor buku tidak valid.   |")
		fmt.Println("+------------------------------+")
		return
	}

	for i := idx; i < *jumlahBuku-1; i++ {
		B[i] = B[i+1]
	}
	if *jumlahBuku > 0 {
		B[*jumlahBuku-1] = Buku{}
	}
	*jumlahBuku--
		fmt.Println("\n+------------------------------+")
		fmt.Println("|    Buku berhasil dihapus.    |")
		fmt.Println("+------------------------------+")
}

// cariBukuByPenulis mencari dan menampilkan buku berdasarkan nama penulis dengan sequential search.
func cariBukuByPenulis(B tabBuku, jumlahBuku int) {
	var penulisCari string
	var found bool = false
	var count int = 0

	if jumlahBuku == 0 {
		fmt.Println("\n+------------------------------+")
		fmt.Println("| Tidak ada buku dalam daftar. |")
		fmt.Println("+------------------------------+")
		return
	}

	fmt.Print("\033[0m")
	fmt.Print("Masukkan nama penulis yang ingin dicari (pencocokan parsial): ")
	penulisCari = scantext()
	fmt.Print("\033[96m")

	fmt.Println("\n+----+------------------------------------------+-------------------------+--------+")
	fmt.Println("| No | Judul                                    | Penulis                 | Halaman|")
	fmt.Println("+----+------------------------------------------+-------------------------+--------+")
	for i := 0; i < jumlahBuku; i++ {
		if strings.Contains(strings.ToLower(B[i].penulis), strings.ToLower(penulisCari)) {
			fmt.Printf("| %-2d | %-40s | %-23s | %-6d |\n", count+1, B[i].judul, B[i].penulis, B[i].halaman)
			found = true
			count++
		}
	}

	if !found {
		fmt.Println("+----+------------------------------------------+-------------------------+--------+")
		fmt.Printf("Buku dengan penulis yang mengandung '%s' tidak ditemukan.\n", penulisCari)
	} else {
		fmt.Println("+----+------------------------------------------+-------------------------+--------+")
	}
}

// cariBukuByJudul mencari dan menampilkan buku berdasarkan kata kunci judul dengan sequential search.
func cariBukuByJudul(dataBuku tabBuku, jumlahBuku int) {
	var judulCari string
	var keywordLower string
	var found bool = false
	var count int = 0

	if jumlahBuku == 0 {
		fmt.Println("\n+------------------------------+")
		fmt.Println("| Tidak ada buku dalam daftar. |")
		fmt.Println("+------------------------------+")
		return
	}

	fmt.Print("\033[0m")
	fmt.Print("Masukkan kata kunci judul buku yang ingin dicari: ")
	judulCari = scantext()
	fmt.Print("\033[96m")
	keywordLower = strings.ToLower(judulCari)

	fmt.Println("\n+----+------------------------------------------+-------------------------+--------+")
	fmt.Println("| No | Judul                                    | Penulis                 | Halaman|")
	fmt.Println("+----+------------------------------------------+-------------------------+--------+")
	for i := 0; i < jumlahBuku; i++ {
		if strings.Contains(strings.ToLower(dataBuku[i].judul), keywordLower) {
			fmt.Printf("| %-2d | %-40s | %-23s | %-6d |\n", count+1, dataBuku[i].judul, dataBuku[i].penulis, dataBuku[i].halaman)
			found = true
			count++
		}
	}
	if !found {
		fmt.Println("+----+------------------------------------------+-------------------------+--------+")
		fmt.Printf("Buku dengan judul yang mengandung '%s' tidak ditemukan.\n", judulCari)
	} else {
		fmt.Println("+----+------------------------------------------+-------------------------+--------+")
	}
}

// cariBukuByHalaman mencari buku dengan jumlah halaman yang tepat setelah mengurutkan data dengan binary search.
func cariBukuByHalaman(B *tabBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku dalam daftar.")
		return
	}

	target := periksaAngka("Masukkan jumlah halaman yang ingin dicari: ")
	tempB := *B
	urutkanBukuTerkecilInternal(&tempB, n)

	low, high := 0, n-1
	foundIndex := -1

	for low <= high && foundIndex == -1 {
		mid := (low + high) / 2
		if tempB[mid].halaman == target {
			foundIndex = mid
		} else if tempB[mid].halaman < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if foundIndex != -1 {
		fmt.Println("\n+----+------------------------------------------+-------------------------+--------+")
		fmt.Println("| No | Judul                                    | Penulis                 | Halaman|")
		fmt.Println("+----+------------------------------------------+-------------------------+--------+")
		fmt.Printf("| %-2d | %-40s | %-23s | %-6d |\n", 1, tempB[foundIndex].judul, tempB[foundIndex].penulis, tempB[foundIndex].halaman)
		fmt.Println("+----+------------------------------------------+-------------------------+--------+")
	} else {
		fmt.Println("\n+------------------------------------------------------+")
		fmt.Printf("| Buku dengan %-3d halaman tidak ditemukan.             |\n", target)
		fmt.Println("+------------------------------------------------------+")
	}
}

// urutkanBukuTerkecilInternal adalah bantuan untuk mengurutkan buku berdasarkan halaman secara menaik, tanpa cetak dengan selection sort.
func urutkanBukuTerkecilInternal(B *tabBuku, jumlahBuku int) {
	var minIdx int
	if jumlahBuku == 0 {
		return
	}
	for i := 0; i < jumlahBuku-1; i++ {
		minIdx = i
		for j := i + 1; j < jumlahBuku; j++ {
			if B[j].halaman < B[minIdx].halaman {
				minIdx = j
			}
		}
		B[i], B[minIdx] = B[minIdx], B[i]
	}
}

// urutanBukuTerbesar mengurutkan array buku B berdasarkan jumlah halaman secara menurun (terbanyak dulu) dengan insertion sort.
func urutanBukuTerbesar(B *tabBuku, jumlahBuku int) {
	var temp Buku
	if jumlahBuku == 0 {
		fmt.Println("\n+------------------------------+")
		fmt.Println("| Tidak ada buku dalam daftar. |")
		fmt.Println("+------------------------------+")
		return
	}
	for i := 1; i < jumlahBuku; i++ {
		temp = B[i]
		j := i - 1
		for j >= 0 && B[j].halaman < temp.halaman {
			B[j+1] = B[j]
			j--
		}
		B[j+1] = temp
	}
	fmt.Println("\n+-------------------------------------------------------------------------------+")
	fmt.Println("|      Buku diurutkan berdasarkan jumlah halaman (terbanyak ke terkecil).       |")
	fmt.Println("+-------------------------------------------------------------------------------+")
	tampilkanBuku(*B, jumlahBuku)
}

// urutanBukuTerkecil mengurutkan array buku berdasarkan jumlah halaman secara menaik (tersedikit dulu) dan menampilkannya dengan selection sort.
func urutanBukuTerkecil(B *tabBuku, jumlahBuku int) {
	if jumlahBuku == 0 {
		fmt.Println("\n+------------------------------+")
		fmt.Println("| Tidak ada buku dalam daftar. |")
		fmt.Println("+------------------------------+")
		return
	}
	urutkanBukuTerkecilInternal(B, jumlahBuku)
	fmt.Println("\n+-------------------------------------------------------------------------------+")
	fmt.Println("|      Buku diurutkan berdasarkan jumlah halaman (tersedikit ke terbanyak).     |")
	fmt.Println("+-------------------------------------------------------------------------------+")
	tampilkanBuku(*B, jumlahBuku)
}

// urutkanBukuByJudulAsc mengurutkan array buku berdasarkan judul secara menaik (A-Z) dengan insertion sort.
func urutkanBukuByJudulAsc(B *tabBuku, jumlahBuku int) {
	var temp Buku
	if jumlahBuku == 0 {
		fmt.Println("\n+------------------------------+")
		fmt.Println("| Tidak ada buku dalam daftar. |")
		fmt.Println("+------------------------------+")
		return
	}
	for i := 1; i < jumlahBuku; i++ {
		temp = B[i]
		j := i - 1
		for j >= 0 && strings.ToLower(B[j].judul) > strings.ToLower(temp.judul) {
			B[j+1] = B[j]
			j--
		}
		B[j+1] = temp
	}
	fmt.Println("\n+----------------------------------------------------------------------+")
	fmt.Println("|           Buku diurutkan berdasarkan judul (A-Z).                    |")
	fmt.Println("+----------------------------------------------------------------------+")
	tampilkanBuku(*B, jumlahBuku)
}

// urutkanBukuByJudulDesc mengurutkan array buku berdasarkan judul secara menurun (Z-A) dengan selection sort.
func urutkanBukuByJudulDesc(B *tabBuku, jumlahBuku int) {
	var maxIdx int
	if jumlahBuku == 0 {
		fmt.Println("\n+------------------------------+")
		fmt.Println("| Tidak ada buku dalam daftar. |")
		fmt.Println("+------------------------------+")
		return
	}
	for i := 0; i < jumlahBuku-1; i++ {
		maxIdx = i
		for j := i + 1; j < jumlahBuku; j++ {
			if strings.ToLower(B[j].judul) > strings.ToLower(B[maxIdx].judul) {
				maxIdx = j
			}
		}
		B[i], B[maxIdx] = B[maxIdx], B[i]
	}
	fmt.Println("\n+----------------------------------------------------------------------+")
	fmt.Println("|           Buku diurutkan berdasarkan judul (Z-A).                    |")
	fmt.Println("+----------------------------------------------------------------------+")
	tampilkanBuku(*B, jumlahBuku)
}

// tampilkanStatistikBuku menghitung dan menampilkan statistik dasar dari data buku yang ada.
func tampilkanStatistikBuku(B tabBuku, jumlahBuku int) {
	var totalHalaman int
	var halamanMin int
	var halamanMax int
	var rataRataHalaman float64

	if jumlahBuku == 0 {
		fmt.Println("\n+--------------------------------------+")
		fmt.Println("| Tidak ada buku untuk dianalisis.     |")
		fmt.Println("+--------------------------------------+")
		return
	}

	if jumlahBuku > 0 {
		halamanMax = B[0].halaman
		halamanMin = B[0].halaman
	}

	for i := 0; i < jumlahBuku; i++ {
		totalHalaman += B[i].halaman
		if B[i].halaman < halamanMin {
			halamanMin = B[i].halaman
		}
		if B[i].halaman > halamanMax {
			halamanMax = B[i].halaman
		}
	}

	if jumlahBuku > 0 {
		rataRataHalaman = float64(totalHalaman) / float64(jumlahBuku)
	}

	fmt.Println("\n+---------------------------------------+")
	fmt.Println("|         STATISTIK HALAMAN BUKU        |")
	fmt.Println("+---------------------------------------+")
	fmt.Printf("| Jumlah Total Buku: %-18d |\n", jumlahBuku)
	fmt.Printf("| Rata-rata Halaman: %-18.2f |\n", rataRataHalaman)
	if jumlahBuku > 0 {
		fmt.Printf("| Halaman Paling Sedikit: %-13d |\n", halamanMin)
		fmt.Printf("| Halaman Paling Banyak: %-14d |\n", halamanMax)
	}
	fmt.Println("+---------------------------------------+")
}
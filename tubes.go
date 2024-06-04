//Kelompok 8
//Muhammad Bayu Satrio dan Gilang Novadly

package main

import (
	"bufio" //untuk scan semua text jika ada spasi nya
	"fmt"
	"os"      //diperlukan untuk library bufio
	"strings" //untuk check tolower strings dengan EqualFold
)

const NMAX int = 100

type dataPenelitian struct {
	ketua          string
	anggota        [6]string
	prodi          string
	judul          string
	sumberDana     string
	luaran         string
	tahunKegiatan  int
	jumlahKegiatan int
}

type arrData [NMAX]dataPenelitian

func main() {
	var D arrData
	var n int
	Data(&D, &n)
	menu(&D, n)
}

// membuat tampilan menu
func header() {
	fmt.Println("----------------------------------")
	fmt.Println("|  Perguruan Tri Darma Tinggi 2  |")
	fmt.Println("----------------------------------")
	fmt.Println("Selamat datang di aplikasi kami")
}

// membuat opsi menu
func menu(D *arrData, n int) {
	header()
	var opsi int
	fmt.Println("\nMenu:")
	fmt.Println("1. Input data penelitian")
	fmt.Println("2. Ubah data penelitian")
	fmt.Println("3. Hapus data penelitian")
	fmt.Println("4. Lihat data penelitian")
	fmt.Println("5. Cari data penelitian")
	fmt.Println("6. Keluar")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&opsi)
	fmt.Println("----------------------------------")
	switch opsi {
	case 1:
		InputData(D, &n)
	case 2:
		UbahData(D, n)
	case 3:
		HapusData(D, &n)
	case 4:
		LihatData(*D, n)
	case 5:
		CariData(*D, n)
	case 6:
		fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
		return
	default:
		fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		fmt.Println()
		menu(D, n)
	}
}

// menambahkan data penelitian
func InputData(D *arrData, n *int) {
	if *n >= len(D) {
		fmt.Println("Data penuh, tidak bisa menambah data lagi.")
		return
	}
	fmt.Println("=== Input Data Penelitian ===")
	fmt.Print("Ketua : ")
	D[*n].ketua = scantext()
	for i := 0; i < 6; i++ {
		fmt.Printf("Anggota %d : ", i+1)
		D[*n].anggota[i] = scantext()
	}
	fmt.Print("Prodi : ")
	D[*n].prodi = scantext()
	fmt.Print("Judul : ")
	D[*n].judul = scantext()
	fmt.Print("Sumber Dana (internal/eksternal) : ")
	D[*n].sumberDana = scantext()
	fmt.Print("Luaran (publikasi/produk/seminar/pelatihan) : ")
	D[*n].luaran = scantext()
	fmt.Print("Tahun Kegiatan : ")
	fmt.Scanln(&D[*n].tahunKegiatan)
	fmt.Print("Jumlah Kegiatan : ")
	fmt.Scanln(&D[*n].jumlahKegiatan)
	*n++
	fmt.Println()
	fmt.Println("Data berhasil ditambahkan.")
	fmt.Println()
	menu(D, *n)
}

// untuk mencari judul (sequential search)
func FindIndexJudul(D arrData, n int, judul string) int {
	for i := 0; i < n; i++ {
		if strings.EqualFold(D[i].judul, judul) {
			return i
		}
	}
	return -1
}

// mengurutkan sesuai abjad
func selection_ketua(D *arrData, n int) {
	for i := 0; i < n-1; i++ {
		min := i
		for x := i + 1; x < n; x++ {
			if strings.ToLower(D[x].ketua) < strings.ToLower(D[min].ketua) {
				min = x
			}
		}
		D[i], D[min] = D[min], D[i]
	}
}

// untuk mencari Nama Ketua (Binary search)
func FindIndexKetua(D arrData, n int, ketua string) int {
	kiri := 0
	kanan := n - 1
	ketua = strings.ToLower(ketua)
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		nama_i := strings.ToLower(D[tengah].ketua)
		if ketua < nama_i {
			kanan = tengah - 1
		} else if ketua > nama_i {
			kiri = tengah + 1
		} else {
			return tengah
		}
	}
	return -1
}

// untuk mengubah data
func UbahData(D *arrData, n int) {
	var judul string
	var idx int
	if n == 0 {
		fmt.Println("Data masih kosong.")
		fmt.Println()
	} else {
		fmt.Println("=== Ubah Data Penelitian ===")
		fmt.Print("Masukkan judul penelitian yang akan diubah : ")
		judul = scantext()
		idx = FindIndexJudul(*D, n, judul)
		if idx != -1 {
			fmt.Println("Data ditemukan. Masukkan data baru.")
			fmt.Print("Ketua : ")
			D[idx].ketua = scantext()
			for j := 0; j < 6; j++ {
				fmt.Printf("Anggota %d : ", j+1)
				D[idx].anggota[j] = scantext()
			}
			fmt.Print("Prodi : ")
			D[idx].prodi = scantext()
			fmt.Print("Judul: ")
			D[idx].judul = scantext()
			fmt.Print("Sumber Dana (internal/eksternal) : ")
			D[idx].sumberDana = scantext()
			fmt.Print("Luaran (publikasi/produk/seminar/pelatihan) : ")
			D[idx].luaran = scantext()
			fmt.Print("Tahun Kegiatan : ")
			fmt.Scanln(&D[idx].tahunKegiatan)
			fmt.Print("Jumlah Kegiatan : ")
			fmt.Scanln(&D[idx].jumlahKegiatan)
			fmt.Println("Data berhasil diubah.")
			fmt.Println()
		} else {
			fmt.Println("Data tidak ditemukan.")
			fmt.Println()
		}
	}
	menu(D, n)
}

// untuk hapus data
func HapusData(D *arrData, n *int) {
	var judul string
	var idx int
	if *n == 0 {
		fmt.Println("Data masih kosong.")
		fmt.Println()
	} else {
		fmt.Println("=== Hapus Data Penelitian ===")
		fmt.Print("Masukkan judul penelitian yang akan dihapus : ")
		judul = scantext()
		idx = FindIndexJudul(*D, *n, judul)
		if idx != -1 {
			for j := idx; j < *n-1; j++ {
				D[j] = D[j+1]
			}
			*n--
			fmt.Println("Data berhasil dihapus.")
			fmt.Println()
		} else {
			fmt.Println("Data tidak ditemukan.")
			fmt.Println()
		}
	}
	menu(D, *n)
}

func LihatData(D arrData, n int) {
	fmt.Println("=== Cari Data Penelitian ===")
	var opsi int
	fmt.Println("\nMenu:")
	fmt.Println("1. Lihat semua data terurut sesuai Tahun Kegiatan")
	fmt.Println("2. Lihat semua data terurut sesuai Jumlah Kegiatan")
	fmt.Println("3. Lihat semua data Nama Ketua") // Print semua nama ketuanya aja
	fmt.Println("4. Lihat semua data Nama Judul") // Print semua nama judul nya aja
	fmt.Println("5. Lihat semua data")
	fmt.Println("6. Back")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&opsi)
	fmt.Println("----------------------------------")
	switch opsi {
	case 1:
		LihatData_Tahun(D, n)
	case 2:
		LihatData_Jumlah(D, n)
	case 3:
		LihatData_Semua_Ketua(D, n)
	case 4:
		LihatData_Semua_Judul(D, n)
	case 5:
		LihatData_Semua(D, n)
	case 6:
		menu(&D, n)
	default:
		fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		fmt.Println()
		LihatData(D, n)
	}
}

func LihatData_Tahun(D arrData, n int) {
	fmt.Println("=== Cari Data Penelitian ===")
	var opsi int
	fmt.Println("\nMenu:")
	fmt.Println("1. Lihat semua data Tahun Terbaru - Terlama") // descending 2024, 2023, 2022, 2021
	fmt.Println("2. Lihat semua data Tahun Terlama - Terbaru") // ascending 2021, 2022, 2023, 2024
	fmt.Println("3. Back")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&opsi)
	fmt.Println("----------------------------------")
	switch opsi {
	case 1:
		selection_tahun(&D, n)
		LihatData_Tahun_All(D, n)
	case 2:
		insertion_tahun(&D, n)
		LihatData_Tahun_All(D, n)
	case 3:
		LihatData(D, n)
	default:
		fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		fmt.Println()
		LihatData_Tahun(D, n)
	}
}

func selection_tahun(D *arrData, n int) { // SELECTION SORTING DESCENDING
	for i := 0; i < n; i++ {
		max := i
		for x := i; x < n; x++ {
			if D[max].tahunKegiatan < D[x].tahunKegiatan {
				max = x
			}
		}
		D[i], D[max] = D[max], D[i]
	}
}

func insertion_tahun(D *arrData, n int) { // INSERTION SORTING ASCENDING
	for i := 1; i < n; i++ {
		save := D[i]
		j := i
		for j > 0 && save.tahunKegiatan < D[j-1].tahunKegiatan {
			D[j] = D[j-1]
			j = j - 1
		}
		D[j] = save
	}
}

func LihatData_Tahun_All(D arrData, n int) {
	if n == 0 {
		fmt.Println("Data masih kosong.")
		fmt.Println()
	} else {
		fmt.Println("=== Lihat Semua Data Penelitian ===")
		for i := 0; i < n; i++ {
			fmt.Printf("Data %d:\n", i+1)
			fmt.Printf("Ketua: %s\n", D[i].ketua)
			fmt.Print("Anggota: ")
			for x := 0; x < len(D[i].anggota); x++ {
				fmt.Print("\n   ", x+1, ". ", D[i].anggota[x])
			}
			fmt.Println()
			fmt.Printf("Prodi: %s\n", D[i].prodi)
			fmt.Printf("Judul: %s\n", D[i].judul)
			fmt.Printf("Sumber Dana: %s\n", D[i].sumberDana)
			fmt.Printf("Luaran: %s\n", D[i].luaran)
			fmt.Printf("Tahun Kegiatan: %d\n", D[i].tahunKegiatan)
			fmt.Printf("Jumlah Kegiatan: %d\n", D[i].jumlahKegiatan)
			fmt.Println("----------------------------------")
		}
	}
	LihatData_Tahun(D, n)
}

func LihatData_Jumlah(D arrData, n int) {
	fmt.Println("=== Cari Data Penelitian ===")
	var opsi int
	fmt.Println("\nMenu:")
	fmt.Println("1. Lihat semua data Jumlah Kegiatan Terbanyak - Tersedikit") // descending 6, 5, 4, 3, 2, 1
	fmt.Println("2. Lihat semua data Jumlah Kegiatan Tersedikit - Terbanyak") // ascending 1, 2, 3, 4, 5, 6
	fmt.Println("3. Back")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&opsi)
	fmt.Println("----------------------------------")
	switch opsi {
	case 1:
		selection_Jumlah(&D, n)
		LihatData_Jumlah_All(D, n)
	case 2:
		insertion_Jumlah(&D, n)
		LihatData_Jumlah_All(D, n)
	case 3:
		LihatData(D, n)
	default:
		fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		fmt.Println()
		LihatData_Jumlah(D, n)
	}
}

func selection_Jumlah(D *arrData, n int) { // SELECTION SORTING DESCENDING
	for i := 0; i < n; i++ {
		max := i
		for x := i; x < n; x++ {
			if D[max].jumlahKegiatan < D[x].jumlahKegiatan {
				max = x
			}
		}
		D[i], D[max] = D[max], D[i]
	}
}

func insertion_Jumlah(D *arrData, n int) { // INSERTION SORTING ASCENDING
	for i := 1; i < n; i++ {
		save := D[i]
		j := i
		for j > 0 && save.jumlahKegiatan < D[j-1].jumlahKegiatan {
			D[j] = D[j-1]
			j = j - 1
		}
		D[j] = save
	}
}

func LihatData_Jumlah_All(D arrData, n int) {
	if n == 0 {
		fmt.Println("Data masih kosong.")
		fmt.Println()
	} else {
		fmt.Println("=== Lihat Semua Data Penelitian ===")
		for i := 0; i < n; i++ {
			fmt.Printf("Data %d:\n", i+1)
			fmt.Printf("Ketua: %s\n", D[i].ketua)
			fmt.Print("Anggota: ")
			for x := 0; x < len(D[i].anggota); x++ {
				fmt.Print("\n   ", x+1, ". ", D[i].anggota[x])
			}
			fmt.Println()
			fmt.Printf("Prodi: %s\n", D[i].prodi)
			fmt.Printf("Judul: %s\n", D[i].judul)
			fmt.Printf("Sumber Dana: %s\n", D[i].sumberDana)
			fmt.Printf("Luaran: %s\n", D[i].luaran)
			fmt.Printf("Tahun Kegiatan: %d\n", D[i].tahunKegiatan)
			fmt.Printf("Jumlah Kegiatan: %d\n", D[i].jumlahKegiatan)
			fmt.Println("----------------------------------")
		}
	}
	LihatData_Jumlah(D, n)
}

func LihatData_Semua_Ketua(D arrData, n int) {
	if n == 0 {
		fmt.Println("Data masih kosong.")
		fmt.Println()
	} else {
		fmt.Println("=== Lihat Semua Data Ketua Penelitian ===")
		for i := 0; i < n; i++ {
			fmt.Printf("Data %d:\n", i+1)
			fmt.Printf("Ketua : %s\n", D[i].ketua)
		}
		fmt.Println("----------------------------------")
	}
	LihatData(D, n)
}

func LihatData_Semua_Judul(D arrData, n int) {
	if n == 0 {
		fmt.Println("Data masih kosong.")
		fmt.Println()
	} else {
		fmt.Println("=== Lihat Semua Data Judul Penelitian ===")
		for i := 0; i < n; i++ {
			fmt.Printf("Data %d:\n", i+1)
			fmt.Printf("Judul : %s\n", D[i].judul)
		}
		fmt.Println("----------------------------------")
	}
	LihatData(D, n)
}

// Buat nampilin semua data yang ada
func LihatData_Semua(D arrData, n int) {
	if n == 0 {
		fmt.Println("Data masih kosong.")
		fmt.Println()
	} else {
		fmt.Println("=== Lihat Semua Data Penelitian ===")
		for i := 0; i < n; i++ {
			fmt.Printf("Data %d:\n", i+1)
			fmt.Printf("Ketua: %s\n", D[i].ketua)
			fmt.Print("Anggota: ")
			for x := 0; x < len(D[i].anggota); x++ {
				fmt.Print("\n   ", x+1, ". ", D[i].anggota[x])
			}
			fmt.Println()
			fmt.Printf("Prodi: %s\n", D[i].prodi)
			fmt.Printf("Judul: %s\n", D[i].judul)
			fmt.Printf("Sumber Dana: %s\n", D[i].sumberDana)
			fmt.Printf("Luaran: %s\n", D[i].luaran)
			fmt.Printf("Tahun Kegiatan: %d\n", D[i].tahunKegiatan)
			fmt.Printf("Jumlah Kegiatan: %d\n", D[i].jumlahKegiatan)
			fmt.Println("----------------------------------")
		}
	}
	LihatData(D, n)
}

func CariData(D arrData, n int) {
	fmt.Println("=== Cari Data Penelitian ===")
	var opsi int
	fmt.Println("\nMenu:")
	fmt.Println("1. Cari semua data berdasarkan tahun kegiatan")
	fmt.Println("2. Cari semua data berdasarkan Prodi")
	fmt.Println("3. Cari data dengan Judul Penelitian")      //Sequential search
	fmt.Println("4. Cari data dengan Nama Ketua Penelitian") //Binary Search
	fmt.Println("5. Back")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&opsi)
	fmt.Println("----------------------------------")
	switch opsi {
	case 1:
		CariData_tahun(D, n)
	case 2:
		CariData_Prodi(D, n)
	case 3:
		CariData_Judul(D, n)
	case 4:
		CariData_Ketua(D, n)
	case 5:
		menu(&D, n)
	default:
		fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		fmt.Println()
		CariData(D, n)
	}
}

func CariData_Ketua(D arrData, n int) {
	var ketua string
	var idx int
	if n == 0 {
		fmt.Println("Data masih kosong.")
		fmt.Println()
	} else {
		fmt.Println("=== Cari Data Penelitian ===")
		fmt.Print("Masukkan Nama Ketua Penelitian yang ingin dicari: ")
		ketua = scantext()
		selection_ketua(&D, n) // Di urutkan sesuai abjad sebelum di cari index ketua dengan binary search
		idx = FindIndexKetua(D, n, ketua)
		if idx != -1 {
			fmt.Printf("Data %d:\n", idx+1)
			fmt.Printf("Ketua : %s\n", D[idx].ketua)
			fmt.Print("Anggota: ")
			for x := 0; x < len(D[idx].anggota); x++ {
				fmt.Print("\n   - ", D[idx].anggota[x])
			}
			fmt.Println()
			fmt.Printf("Prodi : %s\n", D[idx].prodi)
			fmt.Printf("Judul : %s\n", D[idx].judul)
			fmt.Printf("Sumber Dana : %s\n", D[idx].sumberDana)
			fmt.Printf("Luaran : %s\n", D[idx].luaran)
			fmt.Printf("Tahun Kegiatan : %d\n", D[idx].tahunKegiatan)
			fmt.Printf("Jumlah Kegiatan: %d\n", D[idx].jumlahKegiatan)
			fmt.Println()
		} else {
			fmt.Println("Data tidak ditemukan.")
			fmt.Println()
		}
	}
	CariData(D, n)
}

func CariData_Judul(D arrData, n int) {
	var judul string
	var idx int
	if n == 0 {
		fmt.Println("Data masih kosong.")
		fmt.Println()
	} else {
		fmt.Println("=== Cari Data Penelitian ===")
		fmt.Print("Masukkan Judul Penelitian yang akan ingin di cari : ")
		judul = scantext()
		idx = FindIndexJudul(D, n, judul)
		if idx != -1 {
			fmt.Printf("Data %d:\n", idx+1)
			fmt.Printf("Ketua : %s\n", D[idx].ketua)
			fmt.Print("Anggota: ")
			for x := 0; x < len(D[idx].anggota); x++ {
				fmt.Print("\n   - ", D[idx].anggota[x])
			}
			fmt.Println()
			fmt.Printf("Prodi : %s\n", D[idx].prodi)
			fmt.Printf("Judul : %s\n", D[idx].judul)
			fmt.Printf("Sumber Dana : %s\n", D[idx].sumberDana)
			fmt.Printf("Luaran : %s\n", D[idx].luaran)
			fmt.Printf("Tahun Kegiatan : %d\n", D[idx].tahunKegiatan)
			fmt.Printf("Jumlah Kegiatan: %d\n", D[idx].jumlahKegiatan)
			fmt.Println()
		} else {
			fmt.Println("Data tidak ditemukan.")
			fmt.Println()
		}
	}
	CariData(D, n)
}

// untuk mencari data menggunakan tahun kegiatan
func CariData_tahun(D arrData, n int) {
	var tahun, count int
	if n == 0 {
		fmt.Println("Data masih kosong.")
		fmt.Println()
	} else {
		fmt.Println("=== Cari Data Penelitian ===")
		fmt.Print("Masukkan Tahun Kegiatan : ")
		fmt.Scanln(&tahun)
		for i := 0; i < n; i++ {
			if D[i].tahunKegiatan == tahun {
				fmt.Printf("Data %d:\n", i+1)
				fmt.Printf("Ketua : %s\n", D[i].ketua)
				fmt.Print("Anggota : ")
				for x := 0; x < len(D[i].anggota); x++ {
					fmt.Print("\n   - ", D[i].anggota[x])
				}
				fmt.Println()
				fmt.Printf("Prodi : %s\n", D[i].prodi)
				fmt.Printf("Judul : %s\n", D[i].judul)
				fmt.Printf("Sumber Dana : %s\n", D[i].sumberDana)
				fmt.Printf("Luaran : %s\n", D[i].luaran)
				fmt.Printf("Tahun Kegiatan : %d\n", D[i].tahunKegiatan)
				fmt.Printf("Jumlah Kegiatan: %d\n", D[i].jumlahKegiatan)
				fmt.Println()
				fmt.Println("----------------------------------")
				count++
			}
		}
		if count == 0 {
			fmt.Println("Data tidak ditemukan.")
			fmt.Println()
		}
	}
	CariData(D, n)
}

func CariData_Prodi(D arrData, n int) {
	var prodi string
	var count int
	if n == 0 {
		fmt.Println("Data masih kosong.")
		fmt.Println()
	} else {
		fmt.Println("=== Cari Data Penelitian ===")
		fmt.Print("Masukkan Prodi yang ingin di cari : ")
		prodi = scantext()
		for i := 0; i < n; i++ {
			if strings.EqualFold(D[i].prodi, prodi) {
				fmt.Printf("Data %d:\n", i+1)
				fmt.Printf("Ketua : %s\n", D[i].ketua)
				fmt.Print("Anggota : ")
				for x := 0; x < len(D[i].anggota); x++ {
					fmt.Print("\n   - ", D[i].anggota[x])
				}
				fmt.Println()
				fmt.Printf("Prodi : %s\n", D[i].prodi)
				fmt.Printf("Judul : %s\n", D[i].judul)
				fmt.Printf("Sumber Dana : %s\n", D[i].sumberDana)
				fmt.Printf("Luaran : %s\n", D[i].luaran)
				fmt.Printf("Tahun Kegiatan : %d\n", D[i].tahunKegiatan)
				fmt.Printf("Jumlah Kegiatan: %d\n", D[i].jumlahKegiatan)
				fmt.Println()
				fmt.Println("----------------------------------")
				count++
			}
		}
		if count == 0 {
			fmt.Println("Data tidak ditemukan.")
			fmt.Println()
		}
	}
	CariData(D, n)
}

func Data(D *arrData, n *int) {
	D[*n].ketua = "Bayu Supratman"
	D[*n].anggota = [6]string{"Ade Rainhard", "Rizki Nata", "Arief Ridwansyah", "Alif Nelson", "Kenzo Paramarafsya", "Nopal Yusriya"}
	D[*n].prodi = "Teknik Elektro"
	D[*n].judul = "Mekanika Kuantun"
	D[*n].sumberDana = "Internal"
	D[*n].luaran = "Publikasi"
	D[*n].tahunKegiatan = 2023
	D[*n].jumlahKegiatan = 6
	*n += 1
	D[*n].ketua = "Aldo Darma"
	D[*n].anggota = [6]string{"Fyrza", "Bima", "Andhika", "Nyoman", "Edgar", "Gopal"}
	D[*n].prodi = "Teknik Elektro"
	D[*n].judul = "Mobil Listrik"
	D[*n].sumberDana = "Eksternal"
	D[*n].luaran = "Produk"
	D[*n].tahunKegiatan = 2023
	D[*n].jumlahKegiatan = 4
	*n += 1
	D[*n].ketua = "Riyad Mustafid"
	D[*n].anggota = [6]string{"Ada", "Eki Nata", "Dendi", "Dadang", "Pram Saputra", "Rio Maulana"}
	D[*n].prodi = "Sistem Informasi"
	D[*n].judul = "Studi Kasus pada Toko Online"
	D[*n].sumberDana = "Internal"
	D[*n].luaran = "Publikasi"
	D[*n].tahunKegiatan = 2022
	D[*n].jumlahKegiatan = 5
	*n += 1

}

func scantext() string {
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()
	return text
}

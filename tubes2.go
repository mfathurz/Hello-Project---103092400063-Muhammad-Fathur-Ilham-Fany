package main

import "fmt"

const maxPasien = 100
const maxPaket = 10

type PaketMCU struct {
	id    int
	nama  string
	harga int
}

type Pasien struct {
	id      int
	nama    string
	paketID int
	tanggal string
}

var daftarPaket [maxPaket]PaketMCU
var daftarPasien [maxPasien]Pasien
var jumlahPaket, jumlahPasien int

func tambahPaket() {
	var jumlahInput int
	fmt.Println("Ingin Menambahkan Berapa Paket?: ")
	fmt.Scan(&jumlahInput)

	for i := 0; i < jumlahInput; i++ {
		fmt.Printf("\nInput Paket ke-%d", i+1)

		fmt.Println("\nMasukkan Id Paket: ")
		fmt.Scan(&daftarPaket[jumlahPaket].id)

		fmt.Println("Masukkan Nama Paket: ")
		fmt.Scan(&daftarPaket[jumlahPaket].nama)

		fmt.Println("Masukkan Harga Paket: ")
		fmt.Scan(&daftarPaket[jumlahPaket].harga)
		jumlahPaket++
	}

	fmt.Println("Paket berhasil ditambahkan.")
}

func ubahPaket() {
	var id int
	fmt.Print("ID Paket yang ingin diubah: ")
	fmt.Scan(&id)
	for i := 0; i < jumlahPaket; i++ {
		if daftarPaket[i].id == id {
			fmt.Print("Nama Baru: ")
			fmt.Scan(&daftarPaket[i].nama)
			fmt.Print("Harga Baru: ")
			fmt.Scan(&daftarPaket[i].harga)
			fmt.Println("Paket berhasil diubah.")
			return
		}
	}
	fmt.Println("Paket tidak ditemukan.")
}

func hapusPaket() {
	var id int
	fmt.Print("ID Paket yang ingin dihapus: ")
	fmt.Scan(&id)
	for i := 0; i < jumlahPaket; i++ {
		if daftarPaket[i].id == id {
			for j := i; j < jumlahPaket-1; j++ {
				daftarPaket[j] = daftarPaket[j+1]
			}
			jumlahPaket--
			fmt.Println("Paket berhasil dihapus.")
			return
		}
	}
	fmt.Println("Paket tidak ditemukan.")
}

func tampilkanPaket() {
	fmt.Println("\n[Daftar Paket MCU]")
	for i := 0; i < jumlahPaket; i++ {
		fmt.Printf("%d. %s - Rp%d\n", daftarPaket[i].id, daftarPaket[i].nama, daftarPaket[i].harga)
	}
}

func tampilNamaPaket(id int) string {
	for i := 0; i < jumlahPaket; i++ {
		if daftarPaket[i].id == id {
			return daftarPaket[i].nama
		}
	}
	return "Tidak diketahui"
}

func tampilHargaPaket(id int) int {
	for i := 0; i < jumlahPaket; i++ {
		if daftarPaket[i].id == id {
			return daftarPaket[i].harga
		}
	}
	return 0
}

func laporanPemasukan() {
	var total int
	for i := 0; i < jumlahPasien; i++ {
		harga := tampilHargaPaket(daftarPasien[i].paketID)
		total += harga
	}
	fmt.Println("\n=== Laporan Pemasukan ===")
	fmt.Printf("Total pemasukan dari semua pasien: Rp%d\n", total)
}

func tambahPasien() {
	var jumlahPasien int

	fmt.Println("mau ditambah berapa?: ")
	fmt.Scan(&jumlahPasien)

	for i := 0; i < jumlahPasien; i++ {
		fmt.Print("ID Pasien: ")
		fmt.Scan(&daftarPasien[jumlahPasien].id, i+1)
		fmt.Print("Nama Pasien: ")
		fmt.Scan(&daftarPasien[jumlahPasien].nama)
		for i := 0; i < count; i++ {
			if paketsTRUCT == paketID {
				fmt.Scan(&daftarPasien[jumlahPasien].paketID)
			}
			else {
				fmt.Print("ID Paket: ")
		fmt.Scan(&daftarPasien[jumlahPasien].paketID)
			}
		}
		fmt.Print("ID Paket: ")
		fmt.Scan(&daftarPasien[jumlahPasien].paketID)
		fmt.Print("Tanggal MCU (YYYY-MM-DD): ")
		fmt.Scan(&daftarPasien[jumlahPasien].tanggal)
		jumlahPasien++
		fmt.Println("Pasien berhasil ditambahkan.")
	}
}

func ubahPasien() {
	var id int
	fmt.Print("ID Pasien yang ingin diubah: ")
	fmt.Scan(&id)
	for i := 0; i < jumlahPasien; i++ {
		if daftarPasien[i].id == id {
			fmt.Print("Nama Baru: ")
			fmt.Scan(&daftarPasien[i].nama)
			fmt.Print("ID Paket Baru: ")
			fmt.Scan(&daftarPasien[i].paketID)
			fmt.Print("Tanggal Baru (YYYY-MM-DD): ")
			fmt.Scan(&daftarPasien[i].tanggal)
			fmt.Println("Pasien berhasil diubah.")
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan.")
}

func hapusPasien() {
	var id int
	fmt.Print("ID Pasien yang ingin dihapus: ")
	fmt.Scan(&id)
	for i := 0; i < jumlahPasien; i++ {
		if daftarPasien[i].id == id {
			for j := i; j < jumlahPasien-1; j++ {
				daftarPasien[j] = daftarPasien[j+1]
			}
			jumlahPasien--
			fmt.Println("Pasien berhasil dihapus.")
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan.")
}

func tampilkanPasien() {
	fmt.Println("\n[Data Pasien]")
	for i := 0; i < jumlahPasien; i++ {
		fmt.Printf("%d. %s - Paket: %s - Tanggal: %s\n",
			daftarPasien[i].id, daftarPasien[i].nama, tampilNamaPaket(daftarPasien[i].paketID), daftarPasien[i].tanggal)
	}
}

func cariPasienByPaket() {
	var id int
	fmt.Print("ID Paket yang dicari: ")
	fmt.Scan(&id)
	fmt.Println("[Pasien dengan Paket Tersebut]:")
	for i := 0; i < jumlahPasien; i++ {
		if daftarPasien[i].paketID == id {
			fmt.Println(daftarPasien[i].nama)
		}
	}
}

func cariPasienByNama() {
	var nama string
	fmt.Print("Nama pasien (harus sesuai): ")
	fmt.Scan(&nama)
	fmt.Println("[Hasil Pencarian Nama]:")
	for i := 0; i < jumlahPasien; i++ {
		if daftarPasien[i].nama == nama {
			fmt.Println(daftarPasien[i].nama)
		}
	}
}

func cariPasienByPeriode() {
	var periode string
	fmt.Print("Periode (YYYY-MM): ")
	fmt.Scan(&periode)
	fmt.Println("[Pasien pada Periode Tersebut]:")
	for i := 0; i < jumlahPasien; i++ {
		cocok := true
		for j := 0; j < 7; j++ {
			if daftarPasien[i].tanggal[j] != periode[j] {
				cocok = false
			}
		}
		if cocok {
			fmt.Println(daftarPasien[i].nama)
		}
	}
}

func tampilkanPasienByTanggal() {
	var temp [maxPasien]Pasien

	for i := 0; i < jumlahPasien; i++ {
		temp[i] = daftarPasien[i]
	}

	for i := 1; i < jumlahPasien; i++ {
		data := temp[i]
		j := i - 1

		for j >= 0 && temp[j].tanggal > data.tanggal {
			temp[j+1] = temp[j]
			j--
		}
		temp[j+1] = data
	}

	fmt.Println("\n[Data Pasien Urut Berdasarkan Tanggal MCU]")

	for i := 0; i < jumlahPasien; i++ {
		fmt.Printf("%d. %s - Paket ID: %d - Tanggal: %s\n",
			temp[i].id, temp[i].nama, temp[i].paketID, temp[i].tanggal)
	}
}

func tampilkanPasienByPaket() {
	var temp [maxPasien]Pasien

	for i := 0; i < jumlahPasien; i++ {
		temp[i] = daftarPasien[i]
	}

	for i := 0; i < jumlahPasien-1; i++ {
		minIdx := i

		for j := i + 1; j < jumlahPasien; j++ {
			if temp[j].paketID < temp[minIdx].paketID {
				minIdx = j
			}
		}
		temp[i], temp[minIdx] = temp[minIdx], temp[i]
	}

	fmt.Println("\n[Data Pasien Urut Berdasarkan Paket MCU]")

	for i := 0; i < jumlahPasien; i++ {
		fmt.Printf("%d. %s - Paket ID: %d - Tanggal: %s\n",
			temp[i].id, temp[i].nama, temp[i].paketID, temp[i].tanggal)
	}
}

func main() {
	var selesai bool = false
	var pilihan int
	for !selesai {
		fmt.Println("\n=== Menu Utama ===")
		fmt.Println("1. Tambah Paket MCU")
		fmt.Println("2. Ubah Paket MCU")
		fmt.Println("3. Hapus Paket MCU")
		fmt.Println("4. Tampilkan Daftar Paket")
		fmt.Println("5. Tambah Pasien")
		fmt.Println("6. Ubah Pasien")
		fmt.Println("7. Hapus Pasien")
		fmt.Println("8. Tampilkan Daftar Pasien")
		fmt.Println("9. Cari Pasien Berdasarkan Nama")
		fmt.Println("10. Cari Pasien Berdasarkan ID Paket")
		fmt.Println("11. Cari Pasien Berdasarkan Periode")
		fmt.Println("12. Laporan Pemasukan")
		fmt.Println("13. Urut Pasien Berdasarkan Tanggal MCU")
		fmt.Println("14. Urut Pasien Berdasarkan Paket MCU")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahPaket()
		} else if pilihan == 2 {
			ubahPaket()
		} else if pilihan == 3 {
			hapusPaket()
		} else if pilihan == 4 {
			tampilkanPaket()
		} else if pilihan == 5 {
			tambahPasien()
		} else if pilihan == 6 {
			ubahPasien()
		} else if pilihan == 7 {
			hapusPasien()
		} else if pilihan == 8 {
			tampilkanPasien()
		} else if pilihan == 9 {
			cariPasienByNama()
		} else if pilihan == 10 {
			cariPasienByPaket()
		} else if pilihan == 11 {
			cariPasienByPeriode()
		} else if pilihan == 12 {
			laporanPemasukan()
		} else if pilihan == 13 {
			tampilkanPasienByTanggal()
		} else if pilihan == 14 {
			tampilkanPasienByPaket()
		} else if pilihan == 0 {
			fmt.Println("Terima kasih!")
			selesai = true
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

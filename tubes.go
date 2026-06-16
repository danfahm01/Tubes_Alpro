package main

import "fmt"

const NMAX = 100

type Warga struct {
	id          string
	nama        string
	totalSampah float64
}

type Setoran struct {
	idWarga    string
	tanggal    string
	minggu     int
	jenisSampah string
	berat      float64
}

type ArrWarga [NMAX]Warga
type ArrSetoran [NMAX]Setoran

func main() {
	var dataWarga ArrWarga
	var dataSetoran ArrSetoran
	var totalWargaAktif, totalSetoranAktif int
	var pilihan int

	for {
		fmt.Println("\n=== Aplikasi Waste-Track ===")
		fmt.Println("1. Tambah Warga")
		fmt.Println("2. Ubah Data Warga")
		fmt.Println("3. Hapus Data Warga")
		fmt.Println("4. Catat Setoran Sampah")
		fmt.Println("5. Cari Warga (Sequential - Berdasarkan Nama)")
		fmt.Println("6. Cari Warga (Binary - Berdasarkan ID)")
		fmt.Println("7. Urutkan Warga (Selection Sort - Sampah Terbanyak)")
		fmt.Println("8. Urutkan Warga (Insertion Sort - Sampah Terbanyak)")
		fmt.Println("9. Statistik Mingguan")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahWarga(&dataWarga, &totalWargaAktif)
		case 2:
			ubahWarga(&dataWarga, totalWargaAktif)
		case 3:
			hapusWarga(&dataWarga, &totalWargaAktif)
		case 4:
			catatSetoran(&dataWarga, totalWargaAktif, &dataSetoran, &totalSetoranAktif)
		case 5:
			var nama string
			fmt.Print("Masukkan nama warga yang dicari: ")
			fmt.Scan(&nama)
			idx := cariSequentialNama(dataWarga, totalWargaAktif, nama)
			tampilHasilCari(dataWarga, idx)
		case 6:
			var id string
			fmt.Print("Masukkan ID warga yang dicari: ")
			fmt.Scan(&id)
			urutWargaByID(&dataWarga, totalWargaAktif)
			idx := cariBinaryID(dataWarga, totalWargaAktif, id)
			tampilHasilCari(dataWarga, idx)
		case 7:
			selectionSortDescending(&dataWarga, totalWargaAktif)
			tampilSemuaWarga(dataWarga, totalWargaAktif)
		case 8:
			insertionSortDescending(&dataWarga, totalWargaAktif)
			tampilSemuaWarga(dataWarga, totalWargaAktif)
		case 9:
			statistikMingguan(dataSetoran, totalSetoranAktif, dataWarga, totalWargaAktif)
		case 0:
			fmt.Println("Terima kasih telah menggunakan Waste-Track!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahWarga(daftarWarga *ArrWarga, jumlahWarga *int) {
	if *jumlahWarga < NMAX {
		fmt.Print("Masukkan ID Warga: ")
		fmt.Scan(&daftarWarga[*jumlahWarga].id)
		fmt.Print("Masukkan Nama Warga (Tanpa Spasi): ")
		fmt.Scan(&daftarWarga[*jumlahWarga].nama)
		daftarWarga[*jumlahWarga].totalSampah = 0
		*jumlahWarga++
		fmt.Println("Data warga berhasil ditambahkan!")
	} else {
		fmt.Println("Kapasitas penyimpanan penuh!")
	}
}

func ubahWarga(daftarWarga *ArrWarga, jumlahWarga int) {
	var id string
	fmt.Print("Masukkan ID Warga yang akan diubah: ")
	fmt.Scan(&id)

	idx := -1
	for i := 0; i < jumlahWarga; i++ {
		if daftarWarga[i].id == id {
			idx = i
			break
		}
	}

	if idx != -1 {
		fmt.Print("Masukkan Nama Baru: ")
		fmt.Scan(&daftarWarga[idx].nama)
		fmt.Println("Data warga berhasil diubah!")
	} else {
		fmt.Println("Data warga tidak ditemukan.")
	}
}

func hapusWarga(daftarWarga *ArrWarga, jumlahWarga *int) {
	var id string
	fmt.Print("Masukkan ID Warga yang akan dihapus: ")
	fmt.Scan(&id)

	idx := -1
	for i := 0; i < *jumlahWarga; i++ {
		if daftarWarga[i].id == id {
			idx = i
			break
		}
	}

	if idx != -1 {
		for i := idx; i < *jumlahWarga-1; i++ {
			daftarWarga[i] = daftarWarga[i+1]
		}
		*jumlahWarga--
		fmt.Println("Data warga berhasil dihapus!")
	} else {
		fmt.Println("Data warga tidak ditemukan.")
	}
}

func catatSetoran(daftarWarga *ArrWarga, jumlahWarga int, daftarSetoran *ArrSetoran, jumlahSetoran *int) {
	if *jumlahSetoran < NMAX {
		fmt.Print("Masukkan ID Warga: ")
		fmt.Scan(&daftarSetoran[*jumlahSetoran].idWarga)

		ditemukan := false
		idxWarga := -1
		for i := 0; i < jumlahWarga; i++ {
			if daftarWarga[i].id == daftarSetoran[*jumlahSetoran].idWarga {
				ditemukan = true
				idxWarga = i
				break
			}
		}

		if ditemukan {
			var optJenis int
			fmt.Print("Masukkan Tanggal (DD-MM-YYYY): ")
			fmt.Scan(&daftarSetoran[*jumlahSetoran].tanggal)
			fmt.Print("Masukkan Minggu Ke- (angka): ")
			fmt.Scan(&daftarSetoran[*jumlahSetoran].minggu)
			
			for {
				fmt.Println("Pilih Jenis Sampah:")
				fmt.Println("1. Organik")
				fmt.Println("2. Anorganik")
				fmt.Println("3. B3")
				fmt.Print("Pilihan (1-3): ")
				fmt.Scan(&optJenis)
				if optJenis == 1 {
					daftarSetoran[*jumlahSetoran].jenisSampah = "Organik"
					break
				} else if optJenis == 2 {
					daftarSetoran[*jumlahSetoran].jenisSampah = "Anorganik"
					break
				} else if optJenis == 3 {
					daftarSetoran[*jumlahSetoran].jenisSampah = "B3"
					break
				} else {
					fmt.Println("Pilihan tidak valid, coba lagi.")
				}
			}

			fmt.Print("Masukkan Berat Sampah (Kg): ")
			fmt.Scan(&daftarSetoran[*jumlahSetoran].berat)

			daftarWarga[idxWarga].totalSampah += daftarSetoran[*jumlahSetoran].berat
			*jumlahSetoran++
			fmt.Println("Setoran berhasil dicatat!")
		} else {
			fmt.Println("ID Warga tidak terdaftar!")
		}
	} else {
		fmt.Println("Kapasitas log setoran penuh!")
	}
}

func statistikMingguan(daftarSetoran ArrSetoran, jumlahSetoran int, daftarWarga ArrWarga, jumlahWarga int) {
	var mingguCari int
	var total float64 = 0
	var ditemukan bool = false

	fmt.Print("Masukkan Minggu Ke- berapa yang ingin dihitung: ")
	fmt.Scan(&mingguCari)

	fmt.Printf("\n--- Detail Setoran Minggu Ke-%d ---\n", mingguCari)
	fmt.Println("=========================================================================")
	fmt.Printf("%-12s | %-15s | %-15s | %-10s\n", "Tanggal", "Nama Warga", "Jenis Sampah", "Berat (Kg)")
	fmt.Println("=========================================================================")

	for i := 0; i < jumlahSetoran; i++ {
		if daftarSetoran[i].minggu == mingguCari {
			namaWarga := "Tidak Diketahui"
			for j := 0; j < jumlahWarga; j++ {
				if daftarWarga[j].id == daftarSetoran[i].idWarga {
					namaWarga = daftarWarga[j].nama
					break
				}
			}

			fmt.Printf("%-12s | %-15s | %-15s | %-10.2f\n", daftarSetoran[i].tanggal, namaWarga, daftarSetoran[i].jenisSampah, daftarSetoran[i].berat)
			total += daftarSetoran[i].berat
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ada data setoran pada minggu ini.")
	}

	fmt.Println("=========================================================================")
	fmt.Printf("Total akumulasi sampah pada minggu ke-%d adalah %.2f Kg\n", mingguCari, total)
}

func cariSequentialNama(daftarWarga ArrWarga, jumlahWarga int, nama string) int {
	for i := 0; i < jumlahWarga; i++ {
		if daftarWarga[i].nama == nama {
			return i
		}
	}
	return -1
}

func cariBinaryID(daftarWarga ArrWarga, jumlahWarga int, id string) int {
	kiri := 0
	kanan := jumlahWarga - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if daftarWarga[tengah].id == id {
			return tengah
		} else if daftarWarga[tengah].id < id {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func tampilHasilCari(daftarWarga ArrWarga, idx int) {
	if idx != -1 {
		fmt.Printf("Nemu Nih --> ID: %s | Nama: %s | Total Setoran: %.2f Kg\n", daftarWarga[idx].id, daftarWarga[idx].nama, daftarWarga[idx].totalSampah)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func selectionSortDescending(daftarWarga *ArrWarga, jumlahWarga int) {
	for i := 0; i < jumlahWarga-1; i++ {
		idxMax := i
		for j := i + 1; j < jumlahWarga; j++ {
			if daftarWarga[j].totalSampah > daftarWarga[idxMax].totalSampah {
				idxMax = j
			}
		}

		temp := daftarWarga[i]
		daftarWarga[i] = daftarWarga[idxMax]
		daftarWarga[idxMax] = temp
	}
	fmt.Println("Data berhasil diurutkan dengan Selection Sort!")
}

func insertionSortDescending(daftarWarga *ArrWarga, jumlahWarga int) {
	for i := 1; i < jumlahWarga; i++ {
		temp := daftarWarga[i]
		j := i - 1
		for j >= 0 && daftarWarga[j].totalSampah < temp.totalSampah {
			daftarWarga[j+1] = daftarWarga[j]
			j--
		}
		daftarWarga[j+1] = temp
	}
	fmt.Println("Data berhasil diurutkan dengan Insertion Sort!")
}

func urutWargaByID(daftarWarga *ArrWarga, jumlahWarga int) {
	for i := 1; i < jumlahWarga; i++ {
		temp := daftarWarga[i]
		j := i - 1
		for j >= 0 && daftarWarga[j].id > temp.id {
			daftarWarga[j+1] = daftarWarga[j]
			j--
		}
		daftarWarga[j+1] = temp
	}
}

func tampilSemuaWarga(daftarWarga ArrWarga, jumlahWarga int) {
	fmt.Println("\n--- Daftar Warga ---")
	fmt.Println("======================================================")
	fmt.Printf("%-5s | %-10s | %-20s | %-12s\n", "No", "ID Warga", "Nama Warga", "Total (Kg)")
	fmt.Println("======================================================")
	for i := 0; i < jumlahWarga; i++ {
		fmt.Printf("%-5d | %-10s | %-20s | %-12.2f\n", i+1, daftarWarga[i].id, daftarWarga[i].nama, daftarWarga[i].totalSampah)
	}
	fmt.Println("======================================================")
}
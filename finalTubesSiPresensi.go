package main

import "fmt"

const nmax = 100

type mhs struct {
	nim         int64
	nama        string
	kelas       string
	jadwalKelas string
	ket         string
}

type tmhs [nmax]mhs

func main() {
	var mahasiswa tmhs
	var n, tambah int
	var pilih int
	var cariNim int64
	var cariNama, cariKet, cariKelas, cariJadwalKuliah string

	for {
		fmt.Println("\n===== PRESENSI MAHASISWA =====")
		fmt.Println("1. Input Data Mahasiswa")
		fmt.Println("2. Tampilkan Data Mahasiswa")
		fmt.Println("3. Cari Mahasiswa")
		fmt.Println("4. Urutkan Data")
		fmt.Println("5. Statistik")
		fmt.Println("6. Hapus Data")
		fmt.Println("7. Edit Data")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		switch pilih {

		case 1:
			fmt.Print("Jumlah mahasiswa: ")
			fmt.Scan(&tambah)
			inputmhs(n, tambah, mahasiswa[:])
			n = n + tambah

		case 2:
			tampil(n, mahasiswa[:])

		case 3:
			fmt.Println("===== CARI BERDASARKAN =====")
			fmt.Println("1. Cari dengan NIM")
			fmt.Println("2. Cari dengan Nama")
			fmt.Println("3. Cari dengan Keterangan")
			fmt.Println("4. Cari dengan Kelas")
			fmt.Println("5. Cari dengan Mata Kuliah")
			fmt.Print("Pilih: ")
			var pilih2 int
			fmt.Scan(&pilih2)

			switch pilih2 {

			case 1:
				fmt.Print("Masukkan NIM yang dicari: ")
				fmt.Scan(&cariNim)
				selectionSortByNIM(n, mahasiswa[:])
				idx := binarySearchMhs(n, mahasiswa[:], cariNim)
				if idx != -1 {
					fmt.Println("Data ditemukan")
					fmt.Println("Nama      :", mahasiswa[idx].nama)
					fmt.Println("NIM       :", mahasiswa[idx].nim)
					fmt.Println("Kelas     :", mahasiswa[idx].kelas)
					fmt.Println("Keterangan:", mahasiswa[idx].ket)
				} else {
					fmt.Println("Data tidak ditemukan")
				}

			case 2:
				fmt.Print("Masukkan Nama yang dicari: ")
				fmt.Scan(&cariNama)
				idx := sequentialSearchNama(n, mahasiswa[:], cariNama)
				if idx != -1 {
					fmt.Println("Data ditemukan")
					fmt.Println("Nama      :", mahasiswa[idx].nama)
					fmt.Println("NIM       :", mahasiswa[idx].nim)
					fmt.Println("Kelas     :", mahasiswa[idx].kelas)
					fmt.Println("Keterangan:", mahasiswa[idx].ket)
				} else {
					fmt.Println("Data tidak ditemukan")
				}

			case 3:
				fmt.Print("Masukkan Keterangan (Hadir/Izin/Sakit/Alpha): ")
				fmt.Scan(&cariKet)
				sequentialSearchKet(n, mahasiswa[:], cariKet)

			case 4:
				fmt.Print("Masukkan Kelas yang dicari: ")
				fmt.Scan(&cariKelas)
				sequentialSearchKelas(n, mahasiswa[:], cariKelas)

			case 5:
				fmt.Print("Masukkan Mata Kuliah yang dicari: ")
				fmt.Scan(&cariJadwalKuliah)
				sequentialSearchJadwalKelas(n, mahasiswa[:], cariJadwalKuliah)

			default:
				fmt.Println("Pilihan tidak tersedia")
			}

		case 4:
			fmt.Println("===== URUTKAN BERDASARKAN =====")
			fmt.Println("1. Nama Ascending (A-Z)")
			fmt.Println("2. Nama Descending (Z-A)")
			fmt.Println("3. NIM Ascending (0-9)")
			fmt.Println("4. NIM Descending (9-0)")
			fmt.Print("Pilih: ")
			var pilih3 int
			fmt.Scan(&pilih3)

			if pilih3 == 1 {
				selectionSortMhs(n, mahasiswa[:], true)
				fmt.Println("Data berhasil diurutkan Ascending")
			} else if pilih3 == 2 {
				selectionSortMhs(n, mahasiswa[:], false)
				fmt.Println("Data berhasil diurutkan Descending")
			} else if pilih3 == 3 {
				insertionsortbynim(n, mahasiswa[:], true)
				fmt.Println("Data berhasil diurutkan NIM Ascending")
			} else if pilih3 == 4 {
				insertionsortbynim(n, mahasiswa[:], false)
				fmt.Println("Data berhasil diurutkan NIM Descending")
			} else {
				fmt.Println("Pilihan tidak tersedia")
			}
			tampil(n, mahasiswa[:])

		case 5:
			statistik(n, mahasiswa[:])

		case 6:
			fmt.Println("===== HAPUS BERDASARKAN =====")
			fmt.Println("1. Hapus dengan NIM")
			fmt.Println("2. Hapus dengan Nama")
			fmt.Print("Pilih: ")
			var pilih4 int
			fmt.Scan(&pilih4)

			switch pilih4 {
			case 1:
				fmt.Print("Masukkan NIM yang akan dihapus: ")
				fmt.Scan(&cariNim)
				hapusByNim(&n, mahasiswa[:], cariNim)

			case 2:
				fmt.Print("Masukkan Nama yang akan dihapus: ")
				fmt.Scan(&cariNama)
				hapusByNama(&n, mahasiswa[:], cariNama)

			default:
				fmt.Println("Pilihan tidak tersedia")
			}
			tampil(n, mahasiswa[:])

		case 7:
			fmt.Print("Masukan NIM yang akan diedit keterangannya :")
			fmt.Scan(&cariNim)
			editdata(n, mahasiswa[:], cariNim)

		case 0:
			fmt.Println("Program selesai")
			return

		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func inputmhs(mulai int, jumlah int, mahasiswa []mhs) {
	var i int
	for i = mulai; i < mulai+jumlah; i++ {
		fmt.Println("\nInput Data Mahasiswa ke-", i+1)
		fmt.Print("Nama : ")
		fmt.Scan(&mahasiswa[i].nama)
		fmt.Print("NIM : ")
		fmt.Scan(&mahasiswa[i].nim)
		fmt.Print("Kelas : ")
		fmt.Scan(&mahasiswa[i].kelas)
		fmt.Print("Mata Kuliah : ")
		fmt.Scan(&mahasiswa[i].jadwalKelas)
		fmt.Print("Keterangan (Hadir/Izin/Sakit/Alpha) : ")
		fmt.Scan(&mahasiswa[i].ket)
	}
}

func tampil(n int, mahasiswa []mhs) {
	fmt.Println("\n===== DATA MAHASISWA =====")
	for i := 0; i < n; i++ {
		fmt.Println("Nama       :", mahasiswa[i].nama)
		fmt.Println("NIM        :", mahasiswa[i].nim)
		fmt.Println("Kelas      :", mahasiswa[i].kelas)
		fmt.Println("Mata Kuliah:", mahasiswa[i].jadwalKelas)
		fmt.Println("Keterangan :", mahasiswa[i].ket)
		fmt.Println()
	}
}

func sequentialSearchNama(n int, mahasiswa []mhs, nama string) int {
	for i := 0; i < n; i++ {
		if mahasiswa[i].nama == nama {
			return i
		}
	}
	return -1
}

func sequentialSearchKelas(n int, mahasiswa []mhs, kel string) {
	found := false
	for i := 0; i < n; i++ {
		if mahasiswa[i].kelas == kel {
			fmt.Println("Nama  :", mahasiswa[i].nama)
			fmt.Println("NIM   :", mahasiswa[i].nim)
			fmt.Println("Kelas :", mahasiswa[i].kelas)
			fmt.Println()
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada mahasiswa dengan kelas", kel)
	}
}

func sequentialSearchJadwalKelas(n int, mahasiswa []mhs, jadwalKelas string) {
	found := false
	for i := 0; i < n; i++ {
		if mahasiswa[i].jadwalKelas == jadwalKelas {
			fmt.Println("Nama  :", mahasiswa[i].nama)
			fmt.Println("NIM   :", mahasiswa[i].nim)
			fmt.Println("Kelas :", mahasiswa[i].kelas)
			fmt.Println()
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada mahasiswa dengan mata kuliah", jadwalKelas)
	}
}

func sequentialSearchKet(n int, mahasiswa []mhs, ket string) {
	found := false
	for i := 0; i < n; i++ {
		if mahasiswa[i].ket == ket {
			fmt.Println("Nama  :", mahasiswa[i].nama)
			fmt.Println("NIM   :", mahasiswa[i].nim)
			fmt.Println("Kelas :", mahasiswa[i].kelas)
			fmt.Println()
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada mahasiswa dengan keterangan", ket)
	}
}

func selectionSortMhs(n int, mahasiswa []mhs, ascending bool) {
	var pass, i, idx int
	var temp mhs
	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if ascending {
				if mahasiswa[i].nama < mahasiswa[idx].nama {
					idx = i
				}
			} else {
				if mahasiswa[i].nama > mahasiswa[idx].nama {
					idx = i
				}
			}
		}
		temp = mahasiswa[pass]
		mahasiswa[pass] = mahasiswa[idx]
		mahasiswa[idx] = temp
	}
}

func selectionSortByNIM(n int, mahasiswa []mhs) {
	var pass, i, idx int
	var temp mhs
	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if mahasiswa[i].nim < mahasiswa[idx].nim {
				idx = i
			}
		}
		temp = mahasiswa[pass]
		mahasiswa[pass] = mahasiswa[idx]
		mahasiswa[idx] = temp
	}
}

func binarySearchMhs(n int, mahasiswa []mhs, nim int64) int {
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if mahasiswa[mid].nim == nim {
			return mid
		}
		if mahasiswa[mid].nim < nim {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func hapusByNim(n *int, mahasiswa []mhs, nim int64) {
	selectionSortByNIM(*n, mahasiswa)
	idx := binarySearchMhs(*n, mahasiswa, nim)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}
	for i := idx; i < *n-1; i++ {
		mahasiswa[i] = mahasiswa[i+1]
	}
	*n = *n - 1
	fmt.Println("Data berhasil dihapus")
}

func hapusByNama(n *int, mahasiswa []mhs, nama string) {
	idx := sequentialSearchNama(*n, mahasiswa, nama)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}
	for i := idx; i < *n-1; i++ {
		mahasiswa[i] = mahasiswa[i+1]
	}
	*n = *n - 1
	fmt.Println("Data berhasil dihapus")
}

func statistik(n int, mahasiswa []mhs) {
	var totalH, totalS, totalI, totalA int
	for i := 0; i < n; i++ {
		if mahasiswa[i].ket == "Hadir" {
			totalH++
		} else if mahasiswa[i].ket == "Sakit" {
			totalS++
		} else if mahasiswa[i].ket == "Izin" {
			totalI++
		} else {
			totalA++
		}
	}

	fmt.Println("\n===== STATISTIK =====")
	fmt.Println("Total Hadir :", totalH)
	fmt.Println("Total Sakit :", totalS)
	fmt.Println("Total Izin  :", totalI)
	fmt.Println("Total Alpha :", totalA)

	if n > 0 {
		persenHadir := float64(totalH) / float64(n) * 100
		fmt.Printf("Persentase Kehadiran : %.2f%%\n", persenHadir)
	}

	var maxAlpha, alphaCount int
	var namaMax string
	var sudahAda bool

	maxAlpha = -1

	for i := 0; i < n; i++ {
		sudahAda = false
		for k := 0; k < i; k++ {
			if mahasiswa[k].nim == mahasiswa[i].nim {
				sudahAda = true
			}
		}
		if sudahAda {
			continue
		}
		alphaCount = 0
		for j := 0; j < n; j++ {
			if mahasiswa[j].nim == mahasiswa[i].nim && mahasiswa[j].ket == "Alpha" {
				alphaCount++
			}
		}
		if alphaCount > maxAlpha {
			maxAlpha = alphaCount
			namaMax = mahasiswa[i].nama
		} else if alphaCount == maxAlpha {
			if mahasiswa[i].nama < namaMax {
				namaMax = mahasiswa[i].nama
			}
		}
	}

	if maxAlpha > 0 {
		fmt.Println("Mahasiswa dengan Alpha Terbanyak :", namaMax)
		fmt.Println("Jumlah Alpha :", maxAlpha)
	} else {
		fmt.Println("Tidak ada mahasiswa yang Alpha")
	}
}

func editdata(n int, mahasiswa []mhs, nim int64) {
	var idx int
	idx = binarySearchMhs(n, mahasiswa, nim)

	if idx != -1 {
		fmt.Println("Data ditemukan")
		fmt.Println("Nama :", mahasiswa[idx].nama)
		fmt.Println("Keterangan Lama :", mahasiswa[idx].ket)

		fmt.Print("Keterangan Baru (Hadir/Izin/Sakit/Alpha): ")
		fmt.Scan(&mahasiswa[idx].ket)

		fmt.Println("Keterangan berhasil diubah")
		fmt.Println("Nama :", mahasiswa[idx].nama)
		fmt.Println("Keterangan Baru :", mahasiswa[idx].ket)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func insertionsortbynim(n int, mahasiswa []mhs, ascend bool) {
	var pass, i int
	var temp mhs

	for pass = 1; pass < n; pass++ {
		temp = mahasiswa[pass]
		i = pass - 1

		if ascend {
			for i >= 0 && mahasiswa[i].nim > temp.nim {
				mahasiswa[i+1] = mahasiswa[i]
				i--
			}
		} else {
			for i >= 0 && mahasiswa[i].nim < temp.nim {
				mahasiswa[i+1] = mahasiswa[i]
				i--
			}
		}
		mahasiswa[i+1] = temp

	}
}

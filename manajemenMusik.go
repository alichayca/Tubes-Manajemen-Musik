package main

import (
	"fmt"
	"strings"
)

const NMAX int = 100

type Musik struct {
	Judul  string
	Artis  string
	Genre  string
	Rating int
}

type ArrMusik [NMAX]Musik

func inputData(A *ArrMusik, n *int) {
	var jumlah int

	fmt.Print("Jumlah musik yang ingin ditambahkan: ")
	fmt.Scan(&jumlah)

	for i := 0; i < jumlah; i++ {
		fmt.Println("\nData ke-", *n+1)

		fmt.Print("Judul : ")
		fmt.Scan(&A[*n].Judul)

		fmt.Print("Artis : ")
		fmt.Scan(&A[*n].Artis)

		fmt.Print("Genre : ")
		fmt.Scan(&A[*n].Genre)

		fmt.Print("Rating : ")
		fmt.Scan(&A[*n].Rating)

		*n++
	}

	fmt.Println("Data berhasil ditambahkan")
}

func tampilData(A ArrMusik, n int) {
	if n == 0 {
		fmt.Println("Data kosong")
	} else {
		fmt.Println("\n===== DAFTAR MUSIK =====")
		for i := 0; i < n; i++ {
			fmt.Printf("%d. %s | %s | %s | %d\n",
				i+1,
				A[i].Judul,
				A[i].Artis,
				A[i].Genre,
				A[i].Rating)
		}
	}
}

func sequentialSearch(A ArrMusik, n int, judul string) int {
	idx := -1
	i := 0

	for i < n && idx == -1 {
		if strings.EqualFold(A[i].Judul, judul) {
			idx = i
		}
		i++
	}

	return idx
}

func editMusik(A *ArrMusik, n int) {
	var judul string

	fmt.Print("Masukkan judul yang ingin diedit: ")
	fmt.Scan(&judul)

	idx := sequentialSearch(*A, n, judul)

	if idx == -1 {
		fmt.Println("Lagu tidak ditemukan")
	} else {
		fmt.Print("Judul baru : ")
		fmt.Scan(&A[idx].Judul)

		fmt.Print("Artis baru : ")
		fmt.Scan(&A[idx].Artis)

		fmt.Print("Genre baru : ")
		fmt.Scan(&A[idx].Genre)

		fmt.Print("Rating baru : ")
		fmt.Scan(&A[idx].Rating)

		fmt.Println("Data berhasil diubah")
	}
}

func hapusMusik(A *ArrMusik, n *int) {
	var judul string

	fmt.Print("Masukkan judul yang ingin dihapus: ")
	fmt.Scan(&judul)

	idx := sequentialSearch(*A, *n, judul)

	if idx == -1 {
		fmt.Println("Lagu tidak ditemukan")
	} else {
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}

		*n--
		fmt.Println("Data berhasil dihapus")
	}
}

func insertionSortJudul(A *ArrMusik, n int) {
	var temp Musik
	var j int

	for i := 1; i < n; i++ {
		temp = A[i]
		j = i - 1

		for j >= 0 && strings.ToLower(A[j].Judul) > strings.ToLower(temp.Judul) {
			A[j+1] = A[j]
			j--
		}

		A[j+1] = temp
	}
}

func binarySearch(A ArrMusik, n int, judul string) int {
	left := 0
	right := n - 1
	found := -1

	for left <= right && found == -1 {
		mid := (left + right) / 2

		if strings.EqualFold(A[mid].Judul, judul) {
			found = mid
		} else if strings.ToLower(judul) < strings.ToLower(A[mid].Judul) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return found
}

func cariLagu(A ArrMusik, n int) {
	var judul string
	var temp ArrMusik

	for i := 0; i < n; i++ {
		temp[i] = A[i]
	}

	insertionSortJudul(&temp, n)

	fmt.Print("Masukkan judul lagu: ")
	fmt.Scan(&judul)

	idx := binarySearch(temp, n, judul)

	if idx == -1 {
		fmt.Println("Lagu tidak ditemukan")
	} else {
		fmt.Println("\n===== DATA DITEMUKAN =====")
		fmt.Println("Judul  :", temp[idx].Judul)
		fmt.Println("Artis  :", temp[idx].Artis)
		fmt.Println("Genre  :", temp[idx].Genre)
		fmt.Println("Rating :", temp[idx].Rating)
	}
}

func selectionSortAsc(A *ArrMusik, n int) {
	var min int
	var temp Musik

	for i := 0; i < n-1; i++ {
		min = i

		for j := i + 1; j < n; j++ {
			if A[j].Rating < A[min].Rating {
				min = j
			}
		}

		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}
}

func selectionSortDesc(A *ArrMusik, n int) {
	var max int
	var temp Musik

	for i := 0; i < n-1; i++ {
		max = i

		for j := i + 1; j < n; j++ {
			if A[j].Rating > A[max].Rating {
				max = j
			}
		}

		temp = A[i]
		A[i] = A[max]
		A[max] = temp
	}
}

func urutRating(A *ArrMusik, n int) {
	var pilih int

	fmt.Println("\n1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selectionSortAsc(A, n)
		fmt.Println("Data berhasil diurutkan ascending")
	} else if pilih == 2 {
		selectionSortDesc(A, n)
		fmt.Println("Data berhasil diurutkan descending")
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func selectionSortArtisAZ(A *ArrMusik, n int) {
	var min int
	var temp Musik

	for i := 0; i < n-1; i++ {
		min = i

		for j := i + 1; j < n; j++ {
			if strings.ToLower(A[j].Artis) < strings.ToLower(A[min].Artis) {
				min = j
			}
		}

		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}
}

func selectionSortArtisZA(A *ArrMusik, n int) {
	var max int
	var temp Musik

	for i := 0; i < n-1; i++ {
		max = i

		for j := i + 1; j < n; j++ {
			if strings.ToLower(A[j].Artis) > strings.ToLower(A[max].Artis) {
				max = j
			}
		}

		temp = A[i]
		A[i] = A[max]
		A[max] = temp
	}
}

func urutArtis(A *ArrMusik, n int) {
	var pilih int

	fmt.Println("\n1. A-Z")
	fmt.Println("2. Z-A")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selectionSortArtisAZ(A, n)
		fmt.Println("Data berhasil diurutkan A-Z")
	} else if pilih == 2 {
		selectionSortArtisZA(A, n)
		fmt.Println("Data berhasil diurutkan Z-A")
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func isiDataAwal(A *ArrMusik, n *int) {
	A[0] = Musik{"Always", "BonJovi", "BalladRock", 4}
	A[1] = Musik{"FotoKitaBlur", "SalPriadi", "IndiePop", 3}
	A[2] = Musik{"Thunder", "Seventeen", "HiphopKpop", 5}
	A[3] = Musik{"WithoutMe", "Halsey", "Pop", 4}
	A[4] = Musik{"Sofia", "Clairo", "Pop", 5}
	A[5] = Musik{"OneCallAway", "CharliePuth", "Pop", 2}
	A[6] = Musik{"WishYouWereHere", "AvrilLavigne", "PopRock", 3}
	A[7] = Musik{"PlayDate", "MelanieMartinez", "Pop", 1}
	A[8] = Musik{"3Strikes", "TerrorJr", "PopRNB", 2}
	A[9] = Musik{"LookAtMe", "XXXTENTACION", "Hiphop", 2}

	*n = 10
}

func main() {
	var data ArrMusik
	var n int
	var menu int

	isiDataAwal(&data, &n)

	for menu != 8 {
		fmt.Println("\n===== MENU MUSIK =====")
		fmt.Println("1. Tampilkan Musik")
		fmt.Println("2. Tambah Musik")
		fmt.Println("3. Edit Musik")
		fmt.Println("4. Hapus Musik")
		fmt.Println("5. Cari Lagu")
		fmt.Println("6. Urut Rating")
		fmt.Println("7. Urut Artis")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&menu)

		if menu == 1 {
			tampilData(data, n)

		} else if menu == 2 {
			inputData(&data, &n)

		} else if menu == 3 {
			editMusik(&data, n)

		} else if menu == 4 {
			hapusMusik(&data, &n)

		} else if menu == 5 {
			cariLagu(data, n)

		} else if menu == 6 {
			urutRating(&data, n)

		} else if menu == 7 {
			urutArtis(&data, n)

		} else if menu == 8 {
			fmt.Println("Program selesai")

		} else {
			fmt.Println("Menu tidak tersedia")
		}
	}
}

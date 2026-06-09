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
	fmt.Print("Jumlah data musik: ")
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Println("\nData ke-", i+1)

		fmt.Print("Judul : ")
		fmt.Scan(&A[i].Judul)

		fmt.Print("Artis : ")
		fmt.Scan(&A[i].Artis)

		fmt.Print("Genre : ")
		fmt.Scan(&A[i].Genre)

		fmt.Print("Rating : ")
		fmt.Scan(&A[i].Rating)
	}
}

func tampilData(A ArrMusik, n int) {
	if n == 0 {
		fmt.Println("Data kosong")
	} else {
		fmt.Println("\nDAFTAR MUSIK")
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
		fmt.Println("\nData ditemukan")
		fmt.Println("Judul :", temp[idx].Judul)
		fmt.Println("Artis :", temp[idx].Artis)
		fmt.Println("Genre :", temp[idx].Genre)
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

	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selectionSortAsc(A, n)
		fmt.Println("Berhasil diurutkan ascending")
	} else if pilih == 2 {
		selectionSortDesc(A, n)
		fmt.Println("Berhasil diurutkan descending")
	}
}

func main() {
	var data ArrMusik
	var n int
	var menu int

	inputData(&data, &n)

	for menu != 6 {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Tampilkan Musik")
		fmt.Println("2. Edit Musik")
		fmt.Println("3. Hapus Musik")
		fmt.Println("4. Cari Lagu")
		fmt.Println("5. Urut Rating")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menu)

		if menu == 1 {
			tampilData(data, n)
		} else if menu == 2 {
			editMusik(&data, n)
		} else if menu == 3 {
			hapusMusik(&data, &n)
		} else if menu == 4 {
			cariLagu(data, n)
		} else if menu == 5 {
			urutRating(&data, n)
		}
	}
}

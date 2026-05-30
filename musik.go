package main

import "fmt"

const NMAX = 100

type Musik struct {
	judul string
	artis string
	genre string
	rating int
}

type ArrMusik [NMAX]Musik

func tambahMusik(A *ArrMusik, n *int) {
	fmt.Print("Judul Lagu : ")
	fmt.Scan(&A[*n].judul)

	fmt.Print("Artis : ")
	fmt.Scan(&A[*n].artis)

	fmt.Print("Genre : ")
	fmt.Scan(&A[*n].genre)

	fmt.Print("Rating : ")
	fmt.Scan(&A[*n].rating)

	*n = *n + 1
	fmt.Println("Data berhasil ditambah")
}

func tampilMusik(A ArrMusik, n int) {
	var i int

	if n == 0 {
		fmt.Println("Data kosong")
	} else {
		for i = 0; i < n; i++ {
			fmt.Println("----------------")
			fmt.Println("Judul  :", A[i].judul)
			fmt.Println("Artis  :", A[i].artis)
			fmt.Println("Genre  :", A[i].genre)
			fmt.Println("Rating :", A[i].rating)
		}
	}
}

func sequentialSearch(A ArrMusik, n int, x string) int {
	var i, idx int
	idx = -1

	for i = 0; i < n; i++ {
		if A[i].judul == x {
			idx = i
		}
	}

	return idx
}

func selectionSortAsc(A *ArrMusik, n int) {
	var i, j, min int
	var temp Musik

	for i = 0; i < n-1; i++ {
		min = i

		for j = i + 1; j < n; j++ {
			if A[j].rating < A[min].rating {
				min = j
			}
		}

		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}
}

func selectionSortDesc(A *ArrMusik, n int) {
	var i, j, max int
	var temp Musik

	for i = 0; i < n-1; i++ {
		max = i

		for j = i + 1; j < n; j++ {
			if A[j].rating > A[max].rating {
				max = j
			}
		}

		temp = A[i]
		A[i] = A[max]
		A[max] = temp
	}
}

func insertionSortAsc(A *ArrMusik, n int) {
	var pass, i int
	var temp Musik

	for pass = 1; pass < n; pass++ {
		temp = A[pass]
		i = pass

		for i > 0 && temp.artis < A[i-1].artis {
			A[i] = A[i-1]
			i = i - 1
		}

		A[i] = temp
	}
}

func insertionSortDesc(A *ArrMusik, n int) {
	var pass, i int
	var temp Musik

	for pass = 1; pass < n; pass++ {
		temp = A[pass]
		i = pass

		for i > 0 && temp.artis > A[i-1].artis {
			A[i] = A[i-1]
			i = i - 1
		}

		A[i] = temp
	}
}

func binarySearch(A ArrMusik, n, x int) int {
	var left, right, mid, idx int

	left = 0
	right = n - 1
	idx = -1

	for left <= right {
		mid = (left + right) / 2

		if A[mid].rating == x {
			idx = mid
			left = right + 1
		} else if x < A[mid].rating {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return idx
}

func editMusik(A *ArrMusik, n int) {
	var cari string
	var idx int

	fmt.Print("Masukkan judul yang dicari : ")
	fmt.Scan(&cari)

	idx = sequentialSearch(*A, n, cari)

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Print("Judul baru : ")
		fmt.Scan(&A[idx].judul)

		fmt.Print("Artis baru : ")
		fmt.Scan(&A[idx].artis)

		fmt.Print("Genre baru : ")
		fmt.Scan(&A[idx].genre)

		fmt.Print("Rating baru : ")
		fmt.Scan(&A[idx].rating)

		fmt.Println("Data berhasil diubah")
	}
}

func hapusMusik(A *ArrMusik, n *int) {
	var cari string
	var idx, i int

	fmt.Print("Masukkan judul yang dihapus : ")
	fmt.Scan(&cari)

	idx = sequentialSearch(*A, *n, cari)

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}

		*n = *n - 1
		fmt.Println("Data berhasil dihapus")
	}
}

func main() {
	var A ArrMusik
	var n, pilih, idx, ratingCari int
	var cari string

	for pilih != 9 {
		fmt.Println("
===== MENU =====")
		fmt.Println("1. Tambah Musik")
		fmt.Println("2. Tampilkan Musik")
		fmt.Println("3. Edit Musik")
		fmt.Println("4. Hapus Musik")
		fmt.Println("5. Sequential Search")
		fmt.Println("6. Binary Search")
		fmt.Println("7. Selection Sort")
		fmt.Println("8. Insertion Sort")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahMusik(&A, &n)

		} else if pilih == 2 {
			tampilMusik(A, n)

		} else if pilih == 3 {
			editMusik(&A, n)

		} else if pilih == 4 {
			hapusMusik(&A, &n)

		} else if pilih == 5 {
			fmt.Print("Cari judul : ")
			fmt.Scan(&cari)

			idx = sequentialSearch(A, n, cari)

			if idx == -1 {
				fmt.Println("Data tidak ditemukan")
			} else {
				fmt.Println("Data ditemukan")
				fmt.Println("Judul :", A[idx].judul)
			}

		} else if pilih == 6 {
			selectionSortAsc(&A, n)

			fmt.Print("Cari rating : ")
			fmt.Scan(&ratingCari)

			idx = binarySearch(A, n, ratingCari)

			if idx == -1 {
				fmt.Println("Data tidak ditemukan")
			} else {
				fmt.Println("Data ditemukan")
				fmt.Println("Judul :", A[idx].judul)
			}

		} else if pilih == 7 {
			var p int

			fmt.Println("1. Ascending")
			fmt.Println("2. Descending")
			fmt.Print("Pilih : ")
			fmt.Scan(&p)

			if p == 1 {
				selectionSortAsc(&A, n)
			} else {
				selectionSortDesc(&A, n)
			}

			fmt.Println("Data berhasil diurutkan")

		} else if pilih == 8 {
			var p int

			fmt.Println("1. Ascending")
			fmt.Println("2. Descending")
			fmt.Print("Pilih : ")
			fmt.Scan(&p)

			if p == 1 {
				insertionSortAsc(&A, n)
			} else {
				insertionSortDesc(&A, n)
			}

			fmt.Println("Data berhasil diurutkan")

		} else if pilih == 9 {
			fmt.Println("Program selesai")
		}
	}
}
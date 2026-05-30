package main

import (
	"fmt"
	"strings"
)

const MAX = 100

type Musik struct {
	Judul  string
	Artis  string
	Genre  string
	Rating int
}

func tambahMusik(A *[MAX]Musik, n *int) {
	if *n >= MAX {
		fmt.Println("Daftar musik penuh!")
		return
	}
	var m Musik
	fmt.Print("Judul Lagu: ")
	fmt.Scanln(&m.Judul)
	fmt.Print("Artis: ")
	fmt.Scanln(&m.Artis)
	fmt.Print("Genre: ")
	fmt.Scanln(&m.Genre)
	fmt.Print("Rating (1-5): ")
	fmt.Scanln(&m.Rating)
	A[*n] = m
	*n++
}

func tampilMusik(A *[MAX]Musik, n int) {
	fmt.Println("Daftar Musik:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s - %s (%s) Rating: %d\n", i+1, A[i].Judul, A[i].Artis, A[i].Genre, A[i].Rating)
	}
}

func editMusik(A *[MAX]Musik, n int) {
	var idx int
	fmt.Print("Pilih nomor lagu untuk diubah: ")
	fmt.Scanln(&idx)
	idx--
	if idx < 0 || idx >= n {
		fmt.Println("Nomor salah!")
		return
	}
	fmt.Print("Judul Baru: ")
	fmt.Scanln(&A[idx].Judul)
	fmt.Print("Artis Baru: ")
	fmt.Scanln(&A[idx].Artis)
	fmt.Print("Genre Baru: ")
	fmt.Scanln(&A[idx].Genre)
	fmt.Print("Rating Baru: ")
	fmt.Scanln(&A[idx].Rating)
}

func hapusMusik(A *[MAX]Musik, n *int) {
	var idx int
	fmt.Print("Pilih nomor lagu untuk dihapus: ")
	fmt.Scanln(&idx)
	idx--
	if idx < 0 || idx >= *n {
		fmt.Println("Nomor salah!")
		return
	}
	for i := idx; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n--
}

func cariLagu(A *[MAX]Musik, n int) {
	var key string
	fmt.Print("Cari judul: ")
	fmt.Scanln(&key)
	found := false
	for i := 0; i < n; i++ {
		if strings.EqualFold(A[i].Judul, key) {
			fmt.Printf("%s - %s (%s) Rating: %d\n", A[i].Judul, A[i].Artis, A[i].Genre, A[i].Rating)
			found = true
		}
	}
	if !found {
		fmt.Println("Lagu tidak ditemukan!")
	}
}

func urutRating(A *[MAX]Musik, n int) {
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if A[j].Rating > A[maxIdx].Rating {
				maxIdx = j
			}
		}
		A[i], A[maxIdx] = A[maxIdx], A[i]
	}
	fmt.Println("Daftar musik sudah diurutkan berdasarkan rating.")
}

func urutArtis(A *[MAX]Musik, n int) {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && strings.ToLower(A[j].Artis) > strings.ToLower(key.Artis) {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}
	fmt.Println("Daftar musik sudah diurutkan berdasarkan artis.")
}

func main() {
	var A [MAX]Musik
	n := 0
	var pilih int
	for {
		fmt.Println("===== MENU =====")
		fmt.Println("1. Tambah Musik")
		fmt.Println("2. Tampilkan Musik")
		fmt.Println("3. Edit Musik")
		fmt.Println("4. Hapus Musik")
		fmt.Println("5. Cari Lagu")
		fmt.Println("6. Cari Artis")
		fmt.Println("7. Urut Rating")
		fmt.Println("8. Urut Artis")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			tambahMusik(&A, &n)
		case 2:
			tampilMusik(&A, n)
		case 3:
			editMusik(&A, n)
		case 4:
			hapusMusik(&A, &n)
		case 5:
			cariLagu(&A, n)
		case 6:
			fmt.Print("Cari artis: ")
			var artis string
			fmt.Scanln(&artis)
			found := false
			for i := 0; i < n; i++ {
				if strings.EqualFold(A[i].Artis, artis) {
					fmt.Printf("%s - %s (%s) Rating: %d\n", A[i].Judul, A[i].Artis, A[i].Genre, A[i].Rating)
					found = true
				}
			}
			if !found {
				fmt.Println("Artis tidak ditemukan!")
			}
		case 7:
			urutRating(&A, n)
		case 8:
			urutArtis(&A, n)
		case 9:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Menu tidak tersedia!")
		}
	}
}

package main
import "fmt"
const nmax = 1000
type produk struct {
	warna, kategori string
	ukuran, stok int
	harga float64
}
type tabproduk [nmax]produk

func inputukuran(d *produk, i int) {
	var ld, lb, pb, s int
	fmt.Printf("LD (Lingkar Dada) cm[int]> ")
	fmt.Scan(&ld)
	fmt.Printf("LB (Lebar Bahu) cm[int]> ")
	fmt.Scan(&lb)
	fmt.Printf("PB (Panjang Baju) cm[int]> ")
	fmt.Scan(&pb)
	s = i + 1
	d.ukuran = s * 1000000000 + ld * 1000000 + lb * 1000 + pb
	d.harga = float64(ld + lb + pb) / 100.0 * 250000.0 + 35000.0
	fmt.Printf("kode ukuran : %d\nHarga : Rp.%.2f \n", d.ukuran, d.harga)
}
func tambahproduk(banyakdatasaatini *int, d *tabproduk) {
	var i, h, wr, n, kategori int
	fmt.Printf("\n#Tambah Produk \n")
	fmt.Printf(" .../Produk/Tambah_Produk \n \n")
	fmt.Println("Banyak Produk :")
	fmt.Printf("input> ")
	fmt.Scan(&n)
	h = 0
	for i = *banyakdatasaatini; i < *banyakdatasaatini + n; i++ {
		h = h + 1
		fmt.Printf("\nProduk ke-%d\n", h)
		tabelwarnadankategori(0, 1)
		fmt.Printf("\nKategori[1/2/3/4/5]> ")
		fmt.Scan(&kategori)
		for !(kategori < 6 && kategori > 0) {
			fmt.Println("Pilihan tidak valid! Coba lagi...")
			tabelwarnadankategori(0, 1)
			fmt.Printf("\nKategori[1/2/3/4/5]> ")
			fmt.Scan(&kategori)
		}
		d[i].kategori = comp(0, kategori)
		tabelproduk(1, 1)
		inputukuran(&d[i], i)
		tabelproduk(1, 2)
		fmt.Printf("\nWarna[1/2/3/4/5]> ")
		fmt.Scan(&wr)
		if wr < 6 && wr > 0 {
			d[i].warna = comp(wr, 0)
		} else {
			for !(wr < 6 && wr > 0) {
				fmt.Println("Pilihan tidak valid! Coba lagi...")
				tabelwarnadankategori(1, 0)
				fmt.Printf("\nWarna[1/2/3/4/5]> ")
				fmt.Scan(&wr)
			}
			d[i].warna = comp(wr, 0)
		}
		fmt.Printf("Stok> ")
		fmt.Scan(&d[i].stok)
	}
	*banyakdatasaatini = *banyakdatasaatini + n
	fmt.Println("\n \nData Diperbaharui!\n \n")
}

func lihatdataproduk(banyakdatasaatini int, d tabproduk) {
	var i int
	fmt.Printf("\n%-20s | %-20s | %-20s | %-20s | %-20s\n", "Kategori", "Ukuran", "Warna", "Harga", "stok")
	for i = 0; i < banyakdatasaatini; i++ {
		fmt.Printf("%-20s | %-20d | %-20s | %-20.2f | %-20d\n", d[i].kategori, d[i].ukuran, d[i].warna, d[i].harga, d[i].stok)
	}
}

func lihatproduk(banyakdatasaatini int, d *tabproduk) {
    var n string
    var sub string
    fmt.Printf("\n#Lihat Produk \n")
    fmt.Println(".../lihat_Produk \n")
    lihatdataproduk(banyakdatasaatini, *d)
    for n != "bck" {
        tabelproduk(2, 0)
        fmt.Printf("\ninput[hrg/stk/bck]> ")
        fmt.Scan(&n)
        if n == "hrg" {
            fmt.Printf("\ninput[up/dn]> ")
            fmt.Scan(&sub)
            for sub != "up" && sub != "dn" {
                fmt.Println("Pilihan tidak valid! Coba lagi...")
                fmt.Printf("\ninput[up/dn]> ")
                fmt.Scan(&sub)
            }
            if sub == "dn" {
                decending(d, 1, 0, banyakdatasaatini)
                fmt.Println("Data yang diurutkan (descending) : ")
                lihatdataproduk(banyakdatasaatini, *d)
            } else if sub == "up" {
                acending(d, 1, 0, 0, banyakdatasaatini)
                fmt.Println("Data yang diurutkan (ascending) : ")
                lihatdataproduk(banyakdatasaatini, *d)
            }
            
        } else if n == "stk" {
            fmt.Printf("\ninput[up/dn]> ")
            fmt.Scan(&sub)
            for sub != "up" && sub != "dn" {
                fmt.Println("Pilihan tidak valid! Coba lagi...")
                fmt.Printf("\ninput[up/dn]> ")
                fmt.Scan(&sub)
            }
            if sub == "dn" {
                decending(d, 0, 1, banyakdatasaatini)
                fmt.Println("Data yang diurutkan (descending) : ")
                lihatdataproduk(banyakdatasaatini, *d)
            } else if sub == "up" {
                acending(d, 0, 0, 1, banyakdatasaatini)
                fmt.Println("Data yang diurutkan (ascending) : ")
                lihatdataproduk(banyakdatasaatini, *d)
            }
        } else if n != "bck" { 
			fmt.Println("Pilihan tidak valid! Coba lagi...")
        }
    }
}

func acending(d *tabproduk, harga, ukuran, stok, banyakdatasaatini int) {
	var i, idx, pass int
	var temp produk
	if harga == 1 {
		pass = 1
		for pass < banyakdatasaatini {
			temp = d[pass]
			i = pass
			for i > 0 && d[i - 1].harga > temp.harga {
				d[i] = d[i - 1]
				i = i - 1
			}
			d[i] = temp
			pass = pass + 1
		}
	}
	if stok == 1 {
		pass = 1
		for pass < banyakdatasaatini {
			temp = d[pass - 1]
			idx = pass - 1
			i = pass
			for i < banyakdatasaatini {
				if d[i].stok < d[idx].stok {
					idx = i
				}
				i = i + 1
			}
			temp = d[pass - 1]
			d[pass - 1] = d[idx]
			d[idx] = temp
			pass = pass + 1
		}
	}
	if ukuran == 1 {
		pass = 1
		for pass < banyakdatasaatini {
			temp = d[pass - 1]
			idx = pass - 1
			i = pass
			for i < banyakdatasaatini {
				if d[i].ukuran < d[idx].ukuran {
					idx = i
				}
				i = i + 1
			}
			temp = d[pass - 1]
			d[pass - 1] = d[idx]
			d[idx] = temp
			pass = pass + 1
		}
	}
}

func decending(d *tabproduk, harga, stok, banyakdatasaatini int) {
	var i, idx, pass int
	var temp produk
	if harga == 1 {
		pass = 1
		for pass < banyakdatasaatini {
			idx = pass - 1
			i = pass
			for i < banyakdatasaatini {
				if d[idx].harga < d[i].harga {
					idx = i
				}
				i = i + 1
			}
			temp = d[pass - 1]
			d[pass - 1] = d[idx]
			d[idx] = temp
			pass = pass + 1
		}
	}
	if stok == 1 {
		pass = 1
		for pass < banyakdatasaatini {
			temp = d[pass]
			i = pass
			for i > 0 && d[i - 1].stok < temp.stok {
				d[i] = d[i - 1]
				i = i - 1
			}
			d[i] = temp
			pass = pass + 1
		}
	}
}

func editproduk(banyakdatasaatini int, d *tabproduk) {
	var k, namal int
	var n string
	fmt.Println("Cari data produk (input ukuran) : ")
	fmt.Println()
	lihatdataproduk(banyakdatasaatini, *d)
	fmt.Print("\n \nInput[Ukuran]> ")
	fmt.Scan(&k)
	acending(d, 0, 1, 0, banyakdatasaatini)
	namal = cariukuranproduk(*d, banyakdatasaatini, k)
	if namal == -1 {
		fmt.Printf("Data Tidak Ditemukan! Coba lagi atau keluar...[y/any]\n \n")
		fmt.Printf("\nInput[y/any]> ")
		fmt.Scan(&n)
		if n == "y" {
			editproduk(banyakdatasaatini, d)
		}
	} else {
		fmt.Printf("\n \nData ditemukan : \n \n")
		lihatdataproduk(1, tabproduk{d[namal]})
		fmt.Printf("\nconfirm[n/any]> ")
		fmt.Scan(&n)
		if n == "n" {
			editproduk(banyakdatasaatini, d)
		} else {
			tabelproduk(3, 0)
			fmt.Println()
			fmt.Print("input[ktg/uk/wr/stk/ok]> ")
			fmt.Scan(&n)
			for n != "ok" {
				if n == "ktg" {
					tabelwarnadankategori(0, 1)
					fmt.Printf("input[Kategori baru]\n> ")
					fmt.Scan(&k)
					for !(k < 6 && k > 0) {
						fmt.Println("Pilihan tidak valid! Coba lagi...")
						tabelwarnadankategori(0, 1)
						fmt.Scan(&k)
					}
					d[namal].kategori = comp(0, k)
					fmt.Printf("\nconfirm[ok/ktg/uk/wr/stk]> ")
					fmt.Scan(&n)
				} else if n == "uk" {
					inputukuran(&d[namal], namal)
					fmt.Printf("\nconfirm[ok/ktg/uk/wr/stk]> ")
					fmt.Scan(&n)
				} else if n == "wr" {
					tabelwarnadankategori(1, 0)
					fmt.Printf("\ninput[Warna baru]\n> ")
					fmt.Scan(&k)
					for !(k < 6 && k > 0) {
						fmt.Println("Pilihan tidak valid! Coba lagi...")
						tabelwarnadankategori(1, 0)
						fmt.Scan(&k)
					}
					d[namal].warna = comp(k, 0)
					fmt.Printf("\n \nconfirm[ok/ktg/uk/wr/stk]> ")
					fmt.Scan(&n)
				} else if n == "stk" {
					fmt.Printf("input[Stok baru]\n> ")
					fmt.Scan(&d[namal].stok)
					fmt.Printf("\n \nconfirm[ok/ktg/uk/wr/stk]> ")
					fmt.Scan(&n)
				} else if n == "ok" {

				} else {
					fmt.Println("Pilihan tidak valid! Coba lagi...")
					fmt.Printf("\n \nconfirm[ktg/uk/wr/stk/ok]> ")
					fmt.Scan(&n)
				}
			}
		}
	}
}

func hapusproduk(banyakdatasaatini *int, d *tabproduk) {
	var n string
	var i, ad, k int
	fmt.Println("Cari data produk (input ukuran) : ")
	fmt.Println()
	lihatdataproduk(*banyakdatasaatini, *d)
	fmt.Printf("\n \nInput[Ukuran]> ")
	fmt.Scan(&k)
	acending(d, 0, 1, 0, *banyakdatasaatini)
	ad = cariukuranproduk(*d, *banyakdatasaatini, k)
	if ad == -1 {
		fmt.Printf("Data Tidak Ditemukan! Coba lagi atau keluar...[y/any]\n \n")
		fmt.Printf("\nInput[y/any]> ")
		fmt.Scan(&n)
		if n == "y" {
			hapusproduk(banyakdatasaatini, d)
		}
	} else {
		fmt.Printf("\n\nData ditemukan : \n\n")
		fmt.Printf("\n%-20s | %-20s | %-20s | %-20s | %-20s\n", "Kategori", "Ukuran", "Warna", "Harga", "stok")
		fmt.Printf("%-20s | %-20d | %-20s | %-20.2f | %-20d\n", d[ad].kategori, d[ad].ukuran, d[ad].warna, d[ad].harga, d[ad].stok)
		fmt.Printf("\nconfirm[y/any]> ")
		fmt.Scan(&n)
		if n == "y" {
			for i = ad; i < *banyakdatasaatini - 1; i++ {
				d[i] = d[i + 1]
			}
			*banyakdatasaatini = *banyakdatasaatini - 1
			fmt.Println("Data berhasil dihapus!")
		} else {
			fmt.Println("Penghapusan dibatalkan.")
		}
	}
}

func pengaturanproduk(d *tabproduk, banyakdatasaatini *int) {
	var pilih int
	for pilih != 5 {
		tabelproduk(0, 0)
		fmt.Printf("\ninput[1/2/3/4/5]> ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahproduk(banyakdatasaatini, d)
		} else if pilih == 2 {
			lihatproduk(*banyakdatasaatini, d)
		} else if pilih == 3 {
			editproduk(*banyakdatasaatini, d)
		} else if pilih == 4 {
			hapusproduk(banyakdatasaatini, d)
		} else if pilih == 5 {
			fmt.Printf("Kembali ke Main... \n \n")
		} else {
			fmt.Println("Pilihan tidak valid! Coba lagi...")
		}
	}
}

func cariproduk(d tabproduk, banyakdatasaatini int) {
	var i, n, ketemu int
	var pilih, wrn string
	var ketemu2 bool
	ketemu2 = false
	tabelmain(2, 0)
	fmt.Printf("\n\ninput[uk/wr]> ")
	fmt.Scan(&pilih)
	if pilih == "uk" {
		fmt.Printf("\n\nInput ukuran> ")
		fmt.Scan(&n)
		acending(&d, 0, 1, 0, banyakdatasaatini)
		ketemu = cariukuranproduk(d, banyakdatasaatini, n)
		if ketemu == -1 {
			fmt.Printf("\n\nData Tidak Ditemukan! \n \n")
		} else {
			fmt.Printf("\n\nData ditemukan:\n")
			fmt.Printf("\n%-20s | %-20s | %-20s | %-20s | %-20s\n", "Kategori", "Ukuran", "Warna", "Harga", "Stok")
			fmt.Printf("%-20s | %-20d | %-20s | %-20.2f | %-20d\n", d[ketemu].kategori, d[ketemu].ukuran, d[ketemu].warna, d[ketemu].harga, d[ketemu].stok)
		}
	} else if pilih == "wr" {
		tabelmain(3, 2)
		fmt.Printf("\ninput> ")
		fmt.Scan(&pilih)
		for pilih != "1" && pilih != "2" && pilih != "3" && pilih != "4" && pilih != "5" {
			fmt.Println("Pilihan tidak valid! Coba lagi...")
			tabelmain(3, 2)
			fmt.Printf("\ninput> ")
			fmt.Scan(&pilih)
		}
		if pilih == "1" {
			wrn = "Merah"
		} else if pilih == "2" {
			wrn = "Biru"
		} else if pilih == "3" {
			wrn = "Hijau"
		} else if pilih == "4" {
			wrn = "Kuning"
		} else if pilih == "5" {
			wrn = "Putih"
		}
		for i = 0; i < banyakdatasaatini; i++ {
			if d[i].warna == wrn {
				if ketemu2 == false {
					fmt.Printf("\nData ditemukan : \n")
					fmt.Printf("\n%-20s | %-20s | %-20s | %-20s | %-20s\n", "Kategori", "Ukuran", "Warna", "Harga", "Stok")
					ketemu2 = true
				}
				fmt.Printf("%-20s | %-20d | %-20s | %-20.2f | %-20d\n", d[i].kategori, d[i].ukuran, d[i].warna, d[i].harga, d[i].stok)
			}
		}
		if ketemu2 == false {
			fmt.Printf("Data Tidak Ditemukan! \n \n")
		}
	}
	fmt.Printf("\n\nconfirm[any]> ")
	fmt.Scan(&pilih)
}

func cariukuranproduk(d tabproduk, banyakdatasaatini, cariukuran int) int {
	var l, m, h int
	l = 0
	h = banyakdatasaatini - 1
	for l <= h {
		m = l + (h - l) / 2
		if d[m].ukuran == cariukuran {
			return m
		} else if d[m].ukuran < cariukuran {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return -1
}

func ringkasanstatistik(d tabproduk, banyakdatasaatini int) {
	var i int
	var jd, jt, jb, jm, jk, sd, st, sb, sm, sk, max int
	var n, p string
	for i = 0; i < banyakdatasaatini; i++ {
		if d[i].kategori == "Dress" {
			jd = jd + 1
			sd = sd + d[i].stok
		} else if d[i].kategori == "Tuxedo" {
			jt = jt + 1
			st = st + d[i].stok
		} else if d[i].kategori == "Business Suit" {
			jb = jb + 1
			sb = sb + d[i].stok
		} else if d[i].kategori == "Mantel Custom" {
			jm = jm + 1
			sm = sm + d[i].stok
		} else if d[i].kategori == "Kemeja Custom" {
			jk = jk + 1
			sk = sk + d[i].stok
		}
	}
	p = "Dress"
	max = jd
	if jt > max {
		p = "Tuxedo"
		max = jt
	}
	if jb > max {
		p = "Business Suit"
		max = jb
	}
	if jm > max {
		p = "Mantel Custom"
		max = jm
	}
	if jk > max {
		p = "Kemeja Custom"
		max = jk
	}
	fmt.Printf("\n\n#Ringkasan Statistik \n\n")
	fmt.Printf("%-20s | %-20s | %-20s\n", "Kategori", "Jumlah Produk", "Sisa Total Stok")
	fmt.Printf("%-20s | %-20d | %-20d\n", "Dress", jd, sd)
	fmt.Printf("%-20s | %-20d | %-20d\n", "Tuxedo", jt, st)
	fmt.Printf("%-20s | %-20d | %-20d\n", "Business Suit", jb, sb)
	fmt.Printf("%-20s | %-20d | %-20d\n", "Mantel Custom", jm, sm)
	fmt.Printf("%-20s | %-20d | %-20d\n", "Kemeja Custom", jk, sk)
	fmt.Printf("\nKategori paling populer : %s (%d produk)\n\n", p, max)
	fmt.Printf("confirm[any]> ")
	fmt.Scan(&n)
}

func tabelmain(tm, tms int) {
	if tm == 0 {
		fmt.Printf("SiFashion (Sistem Manajemen Inventaris Produk Fashion) \n\n")
		fmt.Printf(" /main \n\n")
		fmt.Println("Pilihan :")
		fmt.Println(" 1 = Pengaturan Produk")
		fmt.Println(" 2 = Cari Data Produk")
		fmt.Println(" 3 = Lihat Produk")
		fmt.Println(" 4 = Ringkasan Statistik")
		fmt.Println(" 5 = Selesai")
	} else if tm == 2 {
		fmt.Printf("\n#Cari Produk \n")
		fmt.Printf("main/Cari_Produk \n\n")
		fmt.Println("Pilihan :")
		fmt.Println(" uk = Ukuran (LD/LB/PB)")
		fmt.Println(" wr = Warna")
		fmt.Println(" stk = Stok")
	} else if tms == 2 && tm == 3 {
		fmt.Println("\nCari Berdasarkan Warna :")
		tabelwarnadankategori(1, 0)
	}
}
func tabelproduk(tpe, tpi int) {
	if tpe == 0 {
		fmt.Printf("\n#Pengaturan Produk \n")
		fmt.Printf(" .../Produk \n\n")
		fmt.Println("Pilihan :")
		fmt.Println(" 1 = Tambah Produk")
		fmt.Println(" 2 = Lihat Produk")
		fmt.Println(" 3 = Edit Produk")
		fmt.Println(" 4 = Hapus Produk")
		fmt.Println(" 5 = Kembali")
	} else if tpe == 1 && tpi == 1 {
		fmt.Printf("\n#Tambah Produk \n")
		fmt.Printf(" .../Produk/Tambah_Produk-Input_Ukuran \n\n")
		fmt.Println("Input Ukuran (dalam cm) :")
	} else if tpe == 1 && tpi == 2 {
		fmt.Printf("\n#Tambah Produk \n")
		fmt.Printf(" .../Produk/Tambah_Produk-Pilihan_Warna \n\n")
		fmt.Println("Pilihan Warna :")
		tabelwarnadankategori(1, 0)
	} else if tpe == 2 {
		fmt.Printf("\n#Lihat Produk \n")
		fmt.Printf(" .../Produk/Lihat_Produk \n\n")
		fmt.Println("Urutkan berdasarkan :")
		fmt.Println(" harga[up/dn] = Berdasarkan Harga (Up/Down)")
		fmt.Println(" stok[up/dn] = Berdasarkan Stok (Up/Down)")
		fmt.Println(" bck = Kembali")
	} else if tpe == 3 {
		fmt.Printf("\n#Daftar Edit Produk \n")
		fmt.Printf(" .../Produk/Edit_Produk/Daftar_Edit_Produk \n\n")
		fmt.Println("ktg  = Kategori")
		fmt.Println("uk  = Ukuran (LD/LB/PB)")
		fmt.Println("wr  = Warna")
		fmt.Println("stk = Stok")
		fmt.Println("ok  = konfirmasi")
	}
}
func tabelwarnadankategori(tw, tk int) {
	if tw == 1 {
		fmt.Println(" 1 = Merah")
		fmt.Println(" 2 = Biru")
		fmt.Println(" 3 = Hijau")
		fmt.Println(" 4 = Kuning")
		fmt.Println(" 5 = Putih")
	} else if tk == 1 {
		fmt.Println(" 1 = Dress")
		fmt.Println(" 2 = Tuxedo")
		fmt.Println(" 3 = Business Suit")
		fmt.Println(" 4 = Mantel Custom")
		fmt.Println(" 5 = Kemeja Custom")
	}
}

func comp(warna, kategori int) string {
	if warna == 1 {
		return "Merah"
	} else if warna == 2 {
		return "Biru"
	} else if warna == 3 {
		return "Hijau"
	} else if warna == 4 {
		return "Kuning"
	} else if warna == 5 {
		return "Putih"
	}
	if kategori == 1 {
		return "Dress"
	} else if kategori == 2 {
		return "Tuxedo"
	} else if kategori == 3 {
		return "Business Suit"
	} else if kategori == 4 {
		return "Mantel Custom"
	} else if kategori == 5 {
		return "Kemeja Custom"
	}
	return ""
}

func main() {
	var pilih int
	var d tabproduk
	var banyakdatasaatini int
	for pilih != 5 {
		tabelmain(0, 0)
		fmt.Printf("\ninput[1/2/3/4/5]> ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			pengaturanproduk(&d, &banyakdatasaatini)
		} else if pilih == 2 {
			cariproduk(d, banyakdatasaatini)
		} else if pilih == 3 {
			lihatproduk(banyakdatasaatini, &d)
		} else if pilih == 4 {
			ringkasanstatistik(d, banyakdatasaatini)
		} else if pilih == 5 {
			fmt.Printf("\n\nProgram selesai.")
		} else {
			fmt.Println("Pilihan tidak valid! Coba lagi...")
		}
	}
}

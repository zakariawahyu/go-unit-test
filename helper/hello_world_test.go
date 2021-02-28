package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

/**
Menggagalkan Unit Test
- Menggagalkan unit test menggunakan panic bukanlah hal yang bagus.
- Golang sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan testing.T
- Terdapat function Fail(), FailNow(), Error(), Fatal()
 */

/**
t.Fail() dan t.FailNow()
- Terdapat dua fintion untuk menggagalkan unit test, yaitu Fail() dan  FailNow()
- Fail() akan menggagalkan unit test, namun akan tetap melanjutkan ekseskusi unit test, tapi diakhir ketika selesai maka unit test tersebut dianggal gagal
- FailNow() akan menggagalkan unit test saat itu juga tanpa melanjutkan eksekusi unit test atau dipaksa untuk berhenti
 */

/**
t.Error(args...) dan t.Fatal(args...)
- Error() function lebih seperti melakukan log print ketika error, namun setelah melakukan log eorror, maka akan secara otomatis memanggil function Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai
- Fatal() mirip dengan Error(), hanya saja setelah melakukan log error akan memanggil funtion FailNow() , sehingga dipaksa berhenti unit testingnya
 */

/**
Assertion
- Melakukan pengecekan di unit test secara manual menggunakan if else sangatlah menyebalkan apalagi result sata yang di cek itu banyak
- Oleh karena itu disarankan menggunakan assertion untuk melakukan pengecekan
- Sayangnya di golang tidak menyediakan package untuk assertion, sehingga kita butuh menambahkan library untuk melakukan assertion

Testify
-Salah satu library yang paling populer di golang adalah Testify
-Kita bisa menggunakan library ini untuk melakukan assertion terhadap result data di unit test
- https://github.com/stretchr/testify

assert vs require
- Saat kita menggunakan assert, jika pengecekan gagal maka assert akan memanggil Fail() artinya eksekusi unit test teap dilanjutkan
- Sedangkan jika kita menggunakan require, jika pengecekan gagal maka  akan memanggil FailNow() artinya dipaksa berhenti dan eksekusi unit test tidak akan dilanjutkan
 */

/**
Skip Test
- Kadang dalam kondisi tertentu , kita ingin membatalkan eksekusi unit test
- Di golang juga kita bisa membatalkan eksekusinya jika kita mau
- Untuk membatalkan atau skip unit test bisa menggunakan fucntion Skip()
 */

/**
Before and After Test
- Biasanya dalam unit test kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit test dieksekusi
- Jikalau kode yang kita lakukan sebelum dan setelah selalu sama antar unit test di setiap funtionnya, maka membuat manual di unit test funtionnya adalah hal yang membosankan dan terlalu banyak kode duplikat jadinya
- Di golang semua itu bisa dihandle dengan fitur yang bernama testing.M
- Fitur ini bernama main, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga bisa kita gunakan untuk melakukan Before dan After unit test
 */

/**
Sub Test
- Golang mendukung fitur pembuatan function unit test di dalam function unit test
- Fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test di bahasa pemograman lainnya
- Utuk membuat unit test kita bisa memanfaatkan function Run() di struct T

Menjalanjan hanya Sub Test saja
- Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa menggunakan perintah: go test -run TestNamaFunction
- Jika kita ingin menjalankan hanya salah satu sub test, kita bisa menggunakan perintah" go test -run TestNamaFunction/NamaSubTest
- Atau untuk semua test semua sub test di test di semua function, kita bisa menggunakan perintah: go test -run /NamaSubTest
 */

/**
Table test
- Jika diperhatikan, sebenarnya dengan sub test kita bisa membuat test secara dinamis
- Dan fitur sub test ini biasa digunakan programmer golang untuk membuat test dengan konsep table test
- Table test yaitu dimana kita menyediakan data berupa slice yang berisi parameter, request dan ekspektasi hasil dari unit test
- Lalu slice tersebut kita literasi menggunakan sub test
 */

/**
Benchmark
- Selain unit test, golang testing package juga mendukung untuk melakukan benchmark
- Benchmark adalah mekanisme menghitung kecepatan performa kode aplikasi kita
- Benchmark di golang dilakukan dengan cara otomatis melakukan iterasi kode yang kita panggil berkali-kali sampai waktu tertentu
- Kita tidak perlu menentukan jumlah iterasi dan lamanya benchmark, karena itu sudah diatur oleh testing.B bawaan dari package testing

testing.B
- testing.B adalah struct yang digunakan untuk melakukan benchamrk
- testing.B mirip dengan testing.T, terdapat function Fail(), FailNow(), Fatal() dll
- yang memberdakan ada beberapa attribute dan function tambahan yan gdigunakan untuk melakukan benchmark
- salah satuny adalah attribute N, ini digunakan untuk melakukan total iterasi sebuah benchmark

Cara Kerja Benhmark
- Cara kerja benchmark di golang sederhana, hanya perlu membuat perulangan sejumlah N attribute
- Nanti secara otomatis lalu mendeteksi beberapa lama proses tersebut berjalan dan disimpulkan performa benchmark nya dalam waktu

Membuat Benchmark
Benchmark Function
- Mirip seperti unit test, untuk benchmark pun di golang sudah ditentukan nama functionnya. harus diawali dengan kata Benchmark, misal BenchmarkHelloWorld, Bencmarkxxx
- Selain itu harus memiliki parameter (b *testing.B)
- Tidak boleh mengembalikan return value
- Untuk nama file benchmark sama seperti unit test, diakhiri dengan _test misal helllo_world_test

Menjalankan Benchmark
- Untuk menjalankan seluruh benchmark di seluruh modul tapi unit test juga berjalan, kita bisa menggunakan perintah sama seperti unit test namun ditambahkan parameter bench. contoh : go test -v -bench=.
- Jika ktia ingin menjalankan semua benchmark tanpa unit test, kita bisa gunakan perintah : go test -v -run=NoMatchUnitTest -bench=.
- Jika ktia ingin menjalankan salah satu benchmark dan tanpa unit test, kita bisa gunakan go test -v -run=NotMatchTest -bench=NamaBenchmarkTest
- Jika kita ingin menjalankan benchmark di root module atau semuanya, kita bisa gunakan perintah go test -v -bench=. ./...

Sub Benchmark
- Sama seperti testing.T, di testing.B kita juga bisa membuat sub benchmark menggunakan function run()

Menjalankan hanya sub benchmark
- Saat kita menjalankan benchmark function, maka semua sub benchmark akan berjalan
- Namun kita ingin menjalankan salah satu sub benchmark saja, kita bisa gunakan perintah : go test -v -bench=NamaBenchamrk/NamaSubBenchamrk
 */

// ketika fail maka akan tetap lanjut tapi diakhir dianggap gagal
func TestHelloZakaria(t *testing.T) {
	result := HelloWorld("Zakaria")

	if result != "Hello Zakaria Test.." {
		// menggagalkan unit test
		t.Fail()
	}

	fmt.Println("Test Hello Zakaria Done....")
}


// ketika failnow akan dipaksa berhenti dan tidak melanjutkan unit test
func TestHelloWahyu(t *testing.T) {
	result := HelloWorld("Wahyu")

	if result != "Hello Wahyu Test.." {
		// menggagalkan unit test
		t.FailNow()
	}

	fmt.Println("Test Hello Wahyu Done....")
}

// ketika Error() maka bisa mencetak pesan error
// kemudian otomatis memanggil Fail(). akan tetap lanjut tapi diakhir dianggap gagal
func TestHelloNur(t *testing.T) {
	result := HelloWorld("Nur")

	if result != "Hello Nur Test.." {
		// menggagalkan unit test
		t.Error("This result not Hello Nur")
	}

	fmt.Println("Test Hello Nur Done....")
}

// ketika Fatal() maka bisa mencetak pesan error
// kemudian otomatis memanggil FailNow(). akan dipaksa berhenti dan tidak melanjutkan unit test
func TestHelloUtomo(t *testing.T) {
	result := HelloWorld("Utomo")

	if result != "Hello Utomo Test.." {
		// menggagalkan unit test
		t.Fatal("This result not Hello Utomo")
	}

	fmt.Println("Test Hello Utomo Done....")
}

// Contoh unit test berhasil
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Zakaria")

	if result != "Hello Zakaria" {
		// untuk memunculkan error
		panic("Result is not Hello Zakaria")
	}
}

// sengaja di errorkan dan menggunakan panic
// tidak disarankan menggagalkan error dengan panic
func TestHelloWorldZakaria(t *testing.T) {
	result := HelloWorld("Error")

	if result != "Hello Error" {
		// untuk memunculkan error
		panic("Result is not Hello Zakaria")
	}
}

// pembandung menggunakan assert
// Jika gagal maka akan memanggil Fail()
func TestHelloWorldAssert(t *testing.T)  {
	result := HelloWorld("Zakaria")
	assert.Equal(t, "Hello Zakaria Cek...", result, "Results must be 'Hello Zakaria Cek...'")
	fmt.Println("Test selesai")
}

// pembandung menggunakan require
// Jika gagal maka akan memanggil FailNow()
func TestHelloWorldRequire(t *testing.T)  {
	result := HelloWorld("Zakaria")
	require.Equal(t, "Hello Zakaria Cek...", result, "Results must be 'Hello Zakaria Cek...'")
	fmt.Println("Test selesai")
}

// Biasanya ada kondisi diawal untuk menanyakan apakah perlu di skip atau tidak
// Jika sudah sampai funstion Skip() maka kode program unit test dibawahnya tidak akan dijalankan
func TestHelloWorldSkip(t *testing.T)  {
	if runtime.GOOS == "windows" {
		t.Skip("Unit test can't run in operating system windows")
	}
	result := HelloWorld("Zakaria")
	require.Equal(t, "Hello Zakaria Cek...", result, "Results must be 'Hello Zakaria Cek...'")
	fmt.Println("Test selesai")
}

func TestMain(m *testing.M)  {
	// Before unit test
	// Akan dieksekusi sebelum semua unit test dieksekusi pada package ini
	fmt.Println("BEFORE UNIT TEST")

	// Akan mengeksekui semua function unit test pada package ini
	m.Run()

	// After unit test
	// Setelah semua function berhasil dieksekusi maka kode program setelah Run() akan jalan
	fmt.Println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T)  {
	t.Run("Zakaria", func(t *testing.T) {
		result := HelloWorld("Zakaria")
		assert.Equal(t, "Hello Zakaria", result, "Result must be Hello Zakaria")
	})

	t.Run("Wahyu", func(t *testing.T) {
		result := HelloWorld("Wahyu")
		assert.Equal(t, "Hello Wahyu", result, "Result must be Hello Wahyu")
	})
}

func TestTableHelloWorld(t *testing.T)  {
	data := []struct{
		Nama string
		Request string
		Ecpected string
	}{
		{
			"Zakaria",
			"Zakaria",
			"Hello Zakaria",
		},
		{
			"Wahyu",
			"Wahyu",
			"Hello Wahyu",
		},
		{
			"Nur",
			"Nur",
			"Hello Nur",
		},
	}

	for _, field := range data{
		t.Run(field.Nama, func(t *testing.T) {
			result := HelloWorld(field.Request)
			assert.Equal(t, field.Ecpected, result )
		})
	}
}

func BenchmarkHelloWorldZakaria(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Zakaria")
	}
}

func BenchmarkHelloWorldWahyu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Wahyu")
	}
}

func BenchmarkHelloWorldSubBenchmark(b *testing.B) {
	b.Run("Nur", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Nur")
		}
	})

	b.Run("Utomo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Utomo")
		}
	})
}

func BenchmarkHelloWorldTableBenchmark(b *testing.B) {
	benchmarks := []struct{
		name string
		request string
	}{
		{
			"Nur",
			"Hello Nur",
		},
		{
			"Utomo",
			"Hello Utomo",
		},
	}

	for _, benchmark := range benchmarks{
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}




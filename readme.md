**Unit Test**
- Unit test fokus menguji bagian kode program terkecil, biasanya menguci sebuah method
- Unit test biasanya dibuat kecil dan cepat, oleh karena itu biasanya kadang kode unit test lebih banyak  dari kode program aslinya, karena semua skenario pengujian akan dicoba untuk unit test
- Unit test baisanya digunakan sebagai cara cepat untuk meningkatkan kualitas kode program kita

**Testing Package**
- Di bahasa pemograman lain, biasanya untuk implementasi unit test kita butuh sebuah library atau framework khusus untuk melakukan unit test. Contoh di php menggunakan libary phpunit
- Berbeda dengan golang, di golang unit test sudah disediakan sebuah package khusus bernama testing
- Selain untuk menjalankan unit test, di golang juga sudah disediakan perintahnya
- Hal ini membuat implementasi unit testing di golang sangatlah mudah dibanding dengan bahasa pemograman lainnya
- https://golang.org/pkg/testing

**testing.T**
- Golang menyediakan sebuah struct yang bernama testing.T
- Struct ini digunakan untuk melakukan unit test di golang

**testing.M**
- testing.M adalah struct yang disediakan golang untuk mengatur life cycle testing

**Aturan File Unit Test**
- Golang memiliki aturan cara membuat file khusus untuk testing
- Untuk membuat file unit test, kita harus menggunakan akhiran _test di setiap nama filenya
- Misal kita membuat file hello_world. artinya untuk membuat unit testnya kita harus menggunakan nama file hello_world_test.go
- Cara ini untuk penamaan lebih mudah

**Aturan Function Unit Test**
- Selain aturan nama file, di golang juga sudah diatur untuk nama function unit test
- Nama function unit test harus diawali dengan nama Test didepan nama functionnya
- Misal kita ingin mengetes function HelloWorld, maka nama function unit testnya adalah TaestHelloWorld
- Memiliki parameter (t *testing.T) dan tidak mengembalikan return value

**Menjalankan Unit Test**
- Untuk menjalankan unit test kita bisa menggunakan perintah : go test
- Jika kita ingin lebih detail untuk mendapatkan hasil unit testnya yang sudah di running, kita bisa gunakan perintah : go test -v
- Jika hanya ingin memilih salah satu function saja, kita bisa gunakan perintah : go test -v -run TestNamaFuntion
- Jika kita ingin menjalakan semua unit test dari folder utama ktia, kita bisa menggunakan perintah : go test ./...

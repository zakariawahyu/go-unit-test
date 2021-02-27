package repository

import (
	"github.com/stretchr/testify/mock"
	"go-unit-test/entity"
)

/**
Mock
- Mock adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil, di akan menghasilkan data yang sudah kita program diawal tadi
- Mock adalah salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari satu object yang memang sulit untuk di testing
- Misal kita ingin membuat unit test, namun ternyata ada kode program kita yang harus memanggil API Call ke third party services dan belum tentu response nya sesuai dengan apa yang kita mau

Testify Mock
- Untuk membuat mock object, tidak ada fitur bawaan di golang. Namun kita bisa menggunakan library testify
- Testify mendukung pembuatan moock object, sehingga cocok untuk kita gunakan ketika ingin membuat mock object
- Namun perlu diperhatikan, jika desain kdoe program kita jelek akan sulit untuk melakukan mocking, jadi pastikan kita melakukan pembuatan desain kode program kita dengan baik

Contoh kasus Aplikasi Quary ke Database
- Kita akan coba contoh kasus dengan membuat contoh aplikasi golang yang melakukan query ke database
- Dimana kita akan buat layer Services sebagai businnes logic dan layer repository sebagai jemabtan dengan database
- Agar kode program kita mudah di test, disarankan agar membuat kontrak berupa interface

 */

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (respository *CategoryRepositoryMock) FindById(id string) *entity.Category  {
	arguments := respository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}

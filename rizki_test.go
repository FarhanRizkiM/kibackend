package kibackend

import (
	"fmt"
	"testing"
)

// func TestInsertMahasiswa(t *testing.T) {
// 	dbname := "presensi"
// 	mahasiswa := Mahasiswa{
// 		ID:           primitive.NewObjectID(),
// 		Nama:         "Farhan Rizki Maulana",
// 		NPM:          "1214020",
// 		Semester:     "4",
// 		Kelas:        "2A",
// 		Prodi_kampus: "D4 Teknik Informatika",
// 	}
// 	insertedID := InsertMahasiswa(dbname, mahasiswa)
// 	if insertedID == nil {
// 		t.Error("Failed to insert user")
// 	}
// }

// func TestInsertPresensi(t *testing.T) {
// 	dbname := "presensi"
// 	presensi := Presensi{
// 		ID:         primitive.NewObjectID(),
// 		Datetime:   primitive.NewDateTimeFromTime(time.Now().UTC()),
// 		Kehadiran:  "HADIR",
// 		Keterangan: "Masuk",
// 	}
// 	insertedID := InsertPresensi(dbname, presensi)
// 	if insertedID == nil {
// 		t.Error("Failed to insert user")
// 	}
// }

// func TestInsertMataKuliah(t *testing.T) {
// 	dbname := "presensi"
// 	matakuliah := MataKuliah{
// 		ID:      primitive.NewObjectID(),
// 		KodeMK:  "TI41264",
// 		NamaMK:  "Pemrograman 3",
// 		SKS:     "3",
// 		Jurusan: "Teknik Informatika",
// 	}
// 	insertedID := InsertMataKuliah(dbname, matakuliah)
// 	if insertedID == nil {
// 		t.Error("Failed to insert user")
// 	}
// }

// func TestInsertJadwalKuliah(t *testing.T) {
// 	dbname := "presensi"
// 	jadwalkuliah := JadwalKuliah{
// 		ID:         primitive.NewObjectID(),
// 		MataKuliah: "Pemrograman 3",
// 		Hari:       "Senin",
// 		JamMulai:   "13:00",
// 		JamSelesai: "18:00",
// 		Ruangan:    "Lab 314",
// 	}
// 	insertedID := InsertJadwalKuliah(dbname, jadwalkuliah)
// 	if insertedID == nil {
// 		t.Error("Failed to insert user")
// 	}
// }

// func TestInsertDosen(t *testing.T) {
// 	dbname := "presensi"
// 	dosen := Dosen{
// 		ID:         primitive.NewObjectID(),
// 		NIDN:       "0410118609",
// 		Nama:       "Rolly Maulana Awangga,S.T.,MT.,CAIP, SFPC.",
// 		MataKuliah: "Pemrograman 3",
// 	}
// 	insertedID := InsertDosen(dbname, dosen)
// 	if insertedID == nil {
// 		t.Error("Failed to insert user")
// 	}
// }

func TestGetDataMahasiswa(t *testing.T) {
	stats := "Farhan Rizki M"
	data := GetDataMahasiswa(stats)
	fmt.Println(data)
}

func TestGetDataPresensi(t *testing.T) {
	stats := "HADIR"
	data := GetDataPresensi(stats)
	fmt.Println(data)
}

func TestGetDataMataKuliah(t *testing.T) {
	stats := "TI41264"
	data := GetDataMataKuliah(stats)
	fmt.Println(data)
}

func TestGetDataJadwalKuliah(t *testing.T) {
	stats := "Senin"
	data := GetDataJadwalKuliah(stats)
	fmt.Println(data)
}

func TestGetDataDosen(t *testing.T) {
	stats := "Rolly Maulana Awangga,S.T.,MT.,CAIP, SFPC."
	data := GetDataDosen(stats)
	fmt.Println(data)
}

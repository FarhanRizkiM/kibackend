package kibackend

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertMahasiswa(db string, mahasiswa Mahasiswa) (InsertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("mahasiswa").InsertOne(context.TODO(), mahasiswa)
	if err != nil {
		fmt.Printf("InsertMahasiswa: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPresensi(db string, presensi Presensi) (InsertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("presensi").InsertOne(context.TODO(), presensi)
	if err != nil {
		fmt.Printf("InsertPresensi: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertMataKuliah(db string, matakuliah MataKuliah) (InsertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("matakuliah").InsertOne(context.TODO(), matakuliah)
	if err != nil {
		fmt.Printf("InsertMataKuliah: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertJadwalKuliah(db string, jadwalkuliah JadwalKuliah) (InsertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("jadwalkuliah").InsertOne(context.TODO(), jadwalkuliah)
	if err != nil {
		fmt.Printf("InsertJadwalKuliah: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDosen(db string, dosen Dosen) (InsertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("dosen").InsertOne(context.TODO(), dosen)
	if err != nil {
		fmt.Printf("InsertDosen: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetDataMahasiswa(stats string) (data []Mahasiswa) {
	user := MongoConnect("presensi").Collection("mahasiswa")
	filter := bson.M{"nama": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataMahasiswa :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataPresensi(stats string) (data []Presensi) {
	user := MongoConnect("presensi").Collection("presensi")
	filter := bson.M{"kehadiran": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataPresensi :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataMataKuliah(stats string) (data []MataKuliah) {
	user := MongoConnect("presensi").Collection("matakuliah")
	filter := bson.M{"kode_mk": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataMataKuliah :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataJadwalKuliah(stats string) (data []JadwalKuliah) {
	user := MongoConnect("presensi").Collection("jadwalkuliah")
	filter := bson.M{"hari": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataJadwalKuliah :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataDosen(stats string) (data []Dosen) {
	user := MongoConnect("presensi").Collection("dosen")
	filter := bson.M{"nama": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataDosen :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

package rizki

import (
	"context"
	"fmt"
	"os"

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

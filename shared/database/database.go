package database

import (
	"encoding/json"
	"log"

	"github.com/boltdb/bolt"
)

var (
	// Database wrapper
	DB *bolt.DB
)

// Database is the details for the database connection
type Database struct {
	Name string
}

func Connect(d Database) {
	var err error
	if DB, err = bolt.Open(d.Name, 0644, nil); err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to database: %v\n", DB)
}

func Update(bucket_name string, key string, dataStruct interface{}) error {
	var err error
	DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket_name))
		if err != nil {
			return err
		}
		encoded_record, err := json.Marshal(dataStruct)
		if err != nil {
			return err
		}
		if err = bucket.Put([]byte(key), encoded_record); err != nil {
			return err
		}
		return err
	})
	return err
}

// func Get(bucket_name string, key string) (interface{}, error) {
// 	var err error
// 	DB.View(func(tx *bolt.Tx) error {
// 		bucket := tx.Bucket([]byte(bucket_name))
// 		v := bucket.Get([]byte(key))
// 		fmt.Printf("%+v\n", v)
// 		return err
// 	})
// 	DB.Update(func(tx *bolt.Tx) error {
// 		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket_name))
// 		if err != nil {
// 			return err
// 		}
// 		encoded_record, err := json.Marshal(dataStruct)
// 		if err != nil {
// 			return err
// 		}
// 		if err = bucket.Put([]byte(key), encoded_record); err != nil {
// 			return err
// 		}
// 		return err
// 	})
// 	return err
// }

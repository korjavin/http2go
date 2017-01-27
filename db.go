package main

import (
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

var (
	db *bolt.DB
)

func init() {
	var err error
	db, err = bolt.Open("my.db", 0666, &bolt.Options{ReadOnly: false})
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}
func loadusers(users map[string]user) error {
	return db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("users"))

		b.ForEach(func(k, v []byte) error {
			log.Printf("key=%s, value=%d\n", k, btoi(v))
			newuser := user{last_name: "", first_name: "", chat_id: btoi(v)}
			users[string(k)] = newuser
			return nil
		})
		return nil
	})
}

func saveuser(u user) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return err
		}
		codename := u.first_name + "_" + u.last_name
		// buf, err := json.Marshal(u.chat_id)
		// if err != nil {
		// 	return err
		// }
		// log.Println(buf)
		log.Printf("key=%s, value=%d\n", codename, u.chat_id)
		return b.Put([]byte(codename), itob(u.chat_id))
	})
}

func itob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
func btoi(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(b))
}

// func addmsg(m *msg) error {
//
// 	return db.Update(func(tx *bolt.Tx) error {
// 		b, err := tx.CreateBucketIfNotExists([]byte("posts"))
// 		if err != nil {
// 			return err
// 		}
// 		return b.Put([]byte("2015-01-01"), []byte("My New Year post"))
//
// 		//	_, err := tx.CreateBucketIfNotExists([]byte("msg"))
// 		//	if err != nil {
// 		//		return err
// 		//	}
//
// 		// buf, err := json.Marshal(&m)
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		// log.Println(buf)
//
// 		return nil
// 		//	return b.Put([]byte(m.regdt.Format(time.RFC3339)), buf)
// 	})
// }

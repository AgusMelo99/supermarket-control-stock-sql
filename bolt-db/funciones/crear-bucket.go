package funciones

import (
	bolt "go.etcd.io/bbolt"
)

func CrearBucket(db *bolt.DB, bucketName string, key []byte, val []byte) error {
	
	//abre transaccion de escritura
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	
	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))
	
	err = b.Put(key, val)
	if err != nil {
		return err
	}
	
	//cierra transaccion
	if err := tx.Commit(); err != nil {
		return err
	}
	
	return nil
}

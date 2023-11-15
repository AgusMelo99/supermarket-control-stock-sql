package funciones

func LeerBucket(db *bolt.DB, bucketName string, key []byte) ([]byte, error) {
	var buf []byte
	
	//abre transaccion de lectura
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		buf = b.Get(key)
		return nil
	})
	
	return buf, err
}

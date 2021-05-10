package helper

import (
	"errors"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

func MojoBrainPath() (string, error) {
	if path, exists := os.LookupEnv("HOME"); exists {
		return path + "/.mojo-brain.db", nil
	}
	return "", errors.New("home path doesnt exists")
}

func SetupMojoBrain() error {
	path, err := MojoBrainPath()
	if err != nil {
		return err
	}

	db, err := bolt.Open(path, 0644, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		return err
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Credentials"))
		if err != nil {
			return err
		}
		// same as commit to apply the transaction
		return nil
	})
	return nil
}

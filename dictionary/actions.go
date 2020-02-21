package dictionary

import (
	"bytes"
	"encoding/gob"
	"github.com/dgraph-io/badger"
	"sort"
	"strings"
	"time"
)

func (d *Dictionary) Add(word string, definition string) error {
	entry := Entry{
		Word:       strings.Title(word),
		Definition: definition,
		CreatedAt:  time.Now(),
	}

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(entry)
	if err != nil {
		return err
	}

	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(word), buffer.Bytes())
	})
}

func (d *Dictionary) Get(word string) (Entry, error) {
	var entry Entry
	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(word))
		if err != nil {
			return err
		}
		entry, err = getEntry(item)
		return err
	})
	return entry, err
}

// List retrieves all the dictionary content
// []string is an alphabetical sorted array with the words
// [string]Entry is a map of the words and their definition
func (d *Dictionary) List() ([]string, map[string]Entry, error) {
	entries := make(map[string]Entry)
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next(){
			item := it.Item()
			entry, err := getEntry(item)
			if err != nil {
				return err
			}
			entries[entry.Word] = entry
		}
		return nil
	})
	return sortedKeys(entries), entries, err
}

func (d *Dictionary) Remove(word string) error{
	err := d.db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(word))
		if err != nil{
			return err
		}
		return err
	})
	return err
}

func sortedKeys(entries map[string]Entry) []string{
	keys := make([]string, len(entries))
	for key := range entries {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func getEntry(item *badger.Item) (Entry, error) {
	var entry Entry
	var buffer bytes.Buffer
	err := item.Value(func(val []byte) error {
		_, err := buffer.Write(val)
		return err
	})
	dec := gob.NewDecoder(&buffer)
	err = dec.Decode(&entry)
	return entry, err
}

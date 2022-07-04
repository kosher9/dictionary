package main

import (
	"dictionary/dict"
	"flag"
	"fmt"
	"os"
)

func main() {

	action := flag.String("action", "list", "Action to perform")

	d, err := dict.New("./badger")
	handleError(err)
	defer d.Close()

	flag.Parse()
	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	default:
		fmt.Printf("No usage for %v\n", *action)
	}
}

func actionAdd(d *dict.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleError(err)
	fmt.Printf("\n'%v' added to the dict\n", word)
}

func actionList(d *dict.Dictionary) {
	words, entries, err := d.List()
	handleError(err)
	fmt.Println("\nDictionary content\n")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}
func actionDefine(d *dict.Dictionary, args []string) {
	word := args[0]
	entry, err := d.Get(word)
	handleError(err)
	fmt.Println(entry)
}

func actionRemove(d *dict.Dictionary, args []string) {
	word := args[0]
	err := d.Remove(word)
	handleError(err)
	fmt.Printf("\n'%v' has been removed\n", word)
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v", err)
		os.Exit(1)
	}
}

package main

import (
	"dictionnary/dictionary"
	"flag"
	"fmt"
	"os"
)

func main(){

	action := flag.String("action", "list", "Action to perform")

	d, err := dictionary.New("./badger")
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

func actionAdd(d *dictionary.Dictionary, args []string)  {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleError(err)
	fmt.Printf("\n'%v' added to the dictionary\n", word)
}

func actionList(d *dictionary.Dictionary){
	words, entries, err := d.List()
	handleError(err)
	fmt.Println("\nDictionary content\n")
	for _, word := range words{
		fmt.Println(entries[word])
	}
}
func actionDefine(d *dictionary.Dictionary, args []string)  {
	word := args[0]
	entry, err := d.Get(word)
	handleError(err)
	fmt.Println(entry)
}

func actionRemove(d *dictionary.Dictionary, args []string)  {
	word := args[0]
	err := d.Remove(word)
	handleError(err)
	fmt.Printf("\n'%v' has been removed\n", word)
}

func handleError(err error){
	if err != nil{
		fmt.Printf("Dictionary error: %v", err)
		os.Exit(1)
	}
}

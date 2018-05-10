package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/tobischo/gokeepasslib"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/user"
	"regexp"
	"strings"
	"syscall"
)

//var kdbx = "/Users/rex/Dropbox/bitbucket/home/de9Eithu2Vaivo5Coh5pes1oop6Out.kdbx"

func main() {
	usr, _ := user.Current()
	kdbx := fmt.Sprintf("%s/.pass.kdbx", usr.HomeDir)
	file, _ := os.Open(kdbx)

	db := gokeepasslib.NewDatabase()

	pwd := getpass("Database Password: ")
	db.Credentials = gokeepasslib.NewPasswordCredentials(pwd)
	_ = gokeepasslib.NewDecoder(file).Decode(db)
	db.UnlockProtectedEntries()

	// entry := db.Content.Root.Groups[0].Groups[0].Entries[0]
	// fmt.Println(entry.GetTitle())
	// fmt.Println(entry.GetPassword())

	search := os.Args[1]
	//fmt.Println("searching for " + search)
	rsearch, _ := regexp.Compile("(?i)" + search)
	found := make(map[string]string)

	for _, top := range db.Content.Root.Groups {
		for _, groups := range top.Groups {
			for _, entry := range groups.Entries {
				entry_path := fmt.Sprintf("%s/%s/%s", top.Name, groups.Name, entry.GetTitle())
				if rsearch.MatchString(entry_path) {
					fmt.Println(entry_path)
					found[entry_path] = entry.GetPassword()
				}
			}
		}
	}

	if len(found) == 1 {
		for _, found_pw := range found {
			if err := clipboard.WriteAll(string(found_pw)); err != nil {
				panic(err)
			} else {
				fmt.Println("entry written to clipboard")
			}
		}
	} else {
		fmt.Printf("found %d\n", len(found))
	}

}

func getpass(msg string) string {
	fmt.Print(msg)
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err == nil {
		fmt.Println("")
	} else {
		fmt.Println("\nError in ReadPassword")
	}
	password := string(bytePassword)

	return strings.TrimSpace(password)
}

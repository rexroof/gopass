package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/tobischo/gokeepasslib"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"os/user"
	"regexp"
	"strings"
	"syscall"
)

func main() {
	usr, _ := user.Current()
	default_kdbx := fmt.Sprintf("%s/.pass.kdbx", usr.HomeDir)

	kdbx_path := flag.String("kdbx", default_kdbx, "path to kdbx file")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "search pattern required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := os.Open(*kdbx_path)
	if err != nil {
		log.Fatal(err)
	}

	db := gokeepasslib.NewDatabase()
	pwd := getpass("Database Password: ")
	db.Credentials = gokeepasslib.NewPasswordCredentials(pwd)
	_ = gokeepasslib.NewDecoder(file).Decode(db)
	db.UnlockProtectedEntries()

	search := flag.Arg(0)
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

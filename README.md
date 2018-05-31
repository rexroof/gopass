# Simple KeePass CLI written in go

This tool for Keepass serves a very simple purpose:  search your KeePass DB and copy
a password into your clipboard.  More features will be added but this is the primary
purpose.

This is also my first project in Go, so the secondary purpose is to teach myself Go.

TODO:

* use exact match if there is one and there are multiple entries
* recognize when db pw is wrong
* set location of keepass db
* allow for changing passwords
* optional search on other elements
* optional read other data from entry into clipboard
* add new entries to database
* toggle case sensitivity

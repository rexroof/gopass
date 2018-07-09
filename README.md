# Simple KeePass CLI written in go

This tool for Keepass serves a very simple purpose:  search your KeePass DB and copy
a password into your clipboard.  More features will be added but this is the primary
purpose.

This is also my first project in Go, so the secondary purpose is to teach myself Go.

TODO:

* store files in db:
   use something like -notes to just output the notes section
   - also change it to be able to write a -notes file to the db.
   ( this could be used to store vpn profiles? )
* set location of keepass db
* allow for changing passwords
* optional search on other elements
* optional read other data from entry into clipboard
* add new entries to database
* toggle case sensitivity
* trap interrupt and fix stty if in password prompt

* make tests

* working with feature branches https://www.atlassian.com/git/tutorials/comparing-workflows/feature-branch-workflow

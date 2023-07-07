# DATA

This directory contains some files with the raw data from the database. 

It does not contain a standalone SQL initialization script, so that's up to the user.

## jepp.db

The jepp.db file is a SQLite3 database file, containing more or less the same data as the CSV files.

This is the current SQLite database file, used by the applications in this repository.

_In the past it was created with the mysql-to-sqlite3 tool, using the following commands:_
```bash
pip3 install mysql-to-sqlite3
mysql2sqlite -f jepp.db -d jeppdb -u jepp -p
<enter password>
```
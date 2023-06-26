# DATA

This directory contains some files with the from data from the database. 

It does not contain a standalone SQL initialization script, so table creation would need to be done manually before importing these files somewhere.

## dump.sql.gz

The dump.sql.gz file contains a `mysqld` dump of the database, including the table creation statements. 

It was created with this command:
```
mysqldump -u ${DB_USER} -p${DB_PASS} -h ${DB_HOST} -P ${DB_PORT} --skip-comments --column-statistics=0 --databases jeppdb > dump.sql

gzip dump.sql
```

It can be imported into a MySQL database with the following command:

```bash
gunzip -c dump.sql.gz

mysql -u <DB_USER> -p<DB_PASSWORD> -h <DB_HOST> -P <DB_PORT> -D <DB_NAME> < dump.sql
```
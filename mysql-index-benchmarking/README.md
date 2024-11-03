### Creating table with different primary key

test_table1 [id (INT) age (INT)]

test_table2 [id (VARCHAR) age (INT)]

Tried inserting a million records into both of these DB and here are some numbers

| **Table**   | **Records Inserted** | **Time** |
|-------------|----------------------|----------|
| test_table1 |        100,000       |  16.63s  |
| test_table2 |        100,000       |  17.00s  |

| **Table**   | **Records Inserted** | **Time** |
|-------------|----------------------|----------|
| test_table1 |       1,000,000      |   2m46s  |
| test_table2 |       1,000,000      |   2m48s  |


### Created Indexes on AGE column
`CREATE INDEX idx_age ON test.test_table1 (age);`

`CREATE INDEX idx_age ON test.test_table2 (age);`

##### Get the size of the indexes
##### Check the size of indexes
`SELECT database_name, table_name, index_name,
ROUND(stat_value * @@innodb_page_size / 1024 / 1024, 2) size_in_mb
FROM mysql.innodb_index_stats
WHERE stat_name = 'size'
ORDER BY size_in_mb DESC;`

| **DB** | **Table**   | **Index** | **Size in MB** |
|--------|-------------|-----------|----------------|
| test   | test_table1 | PRIMARY   | 28.56          |
| test   | test_table2 | PRIMARY   | 64.59          |
| test   | test_table1 | idx_age   | 16.55          |
| test   | test_table2 | idx_age   | 19.55          |
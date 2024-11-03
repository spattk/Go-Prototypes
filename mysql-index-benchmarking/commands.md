##### Creating Index
CREATE INDEX idx_age ON test.test_table1 (age);
CREATE INDEX idx_age ON test.test_table2 (age);

##### Check the size of indexes
SELECT database_name, table_name, index_name,
ROUND(stat_value * @@innodb_page_size / 1024 / 1024, 2) size_in_mb
FROM mysql.innodb_index_stats
WHERE stat_name = 'size'
ORDER BY size_in_mb DESC;
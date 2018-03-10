# Database

API database information.

## MySQL

Documentation about API using MySQL 8.0.

### Data fixture

Send [users CSV](https://s3.amazonaws.com/careers-picpay/users.csv.gz) file to MySQL container

    docker cp users.csv mysql:/var/lib/mysql-files/users.csv

Load data from CSV file (run and go drink a coffee [maybe two, or twenty])

    LOAD DATA INFILE '/var/lib/mysql-files/users.csv'
    IGNORE INTO TABLE `users`
    COLUMNS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"'
    LINES TERMINATED BY '\r\n'
    (@uuid, `name`, `username`)
    SET `id` = UUID_TO_BIN(@uuid);

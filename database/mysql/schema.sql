CREATE TABLE `users` (
  `id`       BINARY(16) NOT NULL,
  `name`     VARCHAR(150) NOT NULL,
  `username` VARCHAR(80) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE username_idx      (`username`),
  INDEX  name_idx          (`name`),
  INDEX  username_name_idx (`username`, `name`)
) ENGINE=InnoDB;

CREATE TABLE `accounts` (
  `username`    VARCHAR(30) NOT NULL,
  `password`    VARCHAR(60) NOT NULL,
  `created_at`  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB;

CREATE TABLE `tokens` (
  `username`    VARCHAR(30) NOT NULL,
  `token`       VARCHAR(64) NOT NULL,
  `created_at`  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`username`),
  FOREIGN KEY (`username`)
    REFERENCES `accounts` (`username`)
      ON DELETE CASCADE
      ON UPDATE CASCADE
) ENGINE=InnoDB;

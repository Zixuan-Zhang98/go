CREATE TABLE IF NOT EXISTS `user`(
    `id` INT UNSIGNED AUTO_INCREMENT KEY COMMENT 'user id',
    `username` VARCHAR(20) NOT NULL UNIQUE,
    `password` CHAR(32) NOT NULL,
    `email` VARCHAR(50) NOT NULL UNIQUE,
    `age` TINYINT UNSIGNED NOT NULL DEFAULT 18,
    `sex` ENUM('male', 'female', 'unknown') NOT NULL DEFAULT 'unknown',
    `tel` CHAR(11) NOT NULL UNIQUE,
    `addr` VARCHAR(50) NOT NULL DEFAULT 'US',
    `card` CHAR(18) NOT NULL UNIQUE,
    `married` BOOL NOT NULL DEFAULT 0 COMMENT '0 stands for not married, 1 stands for married',
    `salary` FLOAT(8) NOT NULL DEFAULT 0
) ENGINE=INNODB DEFAULT CHARSET=UTF8;

-- Specifying number of digits for floating point data types is deprecated and will be removed in a future release.
-- Integer display width is deprecated and will be removed in a future release.
-- 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous.
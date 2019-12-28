-- create database
CREATE DATABASE IF NOT EXISTS imooc DEFAULT CHARACTER SET = 'UTF8';

-- open database
USE imooc;

-- user table
CREATE TABLE IF NOT EXISTS imooc_user(
  id INT,
  username VARCHAR(20),
  password VARCHAR(32),
  email VARCHAR(50),
  age TINYINT,
  card CHAR(18),
  tel CHAR(11),
  salary FLOAT(8,2),
  married TINYINT(1),
  addr VARCHAR(100),
  sex ENUM('male','female','unknown')
)ENGINE=INNODB CHARSET=UTF8;

-- INSERT [INTO] tal_name(id, username, ...) VALUES(1, "K", ...);
INSERT imooc_user(id, username, password, email, age, card, tel, salary, married, addr, sex)
VALUES(1, 'k', 'k', 'abc@ss.com', 24, '230112197809871234', '18635579617', 88888.68, 0, 'Minnesota', 'male');

-- SELECT * FROM tal_name;
SELECT * FROM imooc_user;

INSERT imooc_user(id, username, password, email, age, card, tel, salary, married, addr, sex)
VALUES(-10, 'k', 'k', 'abc@ss.com', 10, '230112197809871234', '18635579617', 2345.68, 1, 'Minnesota', 'female');

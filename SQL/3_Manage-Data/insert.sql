CREATE DATABASE IF NOT EXISTS test_data DEFAULT CHARACTER SET = UTF8;
USE test_data;
CREATE TABLE IF NOT EXISTS user(
    id INT UNSIGNED AUTO_INCREMENT KEY,
    username VARCHAR(20) NOT NULL UNIQUE,
    age TINYINT UNSIGNED DEFAULT 19,
    email VARCHAR(50) NOT NULL
)ENGINE=INNODB CHARSET=UTF8;

-- 不指定字段名称
INSERT user VALUE(1, 'king', 24, 'king@qq.com');
INSERT user VALUE(NULL, 'queen', 22, 'queen@qq.com');
INSERT user VALUE(DEFAULT , 'prince', 4, 'prince@qq.com');

-- 列出指定字段的形式
INSERT user(username, email) VALUES('rose', 'rose@qq.com');

INSERT user SET username='D', age=45, email='d@gmail.com';

CREATE TABLE test(
    a VARCHAR(20)
);
INSERT test VALUES('AA'), ('BB'), ('CC');

ALTER TABLE user ALTER email SET DEFAULT 'xxx@gmail.com';

INSERT user(username) SELECT a FROM test;

UPDATE user SET age=29 WHERE id=1;

UPDATE user SET username='princess', age = 1, email='princess@royal.com' WHERE id=3;
UPDATE user SET age = age + 10;
UPDATE user SET age = age - 2, email=DEFAULT WHERE id <= 5;
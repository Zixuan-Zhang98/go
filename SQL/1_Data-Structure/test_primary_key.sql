-- 测试主键
CREATE TABLE test_primary_key(
    id INT UNSIGNED PRIMARY KEY,
    username VARCHAR(20)
);

INSERT test_primary_key(id, username) VALUES(1, 'Zixuan');
INSERT test_primary_key(id, username) VALUES('K'); -- ERROR
INSERT test_primary_key(id, username) VALUES(1, 'K'); -- ERROR, 主键唯一

CREATE TABLE test_primary_key1(
    id INT UNSIGNED KEY, -- PRIMARY可省
    username VARCHAR(20)
);

CREATE TABLE test_primary_key2(
    id INT UNSIGNED,
    username VARCHAR(20),
    PRIMARY KEY(id) -- 在最后标明主键
);

CREATE TABLE test_primary_key3(
    id INT UNSIGNED,
    courseID INT UNSIGNED,
    username VARCHAR(20),
    email VARCHAR(50),
    PRIMARY KEY(id, courseID) -- 复合主键: id-courseID
);

-- 测试AUTO_INCREMENT

CREATE TABLE test_primary_key4(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(20)
);

INSERT test_primary_key4(username) VALUES('A'); -- 1
INSERT test_primary_key4(username) VALUES('B'); -- 2
INSERT test_primary_key4(username) VALUES('C'); -- 3
INSERT test_primary_key4(id, username) VALUES(NULL, 'E'); -- 4
INSERT test_primary_key4(id, username) VALUES(DEFAULT, 'F'); -- 5
INSERT test_primary_key4(id, username) VALUES('', 'G'); -- ERROR
INSERT test_primary_key4(id, username) VALUES(15, 'G'); -- ERROR
INSERT test_primary_key4(username) VALUES('G'); -- 16, auto_increment = 已有编号最大值 + 1

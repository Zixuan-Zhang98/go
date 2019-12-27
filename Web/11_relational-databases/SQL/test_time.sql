-- HH:MM:SS [D HH:MM:SS] D表示天数 0～34
-- 测试TIME类型
CREATE TABLE test_time(
  a TIME
);

INSERT test_time(a) VALUES('12:23:45');
INSERT test_time(a) VALUES('48:23:45');
INSERT test_time(a) VALUES('2 48:23:45');
INSERT test_time(a) VALUES('2'); -- 等同于00:00:02
INSERT test_time(a) VALUES('22:22'); --等同于22:22:00
INSERT test_time(a) VALUES('2 13'); --等同于 2 13:00:00
INSERT test_time(a) VALUES('121212'); -- 12:12:12 省略冒号
INSERT test_time(a) VALUES('0'); -- 00:00:00
INSERT test_time(a) VALUES(0); -- 00:00:00
INSERT test_time(a) VALUES('9999:99:99'); -- ERROR
INSERT test_time(a) VALUES(NOW());
INSERT test_time(a) VALUES(CURRENT_TIME);


-- 测试DATE类型 YYYY-MM-DD YYYYMMDD
CREATE TABLE test_date(
  a DATE
);

INSERT test_date(a) VALUES('2019-12-25');
INSERT test_date(a) VALUES('2019-2-14');
INSERT test_date(a) VALUES('19980313');
INSERT test_date(a) VALUES('1998313'); -- ERROR

INSERT test_date(a) VALUES('1998%03@13');

-- YY-MM-DD YYMMDD
-- 70~99 -> 1970~1999; 00~69 -> 2000~2069
INSERT test_date(a) VALUES('45%03@13'); -- 2045-03-13
INSERT test_date(a) VALUES(CURRENT_DATE);
INSERT test_date(a) VALUES(NOW());

-- 测试DATETIME
CREATE TABLE test_datetime(
  a DATETIME
);

INSERT test_datetime(a) VALUES('1005-09-12 13:24:59');
INSERT test_datetime(a) VALUES('720305121212'); -- 1972-03-05 12:12:12
INSERT test_datetime(a) VALUES(NOW());

-- 测试TIMESTAMP
CREATE TABLE test_timestamp(
  a TIMESTAMP
);

INSERT test_timestamp(a) VALUES('1978-10-23 12:12:12');

-- 插入CURRENT_TIMESTAMP, NULL, 或者什么也不写 都是得到当前系统日期和时间(错误)
INSERT test_timestamp(a) VALUES(CURRENT_TIMESTAMP);
INSERT test_timestamp(a) VALUES(NULL); -- NULL
INSERT test_timestamp(a) VALUES(); -- ERROR

-- 测试YEAR
CREATE

-- test INT
CREATE TABLE test_int(
  a TINYINT,
  b SMALLINT,
  c MEDIUMINT,
  d INT,
  e BIGINT
);

INSERT test_int(a) VALUES(-128);
INSERT test_int(a) VALUES(-129); -- out of bound

INSERT test_int(a) VALUES(127);
INSERT test_int(a) VALUES(128); -- out of bound

-- 在非严格模式下，如果超过范围，会产生截断的效果： 200 -> 127, -1000 -> -128

-- test UNSIGNED
CREATE TABLE test_unsigned(
  a TINYINT,
  b TINYINT UNSIGNED
);

INSERT test_unsigned(a, b) VALUES(-13, -12); -- out of bound
INSERT test_unsigned(a, b) VALUES(-13, 255);
INSERT test_unsigned(a, b) VALUES(-13, 256); -- out of bound

-- ZEROFILL
CREATE TABLE test_int1(
  a TINYINT ZEROFILL,
  b SMALLINT ZEROFILL,
  c MEDIUMINT ZEROFILL,
  d INT ZEROFILL,
  e BIGINT ZEROFILL
);

INSERT test_int1(a, b, c, d, e) VALUES(1,1,1,1,1);

-- 数据的实际长度超过显示长度是可以的
CREATE TABLE test_int2(
  a TINYINT(2),
  b INT(2)
);

INSERT test_int2(a, b) VALUES(111, 22222);

CREATE TABLE test_float(
  a FLOAT(6,2),
  b DOUBLE(6,2),
  c DECIMAL(6,3)
);

INSERT test_float(a, b, c) VALUES(4.143, 4.146, 4.149);

CREATE TABLE test_float1(
  a FLOAT,
  b DOUBLE,
  c DECIMAL
);

INSERT test_float(a, b, c) VALUES(4.143, 4.146, 4.149);

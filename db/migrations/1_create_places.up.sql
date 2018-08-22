
CREATE TABLE IF NOT EXISTS places (
  id VARCHAR(100) PRIMARY KEY UNIQUE,
  name TEXT,
  zipcode TEXT,
  country TEXT
);

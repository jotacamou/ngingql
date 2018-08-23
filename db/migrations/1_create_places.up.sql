
CREATE TABLE IF NOT EXISTS places (
  id SERIAL PRIMARY KEY,
  name character varying(256) NOT NULL,
  country character varying(256),
  zipcode character varying(32)
);

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name character varying(255) NOT NULL,
  email character varying(255) NOT NULL,
  password character varying(128) NOT NULL
);

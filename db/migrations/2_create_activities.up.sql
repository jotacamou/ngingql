
CREATE TABLE IF NOT EXISTS activities (
  id SERIAL PRIMARY KEY,
  title text NOT NULL,
  description text,
  owner character varying(256),
  price numeric(9,2),
  fee numeric(9,2),
  main_picture text[],
  slideshow_picture text[],
  duration character varying(32),
  place_id integer,
  hour text[],
  tag text[]
);

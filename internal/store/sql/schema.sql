--@block
DROP TABLE IF EXISTS users;
--@block
CREATE TABLE IF NOT EXISTS users(
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  email  text,
  password text,
  created timestamp DEFAULT NOW()
);


--@block
DROP TABLE IF EXISTS appointments;
--@block
CREATE TABLE IF NOT EXISTS appointments(
	id        BIGSERIAL PRIMARY KEY, 
	client_id   bigint,
	appt_time   timestamp,
	status     text,
	note      text,
  created timestamp  DEFAULT NOW()

);


--@block
DROP TABLE IF EXISTS clients;
--@block
CREATE TABLE IF NOT EXISTS clients(
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  created timestamp DEFAULT NOW()

);

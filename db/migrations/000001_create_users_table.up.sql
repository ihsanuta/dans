CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   username VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL
);

INSERT INTO users(`username`,`password`) VALUES ('admin','admin');
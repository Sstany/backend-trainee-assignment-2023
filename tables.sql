CREATE TABLE users (user_id serial PRIMARY KEY,username text NOT NULL);

CREATE TABLE segments (
  segment_id  serial PRIMARY KEY,
  segment     text NOT NULL
);

CREATE TABLE segment_user (
  segment_id    int REFERENCES segments (segment_id) ON UPDATE CASCADE ON DELETE CASCADE,
   user_id int REFERENCES users (user_id) ON UPDATE CASCADE,
   CONSTRAINT segment_user_pkey PRIMARY KEY (segment_id, user_id)  -- explicit pk
);

CREATE TABLE history (
		historyId serial PRIMARY KEY,
		userId int NOT NULL,
		time TIMESTAMP NOT NULL,
		type text NOT NULL,
		segment text NOT NULL)
CREATE DATABASE belajar_go_nbcamp;

DROP DATABASE belajar_go_nbcamp;

CREATE TABLE roles (
  id SERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(20)
);

select * from roles;

create table auths (
  id varchar(50) not null primary key ,
  role_id int,
  email varchar(50),
  password varchar(255),
  created_at timestamp,
  updated_at timestamp
);

select * from auths;

alter table "auths" add foreign key ("role_id") references roles("id");
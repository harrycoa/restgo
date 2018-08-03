create database users(
  id int(6) unsiged auto_increment primary key,
  username varchar(100) not null,
  password varchar(256) not null,
  email varchar(200) null,
  create_date timestamp default current_timestamp
);
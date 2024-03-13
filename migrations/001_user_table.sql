create table if not exists users (
	id uuid primary key,
	name varchar(255),
	email varchar(255) unique,
	role integer
);

---- create above / drop below ----

drop table if exists users;
begin;
create table if not exists clients(
 id serial primary key ,
first_name varchar (255) not null,
last_name varchar (255) not null,
email varchar(255) not  null,
phone_number varchar(13)not null,
password varchar (255) not null,
is_email_verified boolean not null default false,
email_verified_at timestamp,
created_at timestamp not null default now(),
updated_at timestamp ,
updated_by int  ,
deleted_at timestamp,
deleted_by int ,
constraint fk_users_updated_by
	foreign key (updated_by)
	references  clients(id)
	on delete  set null
	on update cascade ,

constraint clients_email_unique unique ("email")
);

create table app_setting(
id serial primary key,
app_status bool default false,
);

commit;

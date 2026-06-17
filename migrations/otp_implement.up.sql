begin;
create table otp_verifications(
id serial  primary key,
client_id int not null,
otp_code varchar(6) not null,
create_at timestamp not null default now(),
expire_at timestamp not null,
type varchar(20) not null,
used boolean not null default false);


commit;

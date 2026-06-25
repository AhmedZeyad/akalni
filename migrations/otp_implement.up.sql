begin;
create table otp_verifications(
id serial  primary key,
client_id int not null,
otp_code varchar(255) not null,
created_at timestamp not null default now(),
expires_at timestamp not null,
type varchar(20) not null,
used boolean not null default false);


commit;

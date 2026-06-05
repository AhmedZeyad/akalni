begin;
create table users (
id serial  primary key ,
name varchar (100) not null default  '',
email varchar (255) not null default '',
password varchar (255) not null ,
status boolean not null default false,
created_at timestamp default now() not null,
created_by int not null ,
updated_at timestamp default null,
updated_by int default null,
constraint users_email_unique unique (email)
);
create table restaurants (
id serial  primary key ,
name varchar (100) not null default  '',
status boolean not null default false,
lon float  not null,
lat float not null,
address varchar (100),
created_at timestamp default now() not null,
created_by int not null ,
updated_at timestamp default null,
updated_by int default null,
deleted_at timestamp default null,
deleted_by int default null
);
create table categories (
id serial primary key ,
name varchar (100) not null default '',
active boolean not null default false,
created_at timestamp default now() not null,
created_by int not null ,
updated_at timestamp default null,
updated_by int default null,
deleted_at timestamp default null,
deleted_by int default null
);


create table products   (
id serial primary key ,
name varchar(100) not null default '',
active boolean default false ,
rest_id int ,
price numeric (10,2),
category_id int ,
created_at timestamp default now() not null,
created_by int not null ,
updated_at timestamp default null,
updated_by int default null,
deleted_at timestamp default null,
deleted_by int default null,
constraint  fk_product_restaurant_id
foreign key (rest_id) references  restaurants(id)
on delete cascade ,
constraint fk_product_category_id
foreign key (category_id) references categories(id)
on delete set null
);

create table orders(
id serial primary key ,
client_id int,
rest_id int,
total_price numeric ,
sub_total_price numeric ,
last_status varchar(100),
created_at timestamp default now() not null,
created_by int not null ,
updated_at timestamp default null,
updated_by int default null,
deleted_at timestamp default null,
deleted_by int default null,
constraint fk_orders_client_id
foreign key  (client_id) references clients(id)
on delete set null,
constraint fk_orders_restaurant_id
foreign key (rest_id) references restaurants (id)
);
create table order_details(
id serial primary key,
order_id int,
product_id int ,
quantity int ,
price numeric (10,2) not null,
constraint fk_order_details_order_id
foreign key (order_id)  references orders(id)
on delete set null
);
commit;

create table if not exists users
(
        user_id    varchar(50) primary key,
        name       varchar(200),
        username   varchar(100) not null,
        password   varchar(200) not null,
        user_type  varchar(100),
        created_at datetime    default CURRENT_TIMESTAMP,
        UNIQUE (username)
) engine = innoDB;


create table if not exists customers (
      customer_id varchar(50) primary key,
      name varchar(200),
      contact varchar(100) unique,
      created_at datetime default CURRENT_TIMESTAMP
);


create table if not exists categories (
      category_id varchar(50) primary key,
      name varchar(200)
);

create table if not exists products (
     product_id varchar(50) primary key,
     name varchar(200),
     price int,
     category_id varchar(50),
     stock int,
     product_condition varchar(50)
);

create table if not exists coupons (
    id int primary key auto_increment,
    coupon_code varchar(50),
    is_used bool default false,
    expire_date datetime,
    customer_id varchar(50)
);

create table if not exists coupons_prefix (
    id int primary key auto_increment,
    prefix_name varchar(200) unique,
    minimum_price bigint,
    discount int,
    expire_date datetime,
    criteria_id int
);

create table if not exists criteria (
    id int primary key auto_increment,
    name varchar(200) unique
);


alter table products add constraint fk_product_category foreign key products(category_id)
    references categories(category_id) on delete set null;


alter table coupons add constraint fk_coupon_customer foreign key coupons(customer_id)
    references customers(customer_id) on delete cascade;

alter table coupons_prefix add constraint fk_coupon_prefix_criteria foreign key coupons_prefix(criteria_id)
    references criteria(id) on delete set null;


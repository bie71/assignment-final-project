create table if not exists users
(
        user_id    varchar(50) primary key,
        name       varchar(200),
        username   varchar(100) unique not null,
        password   varchar(200) not null,
        user_type  varchar(100),
        created_at datetime default CURRENT_TIMESTAMP
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
     price int unsigned,
     category_id varchar(50),
     stock int unsigned default 0
);

create table if not exists coupons (
    id int primary key auto_increment,
    coupon_code varchar(50),
    is_used bool default false,
    expire_date datetime,
    customer_id varchar(50)
);

create table if not exists initial_coupons (
    id int primary key auto_increment,
    prefix_name varchar(200),
    minimum_price bigint,
    discount int unsigned,
    expire_date datetime,
    criteria varchar(200),
    created_at datetime default current_timestamp
);

create table if not exists transaction (
  transaction_id varchar(50) primary key,
  customer_id varchar(50),
  coupon_code varchar(50),
  total_price double unsigned,
  discount float unsigned,
  total_price_after_discount double unsigned,
  purchase_date datetime
);

create table if not exists transaction_items (
    id int primary key auto_increment,
    transaction_id varchar(50),
    product_id varchar(50),
    quantity int unsigned
);


CREATE TRIGGER update_stock_item
    AFTER INSERT
    ON transaction_items
    FOR EACH ROW
    UPDATE products
    SET stock = stock - NEW.quantity
    WHERE product_id = NEW.product_id;


alter table products add constraint fk_product_category foreign key products(category_id)
    references categories(category_id) on delete set null on update cascade ;

alter table initial_coupons add constraint fk_initial_category foreign key initial_coupons(criteria)
    references categories(category_id) on delete set null on update cascade ;

alter table coupons add constraint fk_coupon_customer foreign key coupons(customer_id)
    references customers(customer_id) on delete cascade on update cascade ;

alter table transaction add constraint fk_transaction_customer foreign key transaction(customer_id)
    references customers(customer_id) on delete set null on update cascade ;

alter table transaction_items add constraint fk_items_transaction foreign key transaction_items(transaction_id)
references transaction(transaction_id) on delete cascade on update cascade ;

alter table transaction_items add constraint fk_items_product foreign key transaction_items(product_id)
    references products(product_id) on delete set null on update cascade ;

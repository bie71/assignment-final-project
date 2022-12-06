create table if not exists products (
 product_id varchar(50) primary key,
 name varchar(200),
 price int,
 category_id varchar(50),
 stock int,
 product_condition varchar(50)
)
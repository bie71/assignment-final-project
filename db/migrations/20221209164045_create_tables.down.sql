alter table products drop foreign key fk_product_category;

alter table coupons drop foreign key fk_coupon_customer;

drop table if exists users;

drop table if exists customers;

drop table if exists categories;

drop table if exists products;

drop table if exists coupons;

drop table if exists coupons_prefix;

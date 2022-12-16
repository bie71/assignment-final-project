alter table products drop foreign key fk_product_category;

alter table coupons drop foreign key fk_coupon_customer;

alter table transaction drop foreign key fk_transaction_customer;

alter table transaction_items drop foreign key fk_items_transaction;

alter table transaction_items drop foreign key fk_items_product;

DROP TRIGGER IF EXISTS update_stock_item;

drop table if exists users;

drop table if exists customers;

drop table if exists categories;

drop table if exists products;

drop table if exists coupons;

drop table if exists initial_coupons;

drop table if exists transaction;

drop table if exists transaction_items;
create table if not exists customers (
    customer_id varchar(50) primary key,
    name varchar(200),
    contact varchar(100) unique,
    created_at datetime default CURRENT_TIMESTAMP
)
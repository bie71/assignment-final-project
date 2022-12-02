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
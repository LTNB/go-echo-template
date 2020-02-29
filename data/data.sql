create table account
(
    id        varchar(255) not null
        constraint account_pk
            primary key,
    full_name varchar(255),
    email     varchar(255),
    active    boolean,
    role      varchar(255),
    password varchar (255)
);

create unique index account_id_uindex
    on account (id);

INSERT INTO account (id, full_name, email, active, role, password) values ('1', 'Tạ Ngọc Bảo Lâm', 'baolam0307@gmail.com', true , 'ADMIN', '$2a$12$kwYCP9mj2nzSAJc5mRFCM.bvviMl8yHivYAGRS9SpG91DQ8OTgqla');
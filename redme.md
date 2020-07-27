##数据库采用postgres

使用四张表
```
create table users
(
    id                 serial PRIMARY KEY,
    login_name          character varying(64),
    pwd                  text unique

);

create table video_info
(
    id                 serial PRIMARY KEY,
    author_id int,
    name text,
    display_name text,
    create_time timestamp
);

create table comments
(
     id          character varying(64) PRIMARY KEY not null ,
    video_id          character varying(64),
    author_id int,
    content text,
    time timestamp
);

create table sessions
(
     session_id         text PRIMARY KEY not null ,
    video_id          character varying(64),
    login_name character varying(64)
);
```
## 数据库相关api在dbops中

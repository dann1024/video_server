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
    id                 character varying(64) PRIMARY KEY,
    author_id int,
    name text,
    display_name text,
    display_ctime character varying(64),
    create_time character varying(64)
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
    login_name character varying(64),
    ttl      character varying(64),
);
```
## 数据库相关api在dbops中

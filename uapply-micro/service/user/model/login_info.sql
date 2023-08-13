-- auto-generated definition
create table login_info
(
    uid        int auto_increment,
    open_id    varchar(255) null,
    xd_stu_num varchar(63)  null comment '学号，统一认证登录使用',

    primary key (uid),
    constraint open_id
        unique (open_id),
    constraint xd_stu_num
        unique (xd_stu_num)
);

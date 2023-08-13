-- auto-generated definition
create table inter_conf
(
    id            int auto_increment,
    dep_id        int          not null comment '部门 id',
    turn          int          not null comment '轮次',
    start         bigint       not null comment '开始面试时间戳',
    end           bigint       not null comment '面试结束时间戳',
    inter_address varchar(255) not null comment '面试地点',
    contact_phone varchar(32)  not null comment '负责人电话',

    primary key (id),
    constraint dep_turn_uniq_key
        unique (dep_id, turn)
) comment '面试配置表';

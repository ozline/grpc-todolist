create table todolist.user
(
    id              bigint                              not null comment 'user_id',
    username        text                                not null comment 'username',
    password        text                                not null comment 'password',
    created_at      timestamp default current_timestamp not null comment 'create time',
    updated_at      timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    deleted_at      timestamp default null              null comment 'user delete time',
    constraint id
        primary key (id)
);

create table todolist.task
(
    id              bigint                              not null comment 'task_id',
    user_id         bigint                              not null comment 'userid',
    status          int                                 not null comment 'status',
    title           text                                not null comment 'title',
    content         text                                not null comment 'content',
    created_at      timestamp default current_timestamp not null comment 'create time',
    updated_at      timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    deleted_at      timestamp default null              null comment 'user delete time',
    constraint id
        primary key (id)
);
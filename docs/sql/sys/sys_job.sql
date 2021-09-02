create table sys_job
(
    id               bigint auto_increment comment '编号'
        primary key,
    job_name         varchar(50)                         null comment '职位名称',
    order_num        int                                 null comment '排序',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP  comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime     DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    del_flag         tinyint   default 0                 null comment '是否删除  -1：已删除  0：正常',
    remarks          varchar(68)                         null comment '备注'
)
    comment '职位管理';

INSERT INTO sys_job (id, job_name, order_num, create_by, create_time, last_update_by, last_update_time, del_flag, remarks) VALUES (1, '董事长', 1, 'admin', '2021-04-26 15:50:45', 'admin', '2021-04-26 16:17:07', 0, '测试1');
INSERT INTO sys_job (id, job_name, order_num, create_by, create_time, last_update_by, last_update_time, del_flag, remarks) VALUES (2, '经理', 2, 'admin', '2021-04-26 16:05:11', 'admin', '2021-04-26 16:16:36', 0, '管理人员1');
INSERT INTO sys_job (id, job_name, order_num, create_by, create_time, last_update_by, last_update_time, del_flag, remarks) VALUES (4, 'test', 3, 'admin', '2021-04-27 09:41:48', 'admin', '2021-04-27 09:43:06', 0, '测试1');
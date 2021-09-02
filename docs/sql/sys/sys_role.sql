create table sys_role
(
    id               bigint auto_increment comment '编号'
        primary key,
    name             varchar(100)                        null comment '角色名称',
    remark           varchar(100)                        null comment '备注',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP  comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime  DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    del_flag         tinyint   default 0                 null comment '是否删除  -1：已删除  0：正常',
    status           int       default 1                 null comment '状态  1:启用,0:禁用'
)
    comment '角色管理';

INSERT INTO sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag, status) VALUES (1, 'admin', '超级管理员', 'admin', '2019-01-19 11:11:11', 'admin', '2021-08-10 00:02:51', 0, 1);
INSERT INTO sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag, status) VALUES (2, 'mng', '项目经理', 'admin', '2019-01-19 11:11:11', 'admin', '2021-08-18 15:35:14', 0, 1);
INSERT INTO sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag, status) VALUES (3, 'dev', '开发人员', 'admin', '2019-01-19 11:11:11', 'admin', '2019-01-19 11:39:28', 0, 1);
INSERT INTO sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag, status) VALUES (4, 'test', '测试人员1', 'admin', '2019-01-19 11:11:11', 'admin', '2021-08-17 15:52:18', 0, 1);
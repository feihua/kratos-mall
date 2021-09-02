create table sys_dept
(
    id               bigint auto_increment comment '编号'
        primary key,
    name             varchar(50)                         null comment '机构名称',
    parent_id        bigint                              null comment '上级机构ID，一级机构为0',
    order_num        int                                 null comment '排序',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP  comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime  DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    del_flag         tinyint   default 0                 null comment '是否删除  -1：已删除  0：正常'
)
    comment '机构管理';

INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (1, '轻尘集团', 0, 0, 'admin', '2018-09-23 19:35:22', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (2, '牧尘集团', 0, 1, 'admin', '2018-09-23 19:35:55', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (3, '三国集团', 0, 2, 'admin', '2018-09-23 19:36:24', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (4, '上海分公司', 2, 0, 'admin', '2018-09-23 19:37:03', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (5, '北京分公司', 1, 1, 'admin', '2018-09-23 19:37:17', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (6, '北京分公司', 2, 1, 'admin', '2018-09-23 19:37:28', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (7, '技术部', 5, 0, 'admin', '2018-09-23 19:38:00', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (8, '技术部', 4, 0, 'admin', '2018-09-23 19:38:10', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (9, '技术部', 6, 0, 'admin', '2018-09-23 19:38:17', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (10, '市场部', 5, 0, 'admin', '2018-09-23 19:38:45', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (11, '市场部', 6, 0, 'admin', '2018-09-23 19:39:01', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (12, '魏国', 3, 0, 'admin', '2018-09-23 19:40:42', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (13, '蜀国', 3, 1, 'admin', '2018-09-23 19:40:54', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (14, '吴国', 3, 2, 'admin', '2018-09-23 19:41:04', 'admin', '2018-09-23 19:35:22', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (15, '1', 0, 121, 'admin', '2021-04-27 10:17:19', 'admin', '2021-04-27 10:21:17', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (16, '12', 0, 121, '', '2021-04-27 10:18:11', 'admin', '2021-04-27 10:19:51', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (17, '1', 16, 112, 'admin', '2021-04-27 10:19:37', 'admin', '2021-04-27 10:21:05', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (18, '11', 15, 1, 'admin', '2021-04-27 10:21:29', 'admin', '2021-04-27 10:21:30', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (19, '系统组', 0, 2, 'admin', '2021-06-07 12:23:34', 'admin', '2021-06-07 12:23:34', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (20, 'rd', 19, 1, 'admin', '2021-06-07 12:23:51', 'admin', '2021-06-07 12:23:51', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (21, '1', 1, 1, 'admin', '2021-06-24 11:52:37', 'admin', '2021-06-24 11:52:37', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (22, '444', 0, 444, 'admin', '2021-06-24 11:52:47', 'admin', '2021-06-24 11:52:48', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (23, '1', 10, 1, 'admin', '2021-07-15 17:40:55', 'admin', '2021-07-15 17:40:55', 0);
INSERT INTO sys_dept (id, name, parent_id, order_num, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (24, '111aaa', 0, 2, 'admin', '2021-08-05 20:45:40', 'admin', '2021-08-05 20:45:41', 0);
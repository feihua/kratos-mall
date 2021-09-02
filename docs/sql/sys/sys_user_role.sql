create table sys_user_role
(
    id               bigint auto_increment comment '编号'
        primary key,
    user_id          bigint      null comment '用户ID',
    role_id          bigint      null comment '角色ID',
    create_by        varchar(50) null comment '创建人',
    create_time      datetime    default CURRENT_TIMESTAMP comment '创建时间',
    last_update_by   varchar(50) null comment '更新人',
    last_update_time datetime    DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP comment '更新时间'
)
    comment '用户角色';

INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (2, 2, 1, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (26, 5, 3, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (33, 6, 2, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (34, 4, 2, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (35, 9, 2, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (36, 10, 3, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (37, 11, 2, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (38, 12, 3, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (39, 15, 2, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (41, 16, 3, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (42, 8, 2, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (43, 7, 4, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (45, 18, 2, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (46, 17, 3, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (47, 3, 4, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (48, 21, 2, null, null, null, null);
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (94, 38, 6, 'admin', null, 'admin', '2021-04-27 11:24:14');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (97, 40, 3, 'admin', null, 'admin', '2021-04-27 13:47:43');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (98, 39, 6, 'admin', null, 'admin', '2021-04-27 13:53:40');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (99, 41, 3, 'admin', null, 'admin', '2021-05-02 23:56:19');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (101, 42, 1, 'admin', null, 'admin', '2021-05-13 15:14:16');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (103, 43, 8, 'admin', null, 'admin', '2021-06-07 12:25:56');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (111, 1, 1, 'admin', null, 'admin', '2021-07-25 21:55:46');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (112, 31, 2, 'admin', null, 'admin', '2021-07-25 21:56:00');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (114, 30, 2, 'admin', null, 'admin', '2021-07-25 21:56:52');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (115, 23, 3, 'admin', null, 'admin', '2021-08-05 09:53:56');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (117, 32, 3, 'admin', null, 'admin', '2021-08-06 17:34:49');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (118, 26, 3, 'admin', null, 'admin', '2021-08-07 17:47:31');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (121, 33, 1, 'admin', null, 'admin', '2021-08-07 17:48:21');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (123, 48, 2, 'admin', null, 'admin', '2021-08-07 18:26:47');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (124, 25, 1, 'admin', null, 'admin', '2021-08-09 10:49:38');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (125, 22, 1, 'admin', null, 'admin', '2021-08-09 10:50:11');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (126, 24, 4, 'admin', null, 'admin', '2021-08-10 00:01:54');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (128, 49, 2, 'admin', null, 'admin', '2021-08-11 23:48:29');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (129, 29, 1, 'admin', null, 'admin', '2021-08-12 00:55:33');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (130, 27, 4, 'admin', null, 'admin', '2021-08-12 00:55:57');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (131, 50, 1, 'admin', null, 'admin', '2021-08-12 01:05:31');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (132, 28, 2, 'admin', null, 'admin', '2021-08-12 02:46:30');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (133, 44, 2, 'admin', null, 'admin', '2021-08-12 09:15:04');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (134, 51, 2, 'admin', null, 'admin', '2021-08-13 23:35:28');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (135, 52, 2, 'admin', null, 'admin', '2021-08-13 23:36:09');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (136, 35, 4, 'admin', null, 'admin', '2021-08-16 00:10:30');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (137, 53, 3, 'admin', null, 'admin', '2021-08-18 15:33:41');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (138, 36, 3, 'admin', null, 'admin', '2021-08-18 16:35:52');
INSERT INTO sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (139, 37, 4, 'admin', null, 'admin', '2021-08-20 18:28:16');
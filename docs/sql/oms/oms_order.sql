create table oms_order
(
    id                      bigint auto_increment comment '订单id'
        primary key,
    member_id               bigint           not null,
    coupon_id               bigint           null,
    order_sn                varchar(64)      null comment '订单编号',
    create_time             datetime         null comment '提交时间',
    member_username         varchar(64)      null comment '用户帐号',
    total_amount            decimal(10, 2)   null comment '订单总金额',
    pay_amount              decimal(10, 2)   null comment '应付金额（实际支付金额）',
    freight_amount          decimal(10, 2)   null comment '运费金额',
    promotion_amount        decimal(10, 2)   null comment '促销优化金额（促销价、满减、阶梯价）',
    integration_amount      decimal(10, 2)   null comment '积分抵扣金额',
    coupon_amount           decimal(10, 2)   null comment '优惠券抵扣金额',
    discount_amount         decimal(10, 2)   null comment '管理员后台调整订单使用的折扣金额',
    pay_type                int(1)           null comment '支付方式：0->未支付；1->支付宝；2->微信',
    source_type             int(1)           null comment '订单来源：0->PC订单；1->app订单',
    status                  int(1)           null comment '订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单',
    order_type              int(1)           null comment '订单类型：0->正常订单；1->秒杀订单',
    delivery_company        varchar(64)      null comment '物流公司(配送方式)',
    delivery_sn             varchar(64)      null comment '物流单号',
    auto_confirm_day        int              null comment '自动确认时间（天）',
    integration             int              null comment '可以获得的积分',
    growth                  int              null comment '可以活动的成长值',
    promotion_info          varchar(100)     null comment '活动信息',
    bill_type               int(1)           null comment '发票类型：0->不开发票；1->电子发票；2->纸质发票',
    bill_header             varchar(200)     null comment '发票抬头',
    bill_content            varchar(200)     null comment '发票内容',
    bill_receiver_phone     varchar(32)      null comment '收票人电话',
    bill_receiver_email     varchar(64)      null comment '收票人邮箱',
    receiver_name           varchar(100)     not null comment '收货人姓名',
    receiver_phone          varchar(32)      not null comment '收货人电话',
    receiver_post_code      varchar(32)      null comment '收货人邮编',
    receiver_province       varchar(32)      null comment '省份/直辖市',
    receiver_city           varchar(32)      null comment '城市',
    receiver_region         varchar(32)      null comment '区',
    receiver_detail_address varchar(200)     null comment '详细地址',
    note                    varchar(500)     null comment '订单备注',
    confirm_status          int(1)           null comment '确认收货状态：0->未确认；1->已确认',
    delete_status           int(1) default 0 not null comment '删除状态：0->未删除；1->已删除',
    use_integration         int              null comment '下单时使用的积分',
    payment_time            datetime         null comment '支付时间',
    delivery_time           datetime         null comment '发货时间',
    receive_time            datetime         null comment '确认收货时间',
    comment_time            datetime         null comment '评价时间',
    modify_time             datetime         null comment '修改时间'
)
    comment '订单表';

INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (12, 1, 2, '201809150101000001', '2018-09-15 12:24:27', 'test', 18732.00, 16377.75, 20.00, 2344.25, 0.00, 10.00, 10.00, 0, 1, 4, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '江苏省', '常州市', '天宁区', '东晓街道', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (13, 1, 2, '201809150102000002', '2018-09-15 14:24:29', 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 1, 1, 1, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (14, 1, 2, '201809130101000001', '2018-09-13 16:57:40', 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 2, 1, 2, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (15, 1, 2, '201809130102000002', '2018-09-13 17:03:00', 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 1, 1, 3, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 1, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (16, 1, 2, '201809140101000001', '2018-09-14 16:16:16', 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 2, 1, 4, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (17, 1, 2, '201809150101000003', '2018-09-15 12:24:27', 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 0, 1, 4, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (18, 1, 2, '201809150102000004', '2018-09-15 14:24:29', 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 1, 1, 1, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (19, 1, 2, '201809130101000003', '2018-09-13 16:57:40', 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 2, 1, 2, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (20, 1, 2, '201809130102000004', '2018-09-13 17:03:00', 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 1, 1, 3, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (21, 1, 2, '201809140101000002', '2018-09-14 16:16:16', 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 2, 1, 4, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 1, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (22, 1, 2, '201809150101000005', '2018-09-15 12:24:27', 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 0, 1, 4, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (23, 1, 2, '201809150102000006', '2018-09-15 14:24:29', 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 1, 1, 1, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (24, 1, 2, '201809130101000005', '2018-09-13 16:57:40', 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 2, 1, 2, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (25, 1, 2, '201809130102000006', '2018-09-13 17:03:00', 'test', 18732.00, 16377.75, 10.00, 2344.25, 0.00, 10.00, 5.00, 1, 1, 4, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨22', '18033441849', '518000', '北京市', '北京城区', '东城区', '东城街道', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (26, 1, 2, '201809140101000003', '2018-09-14 16:16:16', 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 2, 1, 4, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '111', 0, 1, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (27, 1, 2, '202002250100000001', '2020-02-25 15:59:20', 'test', 540.00, 540.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0, 1, 0, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '无优惠,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '南山区', '科兴科学园', '111', 0, 1, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (28, 1, 2, '202002250100000002', '2020-02-25 16:05:47', 'test', 540.00, 540.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0, 1, 0, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '无优惠,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '南山区', '科兴科学园', '111', 0, 1, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (29, 1, 2, '202002250100000003', '2020-02-25 16:07:58', 'test', 540.00, 540.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0, 1, 0, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '无优惠,无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '南山区', '科兴科学园', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
INSERT INTO oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, payment_time, delivery_time, receive_time, comment_time, modify_time) VALUES (30, 1, 2, '202002250100000004', '2020-02-25 16:50:13', 'test', 240.00, 240.00, 20.00, 0.00, 0.00, 0.00, 10.00, 0, 1, 2, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '无优惠', 1, '1111', '1111', '18613030352', '1002219331@qq.com', '大梨', '18033441849', '518000', '广东省', '深圳市', '南山区', '科兴科学园', '111', 0, 0, 1000, '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57', '2021-03-16 20:58:57');
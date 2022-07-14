create table user_info
(
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id',
    `nickname` VARCHAR(128) NOT NULL comment '用户昵称',
    `avatar` VARCHAR(128) NOT NULL comment '用户头像',
    `real_name` VARCHAR(128) NULL comment '真实姓名',
    `id_card` VARCHAR(128) NULL comment '身份证号',
    `certified_phone` VARCHAR(128) NULL comment '手机号',
    `certified_time` TIMESTAMP NULL comment '认证时间',
    `is_certified` TINYINT(4) NOT NULL comment '是否实名认证',
    `chain_account_id` VARCHAR(50) NULL comment '链上账户id',
    `chain_account_name` VARCHAR(128) NULL comment '链上账户名',
    `invite_code` VARCHAR(45) NULL comment '邀请码',
    `channel_code` VARCHAR(20) NULL comment '渠道邀请码',
    `invite_user_id` BIGINT(20) NULL comment '邀请人id',
    `credits_cnt` INT NOT NULL DEFAULT 0 comment '用户积分',
    `diamond_cnt` INT NOT NULL DEFAULT 0 comment '用户钻石数',
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_certified_phone` (`certified_phone`),
    KEY `idx_invite_code` (`invite_code`),
    KEY `idx_chain_account_id` (`chain_account_id`)
 ) ENGINE = InnoDB AUTO_INCREMENT = 1655350 DEFAULT CHARSET = utf8mb4 comment '用户基本信息表';


-- create table user_addition_info
-- {
--     `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREAMENT comment '主键id',
--     `user_id` BIGINT(20) UNSIGNED NOT NULL comment '用户id',
--     `credits_cnt` INT NOT NULL DEFAULT 0 comment '用户积分',
--     `diamond_cnt` INT NOT NULL DEFAULT 0 comment '用户钻石数',
--     `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
--     `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
-- } ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '用户附属信息表'

create table props_info
(
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id',
    `props_id` BIGINT(20) UNSIGNED NOT NULL comment '道具id',
    `name` VARCHAR(128) NOT NULL comment '道具名称',
    `image_url` VARCHAR(128) NULL comment '道具图片',
    `type` INT NOT NULL DEFAULT 0 comment '道具类型',
    `desc` TEXT NULL comment '道具描述',
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_props_id` (`props_id`)
 ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '道具流水表';

create table user_props
(
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id',
    `user_id` BIGINT(20) UNSIGNED NOT NULL comment '用户id',
    `props_id` BIGINT(20) UNSIGNED NOT NULL comment '道具id',
    `num` INT NOT NULL DEFAULT 0 comment '道具数量',
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`)
 ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '用户道具表';


create table props_record
(
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id',
    `user_id` BIGINT(20) UNSIGNED NOT NULL comment '用户id',
    `props_id` BIGINT(20) UNSIGNED NOT NULL comment '道具id',
    `diff_num` INT NOT NULL DEFAULT 0 comment '道具变更数量',
    `record_type` INT NOT NULL DEFAULT 0 comment '道具记录类型: 发放、消耗、转移',
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_props_record_user_id` (`user_id`)
 ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '道具流水表';


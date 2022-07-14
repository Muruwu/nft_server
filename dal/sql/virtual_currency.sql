create table virtual_currency_record
(
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id',
    `user_id` BIGINT(20) UNSIGNED NOT NULL comment '用户id',
    `virtual_currency_cnt` INT NOT NULL DEFAULT 0 comment '货币变更数量',
    `source` INT NOT NULL DEFAULT 0 comment '货币来源',
    `record_type` INT NOT NULL DEFAULT 0 comment '货币记录类型: 发放、消耗、转移',
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
     PRIMARY KEY (`id`),
     KEY `idx_user_id` (`user_id`)
 ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '货币流水表';


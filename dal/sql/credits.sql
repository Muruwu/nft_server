create table credits_record
{
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREAMENT comment '主键id',
    `user_id` BIGINT(20) UNSIGNED NOT NULL comment '用户id',
    `credits_cnt` INT NOT NULL DEFAULT 0 comment '用户积分',
    `source` INT NOT NULL DEFAULT 0 comment '积分来源',
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
} ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '积分流水表'


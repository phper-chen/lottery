CREATE DATABASE IF NOT EXISTS test_db;
USE test_db;

CREATE TABLE IF NOT EXISTS `lottery_prizes` (
 `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
 `prize` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖品类型 0-贴纸 1-电话卡 2-手机',
 `total` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '奖品总数 仅做记录 0代表无限 用prize来限制',
 `stock` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '奖品剩余数量 0代表无限用 prize来限制',
 `version` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '版本号 乐观锁 暂时用不到',
 `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
 `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
 PRIMARY KEY (`id`),
 UNIQUE KEY `uk_prize` (`prize`),
 KEY `ix_mtime` (`mtime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='奖品列表';

-- 初始化奖品数据
INSERT IGNORE INTO `lottery_prizes` (`prize`, `total`, `stock`, `version`)
VALUES (0, 0, 0, 0), (1, 100, 100, 0), (2, 5, 5, 0);

CREATE TABLE IF NOT EXISTS `lottery_records` (
 `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
 `phone` varchar(11) NOT NULL DEFAULT '0' COMMENT '手机号',
 `prize_id` int(11) NOT NULL DEFAULT '0' COMMENT 'prize',
 `draw_date` date NOT NULL DEFAULT '0000-00-00' COMMENT '抽奖时间 天 用ctime容易导致范围查询',
 `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
 `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
 PRIMARY KEY (`id`),
 KEY `phone_prize_id_draw_date` (`phone`, `prize_id`, `draw_date`),
 KEY `ix_mtime` (`mtime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='抽奖记录';


CREATE TABLE IF NOT EXISTS `lottery_users` (
 `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
 `phone` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '手机号',
 `draw_right` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否可以抽奖 0否 1是',
 `article` varchar(500) NOT NULL DEFAULT '0' COMMENT '文章',
 `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
 `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
 PRIMARY KEY (`id`),
 UNIQUE KEY `uk_phone` (`phone`),
 KEY `ix_mtime` (`mtime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户参与信息';

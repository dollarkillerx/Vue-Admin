# 用户
CREATE TABLE `d_user` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL DEFAULT "",
    `email` VARCHAR(355) NOT NULL DEFAULT "",
    `password` CHAR(32) NOT NULL DEFAULT "",
    `avatar` VARCHAR(355) NOT NULL DEFAULT "" COMMENT "头像地址",
    `is_admin` TINYINT NOT NULL DEFAULT 0 COMMENT "0普通用户1管理员",
    `create_time` INT UNSIGNED NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    KEY `is_admin`(`is_admin`)
)ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8;

# 导航
CREATE TABLE `d_nav` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `language` VARCHAR(255) NOT NULL DEFAULT "" COMMENT "语言编号缩写",
    `name` VARCHAR(255) NOT NULL DEFAULT "" COMMENT "导航名称",
    `url` VARCHAR(366) NOT NULL DEFAULT "" COMMENT "导航地址",
    `weight` INT NOT NULL DEFAULT 0 COMMENT "权重",
    PRIMARY KEY (`id`),
    KEY `language`(`language`)
)ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8;

# 分类
CREATE TABLE `d_sort`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(266) NOT NULL DEFAULT "" COMMENT "",
    `language` VARCHAR(255) NOT NULL DEFAULT "" COMMENT "语言编号缩写",
    `describe` VARCHAR(366) NOT NULL DEFAULT "" COMMENT "分类描述",
    `create_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "创建时间",
    `update_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "更新时间",
    `delete_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "软删除",
    PRIMARY KEY (`id`),
    KEY `delete_time`(`delete_time`),
    KEY `language`(`language`)
)ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8;

# 轮播图
CREATE TABLE `d_carousel`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `img` VARCHAR(266) NOT NULL DEFAULT "" COMMENT "轮播图img",
    `url` VARCHAR(266) NOT NULL DEFAULT "" COMMENT "推荐位 地址",
    `language` VARCHAR(255) NOT NULL DEFAULT "" COMMENT "语言编号缩写",
    `create_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "创建时间",
    `update_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "更新时间",
    `delete_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "软删除",
    PRIMARY KEY (`id`),
    KEY `delete_time`(`delete_time`)
)ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8;

# 动漫
CREATE TABLE `d_article`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(366) NOT NULL DEFAULT "",
    `img` VARCHAR(366) NOT NULL DEFAULT "" COMMENT "",
    `language` VARCHAR(255) NOT NULL DEFAULT "" COMMENT "语言编号缩写",
    `state` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT "状态0连载中1完结2停更",
    `author` VARCHAR(366) NOT NULL DEFAULT "" COMMENT "作者",
    `describe` TEXT COMMENT "描述",
    `sort_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "分类id",
    `read` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "阅读量",
    `create_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "创建时间",
    `update_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "更新时间",
    `delete_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "软删除",
    PRIMARY KEY (`id`),
    KEY `sort_id`(`sort_id`),
    KEY `delete_time`(`delete_time`),
    KEY `language`(`language`)
)ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8;

# 视频item
CREATE TABLE `d_article_item`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(366) NOT NULL DEFAULT "",
    `language` VARCHAR(255) NOT NULL DEFAULT "" COMMENT "语言编号缩写",
    `collection` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "集",
    `cartoon_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "漫画id",
    `content` TEXT COMMENT "各种类型数据，哈哈哈",
    `read` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "阅读量",
    `create_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "创建时间",
    `update_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "更新时间",
    `delete_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "软删除",
    PRIMARY KEY (`id`),
    KEY `delete_time`(`delete_time`),
    KEY `language`(`language`)
)ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8;
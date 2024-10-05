CREATE DATABASE commodity;
USE commodity;
CREATE TABLE `commodity_tab`
(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `creator_id` bigint unsigned NOT NULL COMMENT 'user_id of the creator of the commodity',
    `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'title of the commodity',
    `detail` varchar(255) NOT NULL DEFAULT '' COMMENT 'details of the commodity',
    `price` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'price of the commodity',
    `cover` varchar(255) NOT NULL COMMENT 'cover image url of the commodity',
    `type` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '0-sell, 1-buy, 2-request',
    `status` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '0-publishing, 1-done, 2-removed',
    `is_deleted` bool NOT NULL DEFAULT 0 COMMENT '0-not deleted, 1-deleted',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY   (`id`),
    KEY           `idx_creator_id` (`creator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `commodity_image_tab`
(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `commodity_id` bigint unsigned NOT NULL,
    `image` varchar(255) NOT NULL COMMENT 'image url of the commodity',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY   (`id`),
    KEY `idx_commodity_id` (`commodity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `like_tab`
(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint unsigned NOT NULL,
    `commodity_id` bigint unsigned NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY   (`id`),
    KEY `idx_commodity_user_id` (`commodity_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE DATABASE account;
USE account;
CREATE TABLE `user_tab`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `email`       varchar(255) NOT NULL DEFAULT '',
    `password`    varchar(255) NOT NULL DEFAULT '' COMMENT 'encrypted by md5',
    `status`      tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '0-available, 1-suspended, 2-deleted',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY   (`id`),
    UNIQUE KEY    `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

USE account;
CREATE TABLE `profile_tab`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`     bigint unsigned NOT NULL,
    `username`    varchar(255) NOT NULL DEFAULT '',
    `email`       varchar(255) NOT NULL DEFAULT '',
    `avatar_url`  varchar(255) NOT NULL DEFAULT '',
    `is_deleted`  tinyint NOT NULL DEFAULT 0,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY   (`id`),
    KEY           `idx_user_id` (`user_id`),
    UNIQUE KEY    `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

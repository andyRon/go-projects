
DROP TABLE IF EXISTS `room`;
CREATE TABLE `room` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `identity` varchar(36) NOT NULL,
    `name` varchar(100) NOT NULL COMMENT '房间名',
    `begin_at` datetime Comment '开始时间',
    `end_at` datetime Comment '结束时间',
    `create_id` int(11) NOT NULL COMMENT '创建人',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_room_identity` (`identity`),
    KEY `idx_room_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(100) NOT NULL COMMENT '用户名',
    `password` varchar(30) NOT NULL COMMENT '密码',
    `sdp` text COMMENT 'sdp',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `room_user`;
CREATE TABLE `room_user` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `rid` int(11) NOT NULL COMMENT '房间ID',
    `uid` int(11) NOT NULL COMMENT '用户ID',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
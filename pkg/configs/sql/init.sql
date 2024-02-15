-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
     `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
     `user_name` varchar(255) NOT NULL COMMENT '用户名',
     `password` varchar(255) NOT NULL COMMENT '用户密码',
     `email` varchar(255) NOT NULL COMMENT '用户邮箱',
     `school` varchar(255) COMMENT '学校',
     `avatar` varchar(255) COMMENT '用户头像',
     `background_image` varchar(255) COMMENT '用户个人页顶部大图',
     `signature` varchar(255) COMMENT '个人简介',
     PRIMARY KEY (`id`),
     KEY `email_password_idx` (`email`,`password`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
    `post_id` bigint NOT NULL AUTO_INCREMENT COMMENT 'post ID',
    `creator_id` varchar(255) NOT NULL COMMENT '发布者ID',
    `content` TEXT NOT NULL COMMENT '内容',
    `type` int COMMENT '类别，校内还是公开',
    `status` int NOT NULL COMMENT '',
    `create_time` varchar(255) COMMENT '创建时间',
    PRIMARY KEY (`post_id`),
    KEY `creator_idx` (`creator_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2000 DEFAULT CHARSET=utf8 COMMENT='博文表';
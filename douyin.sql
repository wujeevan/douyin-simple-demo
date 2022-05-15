CREATE DATABASE IF NOT EXISTS `douyin`;

use `douyin`;

drop table if exists `video`;

CREATE TABLE `video`(
    `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '视频唯一标识id',
    `user_id` BIGINT NOT NULL COMMENT '视频作者id',
    `play_url` VARCHAR(50) NOT NULL COMMENT '视频播放地址',
    `cover_url` VARCHAR(50) NOT NULL COMMENT '视频封面地址',
    `favorite_count` BIGINT NOT NULL DEFAULT 0 COMMENT '点赞总数',
    `comment_count` BIGINT NOT NULL DEFAULT 0 COMMENT '评论总数',
    `create_time` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `update_time` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '视频是否有效'
) DEFAULT CHARSET UTF8 COMMENT '视频表';

INSERT INTO
    `video`(
        user_id,
        play_url,
        cover_url,
        create_time,
        favorite_count,
        comment_count
    )
VALUES
    (
        1,
        '/upload/test1.mp4',
        '/upload/test1_cover.jpg',
        '2022-5-12 08:31',
        1,
        1
    ),
    (
        2,
        '/upload/test2.mp4',
        '/upload/test2_cover.jpg',
        '2022-5-12 12:25',
        1,
        1
    ),
    (
        1,
        '/upload/test3.mp4',
        '/upload/test3_cover.jpg',
        '2022-5-12 18:49',
        0,
        0
    );

use `douyin`;

drop table if exists `user`;

CREATE TABLE `user`(
    `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '用户id',
    `username` VARCHAR(32) NOT NULL COMMENT '用户名',
    `password` VARCHAR(32) NOT NULL COMMENT '用户密码',
    `token` VARCHAR(64) NOT NULL DEFAULT 'abcdefg' COMMENT '用户鉴权',
    `follow_count` BIGINT NOT NULL DEFAULT 0 COMMENT '关注总数',
    `follower_count` BIGINT NOT NULL DEFAULT 0 COMMENT '粉丝总数',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '账号是否有效'
) DEFAULT CHARSET UTF8 COMMENT '用户表';

INSERT INTO
    `user`(
        `username`,
        `password`,
        `token`,
        `follow_count`,
        `follower_count`
    )
VALUES
    ('admin', 'pass', 'adminpass', 1, 0),
    ('user', 'pass', 'userpass', 0, 1);

use `douyin`;

drop table if exists `comment`;

CREATE TABLE `comment`(
    `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '评论id',
    `user_id` BIGINT NOT NULL COMMENT '用户id',
    `video_id` BIGINT NOT NULL COMMENT '视频id',
    `content` VARCHAR(300) NOT NULL COMMENT '评论内容',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否评论',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
) DEFAULT CHARSET UTF8 COMMENT '评论表';

INSERT INTO
    `comment`(user_id, video_id, content)
VALUES
    (1, 1, '第一个评论'),
    (2, 2, '第二个评论');

use `douyin`;

drop table if exists `user_favorite_video`;

CREATE TABLE `user_favorite_video`(
    `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '点赞id',
    `user_id` BIGINT NOT NULL COMMENT '用户id',
    `video_id` BIGINT NOT NULL COMMENT '视频id',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否点赞',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
) DEFAULT CHARSET UTF8 COMMENT '点赞表';

INSERT INTO
    `user_favorite_video`(user_id, video_id, status)
VALUES
    (1, 1, 1),
    (2, 2, 1);

use `douyin`;

drop table if exists `user_follow`;

CREATE TABLE `user_follow` (
    /* 甲关注乙 */
    `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '关注id',
    `user_id` BIGINT NOT NULL COMMENT '用户甲id',
    `followed_user_id` BIGINT NOT NULL COMMENT '用户乙id',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否关注',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
) DEFAULT CHARSET UTF8 COMMENT '关注表';

INSERT INTO `user_follow`(user_id, followed_user_id) VALUES (1, 2);

SELECT
    unix_timestamp(update_time)
FROM
    video
WHERE
    unix_timestamp(update_time) <= 1652501999;
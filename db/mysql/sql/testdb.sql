CREATE TABLE `t_users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'User ID',
  `first_name` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT 'First Name',
  `last_name` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Last Name',
  `email` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'E-Mail Address',
  `password` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Password',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT 'delete flg',
  `create_datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'created date',
  `update_datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='Users Table';

INSERT INTO `t_users` (`first_name`, `last_name`, `email`, `password`, `delete_flg`, `create_datetime`, `update_datetime`)
VALUES
	('kentaro', 'asakura', 'cccc@aa.jp', '97f089f1f1f8a4d48ad4c8d3c5565e06', '0', now(), now());

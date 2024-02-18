CREATE TABLE `userbasic` (
                             `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                             `name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
                             `gender` tinyint(3) unsigned NOT NULL DEFAULT '0',
                             `mobile_phone` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
                             `email` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
                             `password` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
                             `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                             `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             PRIMARY KEY (`id`),
                             UNIQUE KEY `idx_mobile_unique` (`mobile_phone`),
                             UNIQUE KEY `idx_email_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin


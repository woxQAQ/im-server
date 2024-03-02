CREATE TABLE `single_message` (
                                  `message_id` int NOT NULL AUTO_INCREMENT,
                                  `sender_id` int NOT NULL,
                                  `receiver_id` int NOT NULL,
                                  `content` text COLLATE utf8mb4_bin NOT NULL,
                                  `timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                  `is_read` tinyint(1) DEFAULT '0',
                                  PRIMARY KEY (`message_id`),
                                  KEY `idx_receiver_id` (`receiver_id`),
                                  KEY `idx_sender_id` (`sender_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='单聊表'


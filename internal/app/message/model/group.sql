CREATE TABLE `group_users` (
                               `group_id` int NOT NULL,
                               `user_id` int DEFAULT NULL,
                               `last_ack_msg_id` int DEFAULT NULL,
                               PRIMARY KEY (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;



CREATE TABLE `group_msg` (
                             `group_id` int NOT NULL,
                             `sender_id` int DEFAULT NULL,
                             `msg_id` int DEFAULT NULL,
                             `time` timestamp NULL DEFAULT NULL,
                             `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
                             PRIMARY KEY (`group_id`),
                             KEY `idx_sender_id` (`sender_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin




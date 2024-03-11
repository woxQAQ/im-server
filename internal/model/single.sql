CREATE TABLE `message_dtl`  (
  `msg_id` int NOT NULL AUTO_INCREMENT,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `sender_id` int NOT NULL DEFAULT 0,
  `recv_id` int NOT NULL DEFAULT 0,
  `content_type` enum('text','image','file','video','audio') CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'text',
  `sendtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`msg_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;
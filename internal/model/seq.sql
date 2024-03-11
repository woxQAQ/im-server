CREATE TABLE `session_sequence`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `session_id` int NOT NULL,
  `max_seq` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sid`(`session_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

CREATE TABLE `user_sequence`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `cur_seq` int UNSIGNED NOT NULL DEFAULT 0,
  `max_seq` int UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

CREATE TABLE `t_sender_dtl`  (
                                   `id` int NOT NULL,
                                   `sender_id` int NOT NULL,
                                   `session` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                   `pre_id` int NOT NULL,
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

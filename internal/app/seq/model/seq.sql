CREATE TABLE `session_sequence`  (
                                     `session_id` int NOT NULL,
                                     `max_seq` int NOT NULL,
                                     `cur_seq` int NOT NULL,
                                     PRIMARY KEY (`session_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

CREATE TABLE `session_message_sequence`  (
                                             `session_id` int NOT NULL,
                                             `msg_id` int NOT NULL,
                                             `msg_seq_session` int NOT NULL,
                                             PRIMARY KEY (`session_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

CREATE TABLE `session_id`  (
                               `session` int NOT NULL,
                               `user_id_1` int NOT NULL,
                               `user_id_2` int NOT NULL,
                               PRIMARY KEY (`session`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

CREATE TABLE `group_sequence`  (
                                   `group_id` int NOT NULL,
                                   `cur_seq` int NOT NULL,
                                   `max_seq` int NOT NULL,
                                   PRIMARY KEY (`group_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;
CREATE TABLE `group_message_sequence`  (
                                           `group_id` int NOT NULL,
                                           `msg_id` int NOT NULL,
                                           `seq` int NOT NULL,
                                           PRIMARY KEY (`group_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;
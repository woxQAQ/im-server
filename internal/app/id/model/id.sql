/*
 Navicat Premium Data Transfer

 Source Server         : 172.22.141.30_3306
 Source Server Type    : MySQL
 Source Server Version : 80300 (8.3.0)
 Source Host           : 172.22.141.30:3306
 Source Schema         : im_id

 Target Server Type    : MySQL
 Target Server Version : 80300 (8.3.0)
 File Encoding         : 65001

 Date: 13/03/2024 20:03:46
*/
CREATE TABLE `id_generator`  (
                                 `caller_id` int NOT NULL,
                                 `caller_type` tinyint NOT NULL,
                                 `cur_seq` int NOT NULL,
                                 `max_seq` int NOT NULL,
                                 PRIMARY KEY (`caller_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

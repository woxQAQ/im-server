CREATE DEFINER=`root`@`%` PROCEDURE `create_user_tag`()
BEGIN
DECLARE i INT;
DECLARE table_name VARCHAR(20);
SET i = 0;

#控制表名
WHILE i<100 DO
IF i<10 THEN
SET table_name = CONCAT('t_user_tag0',i);
ELSE
SET table_name = CONCAT('t_user_tag',i);
END IF;

#通过CONCAT函数定义每次执行的SQL语句
SET @csql = CONCAT(
'CREATE TABLE ',table_name, '(
  `user_id` bigint NOT NULL COMMENT "主键，标识用户",
  `tag_info_01` bigint NULL DEFAULT NULL COMMENT "标签记录信息",
  `tag_info_02` bigint NULL DEFAULT NULL COMMENT "标签记录信息",
  `tag_info_03` bigint NULL DEFAULT NULL COMMENT "标签记录信息",
  `create_time` datetime NULL DEFAULT NULL COMMENT "创建时间",
  `update_time` datetime NULL DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;'
);

PREPARE create_stmt FROM @csql;
EXECUTE create_stmt;
SET i = i+1;
END WHILE;

END;
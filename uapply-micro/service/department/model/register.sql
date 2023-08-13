/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:3306
 Source Schema         : uapply-zero

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 11/02/2022 15:25:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for register
-- ----------------------------
DROP TABLE IF EXISTS `register`;
CREATE TABLE `register`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL,
  `org_id` int NOT NULL,
  `org_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `dep_id` int NOT NULL,
  `dep_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `first_status` int NOT NULL DEFAULT 0,
  `second_status` int NOT NULL DEFAULT 0,
  `final_status` int NOT NULL DEFAULT 0,
  `first_eva` tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `first_mark` int NOT NULL DEFAULT 0,
  `second_eva` tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `second_mark` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uid`(`uid`, `org_id`) USING BTREE,
  INDEX `org_id`(`org_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;

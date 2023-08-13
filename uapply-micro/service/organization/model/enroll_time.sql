/*
 Navicat Premium Data Transfer

 Source Server         : uapply-micro
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : 82.157.62.56:33066
 Source Schema         : organization_srv

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 02/12/2021 18:13:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for enroll_time
-- ----------------------------
DROP TABLE IF EXISTS `enroll_time`;
CREATE TABLE `enroll_time`  (
  `id` int NOT NULL COMMENT '组织id',
  `start_time` bigint NOT NULL COMMENT '开始的时间戳',
  `end_time` bigint NOT NULL COMMENT '结束的时间戳',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

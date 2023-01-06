/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : renshihoutai

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 06/01/2023 15:53:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for yuan_gong_zi_liao
-- ----------------------------
DROP TABLE IF EXISTS `yuan_gong_zi_liao`;
CREATE TABLE `yuan_gong_zi_liao`  (
  `id` int(0) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `department` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `entrytime` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `salary` int(0) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of yuan_gong_zi_liao
-- ----------------------------
INSERT INTO `yuan_gong_zi_liao` VALUES (1, '小孙', '教研室主任', '2020.8', 8000);
INSERT INTO `yuan_gong_zi_liao` VALUES (2, '小张', '行政', '2019.8', 6000);
INSERT INTO `yuan_gong_zi_liao` VALUES (3, '小王', '行政', '2020.1', 5000);
INSERT INTO `yuan_gong_zi_liao` VALUES (4, '小夏', '教师', '2021.3', 4500);
INSERT INTO `yuan_gong_zi_liao` VALUES (5, '小刘', '教师', '2021.11', 3000);

SET FOREIGN_KEY_CHECKS = 1;

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

 Date: 06/01/2023 15:53:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bu_men_xin_xi
-- ----------------------------
DROP TABLE IF EXISTS `bu_men_xin_xi`;
CREATE TABLE `bu_men_xin_xi`  (
  `id` int(0) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `department` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `level` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bu_men_xin_xi
-- ----------------------------
INSERT INTO `bu_men_xin_xi` VALUES (1, '小张', '行政', '主任');
INSERT INTO `bu_men_xin_xi` VALUES (2, '小王', '行政', '员工');
INSERT INTO `bu_men_xin_xi` VALUES (3, '小夏', '教师', '主任');
INSERT INTO `bu_men_xin_xi` VALUES (4, '小刘', '教师', '员工');
INSERT INTO `bu_men_xin_xi` VALUES (5, '小孙', '教研室主任', '主任');

SET FOREIGN_KEY_CHECKS = 1;

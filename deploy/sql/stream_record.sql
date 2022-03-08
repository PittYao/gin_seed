/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : stream_record

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 08/03/2022 15:26:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for record_one
-- ----------------------------
DROP TABLE IF EXISTS `record_one`;
CREATE TABLE `record_one`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rtsp_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '单画面录制的rtsp流',
  `host` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '录制的服务器ip和端口',
  `save_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '录制文件存放盘符路径',
  `ffmpeg_transform_state` int(3) NULL DEFAULT NULL COMMENT '转流任务状态 1=准备启动 2=启动成功 3=启动失败 4=关闭成功 5=关闭失败 -1=运行中异常',
  `ffmpeg_transform_cmd` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'ffmpeg转流命令',
  `ffmpeg_save_cmd` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'ffmpeg存储命令',
  `ffmpeg_save_state` int(3) NULL DEFAULT NULL COMMENT '任存储流任务状态 1=准备启动 2=启动成功 3=启动失败 4=关闭成功 5=关闭失败 -1=运行中异常',
  `create_time` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  `update_time` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of record_one
-- ----------------------------
INSERT INTO `record_one` VALUES (1, 'rtsp://admin:cebon61332433@192.168.99.215', '192.168.99.128', NULL, 2, '-rtsp_transport tcp -i rtsp://admin:cebon61332433@192.168.99.215 -c:a aac -b:a 128k -ar 16000 -ac 1 -vcodec copy -r 25 -f flv rtmp://192.168.99.128:1935/stream/5ffcdcca55e7b84539f3c5907d73624f', '', 0, '2022-02-28 09:06:47', '2022-02-28 09:06:48');

SET FOREIGN_KEY_CHECKS = 1;

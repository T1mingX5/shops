/*
MySQL Backup
Source Server Version: 8.0.28
Source Database: shops
Date: 2022/6/1 22:05:38
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
--  Table structure for `sh_admin`
-- ----------------------------
DROP TABLE IF EXISTS `sh_admin`;
CREATE TABLE `sh_admin` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '管理员账号',
  `name` varchar(10) NOT NULL COMMENT '管理员名称',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '管理员密码',
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 COMMENT='管理员表';

-- ----------------------------
--  Table structure for `sh_flow`
-- ----------------------------
DROP TABLE IF EXISTS `sh_flow`;
CREATE TABLE `sh_flow` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '用户',
  `cate` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '分类',
  `money` decimal(10,2) NOT NULL COMMENT '操作金额',
  `end_money` decimal(10,2) NOT NULL COMMENT '剩余金额',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `sh_flow_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `sh_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3 COMMENT='流水表';

-- ----------------------------
--  Table structure for `sh_goods`
-- ----------------------------
DROP TABLE IF EXISTS `sh_goods`;
CREATE TABLE `sh_goods` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名字',
  `price` decimal(10,2) NOT NULL COMMENT '价格',
  `number` int NOT NULL COMMENT '库存',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `status` tinyint NOT NULL COMMENT '状态，0为下架，1为上架',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb3 COMMENT='商品表';

-- ----------------------------
--  Table structure for `sh_order`
-- ----------------------------
DROP TABLE IF EXISTS `sh_order`;
CREATE TABLE `sh_order` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `user_id` int NOT NULL COMMENT '用户',
  `goods_id` int NOT NULL COMMENT '商品',
  `status` tinyint NOT NULL COMMENT '状态，0为未发货，1为已发货，2为拒绝，3为已完成',
  `goods_count` int NOT NULL,
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `goods_id` (`goods_id`),
  CONSTRAINT `sh_order_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `sh_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `sh_order_ibfk_2` FOREIGN KEY (`goods_id`) REFERENCES `sh_goods` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb3 COMMENT='订单表';

-- ----------------------------
--  Table structure for `sh_user`
-- ----------------------------
DROP TABLE IF EXISTS `sh_user`;
CREATE TABLE `sh_user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `account` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账号',
  `name` varchar(10) NOT NULL COMMENT '用户名称',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户密码',
  `money` decimal(10,2) NOT NULL COMMENT '用户账户余额',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb3 COMMENT='用户表';

-- ----------------------------
--  Records 
-- ----------------------------
INSERT INTO `sh_admin` VALUES ('1','123','123','123','2022-05-26 12:41:08');
INSERT INTO `sh_flow` VALUES ('1','1','购买computer','2.00','974.00','2022-06-01 18:39:00'), ('2','1','充值','10.00','1055.00','2022-06-01 19:02:58'), ('3','1','充值','0.00','1055.00','2022-06-01 19:03:01'), ('4','1','充值','20.00','1075.00','2022-06-01 19:03:06'), ('5','1','充值','10.00','1085.00','2022-06-01 21:30:08');
INSERT INTO `sh_goods` VALUES ('1','电脑','10000.00','20','笔记本电脑','1','2022-05-23 10:25:00'), ('2','手机','3999.00','20','华为手机','1','2022-05-23 12:08:33'), ('3','相机','23333.00','10','sony相机','1','2022-05-23 12:09:08'), ('4','手机','9999.00','23','iphone','1','2022-05-23 12:09:44'), ('5','电脑','8999.00','88','联想电脑，质量问题下架','0','2022-05-23 12:10:27'), ('6','电脑','3999.00','100','惠普','1','2022-05-23 12:13:58'), ('7','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('9','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('10','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('11','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('12','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('13','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('14','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('15','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('16','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('17','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('18','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('19','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('20','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('21','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('22','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('23','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('24','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('25','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('26','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('27','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('28','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('29','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('30','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('31','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('32','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('33','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('34','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('35','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('36','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('37','电脑','10000.00','3','笔记本电脑','1','2022-05-23 10:25:00'), ('38','computer','10800.00','23','ttt','1','2022-05-23 13:01:00'), ('39','computer','10800.00','23','ttt','1','2022-05-23 13:01:00'), ('40','computer','10800.00','23','ttt','1','2022-05-23 13:01:00'), ('41','computer','10800.00','23','ttt','1','2022-05-23 13:01:00'), ('42','computer','10800.00','23','ttt','1','2022-05-23 13:01:00'), ('43','computer','10800.00','20','ttt','1','2022-05-23 13:01:00'), ('44','computer','10800.00','23','ttt','1','2022-05-23 13:01:00'), ('45','computer','10800.00','10','ttt','1','2022-05-23 13:01:00'), ('46','computer','10800.00','444','ttt','1','2022-05-23 13:01:00'), ('47','computer','10800.00','333','ttt','1','2022-05-23 13:01:00'), ('48','computer','10800.00','11','ttt','1','2022-05-23 13:01:00'), ('49','computer','10800.00','222','ttt','1','2022-05-23 13:01:00'), ('50','computer','2.00','86','ttt','1','2022-05-23 13:01:00');
INSERT INTO `sh_order` VALUES ('1','1','47','0','3','2022-05-30 15:18:20'), ('2','1','45','0','13','2022-05-30 15:18:35'), ('3','1','46','0','1','2022-05-30 16:37:59'), ('4','1','48','0','11','2022-05-30 18:41:24'), ('5','1','4','0','3','2022-05-30 18:59:11'), ('6','1','43','0','1','2022-06-01 11:55:34'), ('7','1','43','0','1','2022-06-01 11:55:42'), ('8','1','43','0','1','2022-06-01 11:55:55'), ('9','1','50','0','1','2022-06-01 18:11:40'), ('10','1','50','0','1','2022-06-01 18:11:44'), ('11','1','50','0','1','2022-06-01 18:11:47'), ('12','1','50','0','1','2022-06-01 18:32:25'), ('13','1','50','0','4','2022-06-01 18:32:27'), ('14','1','50','0','3','2022-06-01 18:32:30'), ('15','1','50','0','1','2022-06-01 18:39:00');
INSERT INTO `sh_user` VALUES ('1','1272329952','111','111','1085.00','2022-05-17 19:06:57'), ('2','123','123','123','0.00','2022-05-18 15:03:53'), ('3','456','1','123321','0.00','2022-05-18 15:04:43'), ('4','655','1','123','0.00','2022-05-18 15:07:01'), ('5','321','321','123qwe','0.00','2022-05-18 15:07:51'), ('6','www','www','123','0.00','2022-05-18 15:10:15'), ('7','qqq','qqq','123','0.00','2022-05-18 15:10:45'), ('8','fff','fff','fff','0.00','2022-05-19 16:13:50'), ('9','rrr','rrr','rrr','0.00','2022-05-19 16:14:50'), ('10','12332','121','121','0.00','2022-05-19 16:19:49'), ('11','121','121','121','0.00','2022-05-19 16:20:58'), ('12','eee','121','121','0.00','2022-05-24 14:40:25'), ('13','111222','111222','111222','0.00','2022-05-27 09:56:21'), ('14','t1ming','t1ming','111','0.00','2022-05-27 12:38:48');

/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 50740 (5.7.40-log)
 Source Host           : localhost:3306
 Source Schema         : ions_project

 Target Server Type    : MySQL
 Target Server Version : 50740 (5.7.40-log)
 File Encoding         : 65001

 Date: 17/03/2025 17:12:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_auth
-- ----------------------------
DROP TABLE IF EXISTS `sys_auth`;
CREATE TABLE `sys_auth`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限名称',
  `url_for` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'url反转',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '父节点id',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `is_active` int(11) NOT NULL DEFAULT 0 COMMENT '1启用，0停用',
  `is_delete` int(11) NOT NULL DEFAULT 0 COMMENT '1删除，0未删除',
  `weight` int(11) NOT NULL DEFAULT 0 COMMENT '权重，数值越大，权重越大',
  `p_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_auth
-- ----------------------------
INSERT INTO `sys_auth` VALUES (1, '用户', '#', 0, '测试auth', '2025-03-14 14:08:07', 1, 0, 0, NULL);
INSERT INTO `sys_auth` VALUES (2, '用户修改', 'UserController.ToUpdate', 1, '新增用户', '2025-03-20 14:26:49', 1, 0, 0, '用户');
INSERT INTO `sys_auth` VALUES (3, '用户新增', 'UserController.ToAdd', 1, 'erg', '2025-02-12 14:27:17', 1, 0, 0, '用户');
INSERT INTO `sys_auth` VALUES (4, '用户展示', 'UserController.List', 1, '展示所有用户', '2025-01-14 13:51:49', 1, 0, 10, '用户');
INSERT INTO `sys_auth` VALUES (5, '车辆', '#', 0, '车辆', '2025-02-10 14:02:33', 1, 0, 0, NULL);
INSERT INTO `sys_auth` VALUES (6, '车辆租借', 'CarsController.List', 5, '车辆租赁展示', '2025-03-04 16:33:55', 1, 0, 0, '车辆');
INSERT INTO `sys_auth` VALUES (11, '角色管理', 'RoleController.List', 0, '角色', '2025-02-06 09:55:02', 1, 0, 0, '用户');
INSERT INTO `sys_auth` VALUES (12, '用户添加角色', 'RoleController.ToRoleUser', 11, '用户添加角色', '2025-03-05 17:51:38', 1, 0, 0, '角色');
INSERT INTO `sys_auth` VALUES (13, '角色添加权限', 'RoleController.ToRoleAuth', 11, '角色添加权限', '2025-03-05 17:52:28', 1, 0, 0, '角色');
INSERT INTO `sys_auth` VALUES (14, '角色展示', 'RoleController.List', 11, '角色展示', '2025-03-06 10:00:59', 1, 0, 0, '角色');
INSERT INTO `sys_auth` VALUES (15, '财务', '#', 0, '财务', '2025-03-18 16:05:39', 0, 0, 0, NULL);

-- ----------------------------
-- Table structure for sys_caiwu_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_caiwu_data`;
CREATE TABLE `sys_caiwu_data`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `caiwu_date` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '财务月份',
  `sales_volume` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '本月销售额',
  `student_incess` int(11) NOT NULL DEFAULT 0 COMMENT '学员增加数',
  `django` int(11) NOT NULL DEFAULT 0 COMMENT 'django课程卖出数量',
  `vue_django` int(11) NOT NULL DEFAULT 0 COMMENT 'vue+django课程卖出数量',
  `celery` int(11) NOT NULL DEFAULT 0 COMMENT 'celery课程卖出数量',
  `create_date` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_caiwu_data
-- ----------------------------

-- ----------------------------
-- Table structure for sys_cars
-- ----------------------------
DROP TABLE IF EXISTS `sys_cars`;
CREATE TABLE `sys_cars`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '车辆名称',
  `car_brand_id` int(11) NOT NULL COMMENT '车辆品牌外键',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT '1:可借,2:不可借',
  `is_active` int(11) NOT NULL DEFAULT 1 COMMENT '启用:1,停用:0',
  `is_delete` int(11) NOT NULL DEFAULT 0 COMMENT '删除:1,未删除:0',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_cars
-- ----------------------------

-- ----------------------------
-- Table structure for sys_cars_apply
-- ----------------------------
DROP TABLE IF EXISTS `sys_cars_apply`;
CREATE TABLE `sys_cars_apply`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `cars_id` int(11) NOT NULL,
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '申请理由',
  `destination` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '目的地',
  `return_date` date NOT NULL COMMENT '归还日期',
  `return_status` int(11) NOT NULL DEFAULT 0,
  `audit_status` int(11) NOT NULL DEFAULT 3 COMMENT '1:同意，2:未同意，3:未审批',
  `audit_option` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '审批意见',
  `is_active` int(11) NOT NULL DEFAULT 1 COMMENT '启用:1,停用:0',
  `is_delete` int(11) NOT NULL DEFAULT 0 COMMENT '删除:1,未删除:0',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `notify_tag` int(11) NOT NULL DEFAULT 0 COMMENT '1:已发送通知，0：未发送通知',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_cars_apply
-- ----------------------------

-- ----------------------------
-- Table structure for sys_cars_brand
-- ----------------------------
DROP TABLE IF EXISTS `sys_cars_brand`;
CREATE TABLE `sys_cars_brand`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '品牌名称',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '品牌描述',
  `is_active` int(11) NOT NULL DEFAULT 1 COMMENT '启用:1,停用:0',
  `is_delete` int(11) NOT NULL DEFAULT 0 COMMENT '删除:1,未删除:0',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_cars_brand
-- ----------------------------

-- ----------------------------
-- Table structure for sys_message_notify
-- ----------------------------
DROP TABLE IF EXISTS `sys_message_notify`;
CREATE TABLE `sys_message_notify`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `flag` int(11) NOT NULL DEFAULT 1 COMMENT '1:车辆逾期，2:所有通知',
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息标题',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '消息内容',
  `user_id` int(11) NOT NULL,
  `read_tag` int(11) NOT NULL DEFAULT 0 COMMENT '1:已读，0:未读',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_message_notify
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `is_active` int(11) NOT NULL DEFAULT 0,
  `is_delete` int(11) NOT NULL DEFAULT 0,
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '职位描述',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '总经理', 1, 0, '2025-03-13 14:34:16', '公司总经理，拥有所有权限');
INSERT INTO `sys_role` VALUES (2, '技术研发部主管', 1, 0, '2025-03-12 17:47:10', '技术研发部经理，拥有技术研发部权限');
INSERT INTO `sys_role` VALUES (3, '运营维护中心员工', 1, 0, '2025-03-04 14:18:55', '运营维护部员工，拥有部分权限');
INSERT INTO `sys_role` VALUES (4, '人事部主管', 1, 0, '2025-03-05 14:20:08', NULL);
INSERT INTO `sys_role` VALUES (5, '人事部员工', 1, 0, '2020-07-11 14:20:33', NULL);
INSERT INTO `sys_role` VALUES (6, '技术研发部员工', 1, 0, '2025-02-13 14:21:02', NULL);
INSERT INTO `sys_role` VALUES (7, '商务拓展部主管', 1, 0, '2025-02-14 14:21:24', NULL);
INSERT INTO `sys_role` VALUES (8, '商务拓展部员工', 1, 0, '2025-03-06 14:21:47', NULL);
INSERT INTO `sys_role` VALUES (9, '运营维护中心员工', 1, 0, '2025-03-06 14:22:31', NULL);
INSERT INTO `sys_role` VALUES (10, '财务部主管', 1, 0, '2025-03-06 14:22:48', NULL);
INSERT INTO `sys_role` VALUES (11, '财务部员工', 1, 0, '2022-03-11 14:23:01', NULL);
INSERT INTO `sys_role` VALUES (24, '后勤保障部主管', 1, 0, '2025-03-12 15:35:01', '后勤保障部主管');
INSERT INTO `sys_role` VALUES (25, '后期保障部员工', 1, 0, '2025-03-12 15:37:14', '后勤保障部员工');

-- ----------------------------
-- Table structure for sys_role_sys_auths
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_sys_auths`;
CREATE TABLE `sys_role_sys_auths`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `sys_role_id` int(11) NOT NULL,
  `sys_auth_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 53 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_sys_auths
-- ----------------------------
INSERT INTO `sys_role_sys_auths` VALUES (20, 3, 4);
INSERT INTO `sys_role_sys_auths` VALUES (21, 3, 3);
INSERT INTO `sys_role_sys_auths` VALUES (22, 3, 2);
INSERT INTO `sys_role_sys_auths` VALUES (23, 3, 1);
INSERT INTO `sys_role_sys_auths` VALUES (28, 2, 14);
INSERT INTO `sys_role_sys_auths` VALUES (29, 2, 13);
INSERT INTO `sys_role_sys_auths` VALUES (30, 2, 12);
INSERT INTO `sys_role_sys_auths` VALUES (31, 2, 11);
INSERT INTO `sys_role_sys_auths` VALUES (32, 2, 6);
INSERT INTO `sys_role_sys_auths` VALUES (33, 2, 5);
INSERT INTO `sys_role_sys_auths` VALUES (34, 2, 4);
INSERT INTO `sys_role_sys_auths` VALUES (35, 2, 3);
INSERT INTO `sys_role_sys_auths` VALUES (36, 2, 2);
INSERT INTO `sys_role_sys_auths` VALUES (37, 2, 1);
INSERT INTO `sys_role_sys_auths` VALUES (44, 1, 14);
INSERT INTO `sys_role_sys_auths` VALUES (45, 1, 13);
INSERT INTO `sys_role_sys_auths` VALUES (46, 1, 12);
INSERT INTO `sys_role_sys_auths` VALUES (47, 1, 11);
INSERT INTO `sys_role_sys_auths` VALUES (48, 1, 4);
INSERT INTO `sys_role_sys_auths` VALUES (49, 1, 3);
INSERT INTO `sys_role_sys_auths` VALUES (50, 1, 2);
INSERT INTO `sys_role_sys_auths` VALUES (51, 1, 1);
INSERT INTO `sys_role_sys_auths` VALUES (52, 1, 15);

-- ----------------------------
-- Table structure for sys_role_sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_sys_users`;
CREATE TABLE `sys_role_sys_users`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `sys_role_id` int(11) NOT NULL,
  `sys_user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_sys_users
-- ----------------------------
INSERT INTO `sys_role_sys_users` VALUES (2, 2, 1);
INSERT INTO `sys_role_sys_users` VALUES (3, 3, 1);
INSERT INTO `sys_role_sys_users` VALUES (4, 2, 2);
INSERT INTO `sys_role_sys_users` VALUES (5, 3, 2);
INSERT INTO `sys_role_sys_users` VALUES (6, 3, 3);
INSERT INTO `sys_role_sys_users` VALUES (7, 2, 3);
INSERT INTO `sys_role_sys_users` VALUES (16, 11, 1);
INSERT INTO `sys_role_sys_users` VALUES (17, 1, 1);
INSERT INTO `sys_role_sys_users` VALUES (18, 1, 9);
INSERT INTO `sys_role_sys_users` VALUES (19, 2, 24);

-- ----------------------------
-- Table structure for sys_salary_slip
-- ----------------------------
DROP TABLE IF EXISTS `sys_salary_slip`;
CREATE TABLE `sys_salary_slip`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `card_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '员工工号',
  `base_pay` decimal(12, 2) NOT NULL DEFAULT 0.00 COMMENT '基本工资',
  `working_days` decimal(3, 1) NOT NULL DEFAULT 0.0 COMMENT '工作天数',
  `days_off` decimal(3, 1) NOT NULL DEFAULT 0.0 COMMENT '请假天数',
  `days_off_no` decimal(3, 1) NOT NULL DEFAULT 0.0 COMMENT '调休天数',
  `reward` decimal(11, 2) NOT NULL DEFAULT 0.00 COMMENT '奖金',
  `rent_subsidy` decimal(11, 2) NOT NULL DEFAULT 0.00 COMMENT '租房补贴',
  `trans_subsidy` decimal(11, 2) NOT NULL DEFAULT 0.00 COMMENT '交通补贴',
  `social_security` decimal(11, 2) NOT NULL DEFAULT 0.00 COMMENT '社保',
  `house_provident_fund` decimal(11, 2) NOT NULL DEFAULT 0.00 COMMENT '住房公积金',
  `personal_pncome_tax` decimal(11, 2) NOT NULL DEFAULT 0.00 COMMENT '个税',
  `fine` decimal(11, 2) NOT NULL DEFAULT 0.00 COMMENT '罚金',
  `net_salary` decimal(11, 2) NOT NULL DEFAULT 0.00 COMMENT '实发工资',
  `pay_date` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '工资月份',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_salary_slip
-- ----------------------------
INSERT INTO `sys_salary_slip` VALUES (1, '1', 4500.00, 22.0, 0.0, 0.0, 12000.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, '2025-03', '2025-04-15 11:26:01');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `card_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '员工工号',
  `user_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `age` int(11) NULL DEFAULT NULL COMMENT '年龄',
  `gender` int(11) NULL DEFAULT NULL COMMENT '1:男,2:女,3:未知',
  `phone` bigint(20) NULL DEFAULT NULL COMMENT '电话号码',
  `addr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
  `is_active` int(11) NOT NULL DEFAULT 1 COMMENT '1启用，0停用',
  `is_delete` int(11) NOT NULL DEFAULT 0 COMMENT '1删除，0未删除',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_name`(`user_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, '1', '葛新', 'ca2f2d88e5b449688a0c8d7a03a216c8', 30, 1, 17812126257, '佰嘉城小区8号楼1单元201', 1, 0, '2025-02-07 15:04:58');
INSERT INTO `sys_user` VALUES (2, '2', '小太阳', '41ac35bff37746dfe319b8d314cccec0', 18, 1, 15673164793, '昌平区云趣园三区', 1, 0, '2024-11-07 19:45:37');
INSERT INTO `sys_user` VALUES (3, '3', '高林飞', '8fb1cda8bd27a8ff78ddac2f8ccc6b63', 40, 1, 1872536253, '北京市', 1, 0, '2021-06-16 14:51:09');
INSERT INTO `sys_user` VALUES (4, '4', '杨鹏伟', '899853b23b085f11f64df4b40c4394b6', 25, 1, 13423456784, '东北滴', 1, 0, '2020-10-06 14:53:38');
INSERT INTO `sys_user` VALUES (5, '5', '崔文珍', '97981f4173c026e64bb16aeab2f2095b', 45, 1, 154678932465, '河北', 1, 0, '2024-06-06 14:55:54');
INSERT INTO `sys_user` VALUES (6, '6', '周航', 'd1b0337dbab095205a80fa160e60cd0f', 34, 1, 18765230978, '顺义区', 1, 0, '2025-03-20 14:58:08');
INSERT INTO `sys_user` VALUES (7, '7', '程镇', '096d5026a0a9047c9731907db892a98f', 26, 1, 15263524274, '北京市海淀区', 1, 0, '2023-06-14 15:09:07');
INSERT INTO `sys_user` VALUES (8, '8', '董男', 'e191052792665afbce9d9dd37b50b416', 23, 1, 15678902365, '朝阳区', 1, 0, '2025-03-12 15:11:37');
INSERT INTO `sys_user` VALUES (9, '9', '王鸿林', '5befbd9439dae6139a5f943fb78480f4', 35, 1, 17625096274, '西城区', 1, 0, '2025-03-11 15:11:42');
INSERT INTO `sys_user` VALUES (10, '10', '崔昭', '228fff960314cd9590bf6b5a7704ea64', 43, 1, 18725390827, '山西省', 1, 0, '2019-06-19 16:09:14');
INSERT INTO `sys_user` VALUES (11, '11', '尉俊杰', '414d783f3e7150e9c5ff749647c32f24', 47, 1, 16726354253, '河北省保定市', 1, 0, '2019-10-06 16:49:02');
INSERT INTO `sys_user` VALUES (12, '12', '乔建康', '8b934cf409ac0baf4c5e134d965a7508', 43, 1, 157893245667, '山东省青岛市', 1, 0, '2020-02-06 08:50:37');
INSERT INTO `sys_user` VALUES (13, '13', '张迪', 'f78dfa141476f3397e53289da944f6cb', 36, 1, 17826542378, '山东省济南市', 1, 0, '2025-03-12 17:19:33');
INSERT INTO `sys_user` VALUES (14, '14', '周一轮', 'fbbd712d1222cd48af3616ad46e27642', 30, 1, 16728365413, '北京市顺义区', 1, 0, '2025-03-06 17:19:06');
INSERT INTO `sys_user` VALUES (18, '18', '刘冬寒', '06f5b251b87f1926082f62813e44a34e', 36, 2, 18726542345, '北京市丰台区', 1, 0, '2025-03-06 17:24:59');
INSERT INTO `sys_user` VALUES (19, '19', '董俊', '9af99420312652f8b5a03bdf3eee6f70', 24, 1, 17625344256, '北京市', 1, 0, '2025-03-06 17:34:36');
INSERT INTO `sys_user` VALUES (20, '20', '张维嘉', 'ac0e76b5df23ba5a88400caa20c78147', 35, 1, 18765243217, '山东省威海市', 1, 0, '2025-03-06 17:46:56');
INSERT INTO `sys_user` VALUES (21, '21', '杨芳', '0f95d306b3421159db7fc4ad2a19d51b', 34, 2, 15782736541, '北京市', 1, 0, '2025-03-06 17:51:22');
INSERT INTO `sys_user` VALUES (22, '22', '马同森', '7f8199dc9c8242840398aef00d6bbdc0', 43, 1, 17625431098, '北京市东城区', 1, 0, '2025-03-06 17:54:17');
INSERT INTO `sys_user` VALUES (23, '23', '小杨同学', '7d00074ec4db88d18dd1262f800bed3d', 18, 1, 18972635426, '北京市延庆区', 1, 0, '2025-03-11 15:49:19');
INSERT INTO `sys_user` VALUES (24, '24', '刘映芳', 'c585e9df568f06a179c0cf1a21995a32', 56, 2, 15292258746, '湖南省湘潭市湘乡市', 1, 0, '2025-03-12 13:43:16');

-- ----------------------------
-- Table structure for tree
-- ----------------------------
DROP TABLE IF EXISTS `tree`;
CREATE TABLE `tree`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `url_for` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `weight` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tree
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `age` int(11) NULL DEFAULT NULL COMMENT '年龄',
  `gender` int(11) NULL DEFAULT NULL COMMENT '性别',
  `phone` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '电话号码',
  `addr` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '住址',
  `create_time` datetime NULL DEFAULT NULL,
  `is_deleted` tinyint(1) UNSIGNED ZEROFILL NULL DEFAULT NULL,
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_active` int(11) NULL DEFAULT 1,
  `nick_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_name`(`user_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '葛新', 'ca2f2d88e5b449688a0c8d7a03a216c8', 30, 1, '17812126257', '佰嘉城小区', '2025-02-25 10:09:53', 0, 'gexin17812126257@163.com', 1, '强大大大我');
INSERT INTO `user` VALUES (2, '王科', '52a5397052960ee20191240d6ba3f896', 29, 1, '18586268625', '昌平区', '2025-02-25 10:02:23', 0, 'wangke@qq.com', 0, '七点额度强大');
INSERT INTO `user` VALUES (3, '杨鹏伟', '899853b23b085f11f64df4b40c4394b6', 35, 3, '17653625642', '北京市', '2025-02-25 10:09:03', 0, 'ypw@qq.com', 1, '全额4答案的');
INSERT INTO `user` VALUES (4, '宋威', '98c70f4c869818343aef30175cdf38f6', 28, 1, '17624359072', '北京人家小区', '2025-01-17 12:54:28', 0, 'sw@qq.com', 0, '去西安方法多少钱啊');
INSERT INTO `user` VALUES (5, '王鸿林', 'b0263ad905833fb08b7f63b7f5af7adb', 40, 1, '15264826522', '北京市', '2025-01-21 16:55:51', 0, 'whl@qq.com', 1, 'QQ群房室传导');
INSERT INTO `user` VALUES (6, '崔文珍', 'c8fafc1be1ca8c20adbdd5dc3eb7f650', 51, 1, '165243178902', '北京市', '2024-12-12 09:59:50', 0, 'cwzh@qq.com', 1, '确定二热请重新');
INSERT INTO `user` VALUES (7, '程镇', 'ffc2eb6976ed805dc00c1f9ee75ada7f', 25, 1, '176253425672', '北京市', '2024-12-26 17:10:31', 0, 'chzh@qq.com', 1, '去超市QFR');
INSERT INTO `user` VALUES (8, '高林飞', '0d2c037cdec6ee00b52ab73793b69e54', 35, 1, '17625346782', '石景山区', '2024-12-11 17:13:19', 0, 'glf@qq.com', 1, '青曲社不能去我');
INSERT INTO `user` VALUES (9, '马司林', 'a06a83536ed38d22d053946248c41c18', 39, 1, '13524387900', '北京市东城区神话大厦', '2023-02-22 17:15:44', 0, 'msl@qq.com', 1, 'QGA145SVC');
INSERT INTO `user` VALUES (10, '张迪', '6cef3b339c1d0a7599df81a6939d16b7', 36, 1, '17625309729', '北京市', '2022-07-15 17:16:53', 0, 'zhd@qq.com', 1, '爱她额度奇热网');
INSERT INTO `user` VALUES (11, '蔚俊杰', '8eeecfe30166ac4b76cf6358c0eb3304', 25, 1, '17324563728', '北京市西城区', '2025-02-20 17:18:41', 0, 'wjj@qq.com', 1, 'DSR分威神V');
INSERT INTO `user` VALUES (12, '张健康', 'c3dc72ed79b0e96991566a531eac7d69', 43, 1, '15672839871', '河北省保定市', '2024-03-05 17:19:57', 0, 'zhjk@qq.com', 1, '亲人43发达');
INSERT INTO `user` VALUES (13, '杨芳', '2879d4f6c6fd321256e20afca49b23b0', 35, 0, '17524359870', '北京市昌平区', '2024-10-09 17:21:29', 0, 'yf@qq.com', 1, '全国范围为 ');
INSERT INTO `user` VALUES (14, '刘映芳', 'c585e9df568f06a179c0cf1a21995a32', 57, 0, '15292258746', '湖南省湘潭市', '2024-09-19 17:23:30', 0, 'lyf@qq.com', 1, '阿V纹挺多的我');
INSERT INTO `user` VALUES (16, '小太阳', '41ac35bff37746dfe319b8d314cccec0', 30, 1, '15673164793', '北京市昌平区龙泽圆街道云趣园三区20-3-501', '2025-02-19 17:15:36', 0, '861439031@qq.com', 1, '单一担任还');
INSERT INTO `user` VALUES (17, 'admin', '106eb6c8fe8442d3c6ee7bff5c16cae5', 12, 0, '76253415267', '深化研究院', '2025-02-19 17:28:59', 0, 'admin@chnenery@ceic.com', 0, 'okanhywv');
INSERT INTO `user` VALUES (19, 'admin1', '106eb6c8fe8442d3c6ee7bff5c16cae5', 56, 1, '16725345266', '北京市东城区鼓楼大街', '2025-02-19 17:46:20', 0, 'admin1@chnenergy.com', 1, 'amdihbga');
INSERT INTO `user` VALUES (20, 'admin2', '58c93d7b0806c98726e3ff089e9814c6', 28, 1, '18726354526', '湖南省长沙市', '2025-02-20 09:31:57', 0, 'admin2@chnenergy.com', 1, 'haihknhc');
INSERT INTO `user` VALUES (21, 'admin3', '7aab6f8c36abef7d469b09f658d6e35c', 13, 1, '17625443526', '天津市', '2025-02-20 09:34:01', 0, 'admin3@qq.com', 1, 'ahihnhaa');
INSERT INTO `user` VALUES (22, 'admin4', 'd6252057d328318fdbf0c2a20d57ce7b', 14, 1, '17826542342', '北京市', '2025-02-20 10:26:33', 0, 'admin4@chnenergy.com', 1, 'ajianhnf');
INSERT INTO `user` VALUES (23, 'admin5', 'e3ecb8603e7ba6078659b741d181cfc0', 13, 1, '19872653452', '武汉市', '2025-02-20 10:28:09', 0, 'admin5@chnenergy.com', 0, 'aggbndvss');

SET FOREIGN_KEY_CHECKS = 1;

-- MySQL dump 10.13  Distrib 8.0.43, for Linux (x86_64)
--
-- Host: localhost    Database: smart_engi
-- ------------------------------------------------------
-- Server version	8.0.43-0ubuntu0.24.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Dumping data for table `client_order_types`
--

LOCK TABLES `client_order_types` WRITE;
/*!40000 ALTER TABLE `client_order_types` DISABLE KEYS */;
INSERT INTO `client_order_types` VALUES ('after_sales','售后服务','success'),('complaint','客户投诉','danger'),('consultation','技术咨询','info'),('maintenance','维护请求','warning'),('suggestion','建议反馈','info'),('technical','技术支持','primary');
/*!40000 ALTER TABLE `client_order_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `client_orders`
--

LOCK TABLES `client_orders` WRITE;
/*!40000 ALTER TABLE `client_orders` DISABLE KEYS */;
INSERT INTO `client_orders` VALUES (1,'2024-02-05 17:30:00.000','2024-02-06 22:20:00.000',1,3,'technical','空调系统制冷效果不佳','酒店36楼的中央空调系统制冷效果明显下降，客房温度无法降到设定值，影响客户体验，请尽快安排技术人员检查。','已安排技术人员现场检查，发现制冷剂泄漏问题，已完成补充和密封修复，系统运行正常。',1),(2,'2024-02-12 22:15:00.000','2024-02-15 18:45:00.000',2,4,'maintenance','写字楼电梯供电异常','A栋写字楼的3号电梯经常出现供电不稳定，导致电梯停运，严重影响办公楼正常使用。','检查发现电源线路老化，已更换相关线路和电气元件，电梯供电恢复正常。',1),(3,'2024-02-20 19:00:00.000','2024-02-23 00:30:00.000',4,7,'after_sales','泳池水质监测设备故障','主泳池的自动水质监测设备显示异常，无法正常读取pH值和余氯含量，需要技术支持。','已更换传感器模块并重新校准设备，水质监测功能恢复正常，建议定期维护保养。',1),(4,'2024-03-01 21:45:00.000','2024-03-01 21:45:00.000',3,NULL,'consultation','新项目消防系统设计咨询','新开发项目需要设计消防给水系统，希望了解最新的消防规范要求和推荐的设备选型。','',0),(5,'2024-03-08 18:20:00.000','2024-03-10 23:00:00.000',5,9,'complaint','施工现场噪音过大','贵公司在园区施工期间噪音过大，影响周边厂房正常生产，希望能够采取降噪措施。','已调整施工时间安排，避开生产高峰期，并采用低噪音设备，感谢您的理解和配合。',1),(6,'2024-03-16 00:30:00.000','2024-03-16 00:30:00.000',1,NULL,'suggestion','建议增加节能控制系统','建议在现有空调系统基础上增加智能节能控制系统，以降低能耗和运营成本。','',0);
/*!40000 ALTER TABLE `client_orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `clients`
--

LOCK TABLES `clients` WRITE;
/*!40000 ALTER TABLE `clients` DISABLE KEYS */;
INSERT INTO `clients` VALUES (1,'2024-01-15 17:00:00.000','2024-01-15 17:00:00.000','香港国际酒店集团','陈志明','+852-2234-5678','chen.zhiming@hkhotel.com','香港中环金融街88号国际大厦36楼','五星级酒店连锁集团，主要业务包括空调系统维护和泳池设备管理',0,0,0,0),(2,'2024-02-03 22:30:00.000','2024-02-03 22:30:00.000','深圳科技园物业管理有限公司','李小红','+86-755-8876-5432','li.xiaohong@sztech.com','深圳市南山区科技园南区高新南一道9号','科技园区物业管理，负责多栋写字楼的机电设备维护',0,0,0,0),(3,'2024-02-20 18:15:00.000','2024-03-06 00:45:00.000','广州恒大地产开发','王大伟','+86-20-3456-7890','wang.dawei@evergrande.com','广州市天河区珠江新城华夏路10号富力中心','大型房地产开发商，主要需求为新建楼盘的水电安装工程',0,0,0,0),(4,'2024-03-08 19:20:00.000','2024-03-08 19:20:00.000','澳门威尼斯人度假村','黄美玲','+853-2882-8888','huang.meiling@venetian.mo','澳门氹仔金光大道威尼斯人度假村','大型综合度假村，包含酒店、娱乐场、购物中心，设备维护需求量大',0,0,0,0),(5,'2024-03-15 16:45:00.000','2024-03-15 16:45:00.000','东莞制造业园区','刘建国','+86-769-2345-6789','liu.jianguo@dgpark.com','广东省东莞市长安镇工业园区创新路168号','工业园区管理方，主要负责园区内厂房的水电气配套设施',0,0,0,0);
/*!40000 ALTER TABLE `clients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `construction_photos`
--

LOCK TABLES `construction_photos` WRITE;
/*!40000 ALTER TABLE `construction_photos` DISABLE KEYS */;
INSERT INTO `construction_photos` VALUES (1,'2024-02-05 20:30:00.000',1,'https://axogc.net:520/file/construction-photos/air_conditioner_1.png'),(2,'2024-02-05 23:45:00.000',1,'https://axogc.net:520/file/construction-photos/air_conditioner_2.png'),(3,'2024-02-06 01:00:00.000',1,'https://axogc.net:520/file/construction-photos/air_conditioner_3.png'),(4,'2024-02-18 19:20:00.000',2,'https://axogc.net:520/file/construction-photos/clean_hotel_1.png'),(5,'2024-02-19 00:30:00.000',2,'https://axogc.net:520/file/construction-photos/clean_hotel_2.png'),(6,'2024-02-26 18:15:00.000',3,'https://axogc.net:520/file/construction-photos/office_building_1.png'),(7,'2024-02-26 22:45:00.000',3,'https://axogc.net:520/file/construction-photos/office_building_2.png'),(8,'2024-02-27 00:30:00.000',3,'https://axogc.net:520/file/construction-photos/office_building_3.png'),(9,'2024-03-12 19:30:00.000',4,'https://axogc.net:520/file/construction-photos/energy_monitor_1.png'),(10,'2024-03-12 23:00:00.000',4,'https://axogc.net:520/file/construction-photos/energy_monitor_2.png'),(11,'2024-03-20 18:45:00.000',5,'https://axogc.net:520/file/construction-photos/basement_pipes_1.png'),(12,'2024-03-20 22:20:00.000',5,'https://axogc.net:520/file/construction-photos/basement_pipes_2.png'),(13,'2024-03-21 00:15:00.000',5,'https://axogc.net:520/file/construction-photos/basement_pipes_3.png'),(14,'2024-03-10 19:00:00.000',6,'https://axogc.net:520/file/construction-photos/swimming_pool_1.png'),(15,'2024-03-10 21:30:00.000',6,'https://axogc.net:520/file/construction-photos/swimming_pool_2.png');
/*!40000 ALTER TABLE `construction_photos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `construction_records`
--

LOCK TABLES `construction_records` WRITE;
/*!40000 ALTER TABLE `construction_records` DISABLE KEYS */;
INSERT INTO `construction_records` VALUES (1,'2024-02-06 02:30:00.000',1,'2024-02-05 16:00:00.000','2024-02-06 01:30:00.000','完成酒店36楼中央空调主机的初步检查，检测制冷剂压力和系统运行参数，更换老化的密封件。','发现制冷剂轻微泄漏，部分管道保温层老化脱落。','已补充制冷剂并修复泄漏点，重新包扎保温层，系统压力恢复正常。',3),(2,'2024-02-19 03:00:00.000',1,'2024-02-18 17:00:00.000','2024-02-19 02:00:00.000','对酒店大堂和餐厅区域的风机盘管进行深度清洁和维护，更换过滤网，调试温控系统。','部分风机盘管噪音较大，温控器反应迟缓。','清洁风机叶轮，润滑轴承，更换故障温控器，噪音问题得到解决。',7),(3,'2024-02-27 01:45:00.000',2,'2024-02-26 16:30:00.000','2024-02-27 01:00:00.000','在A栋写字楼安装智能照明控制系统，铺设控制线路，配置中央控制柜。','原有配电箱空间不足，需要扩容改造。','增设辅助配电箱，重新规划线路布局，确保系统安全可靠运行。',4),(4,'2024-03-13 00:20:00.000',2,'2024-03-12 17:00:00.000','2024-03-13 00:00:00.000','完成B栋写字楼能耗监控系统的传感器安装和数据采集设备调试。','部分传感器信号不稳定，数据传输偶有中断。','检查并重新连接松动的线路接头，更换故障传感器，系统运行正常。',10),(5,'2024-03-21 02:15:00.000',3,'2024-03-20 16:00:00.000','2024-03-21 01:30:00.000','完成1号楼地下室给水主管道的安装和试压测试，安装增压泵房设备。','发现部分管道接口渗水，泵房通风不良。','重新紧固管道接口，增加密封措施，安装排风扇改善通风条件。',9),(6,'2024-03-11 01:30:00.000',4,'2024-03-10 17:30:00.000','2024-03-11 00:45:00.000','完成主泳池现场勘察，测量池体尺寸，评估现有设备状况，制定改造方案。','现有过滤设备处理能力不足，循环泵老化严重。','建议更换大容量过滤器和高效循环泵，优化管路设计提高循环效率。',7);
/*!40000 ALTER TABLE `construction_records` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `contracts`
--

LOCK TABLES `contracts` WRITE;
/*!40000 ALTER TABLE `contracts` DISABLE KEYS */;
INSERT INTO `contracts` VALUES (1,'2024-01-20 18:00:00.000','2024-01-25 22:30:00.000','香港国际酒店集团合同',1,1,1,'https://axogc.net:520/file/contracts/hotel_group_hvac_maintenance.pdf'),(2,'2024-02-05 17:15:00.000','2024-02-11 00:45:00.000','深圳科技园物业管理有限公司合同',2,1,1,'https://axogc.net:520/file/contracts/sztech_electrical_upgrade.pdf'),(3,'2024-02-25 19:30:00.000','2024-03-01 18:00:00.000','广州恒大地产开发合同',3,1,0,'https://axogc.net:520/file/contracts/evergrande_plumbing_installation.pdf'),(4,'2024-03-10 22:00:00.000','2024-03-15 17:20:00.000','澳门威尼斯人度假村合同',4,1,1,'https://axogc.net:520/file/contracts/venetian_pool_system_upgrade.pdf'),(5,'2024-03-18 21:45:00.000','2024-03-18 21:45:00.000','东莞制造业园区合同',5,0,0,'https://axogc.net:520/file/contracts/dongguan_park_infrastructure.pdf');
/*!40000 ALTER TABLE `contracts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `courses`
--

LOCK TABLES `courses` WRITE;
/*!40000 ALTER TABLE `courses` DISABLE KEYS */;
INSERT INTO `courses` VALUES (1,'2024-01-10 17:00:00.000','2024-01-10 17:00:00.000','空调系统安装与维护','学习中央空调、分体空调的安装流程、维护方法和故障排除技巧','//player.bilibili.com/player.html?isOutside=true&aid=923554389&bvid=BV1fT4y1J7T1&cid=1401494818&p=1',0),(2,'2024-01-12 18:30:00.000','2024-01-12 18:30:00.000','水电工程基础','掌握建筑水电工程的基本原理、施工规范和安全操作要求','//player.bilibili.com/player.html?isOutside=true&aid=114895660059137&bvid=BV1rpgJzvEtG&cid=31204181775&p=1',0),(3,'2024-01-15 22:00:00.000','2024-01-15 22:00:00.000','泳池设备管理','学习泳池循环系统、过滤设备、加热系统的操作和维护','//player.bilibili.com/player.html?isOutside=true&aid=553622063&bvid=BV1Ti4y1U7Pr&cid=587841722&p=1',0),(4,'2024-01-18 17:15:00.000','2024-01-18 17:15:00.000','工程安全培训','工程现场安全规范、个人防护用品使用、应急处理程序培训','//player.bilibili.com/player.html?isOutside=true&aid=880846886&bvid=BV1jK4y167Hi&cid=1394441905&p=1',0),(5,'2024-01-20 23:30:00.000','2024-02-05 19:15:00.000','电气工程高级技术','高压电气设备安装、智能控制系统集成、故障诊断与维修','//player.bilibili.com/player.html?isOutside=true&aid=383254970&bvid=BV1DZ4y1y74c&cid=578105646&p=1',0),(6,'2024-02-01 21:45:00.000','2024-02-01 21:45:00.000','项目管理实务','工程项目计划制定、进度控制、质量管理和团队协调技巧','//player.bilibili.com/player.html?isOutside=true&aid=113226863937156&bvid=BV1MNx4e1E4y&cid=26080513645&p=1',0);
/*!40000 ALTER TABLE `courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `device_statuses`
--

LOCK TABLES `device_statuses` WRITE;
/*!40000 ALTER TABLE `device_statuses` DISABLE KEYS */;
INSERT INTO `device_statuses` VALUES ('broken','已损坏','danger'),('heavy_wear','重度使用','warning'),('in_repair','维修中','info'),('light_wear','轻度使用','success'),('new_device','新设备','primary');
/*!40000 ALTER TABLE `device_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `devices`
--

LOCK TABLES `devices` WRITE;
/*!40000 ALTER TABLE `devices` DISABLE KEYS */;
INSERT INTO `devices` VALUES (1,'中央空调主机','大型商用中央空调系统主机，适用于大型建筑物温控','主机房 A1','new_device'),(2,'水泵设备','高压水泵，用于建筑供水系统和循环水系统','水泵房 B2','light_wear'),(3,'电气控制柜','智能电气控制系统，用于设备自动化控制和监测','配电室 C3','heavy_wear'),(4,'泳池过滤器','专业泳池水循环过滤设备，保证水质清洁','泳池设备间 D4','in_repair'),(5,'发电机组','备用柴油发电机，为重要设备提供应急电源','应急电源区 E5','in_repair'),(6,'管道焊接设备','专业管道焊接工具，用于水管、气管的连接和维修','维修工具库 F6','heavy_wear'),(7,'电气测试仪器','万用表、绝缘测试仪等电气检测设备','检测实验室 G7','broken'),(8,'起重设备','小型起重机和滑轮组，用于重型设备的安装和移动','仓库 H8','light_wear');
/*!40000 ALTER TABLE `devices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `invoice_statuses`
--

LOCK TABLES `invoice_statuses` WRITE;
/*!40000 ALTER TABLE `invoice_statuses` DISABLE KEYS */;
INSERT INTO `invoice_statuses` VALUES ('approved','已审核','primary'),('cancelled','已取消','info'),('draft','草稿','info'),('overdue','逾期未付','danger'),('paid','已付款','success'),('pending','待审核','warning'),('sent','已发送','primary');
/*!40000 ALTER TABLE `invoice_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `invoices`
--

LOCK TABLES `invoices` WRITE;
/*!40000 ALTER TABLE `invoices` DISABLE KEYS */;
INSERT INTO `invoices` VALUES (1,'2024-02-20 17:00:00.000',1,'香港国际酒店中央空调系统维护升级',285600.00,'paid','2024-03-21 07:59:59.000'),(2,'2024-03-01 18:30:00.000',2,'深圳科技园电气系统智能化改造',156800.00,'sent','2024-04-01 07:59:59.000'),(3,'2024-03-10 22:15:00.000',1,'香港国际酒店中央空调系统维护升级 (第二期)',198500.00,'approved','2024-04-11 07:59:59.000'),(4,'2024-03-15 19:45:00.000',3,'广州恒大新楼盘水电安装工程',89600.00,'pending','2024-04-16 07:59:59.000'),(5,'2024-03-18 21:20:00.000',4,'澳门威尼斯人泳池系统升级改造',127300.00,'draft','2024-04-19 07:59:59.000'),(6,'2024-02-28 23:30:00.000',2,'深圳科技园电气系统智能化改造 (材料费)',76400.00,'overdue','2024-03-16 07:59:59.000'),(7,'2024-03-06 00:00:00.000',5,'东莞工业园区基础设施配套建设',45600.00,'cancelled','2024-04-06 07:59:59.000');
/*!40000 ALTER TABLE `invoices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `materials`
--

LOCK TABLES `materials` WRITE;
/*!40000 ALTER TABLE `materials` DISABLE KEYS */;
INSERT INTO `materials` VALUES (1,'2024-02-01 18:30:00.000','美的中央空调主机MDV-D280W',28.00,'台','280kW变频多联机外机，适用于大型商业建筑'),(2,'2024-02-05 22:45:00.000','美的风机盘管FP-102WA',850.00,'台','卧式暗装风机盘管，额定制冷量5.6kW'),(3,'2024-02-10 17:20:00.000','施耐德智能照明控制模块C-Bus',380.00,'个','8通道智能照明控制器，支持DALI协议'),(4,'2024-02-16 00:30:00.000','施耐德电能监测仪PM5560',120.00,'台','三相电能质量分析仪，精度0.5S级'),(5,'2024-02-20 19:00:00.000','联塑PVC-U给水管DN200',5600.00,'米','给水用硬聚氯乙烯管材，公称压力1.6MPa'),(6,'2024-02-25 21:45:00.000','联塑PPR热水管DN110',4200.00,'米','耐热聚丙烯管材，适用于热水系统'),(7,'2024-03-01 23:20:00.000','海龙泳池砂缸过滤器HL-1200',35.00,'台','直径1200mm玻璃钢砂缸，流量60m³/h'),(8,'2024-03-05 18:15:00.000','海龙臭氧发生器OZ-50G',18.00,'台','泳池水处理专用，臭氧产量50g/h'),(9,'2024-03-10 22:30:00.000','格兰富水泵CR32-4',65.00,'台','立式多级离心泵，流量32m³/h，扬程92m'),(10,'2024-03-15 19:45:00.000','格兰富潜水泵SEG.40.12.2.50B',88.00,'台','污水潜水泵，带切割装置，功率1.2kW'),(11,'2024-03-21 00:00:00.000','正泰断路器NXB-63 C32',2800.00,'个','小型断路器，额定电流32A，分断能力6kA'),(12,'2024-03-25 17:30:00.000','正泰接触器NC1-32',2100.00,'个','交流接触器，额定电流32A，线圈电压AC220V');
/*!40000 ALTER TABLE `materials` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `project_devices`
--

LOCK TABLES `project_devices` WRITE;
/*!40000 ALTER TABLE `project_devices` DISABLE KEYS */;
INSERT INTO `project_devices` VALUES (1,1),(3,2),(4,2),(5,2),(1,3),(2,3),(4,3),(4,4),(2,5),(5,5),(3,6),(5,6),(1,7),(2,7),(3,7),(3,8),(5,8);
/*!40000 ALTER TABLE `project_devices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `project_materials`
--

LOCK TABLES `project_materials` WRITE;
/*!40000 ALTER TABLE `project_materials` DISABLE KEYS */;
INSERT INTO `project_materials` VALUES (1,1,5.00),(1,2,120.00),(2,3,85.00),(2,4,25.00),(3,5,800.00),(3,6,600.00),(3,11,200.00),(3,12,150.00),(4,7,8.00),(4,8,4.00),(5,9,12.00),(5,10,15.00);
/*!40000 ALTER TABLE `project_materials` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `project_statuses`
--

LOCK TABLES `project_statuses` WRITE;
/*!40000 ALTER TABLE `project_statuses` DISABLE KEYS */;
INSERT INTO `project_statuses` VALUES (1,'项目启动','项目正式启动，完成前期准备工作和人员安排',0,1,'2024-02-06 01:00:00.000','2024-02-04 00:30:00.000'),(2,'设备检测','对现有空调设备进行全面检测和评估',0,1,'2024-02-21 01:00:00.000','2024-02-18 23:45:00.000'),(3,'维护升级','执行设备维护和升级改造工作',0,1,'2024-04-16 01:00:00.000',NULL),(4,'方案设计','完成电气系统智能化改造的详细设计方案',0,2,'2024-02-29 01:00:00.000','2024-02-26 22:20:00.000'),(5,'设备采购','采购智能控制设备和相关材料',0,2,'2024-03-16 01:00:00.000','2024-03-12 19:30:00.000'),(6,'系统安装','安装智能控制系统和配套设备',0,2,'2024-05-01 01:00:00.000',NULL),(7,'现场勘察','完成施工现场勘察和技术交底',0,3,'2024-03-21 01:00:00.000',NULL),(8,'前期准备','项目前期准备工作，包括设计方案和材料准备',0,4,'2024-04-11 01:00:00.000',NULL),(9,'可行性研究','进行项目可行性研究和初步方案制定',0,5,'2024-04-21 01:00:00.000',NULL);
/*!40000 ALTER TABLE `project_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `projects`
--

LOCK TABLES `projects` WRITE;
/*!40000 ALTER TABLE `projects` DISABLE KEYS */;
INSERT INTO `projects` VALUES (1,'2024-01-25 17:00:00.000','2024-03-16 00:30:00.000','香港国际酒店中央空调系统维护升级','对香港国际酒店集团旗下5栋建筑的中央空调系统进行全面维护和升级改造，包括冷水机组更换、风管系统清洁、控制系统智能化升级、节能设备安装等工作，确保酒店空调系统运行稳定、节能环保。','https://axogc.net:520/file/specification/hk_hotel_maintenance.pdf','1. 更换老化的冷水机组，提升制冷效率30%以上；\n2. 清洁所有风管系统，确保空气质量达标；\n3. 升级BMS控制系统，实现智能温控和远程监控；\n4. 安装变频设备和节能控制器；\n5. 提供24个月质保服务；\n6. 施工期间不影响酒店正常营业；\n7. 符合香港机电工程署相关标准。','2024-02-01 16:00:00.000','2024-05-01 02:00:00.000',1),(2,'2024-02-10 18:30:00.000','2024-03-20 22:15:00.000','深圳科技园电气系统智能化改造','深圳科技园区A、B、C三栋办公楼的电气系统智能化改造项目，涵盖配电系统升级、照明系统智能化、安防系统集成、能耗监测系统建设等，打造现代化智慧办公园区。','https://axogc.net:520/file/specification/shenzhen_tech_renovation.pdf','1. 升级10KV配电系统，提高供电可靠性；\n2. 安装智能照明控制系统，节能率达到40%；\n3. 集成门禁、监控、消防等安防系统；\n4. 建设能耗监测平台，实现精细化能源管理；\n5. 提供运维培训和技术支持；\n6. 符合国家建筑电气设计规范；\n7. 通过ISO14001环境管理体系认证。','2024-02-15 17:00:00.000','2024-05-16 01:00:00.000',2),(3,'2024-03-01 19:00:00.000','2024-03-01 19:00:00.000','广州恒大新楼盘水电安装工程','广州恒大新开发楼盘一期工程的水电安装项目，包括6栋高层住宅楼和1个商业综合体的给排水系统、电气系统、弱电系统等基础设施安装，为后续装修和入住做好准备。','https://axogc.net:520/file/specification/guangzhou_installation.pdf','1. 完成所有单元的给排水管道安装，确保供水压力稳定；\n2. 安装配电箱、电缆桥架、照明线路等电气设施；\n3. 布设网络、电话、电视等弱电系统线路；\n4. 安装消防给水系统和自动喷淋系统；\n5. 通过相关质量验收和安全检查；\n6. 提供完整的竣工图纸和技术资料；\n7. 符合国家住宅建筑标准。','2024-03-15 16:00:00.000',NULL,3),(4,'2024-03-15 22:30:00.000','2024-03-15 22:30:00.000','澳门威尼斯人泳池系统升级改造','澳门威尼斯人度假村室内外泳池系统的全面升级改造，包括水处理设备更新、循环系统优化、水质监测系统安装、加热系统改造等，提升泳池水质和客户体验。','https://axogc.net:520/file/specification/macao_pool_renovation.pdf','1. 更换高效水处理设备，确保水质达到国际五星级酒店标准；\n2. 优化水循环系统，提高循环效率；\n3. 安装自动水质监测和调节系统；\n4. 改造恒温加热系统，实现精准温控；\n5. 升级泳池照明系统，营造良好氛围；\n6. 提供操作培训和维护手册；\n7. 符合澳门卫生署泳池管理规定。','2024-04-01 17:00:00.000',NULL,4),(5,'2024-03-20 23:45:00.000','2024-03-20 23:45:00.000','东莞工业园区基础设施配套建设','东莞制造业园区新建厂房的基础设施配套建设项目，涵盖供电系统、给排水系统、消防系统、通风系统等各项基础设施的设计和安装，为入驻企业提供完善的生产环境。','https://axogc.net:520/file/specification/dongguan_infrastructure.pdf','1. 建设35KV变电站和配电网络，满足园区用电需求；\n2. 铺设给排水管网，包括生产用水和生活污水处理系统；\n3. 安装完整的消防系统，包括自动报警和灭火系统；\n4. 建设工业通风和废气处理系统；\n5. 预留光纤网络和通信设施接口；\n6. 符合工业园区规划和环保要求；\n7. 通过相关部门验收和许可。',NULL,NULL,5);
/*!40000 ALTER TABLE `projects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES ('admin','系统管理员','warning'),('engineer','工程师','info'),('finance','财务人员','success'),('hr','人力资源','danger'),('manager','项目经理','primary');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `supplier_order_materials`
--

LOCK TABLES `supplier_order_materials` WRITE;
/*!40000 ALTER TABLE `supplier_order_materials` DISABLE KEYS */;
INSERT INTO `supplier_order_materials` VALUES (1,1,3.00,258000.00),(1,2,85.00,900.00),(2,3,60.00,2800.00),(2,4,15.00,20000.00),(3,5,500.00,280.00),(3,6,400.00,320.00),(4,7,6.00,48000.00),(4,8,3.00,26000.00),(5,9,12.00,8500.00),(5,10,15.00,3200.00),(6,11,200.00,380.00),(6,12,150.00,480.00);
/*!40000 ALTER TABLE `supplier_order_materials` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `supplier_orders`
--

LOCK TABLES `supplier_orders` WRITE;
/*!40000 ALTER TABLE `supplier_orders` DISABLE KEYS */;
INSERT INTO `supplier_orders` VALUES (1,'2024-02-01 17:00:00.000','2024-02-04 00:30:00.000','香港酒店项目-美的空调设备采购订单','采购香港国际酒店中央空调系统维护升级项目所需的美的品牌空调主机及配套风机盘管设备。要求：1.设备需符合香港能效标准；2.提供原厂质保证书；3.包含现场技术指导服务。',1,856000,'2024-02-25 08:00:00.000'),(2,'2024-02-15 18:30:00.000','2024-02-16 22:20:00.000','深圳科技园-施耐德智能电气设备订单','深圳科技园电气系统智能化改造项目采购清单：智能照明控制系统、电能监测设备、配套软件授权。技术要求：支持楼宇自控系统集成，提供Modbus/BACnet通讯协议。',2,425000,'2024-03-10 08:00:00.000'),(3,'2024-03-01 22:45:00.000','2024-03-02 19:15:00.000','恒大楼盘-联塑管材采购合同','广州恒大新楼盘水电安装工程管材采购：PVC-U给水管、PPR热水管及配件。质量标准：符合GB/T10002.1-2006标准，提供产品合格证及检测报告。',3,268000,'2024-03-20 08:00:00.000'),(4,'2024-03-16 00:20:00.000','2024-03-16 18:30:00.000','威尼斯人泳池-海龙水处理设备订单','澳门威尼斯人度假村泳池系统升级所需设备：砂缸过滤器、臭氧发生器、水质监测系统。特殊要求：设备需通过澳门特区相关认证，提供24小时技术支持。',4,380000,'2024-04-05 08:00:00.000'),(5,'2024-03-20 19:00:00.000','2024-03-21 23:45:00.000','东莞工业园-格兰富水泵设备询价单','东莞工业园区基础设施配套建设项目水泵设备询价：立式多级离心泵12台，潜水泵15台。请提供产品规格、技术参数及报价明细。',5,NULL,NULL),(6,'2024-03-25 21:30:00.000','2024-03-26 17:20:00.000','恒大楼盘-正泰低压电器采购订单','广州恒大新楼盘项目配电系统所需：小型断路器200个，交流接触器150个，配电箱及其他低压电器元件。要求提供3C认证证书。',6,156000,'2024-04-10 08:00:00.000');
/*!40000 ALTER TABLE `supplier_orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `suppliers`
--

LOCK TABLES `suppliers` WRITE;
/*!40000 ALTER TABLE `suppliers` DISABLE KEYS */;
INSERT INTO `suppliers` VALUES (1,'2024-01-10 17:00:00.000','2024-02-01 17:00:00.000','美的中央空调设备有限公司','张工','13800138000','zhanggong@midea.com','广东省佛山市顺德区美的大道6号','主营大型中央空调设备及配件'),(2,'2024-01-15 18:00:00.000','2024-02-10 18:30:00.000','施耐德电气（中国）有限公司','李经理','13900139000','li.jingli@schneider.com','北京市朝阳区东三环中路20号','智能楼宇与电气自动化系统供应商'),(3,'2024-01-20 22:00:00.000','2024-02-20 19:45:00.000','联塑管道科技有限公司','王主管','13700137000','wang@lesso.com','广东省佛山市顺德区龙江镇325国道旁','提供全系列塑料管材管件'),(4,'2024-02-01 20:00:00.000','2024-03-01 21:00:00.000','海龙泳池设备工程有限公司','刘工','13600136000','liugong@hailongpool.com','上海市浦东新区新金桥路188号','泳池水处理系统及安装服务'),(5,'2024-02-05 19:30:00.000','2024-03-05 22:30:00.000','格兰富水泵（上海）有限公司','何经理','13500135000','he.manager@grundfos.cn','上海市青浦区华新镇华隆路299号','水泵及供水系统制造商'),(6,'2024-02-10 23:00:00.000','2024-03-10 23:15:00.000','正泰电器股份有限公司','赵工','13400134000','zhao@chint.com','浙江省温州市乐清市柳市镇','低压电器及配电解决方案供应商');
/*!40000 ALTER TABLE `suppliers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `user_applies`
--

LOCK TABLES `user_applies` WRITE;
/*!40000 ALTER TABLE `user_applies` DISABLE KEYS */;
INSERT INTO `user_applies` VALUES (1,'2024-02-10 17:00:00.000','2024-02-12 22:30:00.000',3,'请假回老家探亲。','申请2024年2月15日至2月18日年假4天，计划回老家探亲。','同意申请，请做好工作交接，注意休假期间保持通讯畅通。',2,1,'leave'),(2,'2024-02-20 18:15:00.000','2024-02-22 00:45:00.000',4,'申请加班','因深圳科技园项目进度紧急，申请2月22日、23日周末加班，完成电气系统调试工作。','同意加班申请，注意安全操作，加班费按公司标准计算。',8,1,'overtime'),(3,'2024-03-01 19:30:00.000','2024-03-03 17:20:00.000',7,'申请出差','申请3月5日至3月12日出差至澳门威尼斯人度假村，进行泳池设备现场勘察和方案制定。','同意出差申请，已安排住宿和交通，请按时完成勘察任务并提交详细报告。',2,1,'business'),(4,'2024-03-08 22:00:00.000','2024-03-08 22:00:00.000',9,'申请互换班次','因个人原因需要调整3月15日的班次，希望与同事互换到3月16日。','',NULL,0,'shift'),(5,'2024-03-13 00:30:00.000','2024-03-14 18:15:00.000',10,'因个人发展申请离职','因个人发展原因，申请于2024年4月30日离职，将做好工作交接。','已收到离职申请，人事部门将安排离职手续，请配合完成工作交接。',5,1,'resign'),(6,'2024-03-18 21:45:00.000','2024-03-18 21:45:00.000',6,'申请参加研修班','申请参加2024年4月在广州举办的工程财务管理研修班，为期3天。','',NULL,0,'other');
/*!40000 ALTER TABLE `user_applies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `user_apply_types`
--

LOCK TABLES `user_apply_types` WRITE;
/*!40000 ALTER TABLE `user_apply_types` DISABLE KEYS */;
INSERT INTO `user_apply_types` VALUES ('business','出差申请','success'),('leave','休假申请','success'),('other','其他申请','info'),('overtime','加班申请','warning'),('resign','离职申请','danger'),('shift','调班申请','primary');
/*!40000 ALTER TABLE `user_apply_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `user_courses`
--

LOCK TABLES `user_courses` WRITE;
/*!40000 ALTER TABLE `user_courses` DISABLE KEYS */;
INSERT INTO `user_courses` VALUES (2,4),(2,6),(3,1),(3,4),(4,2),(4,5),(5,4),(7,1),(7,3),(8,6),(9,2),(9,4),(10,5);
/*!40000 ALTER TABLE `user_courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `user_projects`
--

LOCK TABLES `user_projects` WRITE;
/*!40000 ALTER TABLE `user_projects` DISABLE KEYS */;
INSERT INTO `user_projects` VALUES (2,1,120.50,'项目整体规划与进度控制，协调各方资源，确保项目按时完成'),(2,3,85.00,'项目统筹管理，施工进度控制，质量安全监督'),(2,5,45.50,'项目前期规划，与园区管理方对接，制定施工方案'),(3,1,180.25,'负责空调系统技术方案设计，设备选型和现场技术指导'),(3,4,98.75,'水处理系统技术方案制定，设备选型和系统优化设计'),(4,2,156.50,'智能化系统架构设计，编程调试和系统集成工作'),(4,3,167.50,'水电系统图纸设计，技术交底，施工技术支持'),(7,1,95.75,'执行空调设备检修维护，故障诊断和零部件更换工作'),(7,4,145.00,'泳池循环系统拆装，过滤设备更换，管道维修改造'),(8,2,98.00,'项目管理与客户沟通，制定实施计划，监督项目质量'),(8,4,72.25,'项目协调管理，与度假村对接，确保不影响正常运营'),(9,3,210.75,'水电管线安装施工，设备安装定位，系统压力测试'),(9,5,89.25,'管网铺设安装，阀门设备安装，管道连接焊接工作'),(10,2,142.25,'电气线路安装布线，配电箱安装和电气设备调试'),(10,5,76.00,'变电站电气安装，高低压配电设备安装调试');
/*!40000 ALTER TABLE `user_projects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `user_roles`
--

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
INSERT INTO `user_roles` VALUES (1,'admin'),(2,'engineer'),(3,'engineer'),(4,'engineer'),(7,'engineer'),(8,'engineer'),(9,'engineer'),(10,'engineer'),(6,'finance'),(5,'hr'),(2,'manager'),(8,'manager');
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'2025-08-24 22:00:35.830','2025-08-24 22:00:35.830','https://axogc.net:520/file/images/dingzhen.png','$2a$10$.CApx0STgJnT4v0.B7n.HeoTPXEiGlddxGg815EVTyPo1tXQiF8tS','丁真','dingzhen@163.com'),(2,'2024-01-08 18:30:00.000','2024-03-15 22:20:00.000','https://axogc.net:520/file/images/jerry.png','$2a$10$N9qo8uLOickgx2ZMRZoMye.Uo8VtO3UQG8.vPQfh/sX8vOVy3rQ6.','李项目','li.xiangmu@company.com'),(3,'2024-01-10 16:15:00.000','2024-02-29 00:45:00.000','https://axogc.net:520/file/images/mr-crab.png','$2a$10$N9qo8uLOickgx2ZMRZoMye.Uo8VtO3UQG8.vPQfh/sX8vOVy3rQ6.','王工程','wang.gongcheng@company.com'),(4,'2024-01-12 17:45:00.000','2024-01-12 17:45:00.000','https://axogc.net:520/file/images/patrick-star.png','$2a$10$N9qo8uLOickgx2ZMRZoMye.Uo8VtO3UQG8.vPQfh/sX8vOVy3rQ6.','陈技术','chen.jishu@company.com'),(5,'2024-01-15 19:00:00.000','2024-01-15 19:00:00.000','https://axogc.net:520/file/images/plankton.png','$2a$10$N9qo8uLOickgx2ZMRZoMye.Uo8VtO3UQG8.vPQfh/sX8vOVy3rQ6.','刘人事','liu.renshi@company.com'),(6,'2024-01-18 21:30:00.000','2024-03-10 18:15:00.000','https://axogc.net:520/file/images/gg-bond.png','$2a$10$N9qo8uLOickgx2ZMRZoMye.Uo8VtO3UQG8.vPQfh/sX8vOVy3rQ6.','赵财务','zhao.caiwu@company.com'),(7,'2024-02-01 22:00:00.000','2024-02-01 22:00:00.000','https://axogc.net:520/file/images/spike.png','$2a$10$7iljwLDh026iZnJAs1zNVe6zXvk3W9ncd8mUm3IXjZlvaotBbjkh6','孙维修','admin@admin.com'),(8,'2024-02-05 17:20:00.000','2024-02-05 17:20:00.000','https://axogc.net:520/file/images/sponge-bob.png','$2a$10$N9qo8uLOickgx2ZMRZoMye.Uo8VtO3UQG8.vPQfh/sX8vOVy3rQ6.','周管理','zhou.guanli@company.com'),(9,'2024-02-10 23:45:00.000','2024-02-10 23:45:00.000','https://axogc.net:520/file/images/squidward-tentacles.png','$2a$10$N9qo8uLOickgx2ZMRZoMye.Uo8VtO3UQG8.vPQfh/sX8vOVy3rQ6.','吴安装','wu.anzhuang@company.com'),(10,'2024-03-01 16:30:00.000','2024-03-01 16:30:00.000','https://axogc.net:520/file/images/tomcat.png','$2a$10$N9qo8uLOickgx2ZMRZoMye.Uo8VtO3UQG8.vPQfh/sX8vOVy3rQ6.','郑电工','zheng.diangong@company.com');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-08-28  6:48:05

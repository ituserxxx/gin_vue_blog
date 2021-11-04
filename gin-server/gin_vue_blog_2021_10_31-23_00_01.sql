-- MySQL dump 10.13  Distrib 8.0.26, for Linux (x86_64)
--
-- Host: localhost    Database: tcsd
-- ------------------------------------------------------
-- Server version	8.0.26

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
-- Table structure for table `tc_post`
--

DROP TABLE IF EXISTS `tc_post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tc_post` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `create_time` datetime NOT NULL,
  `view_sum` int DEFAULT '0' COMMENT '查看次数',
  `is_anonymous` tinyint DEFAULT '1' COMMENT '默认：1-开放，2-匿名',
  `love_sum` tinyint DEFAULT '0' COMMENT '点赞次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tc_post`
--

LOCK TABLES `tc_post` WRITE;
/*!40000 ALTER TABLE `tc_post` DISABLE KEYS */;
INSERT INTO `tc_post` VALUES (1,2,'情浓时，海誓山盟;情淡时，形同陌路','2021-10-25 17:19:18',0,1,0),(2,2,'如果一个人影响到了你的情绪，你的焦点应该放在控制自己的情绪上，而不是影响你情绪的人身上。只有这样，才能真正自信起来。','2021-10-25 17:20:49',0,1,0),(3,2,'有的节日都不是为了礼物，而是提醒大家别忘了爱与被爱','2021-10-25 17:20:51',0,1,0),(4,2,'这城市风很大，孤独的人总是晚回家','2021-10-25 17:20:53',0,1,0),(5,2,'那些不经意想起的，总是记忆里最深刻的','2021-10-25 17:20:54',0,1,0),(6,2,'带着面具时你说我虚伪，摘下面具你问我是谁。','2021-10-25 17:20:57',0,1,0),(7,2,'我们都生活在阴沟里，但仍有人仰望星空。','2021-10-25 17:20:59',0,1,0),(8,2,'人的一生注定会遇到两个人，一个惊艳了时光，一个温柔了岁月。','2021-10-25 17:21:03',0,1,0),(9,2,'愿我如星君如月，夜夜流光相皎洁','2021-10-25 17:21:05',0,1,0),(10,2,'喜欢你的眼，里面有我的脸。','2021-10-25 17:21:10',0,1,0),(11,2,'等你跟自己和解了，对的人自然就出现了。你连跟自己都相处不下去，世上便无对的人。','2021-10-25 17:21:14',0,1,0),(12,2,'友情是积累的，爱情是突然的','2021-10-25 17:21:17',0,1,0),(13,2,'就这样与你告别，借万里青山，以之为隔，世世不见。　','2021-10-25 17:21:20',0,1,0),(14,2,'别跟三观不一致的人争执，别向不关心你的人诉苦，别对不喜欢自己的人讨好，真的没有必要。','2021-10-25 17:23:59',0,1,0),(15,2,'因为没有退路，所以义无反顾。','2021-10-25 17:25:18',0,1,0),(16,2,'我们学会了告别，却低估了思念','2021-10-25 17:27:14',0,1,0),(17,2,'别急，会有始料不及的运气，会有突如其来的欢喜。','2021-10-25 17:27:37',0,1,0),(18,2,'有时候，不善于表达的人挺吃亏的，付出多，用情深，却没人说你好。','2021-10-25 17:28:02',0,1,0),(19,2,'得不到的总是挂念，共朝夕的总是厌倦　','2021-10-25 17:28:48',0,1,0),(20,2,'本就是匆匆过客，又何必耿耿于怀。　','2021-10-25 17:28:59',0,1,0),(21,2,'有很多事情你当时想不通，别着急，过一段时间你再想，就想不起来了。','2021-10-25 17:32:35',0,1,0);
/*!40000 ALTER TABLE `tc_post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tc_post_love`
--

DROP TABLE IF EXISTS `tc_post_love`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tc_post_love` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_id` int NOT NULL,
  `user_id` int NOT NULL,
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='用户点赞表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tc_post_love`
--

LOCK TABLES `tc_post_love` WRITE;
/*!40000 ALTER TABLE `tc_post_love` DISABLE KEYS */;
/*!40000 ALTER TABLE `tc_post_love` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tc_user`
--

DROP TABLE IF EXISTS `tc_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tc_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `openid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `nickname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `avatar_url` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `create_time` datetime NOT NULL,
  `session_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `last_login_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_openid` (`openid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tc_user`
--

LOCK TABLES `tc_user` WRITE;
/*!40000 ALTER TABLE `tc_user` DISABLE KEYS */;
INSERT INTO `tc_user` VALUES (1,'RIFYDYDTKKLC',NULL,NULL,NULL,'2021-10-25 16:27:46',NULL,NULL),(2,'ooSyt5IN9ruOUKkMqMDG-L8EddYk',NULL,NULL,NULL,'2021-10-25 16:43:31',NULL,NULL),(3,'ooSyt5JHV16xP2KuuKXSvL_7Y_94',NULL,NULL,NULL,'2021-10-26 14:38:54',NULL,NULL);
/*!40000 ALTER TABLE `tc_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-10-31 15:00:02

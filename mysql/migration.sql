-- MySQL dump 10.13  Distrib 8.0.25, for macos11.3 (x86_64)
--
-- Host: 127.0.0.1    Database: auctionkuy
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
-- Dumping data for table `app_settings`
--

LOCK TABLES `app_settings` WRITE;
/*!40000 ALTER TABLE `app_settings` DISABLE KEYS */;
INSERT INTO `app_settings` VALUES ('company_bank_account_id','45324314214141213'),('company_bank_account_owner_name','AuctionKuy'),('company_bank_id','c2fc9ac4-9f5a-4b9a-b31c-d45c5ffd76a0'),('escrow_fee','5000');
/*!40000 ALTER TABLE `app_settings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `banks`
--

LOCK TABLES `banks` WRITE;
/*!40000 ALTER TABLE `banks` DISABLE KEYS */;
INSERT INTO `banks` VALUES ('7c55b0de-1c84-4048-ab24-022e22c381fd','BRI','assets/banks/bri','2021-07-25 13:18:20','2021-07-25 13:18:20'),('bbadfcd7-6ae7-4472-ac97-34ef00871213','BNI','assets/banks/bni','2021-07-25 13:18:20','2021-07-25 13:18:20'),('c2fc9ac4-9f5a-4b9a-b31c-d45c5ffd76a0','Mandiri','assets/banks/mandiri','2021-07-25 13:18:20','2021-07-25 13:18:20'),('ca9f6c55-f57c-432e-bc9e-958b0ddd73fe','BCA','assets/banks/bca','2021-07-25 13:18:20','2021-07-25 13:18:20');
/*!40000 ALTER TABLE `banks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `transaction_status`
--

LOCK TABLES `transaction_status` WRITE;
/*!40000 ALTER TABLE `transaction_status` DISABLE KEYS */;
INSERT INTO `transaction_status` VALUES (1,'Waiting for buyer','Waiting for buyer',1,1),(2,'Waitng for buyer to pay','Waiting for payment',2,1),(3,'Buyer has paid','Waiting seller to ship the item',3,2),(4,'Wating buyer to confirm item arrival','Item Shipped',4,3),(5,'Waiting for money to be received','Transaction Done',5,4),(6,'Transaction Done','Transaction Done',6,4);
/*!40000 ALTER TABLE `transaction_status` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('6c286331-662c-4bfa-a29a-62e9c8fb1f35','wildangb2@gmail.com','Wildan Ghiffarie Budhi','+6285850245410','23321421421312321312','account/profile/avatar/6c286331-662c-4bfa-a29a-62e9c8fb1f35',0,'id','ca9f6c55-f57c-432e-bc9e-958b0ddd73fe','f34q35r3w4fq3d3qr','Wildan Ghiffarie Budhi','2021-07-25 13:03:30','2021-07-26 14:36:44');
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

-- Dump completed on 2021-07-26 22:05:37

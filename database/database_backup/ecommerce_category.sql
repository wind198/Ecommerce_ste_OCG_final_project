-- MySQL dump 10.13  Distrib 8.0.26, for Win64 (x86_64)
--
-- Host: localhost    Database: ecommerce
-- ------------------------------------------------------
-- Server version	8.0.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `categoryID` int NOT NULL AUTO_INCREMENT,
  `categoryName` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`categoryID`),
  KEY `categoryName` (`categoryName`)
) ENGINE=InnoDB AUTO_INCREMENT=9729 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (9599,'\n            Beds\n          '),(9609,'\n            Benches\n          '),(9592,'\n            Butler\'s Trays\n          '),(9607,'\n            Chest of Drawers and Commodes\n          '),(9591,'\n            Children\'s / Kids Rugs\n          '),(9600,'\n            Coffee Tables\n          '),(9597,'\n            Desk  & Office Chairs\n          '),(9605,'\n            Dining Chairs - Indoor\n          '),(9601,'\n            Dining Tables\n          '),(9598,'\n            Drinks Trolleys / Bar Cart\n          '),(9594,'\n            Floor Mirrors\n          '),(9603,'\n            Low Stools\n          '),(9602,'\n            Nest of Tables\n          '),(9608,'\n            Occasional Tables\n          '),(9606,'\n            Ottomans, Poufs & Footstools\n          '),(9595,'\n            Outdoor Rugs\n          '),(9596,'\n            Room Dividers\n          '),(9610,'\n            Round Mirrors\n          '),(9611,'\n            Shelving Units and Libraries\n          '),(9593,'\n            Trunks\n          '),(9590,'\n            Wall & Table Clocks\n          '),(9604,'\n            Wall Mirrors\n          '),(9693,'Adorabella Cushions'),(9690,'African Furniture & Décor'),(9719,'Albi'),(9633,'Armoires and Cupboards'),(9645,'Artificial Flowers and Trees'),(9692,'Asian Furniture & Oriental Décor'),(9677,'Bandhini'),(9625,'Bar Stools & Kitchen Stools'),(9635,'Bar Tables'),(9639,'Bed Heads'),(9672,'Beds'),(9726,'Bedside Table'),(9644,'Benches'),(9727,'Bisque Interiors'),(9681,'Bone Inlay Furniture'),(9622,'Butler\'s Trays'),(9621,'Cabinets and Sideboards'),(9670,'CAFE LIGHTING'),(9679,'Calibre Furniture'),(9704,'Canvas and Sasson'),(9671,'Chaise Lounge Daybeds'),(9638,'Chest of Drawers and Commodes'),(9700,'Children\'s / Kids Rugs'),(9629,'Coffee Tables'),(9627,'Console Tables'),(9673,'Contemporary Furniture'),(9705,'Cooper Black'),(9663,'Cotton - Jute Rugs'),(9668,'Craft Enterprises Home Décor'),(9640,'Cupboards and Armoires (Bedroom Furniture)'),(9712,'Cushions, Throws and Quilts'),(9720,'Darcy and Duke'),(9717,'Dasch Design'),(9636,'Decorator Pieces'),(9699,'Designer Door Mats'),(9723,'Designer Rugs'),(9659,'Desk  & Office Chairs'),(9628,'Desks'),(9626,'Dining Chairs - Indoor'),(9631,'Dining Tables'),(9634,'Drinks Trolleys / Bar Cart'),(9713,'Emac & Lawton'),(9642,'Floor Mirrors'),(9701,'Framed Prints & Collages'),(9721,'French Country / Provincial Furniture'),(9666,'FUTURE CLASSICS FURNITURE'),(9725,'Hamptons Style Furniture & Décor'),(9697,'Hide Rugs'),(9711,'Horgans'),(9686,'Indian Furniture & Décor'),(9698,'Industrial Furniture'),(9675,'Island Resort Furniture'),(9724,'La Grolla'),(9623,'Lounges and Daybeds'),(9662,'Low Stools'),(9703,'Manhattan Apartment'),(9716,'Mediterranean Furniture & Decor'),(9685,'Mirrored Furniture'),(9684,'Moroccan Furniture & Décor'),(9632,'Nest of Tables'),(9624,'Occasional Chairs / Armchairs'),(9630,'Occasional Tables'),(9680,'One World Furniture'),(9691,'Organic Style Furniture'),(9718,'Ottomans, Poufs & Footstools'),(9649,'Outdoor / Garden Benches'),(9650,'Outdoor Artwork'),(9652,'Outdoor Coffee Tables'),(9660,'Outdoor Cushions'),(9653,'Outdoor Dining and Bar Tables'),(9687,'Outdoor Dining Chairs'),(9654,'Outdoor Dining Sets'),(9646,'Outdoor Heaters'),(9678,'Outdoor Living Furniture'),(9657,'Outdoor Occasional Chairs / Armchairs'),(9674,'Outdoor Ottomans'),(9694,'Outdoor Rugs'),(9655,'Outdoor Settings'),(9708,'Outdoor Side Tables'),(9688,'Outdoor Sofas'),(9714,'Outdoor Stools'),(9689,'Paris Salon / Parisian Furniture'),(9682,'Phil Bee Interiors'),(9661,'Photographs'),(9656,'Planters, Pots & Urns'),(9658,'Rattan & Bamboo Furniture'),(9710,'Room Dividers'),(9643,'Round Mirrors'),(9669,'Satara'),(9722,'Scandi Style / Scandinavian Furniture'),(9647,'Sculpture and Statues'),(9651,'Shelving Units and Libraries'),(9696,'Sofas'),(9707,'Sounds Like Home\n\n            \n          Sounds Like Home'),(9667,'Tantra Mirrors'),(9664,'Tapestries'),(9683,'Traditional Furniture & Decor'),(9637,'Trunks'),(9706,'TV Cabinet / TV Units'),(9728,'Uniqwa Furniture'),(9676,'United Interiors Canvas Prints'),(9665,'Unitex Rugs'),(9695,'Wall & Table Clocks'),(9715,'Wall Mirrors'),(9702,'Wall Paintings / Framed Canvas'),(9641,'Wall Panels'),(9648,'Water Features'),(9709,'Wrought Iron Furniture');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-08-30  1:13:52

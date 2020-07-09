/*
SQLyog Ultimate v12.4.3 (64 bit)
MySQL - 5.6.21 : Database - northwind
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`golang_task4` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `golang_task4`;

/*Table structure for table `employees` */

DROP TABLE IF EXISTS `profiles`;

CREATE TABLE `profiles` (
  `ID` varchar(30) NOT NULL,
  `Name` varchar(20) DEFAULT NULL,
  `Age` int(3) DEFAULT NULL,
  `Gender` varchar(20) DEFAULT NULL,
  `Address` varchar(60) DEFAULT NULL,
  `City` varchar(20) NOT NULL,
  `Region` varchar(15) DEFAULT NULL,
  `PostalCode` varchar(10) DEFAULT NULL,
  `Company` varchar(30) DEFAULT NULL,
  `isActive` varchar(10) NOT NULL,
  `Email` varchar(60) DEFAULT NULL,
  `Phone` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `Name` (`Name`),
  KEY `PostalCode` (`PostalCode`)
) ENGINE=MyISAM AUTO_INCREMENT=29 DEFAULT CHARSET=utf8;

/*Table structure for table `colors` */

DROP TABLE IF EXISTS `colors`;

CREATE TABLE `colors` (
  `Color` varchar(20) NOT NULL,
  `Category` varchar(50) DEFAULT NULL,
  `Type` varchar(20) NOT NULL,
  `Hex` varchar(20) NOT NULL,
  `Rgba` varchar(60) NOT NULL,
  `Height` int(11) DEFAULT NULL,
  `URL` varchar(30) DEFAULT NULL,
  `Width` int(11) DEFAULT NULL,
  PRIMARY KEY (`Color`),
  KEY `Category` (`Category`),
  KEY `Type` (`Type`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

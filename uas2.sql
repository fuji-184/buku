-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Dec 21, 2023 at 01:14 AM
-- Server version: 11.1.2-MariaDB
-- PHP Version: 8.2.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `uas2`
--

-- --------------------------------------------------------

--
-- Table structure for table `Buku`
--

CREATE TABLE `Buku` (
  `Id` int(11) NOT NULL,
  `Nama` varchar(300) NOT NULL,
  `Gambar` varchar(300) NOT NULL,
  `Link` varchar(300) DEFAULT NULL,
  `Tahun` varchar(150) NOT NULL,
  `Kategori` int(150) NOT NULL,
  `Penulis` int(11) NOT NULL,
  `Harga` double NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `Buku`
--

INSERT INTO `Buku` (`Id`, `Nama`, `Gambar`, `Link`, `Tahun`, `Kategori`, `Penulis`, `Harga`) VALUES
(1, 'Pemrograman Otodidak JavaScript', 'js.jpg', NULL, '2017', 5, 3, 30000),
(2, 'Pemrograman Otodidak Java', 'js.jpg', NULL, '2023', 6, 2, 50000),
(3, 'Pemrograman Otodidak C++', 'js.jpg', NULL, '2023', 7, 4, 100000);

-- --------------------------------------------------------

--
-- Table structure for table `Kategori`
--

CREATE TABLE `Kategori` (
  `Id` int(11) NOT NULL,
  `Nama` varchar(150) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `Kategori`
--

INSERT INTO `Kategori` (`Id`, `Nama`) VALUES
(1, 'Matematika'),
(2, 'Fisika'),
(3, 'Kimia'),
(4, 'Biologi'),
(5, 'JavaScript'),
(6, 'Java'),
(7, 'C++');

-- --------------------------------------------------------

--
-- Table structure for table `Penulis`
--

CREATE TABLE `Penulis` (
  `Id` int(11) NOT NULL,
  `Nama` varchar(150) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `Penulis`
--

INSERT INTO `Penulis` (`Id`, `Nama`) VALUES
(1, 'Budi'),
(2, 'Doni'),
(3, 'Alya'),
(4, 'Rena');

-- --------------------------------------------------------

--
-- Table structure for table `Rating`
--

CREATE TABLE `Rating` (
  `Id` int(11) NOT NULL,
  `Id_Buku` int(11) NOT NULL,
  `Rating` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `Rating`
--

INSERT INTO `Rating` (`Id`, `Id_Buku`, `Rating`) VALUES
(1, 3, 4.5),
(2, 1, 4),
(3, 3, 4),
(9, 1, 3),
(10, 2, 2);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Buku`
--
ALTER TABLE `Buku`
  ADD PRIMARY KEY (`Id`),
  ADD UNIQUE KEY `Id` (`Id`),
  ADD KEY `Kategori` (`Kategori`),
  ADD KEY `Penulis` (`Penulis`);

--
-- Indexes for table `Kategori`
--
ALTER TABLE `Kategori`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `Penulis`
--
ALTER TABLE `Penulis`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `Rating`
--
ALTER TABLE `Rating`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `Buku` (`Id_Buku`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `Buku`
--
ALTER TABLE `Buku`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `Kategori`
--
ALTER TABLE `Kategori`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `Penulis`
--
ALTER TABLE `Penulis`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `Rating`
--
ALTER TABLE `Rating`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `Buku`
--
ALTER TABLE `Buku`
  ADD CONSTRAINT `Kategori` FOREIGN KEY (`Kategori`) REFERENCES `Kategori` (`Id`),
  ADD CONSTRAINT `Penulis` FOREIGN KEY (`Penulis`) REFERENCES `Penulis` (`Id`);

--
-- Constraints for table `Rating`
--
ALTER TABLE `Rating`
  ADD CONSTRAINT `Buku` FOREIGN KEY (`Id_Buku`) REFERENCES `Buku` (`Id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

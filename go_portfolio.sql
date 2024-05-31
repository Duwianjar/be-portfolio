-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 30, 2024 at 05:29 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_portfolio`
--

-- --------------------------------------------------------

--
-- Table structure for table `about`
--

CREATE TABLE `about` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `value` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `about`
--

INSERT INTO `about` (`id`, `name`, `value`, `created_at`, `updated_at`) VALUES
(1, 'experience', '2+ Years working', '2024-05-16 13:24:50', '2024-05-16 13:24:50'),
(2, 'clients', 'around 3+ clients', '2024-05-16 06:36:23', '2024-05-16 06:36:23'),
(3, 'project', '12+ Complete', '2024-05-16 06:36:41', '2024-05-16 06:46:45'),
(4, 'description', 'I am a computer science student at Ahmad Dahlan University, who started my studies in 2021 and am currently nearing the completion of my sixth semester. Before becoming a student, I worked at Sinarmas company as a machine operator for more than a year fro', '2024-05-16 06:52:13', '2024-05-16 06:52:13'),
(7, 'facebook', 'https://web.facebook.com/anjar.gawol', '2024-05-30 06:29:59', '2024-05-30 06:29:59'),
(8, 'instagram', 'https://www.instagram.com/', '2024-05-30 06:30:31', '2024-05-30 06:30:31'),
(10, 'tiktok', 'https://www.tiktok.com/@duwiaaw', '2024-05-30 06:34:45', '2024-05-30 06:34:45'),
(11, 'telegram', 'https://t.me/Duwi_Anjar', '2024-05-30 06:35:06', '2024-05-30 06:37:28');

-- --------------------------------------------------------

--
-- Table structure for table `files`
--

CREATE TABLE `files` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `address` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `files`
--

INSERT INTO `files` (`id`, `name`, `address`, `created_at`, `updated_at`) VALUES
(1, 'cv', 'cv_20240514201113.pdf', '2024-05-12 07:44:34', '2024-05-14 06:11:13'),
(2, 'photoprofile', 'pp_20240530203759.png', '2024-05-13 14:46:23', '2024-05-30 06:37:59'),
(3, 'photoabout', 'photoabout_20240530203827.png', '2024-05-16 13:13:55', '2024-05-30 06:38:27');

-- --------------------------------------------------------

--
-- Table structure for table `profile`
--

CREATE TABLE `profile` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `value` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `profile`
--

INSERT INTO `profile` (`id`, `name`, `value`, `created_at`, `updated_at`) VALUES
(14, 'nama', 'Duwi Anjar Ari Wibowo', '2024-05-11 07:13:26', '2024-05-11 07:13:26'),
(21, 'job', 'Web Developer', '2024-05-11 07:48:55', '2024-05-11 07:48:55'),
(24, 'bio', 'Just a human being interested in the world of coding. Let\'s learn together!', '2024-05-14 07:42:21', '2024-05-14 07:42:21'),
(25, 'wa', 'https://wa.me/6285157993801', '2024-05-14 07:43:38', '2024-05-14 07:43:38');

-- --------------------------------------------------------

--
-- Table structure for table `skill`
--

CREATE TABLE `skill` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `level` varchar(100) NOT NULL,
  `category` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `skill`
--

INSERT INTO `skill` (`id`, `name`, `level`, `category`, `created_at`, `updated_at`) VALUES
(1, 'HTML', 'Experienced', 'Frontend', '2024-05-30 07:10:48', '2024-05-30 07:31:50'),
(2, 'PHP', 'Experienced', 'Backend', '2024-05-30 07:12:52', '2024-05-30 07:32:00'),
(3, 'CSS', 'Experienced', 'Frontend', '2024-05-30 07:22:29', '2024-05-30 07:32:10'),
(4, 'JavaScript', 'Intermediate', 'Frontend', '2024-05-30 07:32:43', '2024-05-30 07:32:43'),
(5, 'Bootstrap', 'Experienced', 'Frontend', '2024-05-30 07:33:07', '2024-05-30 07:33:07'),
(6, 'Tailwind', 'Intermediate', 'Frontend', '2024-05-30 07:33:20', '2024-05-30 07:33:20'),
(7, 'React JS', 'Base', 'Frontend', '2024-05-30 07:33:39', '2024-05-30 07:33:39'),
(9, 'MySQL', 'Experienced', 'Backend', '2024-05-30 08:10:59', '2024-05-30 08:10:59'),
(10, 'PostgreSQL', 'Intermediate', 'Backend', '2024-05-30 08:11:21', '2024-05-30 08:11:21'),
(11, 'Python', 'Base', 'Backend', '2024-05-30 08:11:38', '2024-05-30 08:11:38'),
(12, 'C++', 'Intermediate', 'Backend', '2024-05-30 08:11:51', '2024-05-30 08:11:51'),
(13, 'JAVA', 'Intermediate', 'Backend', '2024-05-30 08:12:06', '2024-05-30 08:12:06'),
(14, 'AJAX', 'Base', 'Frontend', '2024-05-30 08:13:34', '2024-05-30 08:13:34');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `about`
--
ALTER TABLE `about`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `files`
--
ALTER TABLE `files`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `profile`
--
ALTER TABLE `profile`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `skill`
--
ALTER TABLE `skill`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `about`
--
ALTER TABLE `about`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `files`
--
ALTER TABLE `files`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `profile`
--
ALTER TABLE `profile`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- AUTO_INCREMENT for table `skill`
--
ALTER TABLE `skill`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

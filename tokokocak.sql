-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 22, 2021 at 08:07 AM
-- Server version: 10.4.11-MariaDB
-- PHP Version: 7.4.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `tokokocak`
--

-- --------------------------------------------------------

--
-- Table structure for table `tk_categories`
--

CREATE TABLE `tk_categories` (
  `category_id` varchar(100) NOT NULL,
  `category_name` varchar(255) NOT NULL,
  `isactive` enum('1','0') NOT NULL DEFAULT '1',
  `user_id` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tk_categories`
--

INSERT INTO `tk_categories` (`category_id`, `category_name`, `isactive`, `user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
('14391383-5b2a-4e28-92b6-45f9d9de2084', 'Makanan', '1', '9cc6b33a-f867-4b28-8ef1-0c1913f05152', '2021-11-17 14:22:13', '2021-11-17 14:22:13', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `tk_locations`
--

CREATE TABLE `tk_locations` (
  `location_id` varchar(100) NOT NULL,
  `user_id` varchar(100) NOT NULL,
  `location_name` varchar(255) NOT NULL,
  `location_address` text NOT NULL,
  `location_phone` varchar(20) NOT NULL,
  `location_city` varchar(100) NOT NULL,
  `location_province` varchar(100) NOT NULL,
  `location_country` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `tk_modifier_childs`
--

CREATE TABLE `tk_modifier_childs` (
  `modifier_child_id` varchar(100) NOT NULL,
  `modifier_child_name` varchar(255) NOT NULL,
  `modifier_child_price` decimal(15,2) DEFAULT NULL,
  `modifier_child_desc` text DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `tk_modifier_parents`
--

CREATE TABLE `tk_modifier_parents` (
  `modifier_parent_id` varchar(100) NOT NULL,
  `modifier_parent_name` varchar(255) NOT NULL,
  `modifier_parent_type` enum('multiple','justone') NOT NULL DEFAULT 'justone',
  `selected_min` int(5) NOT NULL DEFAULT 0,
  `selected_max` int(5) NOT NULL DEFAULT 1,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `tk_parent_child_modifiers`
--

CREATE TABLE `tk_parent_child_modifiers` (
  `modifier_parent_id` varchar(100) NOT NULL,
  `modifier_child_id` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `tk_payments`
--

CREATE TABLE `tk_payments` (
  `payment_id` varchar(100) NOT NULL,
  `transaction_id` varchar(100) NOT NULL,
  `total_payment` decimal(15,2) NOT NULL,
  `paidchange` decimal(15,2) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `tk_products`
--

CREATE TABLE `tk_products` (
  `product_id` varchar(100) NOT NULL,
  `product_name` varchar(255) NOT NULL,
  `product_desc` text DEFAULT NULL,
  `product_price` decimal(15,2) NOT NULL,
  `product_image` varchar(255) DEFAULT NULL,
  `category_id` varchar(100) NOT NULL,
  `user_id` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tk_products`
--

INSERT INTO `tk_products` (`product_id`, `product_name`, `product_desc`, `product_price`, `product_image`, `category_id`, `user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
('beb0fabf-16d7-474e-a30b-1cdd79f081d4', 'Kopi', 'Ini Kopi', '10000.00', NULL, '72ada7d1-2c55-40bc-a994-64a0d865b195', 'fc6b4960-e472-4c3b-af71-95efa937b402', '2021-07-08 11:18:49', '2021-07-08 11:18:49', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `tk_product_locations`
--

CREATE TABLE `tk_product_locations` (
  `product_id` varchar(100) NOT NULL,
  `location_id` varchar(100) NOT NULL,
  `user_id` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `tk_product_modifiers`
--

CREATE TABLE `tk_product_modifiers` (
  `product_id` varchar(100) NOT NULL,
  `modifier_id` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `tk_transaction_childs`
--

CREATE TABLE `tk_transaction_childs` (
  `transaction_c_id` varchar(100) NOT NULL,
  `transaction_id` varchar(100) NOT NULL,
  `product_id` varchar(100) NOT NULL,
  `stotal` decimal(15,2) NOT NULL,
  `modifier_total_amount` decimal(15,2) DEFAULT NULL,
  `modifier_detail` text DEFAULT NULL,
  `gtotal` decimal(15,2) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `tk_transaction_parents`
--

CREATE TABLE `tk_transaction_parents` (
  `transaction_id` varchar(100) NOT NULL,
  `location_id` varchar(100) NOT NULL,
  `user_id` varchar(100) NOT NULL,
  `gtotal` decimal(15,2) NOT NULL,
  `ispaidoff` enum('0','1') NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `tk_users`
--

CREATE TABLE `tk_users` (
  `user_id` varchar(100) NOT NULL,
  `user_email` varchar(255) NOT NULL,
  `user_password` varchar(255) NOT NULL,
  `user_firstname` varchar(255) NOT NULL,
  `user_lastname` varchar(255) NOT NULL,
  `user_role` enum('owner','cashier') NOT NULL,
  `user_image_profile` varchar(255) DEFAULT NULL,
  `date_verification` datetime DEFAULT NULL,
  `is_active` enum('1','0') NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tk_users`
--

INSERT INTO `tk_users` (`user_id`, `user_email`, `user_password`, `user_firstname`, `user_lastname`, `user_role`, `user_image_profile`, `date_verification`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES
('9cc6b33a-f867-4b28-8ef1-0c1913f05152', 'ramadhansalmanalfarisi8@gmail.com', 'ac7490ff64d4dce0768069c8be203db8', 'Ramadhan', 'Salman Alfarisi', 'owner', '', '2021-11-17 20:58:55', '1', NULL, NULL, NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tk_categories`
--
ALTER TABLE `tk_categories`
  ADD PRIMARY KEY (`category_id`),
  ADD KEY `location_id` (`user_id`);

--
-- Indexes for table `tk_locations`
--
ALTER TABLE `tk_locations`
  ADD PRIMARY KEY (`location_id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `tk_modifier_childs`
--
ALTER TABLE `tk_modifier_childs`
  ADD PRIMARY KEY (`modifier_child_id`);

--
-- Indexes for table `tk_modifier_parents`
--
ALTER TABLE `tk_modifier_parents`
  ADD PRIMARY KEY (`modifier_parent_id`);

--
-- Indexes for table `tk_parent_child_modifiers`
--
ALTER TABLE `tk_parent_child_modifiers`
  ADD KEY `modifier_parent_id` (`modifier_parent_id`),
  ADD KEY `modifier_child_id` (`modifier_child_id`);

--
-- Indexes for table `tk_payments`
--
ALTER TABLE `tk_payments`
  ADD PRIMARY KEY (`payment_id`),
  ADD KEY `transaction_id` (`transaction_id`);

--
-- Indexes for table `tk_products`
--
ALTER TABLE `tk_products`
  ADD PRIMARY KEY (`product_id`),
  ADD KEY `category_id` (`category_id`),
  ADD KEY `location_id` (`user_id`);

--
-- Indexes for table `tk_product_locations`
--
ALTER TABLE `tk_product_locations`
  ADD KEY `product_id` (`product_id`),
  ADD KEY `location_id` (`location_id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `tk_product_modifiers`
--
ALTER TABLE `tk_product_modifiers`
  ADD KEY `product_id` (`product_id`),
  ADD KEY `modifier_id` (`modifier_id`);

--
-- Indexes for table `tk_transaction_childs`
--
ALTER TABLE `tk_transaction_childs`
  ADD PRIMARY KEY (`transaction_c_id`),
  ADD KEY `transaction_id` (`transaction_id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `tk_transaction_parents`
--
ALTER TABLE `tk_transaction_parents`
  ADD PRIMARY KEY (`transaction_id`),
  ADD KEY `location_id` (`location_id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `tk_users`
--
ALTER TABLE `tk_users`
  ADD PRIMARY KEY (`user_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

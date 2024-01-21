-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Hôte : database
-- Généré le : dim. 21 jan. 2024 à 19:48
-- Version du serveur : 10.8.6-MariaDB-1:10.8.6+maria~ubu2204
-- Version de PHP : 8.0.24

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `data`
--

-- --------------------------------------------------------

--
-- Structure de la table `Allergens`
--

CREATE TABLE `Allergens` (
                             `id` int(11) NOT NULL,
                             `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Orders`
--

CREATE TABLE `Orders` (
                          `id` int(11) NOT NULL,
                          `user_id` int(11) NOT NULL,
                          `restaurant_id` int(11) NOT NULL,
                          `status` varchar(255) NOT NULL,
                          `total_price` float NOT NULL,
                          `number` varchar(255) NOT NULL,
                          `created_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                          `closed_date` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Order_has_products`
--

CREATE TABLE `Order_has_products` (
                                      `order_id` int(11) NOT NULL,
                                      `product_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Products`
--

CREATE TABLE `Products` (
                            `id` int(11) NOT NULL,
                            `product_type_id` int(11) NOT NULL,
                            `restaurant_id` int(11) NOT NULL,
                            `name` varchar(255) NOT NULL,
                            `price` float NOT NULL,
                            `image` varchar(255) NOT NULL,
                            `description` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Product_has_allergens`
--

CREATE TABLE `Product_has_allergens` (
                                         `product_id` int(11) NOT NULL,
                                         `allergen_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Product_type`
--

CREATE TABLE `Product_type` (
                                `id` int(11) NOT NULL,
                                `name` varchar(255) NOT NULL,
                                `icon` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
... (247 lignes restantes)

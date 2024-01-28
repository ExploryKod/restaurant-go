-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- HÃ´te : database
-- GÃ©nÃ©rÃ© le : lun. 22 jan. 2024 Ã  10:44
-- Version du serveur : 10.8.8-MariaDB-1:10.8.8+maria~ubu2204
-- Version de PHP : 8.2.15

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de donnÃ©es : `restaurantbdd`
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

-- --------------------------------------------------------

--
-- Structure de la table `Restaurants`
--

CREATE TABLE `Restaurants` (
                               `id` int(11) NOT NULL,
                               `name` varchar(255) NOT NULL,
                               `logo` varchar(255) NOT NULL,
                               `image` varchar(255) NOT NULL,
                               `phone` varchar(255) NOT NULL,
                               `mail` varchar(255) NOT NULL,
                               `is_open` tinyint(1) NOT NULL,
                               `opening_time` time NOT NULL,
                               `closing_time` time NOT NULL,
                               `grade` int(11) NOT NULL,
                               `is_validated` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Restaurant_has_tags`
--

CREATE TABLE `Restaurant_has_tags` (
                                       `restaurant_id` int(11) NOT NULL,
                                       `tag_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Restaurant_has_users`
--

CREATE TABLE `Restaurant_has_users` (
                                        `restaurant_id` int(11) NOT NULL,
                                        `user_id` int(11) NOT NULL,
                                        `is_admin` tinyint(1) NOT NULL,
                                        `role` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Tags`
--

CREATE TABLE `Tags` (
                        `id` int(11) NOT NULL,
                        `name` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Users`
--

CREATE TABLE `Users` (
                         `id` int(11) NOT NULL,
                         `username` varchar(250) NOT NULL,
                         `password` varchar(250) NOT NULL,
                         `name` varchar(255) DEFAULT NULL,
                         `firstname` varchar(255) DEFAULT NULL,
                         `mail` varchar(255) DEFAULT NULL,
                         `phone` varchar(255) DEFAULT NULL,
                         `is_superadmin` tinyint(1) NOT NULL DEFAULT 0,
                         `birthday` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Index pour les tables dÃ©chargÃ©es
--

--
-- Index pour la table `Allergens`
--
ALTER TABLE `Allergens`
    ADD PRIMARY KEY (`id`);

--
-- Index pour la table `Orders`
--
ALTER TABLE `Orders`
    ADD PRIMARY KEY (`id`),
  ADD KEY `order_user_id` (`user_id`),
  ADD KEY `order_restaurant_id` (`restaurant_id`);

--
-- Index pour la table `Order_has_products`
--
ALTER TABLE `Order_has_products`
    ADD KEY `orders_order_id` (`order_id`),
  ADD KEY `orders_product_id` (`product_id`);

--
-- Index pour la table `Products`
--
ALTER TABLE `Products`
    ADD PRIMARY KEY (`id`),
  ADD KEY `product_type` (`product_type_id`),
  ADD KEY `product_restaurant_id` (`restaurant_id`);

--
-- Index pour la table `Product_has_allergens`
--
ALTER TABLE `Product_has_allergens`
    ADD KEY `allergens_product_id` (`product_id`),
  ADD KEY `allergens_allergen_id` (`allergen_id`);

--
-- Index pour la table `Product_type`
--
ALTER TABLE `Product_type`
    ADD PRIMARY KEY (`id`);

--
-- Index pour la table `Restaurants`
--
ALTER TABLE `Restaurants`
    ADD PRIMARY KEY (`id`);

--
-- Index pour la table `Restaurant_has_tags`
--
ALTER TABLE `Restaurant_has_tags`
    ADD KEY `tags_restaurant_id` (`restaurant_id`),
  ADD KEY `tags_tag_id` (`tag_id`);

--
-- Index pour la table `Restaurant_has_users`
--
ALTER TABLE `Restaurant_has_users`
    ADD KEY `restaurant_id` (`restaurant_id`),
  ADD KEY `user_id` (`user_id`);

--
-- Index pour la table `Tags`
--
ALTER TABLE `Tags`
    ADD PRIMARY KEY (`id`);

--
-- Index pour la table `Users`
--
ALTER TABLE `Users`
    ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT pour les tables dÃ©chargÃ©es
--

--
-- AUTO_INCREMENT pour la table `Allergens`
--
ALTER TABLE `Allergens`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Orders`
--
ALTER TABLE `Orders`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Products`
--
ALTER TABLE `Products`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Product_type`
--
ALTER TABLE `Product_type`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Restaurants`
--
ALTER TABLE `Restaurants`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Tags`
--
ALTER TABLE `Tags`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Users`
--
ALTER TABLE `Users`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Contraintes pour les tables dÃ©chargÃ©es
--

--
-- Contraintes pour la table `Orders`
--
ALTER TABLE `Orders`
    ADD CONSTRAINT `order_restaurant_id` FOREIGN KEY (`restaurant_id`) REFERENCES `Restaurants` (`id`),
  ADD CONSTRAINT `order_user_id` FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

--
-- Contraintes pour la table `Order_has_products`
--
ALTER TABLE `Order_has_products`
    ADD CONSTRAINT `orders_order_id` FOREIGN KEY (`order_id`) REFERENCES `Orders` (`id`),
  ADD CONSTRAINT `orders_product_id` FOREIGN KEY (`product_id`) REFERENCES `Products` (`id`);

--
-- Contraintes pour la table `Products`
--
ALTER TABLE `Products`
    ADD CONSTRAINT `product_restaurant_id` FOREIGN KEY (`restaurant_id`) REFERENCES `Restaurants` (`id`),
  ADD CONSTRAINT `product_type` FOREIGN KEY (`product_type_id`) REFERENCES `Product_type` (`id`);

--
-- Contraintes pour la table `Product_has_allergens`
--
ALTER TABLE `Product_has_allergens`
    ADD CONSTRAINT `allergens_allergen_id` FOREIGN KEY (`allergen_id`) REFERENCES `Allergens` (`id`),
  ADD CONSTRAINT `allergens_product_id` FOREIGN KEY (`product_id`) REFERENCES `Products` (`id`);

--
-- Contraintes pour la table `Restaurant_has_tags`
--
ALTER TABLE `Restaurant_has_tags`
    ADD CONSTRAINT `tags_restaurant_id` FOREIGN KEY (`restaurant_id`) REFERENCES `Restaurants` (`id`),
  ADD CONSTRAINT `tags_tag_id` FOREIGN KEY (`tag_id`) REFERENCES `Tags` (`id`);

--
-- Contraintes pour la table `Restaurant_has_users`
--
ALTER TABLE `Restaurant_has_users`
    ADD CONSTRAINT `restaurant_id` FOREIGN KEY (`restaurant_id`) REFERENCES `Restaurants` (`id`),
  ADD CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);
COMMIT;

INSERT INTO `Users` (
    `id`,
    `username`,
    `password`,
    `name`,
    `firstname`,
    `mail`,
    `phone`,
    `is_superadmin`,
    `birthday`
) VALUES
      (1, 'john', 'password', 'Doe', 'John', 'user1@example.com', '+1234567890', 0, '1990-01-15'),
      (2, 'jane', 'password', 'Smith', 'Jane', 'user2@example.com', '+9876543210', 0, '1985-08-22'),
      (3, 'admin', 'password', 'Amaury', 'LeGrandMaître', 'admin@example.com', '+5555555555', 1, '1980-05-10'),
      (4, 'alice', 'password', 'Johnson', 'Alice', 'user3@example.com', '+1111111111', 0, '1995-03-28'),
      (5, 'bob', 'password', 'Williams', 'Bob', 'user4@example.com', '+9999999999', 0, '1992-11-18'),
      (6, 'administrateur', 'password', 'Nassim', 'LeGrandMaître', 'admin@example.com', '+5555555555', 1, '1980-05-10'),
      (6, 'trésorier', 'password', 'Justin', 'LeGrandMaître', 'admin@example.com', '+5555555555', 1, '1980-05-10'),
      (6, 'cto', 'password', 'Reewas', 'LeGrandMaître', 'admin@example.com', '+5555555555', 1, '1980-05-10');

INSERT INTO `Restaurants` ( `id`, `name`, `logo`, `image`, `phone`, `mail`, `is_open`, `opening_time`, `closing_time`, `grade`, `is_validated` ) VALUES ( 1, 'Pegasus', 'logo/restaurant-1.jpg', 'image/restaurant-1.jpg', '+1234567890', 'pegasus@example.com', 1, '08:00:05', '22:06:00', 4, 1 );
INSERT INTO `Restaurants` ( `id`, `name`, `logo`, `image`, `phone`, `mail`, `is_open`, `opening_time`, `closing_time`, `grade`, `is_validated` ) VALUES ( 2, 'Italica', 'logo/restaurant-2.jpg', 'image/restaurant-2.jpg', '+1274566890', 'italica@example.com', 1, '08:00:00', '22:00:00', 2, 0 );
INSERT INTO `Restaurants` ( `id`, `name`, `logo`, `image`, `phone`, `mail`, `is_open`, `opening_time`, `closing_time`, `grade`, `is_validated` ) VALUES ( 3, 'Greca', 'logo/restaurant-3.jpg', 'image/restaurant-3.jpg', '+1274566890', 'greca@example.com', 1, '08:00:00', '22:00:00', 3, 1 );
INSERT INTO `Restaurants` ( `id`, `name`, `logo`, `image`, `phone`, `mail`, `is_open`, `opening_time`, `closing_time`, `grade`, `is_validated` ) VALUES ( 4, 'Algeria', 'logo/restaurant-4.jpg', 'image/restaurant-4.jpg', '+1274769890', 'algeria@example.com', 1, '08:00:00', '22:00:00', 4, 1 );

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

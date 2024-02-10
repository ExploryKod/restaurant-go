-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Hôte : database
-- Généré le : sam. 10 fév. 2024 à 10:26
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
-- Base de données : `restaurantbdd`
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
  `created_date` timestamp NOT NULL DEFAULT current_timestamp(),
  `closed_date` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `Orders`
--

INSERT INTO `Orders` (`id`, `user_id`, `restaurant_id`, `status`, `total_price`, `number`, `created_date`, `closed_date`) VALUES
(21, 10, 1, 'pending', 70, '1', '2024-02-10 00:32:56', NULL);

-- --------------------------------------------------------

--
-- Structure de la table `Order_has_products`
--

CREATE TABLE `Order_has_products` (
  `order_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `Order_has_products`
--

INSERT INTO `Order_has_products` (`order_id`, `product_id`) VALUES
(21, 3),
(21, 13),
(21, 15),
(21, 4);

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

--
-- Déchargement des données de la table `Products`
--

INSERT INTO `Products` (`id`, `product_type_id`, `restaurant_id`, `name`, `price`, `image`, `description`) VALUES
(1, 1, 1, 'Chocolate Cake', 10, 'https://source.unsplash.com/300x150/?ChocolateCake', 'Delicious chocolate cake'),
(2, 2, 1, 'Coca-Cola', 2, 'https://source.unsplash.com/300x150/?Coca-Cola', 'Coca-Cola 33cl'),
(3, 3, 1, 'Cheese Burger', 15, 'https://source.unsplash.com/300x150/?cheeseburger', 'Cheese Burger with fries'),
(4, 4, 1, 'Pizza Margarita', 20, 'https://source.unsplash.com/300x150/?pizza', 'Pizza Margarita 4 seasons'),
(5, 5, 1, 'Sushi Mix', 30, 'https://source.unsplash.com/300x150/?sushi', 'Sushi Mix 24 pieces'),
(6, 6, 1, 'Salad', 8, 'https://source.unsplash.com/300x150/?salad', 'Salad with vegetables'),
(7, 7, 1, 'Pasta Carbonara', 12, 'https://source.unsplash.com/300x150/?Pasta', 'Pasta Carbonara with beef'),
(8, 8, 1, 'Sandwich', 6, 'https://source.unsplash.com/300x150/?Sandwich', 'Sandwich with chicken and cheese'),
(9, 1, 1, 'Cheese Cake', 12, 'https://source.unsplash.com/300x150/?CheeseCake', 'Delicious cheese cake'),
(10, 1, 1, 'Apple pie', 12, 'https://source.unsplash.com/300x150/?ApplePie', 'Delicious apple pie'),
(11, 2, 1, 'Fanta', 2, 'https://source.unsplash.com/300x150/?Fanta', 'Fanta 33cl'),
(12, 2, 1, 'Sprite', 2, 'https://source.unsplash.com/300x150/?Sprite', 'Sprite 33cl'),
(13, 3, 1, 'Chicken Burger', 15, 'https://source.unsplash.com/300x150/?chickenburger', 'Chicken Burger with fries'),
(14, 3, 1, 'Fish Burger', 15, 'https://source.unsplash.com/300x150/?fish burger', 'Fish Burger with fries'),
(15, 4, 1, 'Pizza 4 seasons', 20, 'https://source.unsplash.com/300x150/?pizza', 'Pizza 4 seasons'),
(16, 4, 1, 'Pizza 4 cheese', 20, 'https://source.unsplash.com/300x150/?pizza', 'Pizza 4 cheese'),
(17, 5, 1, 'Sushi californian roll', 30, 'https://source.unsplash.com/300x150/?sushi', 'Sushi Mix 24 pieces'),
(18, 5, 1, 'Sushi Mix 24 pieces', 30, 'https://source.unsplash.com/300x150/?sushi', 'Sushi Mix 24 pieces'),
(19, 6, 1, 'Salad with vegetables', 8, 'https://source.unsplash.com/300x150/?salad', 'Salad with vegetables'),
(20, 7, 1, 'Pasta with tomato sauce', 12, 'https://source.unsplash.com/300x150/?Pasta', 'Pasta with tomato sauce'),
(21, 7, 1, 'Pasta with cheese', 12, 'https://source.unsplash.com/300x150/?Pasta', 'Pasta with cheese'),
(22, 8, 1, 'Sandwich with fish', 6, 'https://source.unsplash.com/300x150/?Sandwich', 'Sandwich with fish');

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
  `restaurant_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `icon` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `Product_type`
--

INSERT INTO `Product_type` (`id`, `restaurant_id`, `name`, `icon`) VALUES
(1, 1, 'Deserts', 'fas fa-ice-cream'),
(2, 1, 'Drinks', 'fas fa-cocktail'),
(3, 1, 'Burgers', 'fas fa-hamburger'),
(4, 1, 'Pizza', 'fas fa-pizza-slice'),
(5, 1, 'Sushi', 'fas fa-fish'),
(6, 1, 'Salads', 'fas fa-carrot'),
(7, 1, 'Pasta', 'fas fa-pepper-hot'),
(8, 1, 'Sandwiches', 'fas fa-bread-slice');

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

--
-- Déchargement des données de la table `Restaurants`
--

INSERT INTO `Restaurants` (`id`, `name`, `logo`, `image`, `phone`, `mail`, `is_open`, `opening_time`, `closing_time`, `grade`, `is_validated`) VALUES
(1, 'Pegasus', 'logo/restaurant-1.jpg', 'image/restaurant-1.jpg', '+1234567890', 'pegasus@example.com', 1, '08:00:05', '22:06:00', 4, 1),
(2, 'Italica', 'logo/restaurant-2.jpg', 'image/restaurant-2.jpg', '+1274566890', 'italica@example.com', 1, '08:00:00', '22:00:00', 2, 0),
(3, 'Greca', 'logo/restaurant-3.jpg', 'image/restaurant-3.jpg', '+1274566890', 'greca@example.com', 1, '08:00:00', '22:00:00', 3, 1),
(4, 'Algeria', 'logo/restaurant-4.jpg', 'image/restaurant-4.jpg', '+1274769890', 'algeria@example.com', 1, '08:00:00', '22:00:00', 4, 1);

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
  `birthday` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `Users`
--

INSERT INTO `Users` (`id`, `username`, `password`, `name`, `firstname`, `mail`, `phone`, `is_superadmin`, `birthday`) VALUES
(1, 'nassnass', '$2a$14$r6pEamBdmQLCqNbMMsN10uyZn51oycF4HKljz6KxRHkaRM7kFDXRy', 'dev', 'nassim', 'nass@nass.fr', '06796868654', 0, NULL),
(2, 'amaury', '$2a$14$r6pEamBdmQLCqNbMMsN10uyZn51oycF4HKljz6KxRHkaRM7kFDXRy', 'Franss', 'Amaury', 'amau@amau.fr', '0649494949', 0, NULL),
(3, 'justin', '$2a$14$r6pEamBdmQLCqNbMMsN10uyZn51oycF4HKljz6KxRHkaRM7kFDXRy', 'juju', 'justin', 'juju@juju.fr', '0649494949', 0, NULL),
(4, 'reewaz', '$2a$14$f8iRGdLfTXOr3f4vXfDSxePHLmFLkAMD8ouUBd9Bv/UBl7MPQjR.6', 'maskey', 'reewaz', 'ree@ree.fr', '0649494949', 0, NULL),
(10, 'grotest', '$2a$14$F.p.tq.IkqE1jl8O/foDWen8aRrT2kpH9CU7h7oQ4.7HrdTcBo.oq', 'gros', 'test', 'gro@gro.fr', '06499292933', 0, NULL);

--
-- Index pour les tables déchargées
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
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_restaurant_id` (`restaurant_id`);

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
-- AUTO_INCREMENT pour les tables déchargées
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
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- AUTO_INCREMENT pour la table `Products`
--
ALTER TABLE `Products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

--
-- AUTO_INCREMENT pour la table `Product_type`
--
ALTER TABLE `Product_type`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT pour la table `Restaurants`
--
ALTER TABLE `Restaurants`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT pour la table `Tags`
--
ALTER TABLE `Tags`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Users`
--
ALTER TABLE `Users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- Contraintes pour les tables déchargées
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
  ADD CONSTRAINT `orders_order_id` FOREIGN KEY (`order_id`) REFERENCES `Orders` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `orders_product_id` FOREIGN KEY (`product_id`) REFERENCES `Products` (`id`) ON DELETE CASCADE;

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
-- Contraintes pour la table `Product_type`
--
ALTER TABLE `Product_type`
  ADD CONSTRAINT `fk_restaurant_id` FOREIGN KEY (`restaurant_id`) REFERENCES `Restaurants` (`id`) ON DELETE CASCADE;

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

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

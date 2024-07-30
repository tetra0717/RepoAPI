CREATE TABLE `reports` (
    `id` VARCHAR(255) PRIMARY KEY,
    `author_id` VARCHAR(255) NOT NULL,
    `count` INT NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `style` VARCHAR(100) NOT NULL,
    `language` VARCHAR(100) NOT NULL
);

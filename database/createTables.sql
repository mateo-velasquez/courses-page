-- Table for Users
CREATE TABLE IF NOT EXISTS users (
    id_user INT PRIMARY KEY AUTO_INCREMENT,
    create_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    lastupdate_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    dni VARCHAR(9) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    user_password VARCHAR(100) NOT NULL,
    access_level ENUM('User', 'Admin') NOT NULL DEFAULT 'User'
);

-- Table for Images (of courses)
CREATE TABLE IF NOT EXISTS images (
    id_image INT PRIMARY KEY AUTO_INCREMENT,
    image_path VARCHAR(300) NOT NULL UNIQUE
);

-- Table for Categories
CREATE TABLE IF NOT EXISTS categories (
    id_category INT PRIMARY KEY AUTO_INCREMENT,
    category_name VARCHAR(100) NOT NULL UNIQUE
);

-- Table for Courses
CREATE TABLE IF NOT EXISTS courses (
    id_course INT PRIMARY KEY AUTO_INCREMENT,
    id_image INT,
    create_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    lastupdate_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    course_name VARCHAR(300) NOT NULL,
    price DECIMAL(10,2) NOT NULL CHECK (price > 0),
    init_date DATE NOT NULL,
    course_description VARCHAR(1000),
    duration VARCHAR(100) NOT NULL,
    rating DECIMAL(3,2) DEFAULT 0 CHECK (rating >= 0 AND rating <= 5),
    FOREIGN KEY (id_image) REFERENCES images(id_image) ON DELETE CASCADE
);

-- Table for Subscriptions
CREATE TABLE IF NOT EXISTS subscriptions (
    id_subscription INT PRIMARY KEY AUTO_INCREMENT,
    id_course INT NOT NULL,
    id_user INT NOT NULL,
    create_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    individual_rating DECIMAL(3,2) CHECK(individual_rating >= 0 AND individual_rating <= 5),
    comment VARCHAR(300),
    lastupdate_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_course) REFERENCES courses(id_course) ON DELETE CASCADE,
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE,
    UNIQUE (id_course, id_user)
);

-- Table for Files
CREATE TABLE IF NOT EXISTS files (
    id_file INT PRIMARY KEY AUTO_INCREMENT,
    id_user INT NOT NULL,
    id_course INT NOT NULL,
    file_name VARCHAR(300) NOT NULL,
    file_path VARCHAR(255) NOT NULL,
    create_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    upload_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_course) REFERENCES courses(id_course) ON DELETE CASCADE,
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE
);

-- Tabla pasarela entre cursos y categorias
CREATE TABLE IF NOT EXISTS course_categories (
    id_course INT,
    id_category INT,
    FOREIGN KEY (id_course) REFERENCES courses(id_course) ON DELETE CASCADE,
    FOREIGN KEY (id_category) REFERENCES categories(id_category) ON DELETE CASCADE,
    PRIMARY KEY (id_course, id_category)
);

-- Triggers for Update "lastupdate_date"
DELIMITER //

CREATE TRIGGER update_lastupdate_date_subscriptions
BEFORE UPDATE ON subscriptions
FOR EACH ROW
BEGIN
    SET NEW.lastupdate_date = NOW(); 
END //

CREATE TRIGGER update_lastupdate_date_users
BEFORE UPDATE ON users
FOR EACH ROW
BEGIN
    SET NEW.lastupdate_date = NOW(); 
END //

CREATE TRIGGER update_lastupdate_date_course
BEFORE UPDATE ON courses
FOR EACH ROW
BEGIN
    SET NEW.lastupdate_date = NOW(); 
END //

DELIMITER ;

DELIMITER //

-- Trigger to update course rating after INSERT on subscriptions
CREATE TRIGGER update_course_rating_after_insert
AFTER INSERT ON subscriptions
FOR EACH ROW
BEGIN
    DECLARE avg_rating DECIMAL(3,2);

    -- Calculate the average rating for the course
    SELECT AVG(individual_rating) INTO avg_rating
    FROM subscriptions
    WHERE id_course = NEW.id_course;

    -- Update the course rating (0 if no ratings)
    UPDATE courses
    SET rating = IFNULL(avg_rating, 0)
    WHERE id_course = NEW.id_course;
END //

-- Trigger to update course rating after UPDATE on subscriptions
CREATE TRIGGER update_course_rating_after_update
AFTER UPDATE ON subscriptions
FOR EACH ROW
BEGIN
    DECLARE avg_rating DECIMAL(3,2);

    IF NEW.individual_rating <> OLD.individual_rating THEN
        SELECT AVG(individual_rating) INTO avg_rating
        FROM subscriptions
        WHERE id_course = NEW.id_course;
        UPDATE courses
        SET rating = IFNULL(avg_rating, 0)
        WHERE id_course = NEW.id_course;
    END IF;
END //

-- Trigger to update course rating after DELETE on subscriptions
CREATE TRIGGER update_course_rating_after_delete
AFTER DELETE ON subscriptions
FOR EACH ROW
BEGIN
    DECLARE avg_rating DECIMAL(3,2);

    -- Calculate the average rating for the course
    SELECT AVG(individual_rating) INTO avg_rating
    FROM subscriptions
    WHERE id_course = OLD.id_course;

    -- Update the course rating (0 if no ratings)
    UPDATE courses
    SET rating = IFNULL(avg_rating, 0)
    WHERE id_course = OLD.id_course;
END //

DELIMITER ;

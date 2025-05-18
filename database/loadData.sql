-- Cargamos algunos usuarios
INSERT INTO users (name, last_name, dni, email, user_password, access_level) VALUES
('Mateo', 'Veda', '45407199', '2214715@ucc.edu.ar', '$2a$10$rj4xjdzIu/Gcw3JViJDnneDl7t1U6sQi0bYq1hpY/9t25WIMSg53q', 'Admin'), -- Password: admin
('Tobias', 'Nicolás', '45496307', '2215878@ucc.edu.ar', '$2a$10$dOU4sF4pMsgy/knb.yIM9uMnwY98K97ZczVFjYlhBDo4AG1Ezl.SS', 'Admin'), -- Password: pass
('Agustina', 'Perez', '44789625', '2201488@ucc.edu.ar', '$2a$10$dOU4sF4pMsgy/knb.yIM9uMnwY98K97ZczVFjYlhBDo4AG1Ezl.SS', 'User'); -- Password: pass

-- Creamos algunas categorías
INSERT INTO category (category_name) VALUES
('Panadería'), 
('Pastelería'), 
('Carnes'),
('Vegetariano'),
('Comida China'),
('Sushi'),
('Pescado'),
('Pastas'),
('Cómida rápida'),
('Ensaladas'),
('Proteico'),
('Saludable'),
('Desayunos');

-- Insertamos las imagenes cargadas en la carpeta de back
INSERT INTO images (image_path) VALUES
('images/Image-1'),
('images/Image-2'),
('images/Image-3'),
('images/Image-4'),
('images/Image-5'),
('images/Image-6'),
('images/Image-7'),
('images/Image-8'),
('images/Image-9'),
('images/Image-10'),
('images/Image-11'),
('images/Image-12'),
('images/Image-13'),
('images/Image-14'),
('images/Image-15'),
('images/Image-16'),
('images/Image-17'),
('images/Image-18'),
('images/Image-19'),
('images/Image-20'),
('images/Image-21');

-- Inserción de cursos (formato de fecha corregido a YYYY-MM-DD)
INSERT INTO course (id_image, course_name, price, init_date, description, duration) VALUES 
(1, 'Introducción a los cortes finos', 45.99, '2025-01-16', 'Este es un curso que abre puertas para todo público interesado en la cocina básica de cortes de carne', '6 meses'),
(2,'Introducción a los desayunos saludables',36.99,'2025-07-24','Este va a ser un curso que te va a enseñar la importancia de comer saludablemente','4 meses'),
(3,'Introducción a los cortes para Horno',45.99,'2025-06-20','En este curso aprenderas la importancia de las temperaturas según los cortes necesarios a usar en el horno','6 meses'),
(4,'Pasta Vegana Avanzada',72.99,'2025-02-27','Chiques, en este curso aprenderan las mejores técnicas para preparar pastas balanceadas, sabrosas y libres de carne','8 meses'),
(5,'Pasta Italiana Intermedia',36.99,'2025-05-09','Este curso es para aquellos con espíritu de Leones, pero manos de doncellas, que sean capaces de replicar los mejores platillos que tiene Italia','6 meses');

-- Relación curso-categoría
INSERT INTO course_category(id_course, id_category) VALUES
(1,3),
(2,11),
(2,12),
(2,13),
(3,3),
(3,11),
(3,12),
(4,4),
(4,8),
(5,8);

-- Inserción de las Subscripciones
INSERT INTO subscription (id_course, id_user) VALUES
(1,3),
(1,1),
(2,3),
(3,3),
(4,3),
(5,3);

CREATE DATABASE avito;
\c avito;

CREATE TABLE employee (
                          id SERIAL PRIMARY KEY,
                          username VARCHAR(50) UNIQUE NOT NULL,
                          first_name VARCHAR(50),
                          last_name VARCHAR(50),
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE organization_type AS ENUM (
    'IE',
    'LLC',
    'JSC'
);

CREATE TABLE organization (
                              id SERIAL PRIMARY KEY,
                              name VARCHAR(100) NOT NULL,
                              description TEXT,
                              type organization_type,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE organization_responsible (
                                          id SERIAL PRIMARY KEY,
                                          organization_id INT REFERENCES organization(id) ON DELETE CASCADE,
                                          user_id INT REFERENCES employee(id) ON DELETE CASCADE
);


-- Вставка данных в таблицу employee
INSERT INTO employee (username, first_name, last_name)
VALUES
    ('john_doe', 'John', 'Doe'),
    ('jane_smith', 'Jane', 'Smith'),
    ('alex_brown', 'Alex', 'Brown'),
    ('emily_jones', 'Emily', 'Jones'),
    ('michael_white', 'Michael', 'White');

-- Вставка данных в таблицу organization
INSERT INTO organization (name, description, type)
VALUES
    ('Tech Innovations', 'A company focused on technological advancements.', 'LLC'),
    ('Green Solutions', 'An organization dedicated to environmental sustainability.', 'JSC'),
    ('Fast Delivery Services', 'Logistics and delivery company.', 'IE'),
    ('Creative Designs', 'A design agency specializing in branding.', 'LLC'),
    ('Health First', 'A health and wellness organization.', 'JSC');


-- Вставка данных в таблицу organization_responsible
INSERT INTO organization_responsible (organization_id, user_id)
VALUES
    (1, 1),  -- John Doe is responsible for Tech Innovations
    (1, 2),  -- Jane Smith is also responsible for Tech Innovations
    (2, 3),  -- Alex Brown is responsible for Green Solutions
    (3, 4),  -- Emily Jones is responsible for Fast Delivery Services
    (4, 5);  -- Michael White is responsible for Creative Designs
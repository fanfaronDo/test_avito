CREATE DATABASE avito;
\c avito;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE employee (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
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
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    type organization_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE organization_responsible (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    user_id UUID REFERENCES employee(id) ON DELETE CASCADE
);

-- Вставка данных в таблицу employee
INSERT INTO employee (username, first_name, last_name)
VALUES
    ('john_doe', 'John', 'Doe'),
    ('jane_smith', 'Jane', 'Smith'),
    ('alex_brown', 'Alex', 'Brown'),
    ('emily_jones', 'Emily', 'Jones'),
    ('michael_white', 'Michael', 'White'),
    ('jdoe', 'John', 'Doe'),
    ('asmith', 'Alice', 'Smith'),
    ('bjackson', 'Bob', 'Jackson'),
    ('cjohnson', 'Charlie', 'Johnson'),
    ('dlee', 'Diana', 'Lee');

-- Вставка данных в таблицу organization
INSERT INTO organization (name, description, type)
VALUES
    ('Tech Innovations', 'A company focused on technological advancements.', 'LLC'),
    ('Green Solutions', 'An organization dedicated to environmental sustainability.', 'JSC'),
    ('Fast Delivery Services', 'Logistics and delivery company.', 'IE'),
    ('Creative Designs', 'A design agency specializing in branding.', 'LLC'),
    ('Health First', 'A health and wellness organization.', 'JSC'),
    ('Tech Innovations LLC', 'A company focused on technological advancements.', 'LLC'),
    ('Green Energy JSC', 'A joint stock company specializing in renewable energy.', 'JSC'),
    ('Health Solutions IE', 'An individual enterprise providing health services.', 'IE'),
    ('Global Logistics LLC', 'Logistics and supply chain management services.', 'LLC'),
    ('Creative Media JSC', 'A joint stock company in the media and entertainment industry.', 'JSC');;

INSERT INTO organization_responsible (organization_id, user_id) VALUES
((SELECT id FROM organization WHERE name = 'Tech Innovations LLC'), (SELECT id FROM employee WHERE username = 'jdoe')),
((SELECT id FROM organization WHERE name = 'Green Energy JSC'), (SELECT id FROM employee WHERE username = 'asmith')),
((SELECT id FROM organization WHERE name = 'Health Solutions IE'), (SELECT id FROM employee WHERE username = 'bjackson')),
((SELECT id FROM organization WHERE name = 'Global Logistics LLC'), (SELECT id FROM employee WHERE username = 'cjohnson')),
((SELECT id FROM organization WHERE name = 'Creative Media JSC'), (SELECT id FROM employee WHERE username = 'dlee')),
((SELECT id FROM organization WHERE name = 'Tech Innovations'), (SELECT id FROM employee WHERE username = 'john_doe')),
((SELECT id FROM organization WHERE name = 'Green Solutions'), (SELECT id FROM employee WHERE username = 'jane_smith')),
((SELECT id FROM organization WHERE name = 'Fast Delivery Services'), (SELECT id FROM employee WHERE username = 'alex_brown')),
((SELECT id FROM organization WHERE name = 'Creative Designs'), (SELECT id FROM employee WHERE username = 'emily_jones')),
((SELECT id FROM organization WHERE name = 'Health First'), (SELECT id FROM employee WHERE username = 'michael_white'));
-- DML

-- INSERT
INSERT INTO 
    users (email, name, dob)
VALUES 
    ('calman@gmail.com', 'calman', '2010-12-20');
INSERT INTO 
    users (email, name, dob)
VALUES 
    ('calmanlagi@gmail.com', 'calman', '2010-12-20'),
    ('tara@gmail.com', 'tara', '2010-12-20');

INSERT INTO
    user_photos(url, user_id)
VALUES
    ('https://www.google.com', '1'),
    ('https://www.google.com/1', '1'),
    ('https://www.google.com/2', '1'),
    ('https://www.facebook.com/1', '3');

INSERT INTO
    user_photos_no_fk(url, user_id)
VALUES
    ('https://www.google.com', '1'),
    ('https://www.google.com/1', '1'),
    ('https://www.google.com/2', '1'),
    ('https://www.facebook.com/1', '3');

-- UPDATE
UPDATE users
SET 
    name = 'calman lagi'
WHERE
    email = 'calmanlagi@gmail.com';
-- update harus selalu dibarengin dengan where
UPDATE users
SET 
    name = 'calman lagi';
-- ORM -> GORM bisa kalau ga ada where, error

-- DELETE
-- hard delete
DELETE FROM users
WHERE id = 4;

-- soft delete
UPDATE users
SET
    deleted_at = now()
WHERE email = 'tara@gmail.com';

-- SELECT -> bisa menjadi bagian dari DML
SELECT id, dob, name, email FROM users
WHERE deleted_at is null;
select * from user_photos;
select * from user_photos_no_fk;

-- join with FK
-- JOIN
-- left join
-- right join
-- inner join
-- outer join
select * from users u
JOIN user_photos up 
    on up.user_id = u.id 
    and up.deleted_at is null
where u.deleted_at is null
    and u.id = 1; 

select * from users u
JOIN user_photos_no_fk up 
    on up.user_id = u.id 
    and up.deleted_at is null
where u.deleted_at is null
    and u.id = 1; 
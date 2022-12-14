# Write your MySQL query statement below
select user_id,  concat(UPPER(left(name, 1)),LOWER(right(name,CHAR_LENGTH(name)-1))) as name 
FROM users ORDER BY user_id

# Write your MySQL query statement below
SELECT p.product_name, S.year, s.price FROM sales as s

left join  product AS p on p.product_id=s.product_id;
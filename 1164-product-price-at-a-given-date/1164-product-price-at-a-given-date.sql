# Write your MySQL query statement below
SELECT p1.product_id, p1.new_price as price
FROM products p1
WHERE (p1.product_id , p1.change_date) IN (
  SELECT p2.product_id,  MAX(p2.change_date)
  FROM products p2
  where change_date <= "2019-08-16" 
  GROUP BY product_id
)
union ALL
SELECT Distinct p3.product_id, 10 as price
FROM products p3
WHERE p3.product_id not in (
    SELECT p1.product_id
    FROM products p1
    WHERE (p1.product_id , p1.change_date) IN (
      SELECT p2.product_id,  MAX(p2.change_date)
      FROM products p2
      where change_date <= "2019-08-16" 
      GROUP BY product_id
  )
);
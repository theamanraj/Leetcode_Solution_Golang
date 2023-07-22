# Write your MySQL query statement below
WITH cte AS(SELECT customer_id ,
MIN(order_date) first_time_date,
MIN(customer_pref_delivery_date) AS customer_pref_delivery_date
FROM Delivery
GROUP BY customer_id )

SELECT 
ROUND((100.0)*AVG(CASE WHEN first_time_date=customer_pref_delivery_date THEN 1 ELSE 0 END) ,2)immediate_percentage 
FROM cte
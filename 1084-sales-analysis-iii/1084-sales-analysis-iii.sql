# Write your MySQL query statement below
select product_id,product_name from product p1
where (select count(*) from Sales s1 where sale_date between "2019-01-01" and "2019-03-31" and s1.product_id=p1.product_id) = (select count(*) from Sales s3 where s3.product_id=p1.product_id) and product_id in (select product_id from Sales)



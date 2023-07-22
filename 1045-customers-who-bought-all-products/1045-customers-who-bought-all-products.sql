# Write your MySQL query statement below
with Cust as (
select c.customer_id,count(distinct c.product_key) as ct,p.product_key from Customer c left join Product P on p.product_key=c.product_key group by c.customer_id
),
prod as (
  select count(distinct product_key) as cnt from Product
)
select c.customer_id from Cust c, prod p where 
c.ct=cnt

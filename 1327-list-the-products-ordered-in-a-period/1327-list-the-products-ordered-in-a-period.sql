# Write your MySQL query statement below
select product_name,sum(unit)as unit from products p  join orders o
on p.product_id=o.product_id
where extract(month from o.order_date)=2
group by p.product_id
having sum(unit)>=100
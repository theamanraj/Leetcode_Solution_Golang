# Write your MySQL query statement below
select * from (
select product_id, 'store1' store, store1 price from products
union all
select product_id, 'store2' store, store2 price from products
union all
select product_id, 'store3' store, store3 price from products) t
where product_id is not null and price is not null
order by product_id, store;
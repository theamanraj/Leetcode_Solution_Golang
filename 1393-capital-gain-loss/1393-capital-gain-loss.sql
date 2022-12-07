# Write your MySQL query statement below
SELECT
  stock_name
  , sum(case when operation = "Sell" then 1*price else 0 end) 
  - sum(case when operation = "Buy" then 1*price else 0 end) as capital_gain_loss
FROM
  Stocks
GROUP BY
  stock_name
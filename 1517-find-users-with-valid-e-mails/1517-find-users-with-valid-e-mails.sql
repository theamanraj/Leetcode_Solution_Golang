# Write your MySQL query statement below
SELECT * FROM users WHERE 
        mail REGEXP '^[a-zA-Z][a-zA-Z-._0-9]*@leetcode[.]{1}com$';
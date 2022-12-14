# Write your MySQL query statement below
SELECT
    id
    , COALESCE( # we need COALESCE() statement for case when total number of seats is odd
        CASE
            WHEN id MOD 2 = 1 THEN LEAD(student) OVER () # i + 1 --> i
            WHEN id MOD 2 = 0 THEN LAG(student) OVER () # i - 1 --> i
        END,
        student
    ) AS student
FROM
    Seat;
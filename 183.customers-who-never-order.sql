--
-- @lc app=leetcode id=183 lang=mysql
--
-- [183] Customers Who Never Order
--
# Write your MySQL query statement below
SELECT c.Name AS Customers
  FROM Customers c
 WHERE c.Id NOT IN (SELECT o.CustomerId FROM Orders o);

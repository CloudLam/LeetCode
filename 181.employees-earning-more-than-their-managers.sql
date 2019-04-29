--
-- @lc app=leetcode id=181 lang=mysql
--
-- [181] Employees Earning More Than Their Managers
--
# Write your MySQL query statement below
    SELECT a.Name AS Employee 
      FROM Employee a
INNER JOIN Employee b
        ON a.ManagerId = b.Id
       AND a.Salary > b.Salary;

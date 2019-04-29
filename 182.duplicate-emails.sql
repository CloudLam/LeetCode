--
-- @lc app=leetcode id=182 lang=mysql
--
-- [182] Duplicate Emails
--
# Write your MySQL query statement below
    SELECT DISTINCT a.Email AS Email 
      FROM Person a
INNER JOIN Person b
        ON a.Id <> b.Id
       AND a.Email = b.Email;

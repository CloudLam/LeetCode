--
-- @lc app=leetcode id=197 lang=mysql
--
-- [197] Rising Temperature
--

-- @lc code=start
# Write your MySQL query statement below

SELECT a.Id 
  FROM Weather a, Weather b 
 WHERE a.Temperature > b.Temperature
   AND DATEDIFF(a.RecordDate, b.RecordDate) = 1

-- @lc code=end

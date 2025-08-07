-- 为game_user_auth_codes表的role_game_id字段添加唯一索引
-- 执行此SQL前，请确保role_game_id字段没有重复值

-- 检查是否有重复的role_game_id
SELECT role_game_id, COUNT(*) as count
FROM game_user_auth_codes
WHERE role_game_id IS NOT NULL AND role_game_id != ''
GROUP BY role_game_id
HAVING COUNT(*) > 1;

-- 添加唯一索引（如果role_game_id有重复值，此操作会失败）
CREATE UNIQUE INDEX idx_game_user_auth_codes_role_game_id ON game_user_auth_codes(role_game_id);

-- 或者使用ALTER TABLE方式添加唯一约束
-- ALTER TABLE game_user_auth_codes ADD CONSTRAINT uk_role_game_id UNIQUE (role_game_id);

-- 如果需要删除重复数据后再添加唯一索引，可以使用以下SQL：
-- DELETE FROM game_user_auth_codes 
-- WHERE id NOT IN (
--     SELECT MIN(id) 
--     FROM game_user_auth_codes 
--     WHERE role_game_id IS NOT NULL AND role_game_id != ''
--     GROUP BY role_game_id
-- );
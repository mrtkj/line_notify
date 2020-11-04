-- m_users
INSERT INTO m_users (id, name) VALUES (1, 'test user');
INSERT INTO m_users (id, name) VALUES (2, 'test user2');
INSERT INTO m_users (id, name) VALUES (3, 'test user3');

-- t_schedule
INSERT INTO t_schedules (user_id, execute_date, execute_time, task)
    VALUES(1, '20201231', '080000', 'test task');
INSERT INTO t_schedules (user_id, execute_date, execute_time, task)
    VALUES(2, '20201230', '090000', 'test task2');
INSERT INTO t_schedules (user_id, execute_date, execute_time, task)
    VALUES(3, '20201230', '100000', 'test task3');


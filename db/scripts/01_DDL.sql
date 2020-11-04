CREATE TABLE m_users(
	id SERIAL,
	name VARCHAR(255) NOT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE t_schedules(
	id SERIAL,
    user_id INT NOT NULL,
	execute_date VARCHAR(8) NOT NULL,
	execute_time VARCHAR(6) NOT NULL,
	task VARCHAR(1000) NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT t_schedules_user_id_fkey FOREIGN KEY (user_id) REFERENCES m_users(id)
);

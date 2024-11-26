copy role(id, name)
    from '/fixtures/000002/role.csv'
    delimiter ';' csv header;

copy account(id, email, password, hash, is_deleted, role_id)
    from '/fixtures/000002/account.csv'
    delimiter ';' csv header;

copy task(id, title, description, status, created_at)
    from '/fixtures/000002/task.csv'
    delimiter ';' csv header;
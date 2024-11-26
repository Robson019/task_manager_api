-- name: SelectAccountByEmail :one
select a.id as account_id,
       a.email as account_email,
       a.password as account_password,
       a.hash as account_hash,
       r.id as role_id,
       r.name as role_name
    from account a
        inner join role r on r.id = a.role_id
    where a.email = @email and a.is_deleted = false;

-- name: SelectAccountByID :one
select a.id as account_id,
       a.email as account_email,
       a.password as account_password,
       a.hash as account_hash,
       r.id as role_id,
       r.name as role_name
    from account a
        inner join role r on r.id = a.role_id
    where a.id = @account_id and a.is_deleted = false;

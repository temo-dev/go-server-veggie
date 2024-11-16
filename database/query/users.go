package query

const QueryCreateUser = `insert into users (created_at, name_account, password, status, role, user_id) values (?,?,?,?,?,?)`
const QueryUpdateUser = `update users set name_account = ?, status = ?, role = ?, updated_at = ? where id = ?`
const QueryDeleteUserById = `delete from users where id = ?`
const QueryAllUsers = `select id, status, name_account, role, status, created_at, updated_at from users`
const QueryUserById = `select id, status, name_account, role, status, created_at, updated_at from users where id = ?`

const QueryLogin = `select id, name_account, password, status, role from users where name_account = ?`

package query

const QueryCreateUser = `insert into users (created_at,user_name, password, status, user_id) values (?,?,?,?,?)`
const QueryUpdateUser = `update users set user_name = ?, status = ?, updated_at = ? where user_id = ?`
const QueryDeleteUserById = `delete from users where user_id = ?`
const QueryAllUsers = `select id, status, user_name,user_id, role_id, status, created_at, updated_at from users`
const QueryUserById = `select id, status, user_name,user_id, role_id, status, created_at, updated_at from users where user_id = ?`

const QueryLogin = `select id, user_name, user_id, password, status, role_id from users where user_name = ?`
const QueryUserFromToken = `select id from users where user_name = ?`

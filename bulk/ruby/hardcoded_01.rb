# CWE-798: Hardcoded database password in configuration
DB_CONFIG = {
  adapter: 'postgresql',
  host: 'db.internal',
  username: 'app_user',
  password: 'hardcoded-db-password-12345'
}

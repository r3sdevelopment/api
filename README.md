Install PSQL extensions

psql -h localhost -U postgres -d postgres -c 'CREATE EXTENSION IF NOT EXISTS "uuid-ossp";'
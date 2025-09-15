CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS tasks (
    id uuid primary key DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    description text NOT NULL,
    status task_status DEFAULT 'Назначена',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deadline TIMESTAMP,
    version INT NOT NULL DEFAULT 1
)
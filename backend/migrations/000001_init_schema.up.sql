CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE task_status AS ENUM ('wait','running', 'finished');

CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title text NOT NULL,
    detail text,
    status task_status NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
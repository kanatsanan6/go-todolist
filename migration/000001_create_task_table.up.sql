CREATE TABLE tasks (
    id bigserial PRIMARY KEY,
    title varchar(256) NOT null,
    completed boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);

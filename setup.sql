CREATE TABLE todos (
    id VARCHAR(255) PRIMARY KEY,
    note VARCHAR(255),
    created_at timestamp,
    updated_at timestamp
);

INSERT INTO todos (
    id,
    note,
    created_at,
    updated_at
) VALUES (
    "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
    "test",
    "2013-02-03 00:00:00 +0000 UTC",
    "2013-02-03 00:00:00 +0000 UTC"
);

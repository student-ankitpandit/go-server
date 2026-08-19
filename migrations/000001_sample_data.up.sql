CREATE TABLE IF NOT EXISTS post (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    likes BIGINT NOT NULL,
    shares BIGINT NOT NULL,
    comments BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)
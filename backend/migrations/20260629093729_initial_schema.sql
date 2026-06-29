-- +goose Up
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    clerk_id TEXT UNIQUE NOT NULL,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE links (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    short_code TEXT NOT NULL UNIQUE
        CHECK (length(short_code) BETWEEN 6 AND 10),
    long_url TEXT NOT NULL
        CHECK (
                long_url ~ '^https?://'
        ),
    title TEXT,
    click_count BIGINT NOT NULL DEFAULT 0,
    last_clicked_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    password_hash TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE click_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    link_id UUID NOT NULL REFERENCES links(id) ON DELETE CASCADE,
    browser TEXT,
    os TEXT,
    device TEXT,
    referer TEXT,
    ip_hash TEXT,
    clicked_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE TABLE api_keys (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    key_hash TEXT NOT NULL UNIQUE,
    last_used_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    revoked_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_links_user_id
ON links(user_id);
CREATE INDEX IF NOT EXISTS idx_links_expires_at
ON links(expires_at);
CREATE INDEX IF NOT EXISTS idx_click_events_link_id
ON click_events(link_id);
CREATE INDEX IF NOT EXISTS idx_click_events_clicked_at
ON click_events(clicked_at);

-- +goose Down

DROP INDEX IF EXISTS idx_click_events_clicked_at;
DROP INDEX IF EXISTS idx_click_events_link_id;
DROP INDEX IF EXISTS idx_links_expires_at;
DROP INDEX IF EXISTS idx_links_user_id;

DROP TABLE IF EXISTS api_keys;
DROP TABLE IF EXISTS click_events;
DROP TABLE IF EXISTS links;
DROP TABLE IF EXISTS users;

DROP EXTENSION IF EXISTS "pgcrypto";
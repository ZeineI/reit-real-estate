CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    role VARCHAR(10) NOT NULL DEFAULT 'unknown',
    created_at TIMESTAMP(6) DEFAULT now()
);

CREATE TABLE IF NOT EXISTS wallets (
    id UUID PRIMARY KEY,
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    address TEXT NOT NULL,
    created_at TIMESTAMP(6) DEFAULT now()
);

CREATE TABLE IF NOT EXISTS properties (
    id UUID PRIMARY KEY,
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    token_total BIGINT NOT NULL,
    created_at TIMESTAMP(6) DEFAULT now()
);

CREATE TABLE IF NOT EXISTS tokens (
    id UUID PRIMARY KEY,
    property_id UUID NOT NULL UNIQUE REFERENCES properties(id) ON DELETE CASCADE,
    symbol TEXT NOT NULL,
    price NUMERIC(20,6) NOT NULL,
    created_at TIMESTAMP(6) DEFAULT now()
);


CREATE TABLE IF NOT EXISTS user_tokens (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_id UUID NOT NULL REFERENCES tokens(id) ON DELETE CASCADE,
    quantity BIGINT NOT NULL,
    created_at TIMESTAMP(6) DEFAULT now(),
    UNIQUE(user_id, token_id)
);

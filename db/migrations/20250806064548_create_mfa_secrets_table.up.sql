CREATE TABLE mfa_secrets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    secret TEXT NOT NULL,           -- base32-encoded secret
    is_verified BOOLEAN DEFAULT FALSE, -- true after user scans and confirms
    method VARCHAR DEFAULT 'totp',  -- 'totp' or other (for future use)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user_mfa FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

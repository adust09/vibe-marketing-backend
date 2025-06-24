CREATE TABLE user_demographics (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    age_range VARCHAR(20) DEFAULT 'unknown',
    gender VARCHAR(20) DEFAULT 'unknown',
    last_updated_from_api TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_source VARCHAR(50) DEFAULT 'google_ads',
    confidence DECIMAL(5,4),
    privacy_compliant BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_user_demographics_user_id ON user_demographics(user_id);
CREATE INDEX idx_user_demographics_age_range ON user_demographics(age_range);
CREATE INDEX idx_user_demographics_gender ON user_demographics(gender);
CREATE INDEX idx_user_demographics_deleted_at ON user_demographics(deleted_at);

CREATE UNIQUE INDEX idx_user_demographics_unique_user ON user_demographics(user_id) WHERE deleted_at IS NULL;
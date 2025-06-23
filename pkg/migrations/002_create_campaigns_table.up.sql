CREATE TABLE campaigns (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'draft',
    budget_amount DECIMAL(12,2) NOT NULL,
    daily_budget DECIMAL(12,2) NOT NULL,
    start_date DATE,
    end_date DATE,
    objective VARCHAR(100) NOT NULL,
    target_audience JSONB,
    google_ads_campaign_id VARCHAR(50),
    platform VARCHAR(50) DEFAULT 'google',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_campaigns_user_id ON campaigns(user_id);
CREATE INDEX idx_campaigns_status ON campaigns(status);
CREATE INDEX idx_campaigns_deleted_at ON campaigns(deleted_at);
CREATE INDEX idx_campaigns_google_ads_campaign_id ON campaigns(google_ads_campaign_id);
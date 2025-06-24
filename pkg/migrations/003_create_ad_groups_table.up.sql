CREATE TABLE ad_groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    campaign_id UUID REFERENCES campaigns(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'active',
    targeting JSONB,
    google_ads_ad_group_id VARCHAR(50),
    cpc DECIMAL(10,4),
    average_cpc DECIMAL(10,4),
    max_cpc DECIMAL(10,4),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_ad_groups_campaign_id ON ad_groups(campaign_id);
CREATE INDEX idx_ad_groups_status ON ad_groups(status);
CREATE INDEX idx_ad_groups_deleted_at ON ad_groups(deleted_at);
CREATE INDEX idx_ad_groups_google_ads_ad_group_id ON ad_groups(google_ads_ad_group_id);
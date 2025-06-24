CREATE TABLE keywords (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ad_group_id UUID REFERENCES ad_groups(id) ON DELETE CASCADE,
    text VARCHAR(255) NOT NULL,
    match_type VARCHAR(50) NOT NULL,
    status VARCHAR(50) DEFAULT 'active',
    google_ads_keyword_id VARCHAR(50),
    cpc DECIMAL(10,4),
    average_cpc DECIMAL(10,4),
    max_cpc DECIMAL(10,4),
    quality_score INTEGER,
    impressions BIGINT,
    clicks BIGINT,
    cost DECIMAL(12,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_keywords_ad_group_id ON keywords(ad_group_id);
CREATE INDEX idx_keywords_status ON keywords(status);
CREATE INDEX idx_keywords_deleted_at ON keywords(deleted_at);
CREATE INDEX idx_keywords_google_ads_keyword_id ON keywords(google_ads_keyword_id);
CREATE INDEX idx_keywords_match_type ON keywords(match_type);
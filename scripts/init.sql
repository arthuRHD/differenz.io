-- SET ANSI_NULLS ON;
-- SET NOCOUNT ON;
-- SET QUOTED_IDENTIFIER ON;
-- SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;

DROP TABLE IF EXISTS content;
CREATE TABLE IF NOT EXISTS content (
    id AUTO_INCREMENT PRIMARY KEY,
    software VARCHAR(50) NOT NULL,
    bucket_key VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

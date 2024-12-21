CREATE TABLE events (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    postal_code VARCHAR(10) NULL,
    location TEXT NULL,
    proposed_dates VARCHAR(255) NOT NULL,
    created_by VARCHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
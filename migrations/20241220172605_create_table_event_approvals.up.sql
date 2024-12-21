CREATE TABLE event_approvals (
    id VARCHAR(36) PRIMARY KEY,
    event_id VARCHAR(36),
    status ENUM('Pending', 'Approved', 'Rejected') NOT NULL,
    vendor_id VARCHAR(36),
    confirmed_date DATE NULL,
    remarks VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_event_id FOREIGN KEY (event_id) REFERENCES events(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
    CONSTRAINT fk_vendor_id FOREIGN KEY (vendor_id) REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
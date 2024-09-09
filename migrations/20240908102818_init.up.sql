CREATE TABLE tenders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500),
    service_type VARCHAR(50) CHECK (service_type IN ('Construction', 'Delivery', 'Manufacture')),
    status VARCHAR(20) CHECK (status IN ('Created', 'Published', 'Closed')),
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    creator_id UUID REFERENCES employee(id) ON DELETE CASCADE,
    version INT DEFAULT 1 CHECK (version >= 1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bids (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500),
    status VARCHAR(20) CHECK (status IN ('Created', 'Published', 'Canceled', 'Approved', 'Rejected')),
    tender_id UUID REFERENCES tenders(id) ON DELETE CASCADE,
    author_type VARCHAR(50),
    author_id UUID REFERENCES employee(id) ON DELETE CASCADE,
    version INT DEFAULT 1 CHECK (version >= 1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
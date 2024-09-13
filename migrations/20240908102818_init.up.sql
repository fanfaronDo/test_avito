CREATE TABLE tenders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500),
    service_type VARCHAR(50) CHECK (service_type IN ('Construction', 'Delivery', 'Manufacture')),
    status VARCHAR(20) CHECK (status IN ('Created', 'Published', 'Closed')),
    organization_id UUID NOT NULL,
    creator_id UUID NOT NULL,
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
    author_id UUID,
    version INT DEFAULT 1 CHECK (version >= 1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tenders_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tender_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500),
    service_type VARCHAR(50) CHECK (service_type IN ('Construction', 'Delivery', 'Manufacture')),
    status VARCHAR(20) CHECK (status IN ('Created', 'Published', 'Closed')),
    organization_id UUID NOT NULL,
    creator_id UUID NOT NULL,
    version INT DEFAULT 1 CHECK (version >= 1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION tender_versions()
RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'INSERT') THEN
       INSERT INTO tenders_history (tender_id, name, description, service_type, status, organization_id, creator_id)
       VALUES (NEW.id, NEW.name, NEW.description, NEW.service_type, NEW.status, NEW.organization_id, NEW.creator_id);
    RETURN NEW;
    ELSIF (TG_OP = 'UPDATE') THEN
        INSERT INTO tenders_history (tender_id, name, description, service_type, status, organization_id, creator_id, version, created_at)
        VALUES (NEW.id, NEW.name, NEW.description, NEW.service_type, NEW.status, NEW.organization_id, NEW.creator_id, NEW.version, NEW.created_at);
    RETURN NEW;
END IF;
END;
$$
LANGUAGE plpgsql;


CREATE TRIGGER tender_versions_trigger
    AFTER INSERT OR UPDATE ON tenders
    FOR EACH ROW
    EXECUTE PROCEDURE tender_versions();
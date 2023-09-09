CREATE TABLE IF NOT EXISTS Individual (
    
    id serial PRIMARY KEY,
    cim_person_id VARCHAR ( 50 ) UNIQUE NOT NULL,
    cim_organisation_id VARCHAR( 50 ) NOT NULL,
    org_id VARCHAR ( 50 ) NOT NULL, -- TODO: Foreign key?
    country VARCHAR( 30 ) NOT NULL,
    first_name VARCHAR ( 50 ) NOT NULL,
    last_name VARCHAR ( 50 ) NOT NULL,
    picture_url VARCHAR ( 200 ) NOT NULL,
    created_at TIMESTAMP,
    date_of_birth DATE,
    org_name VARCHAR ( 20 )
    -- id serial PRIMARY KEY,
    -- linked_in_profile_id VARCHAR ( 50 ) UNIQUE NOT NULL,
    -- cim_person_id VARCHAR ( 50 ) UNIQUE NOT NULL,
    -- cim_organisation_id VARCHAR( 50 ) NOT NULL,
    -- org_id VARCHAR ( 50 ) NOT NULL, -- TODO: Foreign key?
    -- country VARCHAR( 30 ) NOT NULL,
    -- first_name VARCHAR ( 50 ) NOT NULL,
    -- last_name VARCHAR ( 50 ) NOT NULL,
    -- email VARCHAR ( 50 ) NOT NULL,
    -- picture_url VARCHAR ( 200 ) NOT NULL,
    -- available_for_work BOOLEAN NOT NULL,
    -- status VARCHAR ( 20 ) NOT NULL,
    -- account_status VARCHAR ( 20 ) NOT NULL, -- TODO: AccountStatus object
    -- title VARCHAR ( 10 ) NOT NULL,
    -- created_at TIMESTAMP
    -- is_social_login BOOLEAN NOT NULL,
    -- phone VARCHAR ( 20 ) NOT NULL, -- TODO: Phone object
    -- company_info VARCHAR ( 20 ) NOT NULL, -- TODO: Company info
    -- profile_updated BOOLEAN NOT NULL,
    -- date_of_birth DATE,
    -- gender VARCHAR ( 20 ),
    -- org_name VARCHAR ( 20 )
);

CREATE TABLE IF NOT EXISTS Wallet (
    id serial PRIMARY KEY,
    cim_person_id VARCHAR (50), 
    msp_id VARCHAR ( 50 ) NOT NULL,
    public_key BYTEA UNIQUE NOT NULL,
    private_key BYTEA UNIQUE NOT NULL,
    created_at TIMESTAMP,
    CONSTRAINT FK_wallet_individual FOREIGN KEY (cim_person_id)
        REFERENCES Individual(cim_person_id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Certificates (
    id serial PRIMARY KEY,
    course VARCHAR (50), 
    recipient_id VARCHAR ( 50 ) NOT NULL, -- FK
    event_name VARCHAR ( 50 ) UNIQUE NOT NULL,
    recipient_email VARCHAR ( 50 ) UNIQUE NOT NULL,
    recipient_name VARCHAR ( 200 ), -- first name and last name
    individual_person_id VARCHAR ( 50 ) NOT NULL,
    individual_public_key BYTEA UNIQUE NOT NULL,
    certificate_type INT NOT NULL,
    template_name VARCHAR ( 50 ) NOT NULL,
    custom_template_url VARCHAR ( 200 ) NOT NULL,
    issuer_id VARCHAR ( 50 ) NOT NULL, -- organisation who issued the cert (FK)
    issuer_name VARCHAR ( 200 ) NOT NULL, -- organisation name
    issuer_date TIMESTAMP NOT NULL,
    expiry_date TIMESTAMP NOT NULL,
    created_date TIMESTAMP NOT NULL,
    issued_person_id VARCHAR ( 50 ) NOT NULL, -- organisation admin (FK)
    issuer_public_key BYTEA UNIQUE NOT NULL,
    uid VARCHAR ( 200 ) UNIQUE NOT NULL,
    description VARCHAR ( 1000 ) NOT NULL,
    certificate_name VARCHAR ( 50 ) NOT NULL,
    owner_id VARCHAR ( 50 ) NOT NULL, -- organisationId or individualId (FK)
    owner_name VARCHAR ( 200 ) NOT NULL,
    owner_email VARCHAR ( 200 ) NOT NULL,
    shared_from VARCHAR ( 50 ) NOT NULL, -- certificate id
    master_id VARCHAR ( 50 ) NOT NULL, -- certificate id hold by organisation
    type_of_copy INT NOT NULL,
    is_valid BOOLEAN NOT NULL,
    revoked_reason VARCHAR ( 50 ),
    badge_url VARCHAR ( 200 ),
    template_type VARCHAR ( 50 ),
    certificate_desc VARCHAR ( 1000 ),
    status VARCHAR ( 20 ),
    template_json VARCHAR ( 2000 )
);
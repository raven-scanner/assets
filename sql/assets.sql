CREATE TABLE IF NOT EXISTS assets (
        etag VARCHAR(512),
        id VARCHAR PRIMARY KEY,
        created TIMESTAMP,
        updated TIMESTAMP,

        namespace VARCHAR(512) REFERENCES private.namespaces (id),

        type VARCHAR(128),
        value VARCHAR(512),
        env_cvss VARCHAR(128),
        ttl INTEGER,
        deleteat TIMESTAMP,
        active BOOLEAN,

        CONSTRAINT asset_in_namespace UNIQUE (namespace, type, value)
);


CREATE TABLE IF NOT EXISTS assets_metadata (
    etag VARCHAR(512),
    id VARCHAR PRIMARY KEY,
    created TIMESTAMP,
    updated TIMESTAMP,

    asset VARCHAR REFERENCES assets (id),

    key VARCHAR(512),
    value VARCHAR(512),

    CONSTRAINT metadata_in_asset UNIQUE (asset, key)
);
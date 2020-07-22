CREATE TABLE IF NOT EXISTS namespaces (
    etag VARCHAR(512),
    id VARCHAR PRIMARY KEY,
    created TIMESTAMP,
    updated TIMESTAMP,

    name VARCHAR(512),

    CONSTRAINT namespace_id UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS namespaces_settings (
    etag VARCHAR(512),
    id VARCHAR PRIMARY KEY,
    created TIMESTAMP,
    updated TIMESTAMP,

    namespace VARCHAR REFERENCES namespaces (id),

    key VARCHAR(512),
    value VARCHAR(512),

    CONSTRAINT setting_in_namespace UNIQUE (namespace, key)
);

ALTER TABLE namespaces SET SCHEMA private;
ALTER TABLE namespaces_settings SET SCHEMA private;

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
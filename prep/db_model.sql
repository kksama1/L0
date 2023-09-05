
 --  string  --
CREATE TABLE orders (
    order_uid VARCHAR (19) PRIMARY KEY,
    track_number VARCHAR NOT NULL,
    entry VARCHAR NOT NULL,
    locale VARCHAR (9) NOT NULL,
    internal_signature VARCHAR,
    customer_id VARCHAR NOT NULL,
    delivery_service VARCHAR (100) NOT NULL,
    shardkey VARCHAR NOT NULL,
    sm_id INTEGER UNIQUE NOT NULL,
    date_created TIMESTAMPTZ,
    oof_shard VARCHAR
);

CREATE TABLE delivery (
    name VARCHAR(100) NOT NULL,
    phone VARCHAR (11) NOT NULL,
    zip VARCHAR NOT NULL,
    city VARCHAR(100) NOT NULL,
    address VARCHAR(200) NOT NULL,
    region VARCHAR(150) NOT NULL,
    email VARCHAR(150) NOT NULL
);

CREATE TABLE payment (
    transaction VARCHAR PRIMARY KEY,
    request_id VARCHAR,
    currency VARCHAR(3) NOT NULL,
    provider VARCHAR(50) NOT NULL,
    amount INTEGER NOT NULL,
    payment_dt INTEGER NOT NULL,
    bank VARCHAR(50) NOT NULL,
    delivery_cost INTEGER NOT NULL,
    goods_total INTEGER NOT NULL,
    custom_fee INTEGER NOT NULL
);

CREATE TABLE item (
    chrt_id BIGSERIAL PRIMARY KEY,
    track_number VARCHAR NOT NULL,
    price INTEGER NOT NULL,
    rid VARCHAR(21) UNIQUE NOT NULL ,
    name VARCHAR NOT NULL,
    sale INTEGER,
    size VARCHAR NOT NULL,
    total_price INTEGER NOT NULL,
    nm_id INTEGER UNIQUE NOT NULL,
    brand VARCHAR,
    status INTEGER NOT NULL
);


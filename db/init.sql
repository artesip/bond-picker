CREATE TABLE IF NOT EXISTS t_user
(
    id         UUID                     DEFAULT uuidv7() PRIMARY KEY,
    username   TEXT UNIQUE NOT NULL,
    "password" TEXT        NOT NULL,
    salt       TEXT        NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE INDEX username_indx ON t_user USING HASH (username);

CREATE TABLE IF NOT EXISTS t_portfolio
(
    id         UUID                     DEFAULT uuidv7() PRIMARY KEY,
    user_id    UUID references t_user (id),
    "name"     TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE INDEX user_id_indx ON t_portfolio (user_id);

CREATE TABLE IF NOT EXISTS t_bond
(
    id             TEXT PRIMARY KEY,
    "name"         TEXT,
    "type"         VARCHAR(100),
    sub_type       VARCHAR(100),
    curency_id     VARCHAR(10),

    ytm            real,
    duration       real,
    val_today      real,

    face_value     real,
    price          real,
    lot_size       bigint,
    coupon_period  bigint,
    coupon_percent real,
    issue_size     real,
    acruedint      real,

    next_coupon    DATE,
    put_option     DATE,
    call_option    DATE,

    created_at     TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at     TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS t_portfolio_to_bond
(
    portfolio_id UUID references t_portfolio (id),
    bond_id      TEXT references t_bond (id),
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT now(),

    PRIMARY KEY (portfolio_id, bond_id)
);

CREATE TABLE IF NOT EXISTS t_update_status
(
    id         UUID                     DEFAULT uuidv7() PRIMARY KEY,
    "status"   VARCHAR(100),
    msg        TEXT,
    "start"    TIMESTAMP WITH TIME ZONE,
    "end"      TIMESTAMP WITH TIME ZONE
);
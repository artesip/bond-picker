CREATE TABLE IF NOT EXISTS t_user
(
    id         UUID                     DEFAULT uuidv7() PRIMARY KEY,
    username   TEXT UNIQUE NOT NULL,
    "password" TEXT        NOT NULL,
    salt       TEXT        NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS t_portfolio
(
    id         UUID                     DEFAULT uuidv7() PRIMARY KEY,
    user_id    UUID references t_user (id),
    "name"     TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS t_company
(
    id   TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS t_bond
(
    id             UUID                     DEFAULT uuidv4() PRIMARY KEY,
    company_id     TEXT references t_company (id),

    isin           TEXT,
    "name"         TEXT,
    "type"         VARCHAR(100),
    sub_type       VARCHAR(100),
    currency_id    VARCHAR(10),
    board_id       VARCHAR(10),

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
    mat_date       DATE,

    created_at     TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at     TIMESTAMP WITH TIME ZONE DEFAULT now(),

    CONSTRAINT unique_isin_board UNIQUE (isin, board_id)
);

CREATE TABLE IF NOT EXISTS t_portfolio_to_bond
(
    portfolio_id UUID references t_portfolio (id),
    bond_id      UUID references t_bond (id),
    count        bigint,
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT now(),

    PRIMARY KEY (portfolio_id, bond_id)
);

CREATE TABLE IF NOT EXISTS t_events
(
    id       UUID DEFAULT uuidv7() PRIMARY KEY,
    type     VARCHAR(100) NOT NULL,
    "status" VARCHAR(100),
    msg      TEXT,
    "start"  TIMESTAMP WITH TIME ZONE,
    "end"    TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS t_rating_change
(
    id          UUID DEFAULT uuidv7() PRIMARY KEY,
    company_id  TEXT references t_company (id),
    agency_name TEXT        NOT NULL,
    rating      VARCHAR(50) NOT NULL,
    object_name TEXT        NOT NULL,
    url         TEXT,
    date        DATE        NOT NULL,
    is_revoked  BOOLEAN     NOT NULL,

    CONSTRAINT unique_rating_change UNIQUE (company_id, agency_name, rating, date, object_name)
);

CREATE INDEX IF NOT EXISTS user_id_indx ON t_portfolio (user_id);
CREATE INDEX IF NOT EXISTS username_indx ON t_user USING HASH (username);
CREATE INDEX IF NOT EXISTS company_id_indx ON t_rating_change (company_id);
CREATE INDEX IF NOT EXISTS bond_company_id_indx ON t_bond (company_id);
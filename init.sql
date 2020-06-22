CREATE TABLE countries
(
    id         integer auto_increment unique,
    code       char(2)                               NOT NULL,
    name       varchar(255)                          NOT NULL,
    created_at timestamp default current_timestamp   NOT NULL,
    updated_at timestamp on update current_timestamp NOT NULL,
    primary key (id)
);

CREATE TABLE pilots
(
    id         integer auto_increment unique,
    name       varchar(255)                          NOT NULL,
    surname    varchar(255)                          NOT NULL,
    birthday   datetime                              NOT NULL,
    country_id integer                               NOT NULL,
    created_at timestamp default current_timestamp   NOT NULL,
    updated_at timestamp on update current_timestamp NOT NULL,
    primary key (id)
);

CREATE TABLE airports
(
    id         integer auto_increment unique,
    code       varchar(10)                           NOT NULL,
    country_id integer                               NOT NULL,
    created_at timestamp default current_timestamp   NOT NULL,
    updated_at timestamp on update current_timestamp NOT NULL,
    primary key (id),
    foreign key(country_id) references countries(id) on update cascade on delete cascade
);

CREATE TABLE jets
(
    id              integer auto_increment unique,
    type            enum ('A310','A320','Boeing-737')     NOT NULL,
    color           varchar(255)                          NOT NULL,
    home_airport_id integer,
    created_at      timestamp default current_timestamp   NOT NULL,
    updated_at      timestamp on update current_timestamp NOT NULL,
    primary key (id),
    foreign key (home_airport_id) references airports (id) on update cascade on delete cascade
);

CREATE TABLE flights
(
    id              bigint unsigned auto_increment unique,
    origin_id       integer                               not null,
    destination_id  integer                               not null,
    jet_id          integer                               not null,
    pilot_id        integer                               not null,
    second_pilot_id integer,
    created_at      timestamp default current_timestamp   NOT NULL,
    updated_at      timestamp on update current_timestamp NOT NULL,
    primary key (id),
    foreign key (origin_id) references airports (id) on update cascade on delete cascade,
    foreign key (destination_id) references airports (id) on update cascade on delete cascade,
    foreign key (jet_id) references jets (id) on update cascade on delete cascade,
    foreign key (pilot_id) references pilots (id) on update cascade on delete cascade,
    foreign key (second_pilot_id) references pilots (id) on update cascade on delete cascade
);

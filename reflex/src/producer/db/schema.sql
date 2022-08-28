create table mytable (
 id bigint not null auto_increment,
 i int,

  primary key (id)
);

create table producer_events (
    id bigint not null auto_increment,
    `timestamp` datetime(3) not null,
    foreign_id bigint not null,
    `type` int not null,

    primary key (id)
);

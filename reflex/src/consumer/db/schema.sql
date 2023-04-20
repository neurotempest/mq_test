create table myothertable (
  id bigint not null auto_increment,
  i int,

  primary key (id)
);

create table reflex_cursors (
  `id` varchar(255) not null,
  `cursor` int not null,
  `updated_at` datetime not null,

  primary key (`id`)
);

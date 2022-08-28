create user 'producer'@'%' identified by '';

grant insert,select,update on mytable to 'producer'@'%';
grant insert,select on producer_events to 'producer'@'%';

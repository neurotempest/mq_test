create user 'consumer'@'%' identified by '';

grant insert,select,update on myothertable to 'consumer'@'%';
grant insert,select,update on reflex_cursors to 'consumer'@'%';


create user 'consumer'@'%' identified by '';

grant insert,select,update on myothertable to 'consumer'@'%';


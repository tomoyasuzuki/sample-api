insert into users (id,created_at,updated_at,deleted_at,name) values (1,NOW(),NOW(),NULL,"oshibori");
insert into users (id,created_at,updated_at,deleted_at,name) values (2,NOW(),NOW(),NULL,"sample-user");
insert into tags (id,created_at,updated_at,deleted_at,name) values (1,NOW(),NOW(),NULL,"test-tag-1");
insert into tags (id,created_at,updated_at,deleted_at,name) values (2,NOW(),NOW(),NULL,"test-tag-2");
insert into tasks (id,created_at,updated_at,deleted_at,title,description) values (2,NOW(),NOW(),NULL,"Task 2", "This is a second task.");
insert into tasks (id,created_at,updated_at,deleted_at,title,description) values (1,NOW(),NOW(),NULL,"Task 1", "This is a first task.");
insert into tasks (id,created_at,updated_at,deleted_at,title,description) values (2,NOW(),NOW(),NULL,"Task 2", "This is a second task.");
insert into task_users (user_id,task_id) values (1,1);
insert into task_users (user_id,task_id) values (1,2);
insert into task_users (user_id,task_id) values (2,2);
insert into task_tags (tag_id,task_id) values (1,1);
insert into task_tags (tag_id,task_id) values (1,2);
insert into task_tags (tag_id,task_id) values (2,2);


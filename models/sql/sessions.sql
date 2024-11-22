create table sessions (
    id serial primary key,
    user_id int unique references users (id) on delete cascade,
    token_hash text unique not null 
);



-- alter table sessions add constraint sessions_user_id_fkey 
-- foreign key (user_id) references users (id);

CREATE SCHEMA IF NOT EXISTS almanax;

CREATE TABLE IF NOT EXISTS almanax.almanaxes (
	id serial primary key not null,
	date date not null,
        merydes varchar not null,
        type varchar not null,
        bonus varchar not null,
        offerings varchar not null,
        quantity_offered int not null,
        kamas int not null
);

CREATE TABLE IF NOT EXISTS almanax.classes (
	id serial primary key not null,
	name varchar not null
);

INSERT INTO almanax.classes (name) VALUES
	('Feca'),
	('Osamodas'),
	('Enutrof'),
	('Sram'),
	('Xelor'),
	('Ecaflip'),
	('Eniripsa'),
	('Iop'),
	('Cra'),
	('Sadida'),
	('Panda'), 
	('Roublard'),
	('Zobal'),
	('Steamer'),
	('Eliotrope'),
	('Huppermage'),
	('Sacrieur'),
	('Forgelance'),
	('Ouginak')


CREATE TABLE IF NOT EXISTS almanax.jobs (
	id serial primary key not null,
	name varchar not null,
	type varchar not null
);

INSERT INTO almanax.jobs (name, type) VALUES
	('alchimiste', 'récolte'),
	('Bricoleur', 'fabrication'),
	('Paysan', 'récolte'),
	('Mineur', 'récolte'),
	('Pêcheur', 'récolte'),
	('Bûcheron', 'récolte'),
	('Bijoutier', 'fabrication'),
	('Chasseur', 'récolte'),
	('Forgeron', 'fabrication'),
	('Tailleur', 'fabrication'),
	('Sculteur', 'fabrication'),
	('Façonneur' 'fabrication'),
	('Cordomage', 'forgemagie'),
	('Scultomage', 'forgemagie'),
	('Costumage', 'forgemagie'),
	('Façomage', 'forgemagie'),
	('Joaillomage', 'forgemagie'),
	('Forgemage', 'forgemagie')

	

CREATE TABLE IF NOT EXISTS almanax.users (
	id serial primary key not null,
	name varchar not null,
	username varchar not null,
	class_id serial REFERENCES almanax.classes (id)
);

CREATE TABLE IF NOT EXISTS almanax.users_jobs (
	user_id serial REFERENCES almanax.users,
	job_id serial REFERENCES almanax.jobs
);

CREATE TYPE student_class_status AS ENUM ('active', 'inactive', 'suspended');
CREATE TYPE permission_type AS ENUM ('read', 'write', 'delete');
CREATE TYPE permission_name AS ENUM ('user', 'exam', 'class', 'role', 'permission', 'class_student');

CREATE TABLE role (
  id SERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE permission (
  id SERIAL PRIMARY KEY NOT NULL,
  name permission_name NOT NULL,
  type permission_type NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE role_permission (
  id SERIAL PRIMARY KEY NOT NULL,
  role_id INTEGER NOT NULL,
  permission_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE user_role (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id INTEGER NOT NULL,
  role_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE class (
  id SERIAL PRIMARY KEY NOT NULL,
  teacher_id INTEGER NOT NULL,
  schedule_id INTEGER NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  semester VARCHAR(8) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE class_student (
  id SERIAL PRIMARY KEY NOT NULL,
  class_id INTEGER NOT NULL,
  student_id INTEGER NOT NULL,
  status student_class_status NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE exam (
  id SERIAL PRIMARY KEY NOT NULL,
  class_id INTEGER NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE exam_grade (
  id SERIAL PRIMARY KEY NOT NULL,
  exam_id INTEGER NOT NULL,
  student_id INTEGER NOT NULL,
  grade INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE "user" (
  id SERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  surname VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  document_number VARCHAR(255) NOT NULL,
  document_type VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE teacher_info (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id INTEGER NOT NULL,
  profile_image VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE class_schedule (
  id SERIAL PRIMARY KEY NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE class_schedule_day (
  id SERIAL PRIMARY KEY NOT NULL,
  class_schedule_id INTEGER NOT NULL,
  day VARCHAR(255) NOT NULL,
  start_time TIME NOT NULL,
  end_time TIME NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

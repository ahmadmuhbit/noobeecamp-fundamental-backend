CREATE TABLE "employees" (
  "id" varchar PRIMARY KEY,
  "nip" varchar,
  "name" varchar,
  "address" varchar,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "menus" (
  "id" varchar PRIMARY KEY,
  "name" varchar,
  "category" varchar,
  "desc" varchar,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "transactions" (
  "id" varchar PRIMARY KEY,
  "employee_id" varchar,
  "menu_id" varchar,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("menu_id") REFERENCES "menus" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("employee_id") REFERENCES "employees" ("id");

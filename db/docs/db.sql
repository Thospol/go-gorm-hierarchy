-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-03-07T13:01:00.818Z

CREATE TABLE "charts" (
  "id" BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp DEFAULT null,
  "name" varchar,
  "chart_id" bigint
);

COMMENT ON TABLE "charts" IS 'ข้อมูลผัง';

COMMENT ON COLUMN "charts"."id" IS 'ไอดี';

COMMENT ON COLUMN "charts"."created_at" IS 'วันเวลาที่สร้าง';

COMMENT ON COLUMN "charts"."updated_at" IS 'วันเวลาที่อัพเดตล่าสุด';

COMMENT ON COLUMN "charts"."deleted_at" IS 'วันเวลาที่ลบ';

COMMENT ON COLUMN "charts"."name" IS 'ชื่อ';

COMMENT ON COLUMN "charts"."chart_id" IS 'ref';

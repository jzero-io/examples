CREATE TABLE "products" (
                                   "id" BIGINT NOT NULL AUTO_INCREMENT,
                                   "created_at" DATETIME WITH TIME ZONE NULL,
                                   "updated_at" DATETIME WITH TIME ZONE NULL,
                                   "deleted_at" DATETIME WITH TIME ZONE NULL,
                                   "code" VARCHAR(8188) NULL,
                                   "price" BIGINT NULL,
                                   "remark" VARCHAR(8188) NULL,
                                   CONSTRAINT CONS134218790 PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX INDEX33555460 ON "products" ("id");
CREATE INDEX "idx_products_deleted_at" ON "products" ("deleted_at");
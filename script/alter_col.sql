ALTER TABLE `commodity_tab` ADD cover varchar(255) NOT NULL;

USE commodity;
ALTER TABLE `commodity_tab` ADD latitude DECIMAL(9,6) NOT NULL DEFAULT 0.0;
ALTER TABLE `commodity_tab` ADD longitude DECIMAL(9,6) NOT NULL DEFAULT 0.0;

# ALTER TABLE `commodity_tab` DROP COLUMN latitude;
# ALTER TABLE `commodity_tab` DROP COLUMN longitude;
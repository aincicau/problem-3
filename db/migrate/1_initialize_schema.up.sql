BEGIN;

CREATE TABLE vehicles (
    id int NOT NULL,
    name varchar(100) NOT NULL,
    brand varchar(100) NOT NULL,
    typeOfVehicle varchar(10) NULL,
    canDrive boolean NOT NULL,
    CONSTRAINT vehicle_key PRIMARY KEY (id)
);

COMMIT;
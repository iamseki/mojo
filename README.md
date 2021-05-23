# Mojo

<p align="center">Mojo is a helper cli, idealized to help you with repetitive tasks but can becomes very lazy due to adpotion of Homer's unhealthy lifestyle</p>

## Usage :computer:

- **`mojo collect s3 --bucket=etl-pipeline/uuid --file=6.json --output=output.csv --contract=b3prices`**

## Build and Running :scroll:

## Directory Structure :books:

### CMD (Command)

All commands available in mojo must be created in this directory, the name of go file has the name of command it self, e.g `version.go`.

### Logic

### Contract

### MMI (Man-Machine Interface)

The code that interacts someway with the user, belongs to this folder.

### Helper

DRY, you can put any helper function in this folder, following one rule:
  1. Use the right place for that, i.e `csv.go` contain functions related to csv, `aws.go` related to aws and so on.





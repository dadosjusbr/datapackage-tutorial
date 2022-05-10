library(jsonlite)
library(tidyverse)
library(datapackage.r)
library(tableschema.r)
library(here)

install.packages("frictionless")

# URL package
pkgURL <- "https://dadosjusbr.org/download/datapackage/tjal/tjal-2021-12.zip"

# Download file
download.file(pkgURL, mode = "wb", destfile = here("R/data/tjal-2021-12.zip"))

# Unzip file
unzip(here("R/data/tjal-2021-12.zip"), exdir = here::here("R/data"))

dataPackage = Package.load(here('R/data/datapackage.json'))
dataPackage = dataPackage$infer(here('R/data/*.csv'))
dataPackage$descriptor
dataPackage$resources
dataPackage$valid
remun = dataPackage$getResource('remuneracao')
remun$table$read()

install.packages("jsonlite")
install.packages("tidyverse")
install.packages("devtools")
devtools::install_github("frictionlessdata/datapackage-r")
print("hello world")

library(jsonlite)
library(tidyverse)
library(datapackage-r)

pkgURL <- "https://dadosjusbr.org/download/datapackage/tjal/tjal-2021-12.zip"


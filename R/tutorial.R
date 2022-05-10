library(tidyverse)
library(frictionless)
library(here)

# read from zip
read_dadosjus <- function(pkgURL) {

    pkgDest <- str_extract(pkgURL, "(?<=datapackage\\/).*")
    pkgDir <- str_remove(pkgDest, "\\/.*$")

    pkgDest <- here("R/data", pkgDest)
    pkgDir <- here("R/data", pkgDir)

    dir.create(pkgDir, recursive = TRUE)

    message(str_glue("get {pkgURL}\n"))
    download.file(pkgURL, mode = "wb", destfile = pkgDest)
    Sys.sleep(2)

    message(str_glue("unzip {pkgDest} to {pkgDir}\n"))
    unzip(pkgDest, exdir = pkgDir)
    Sys.sleep(2)

    message("read DadosJusBr datapackage\n")
    package <- read_package(here(pkgDir, 'datapackage.json'))

    unlink(here("R/data"), recursive = TRUE)

    return(package)

}

# URL package
pkgURL <- "https://dadosjusbr.org/download/datapackage/tjal/tjal-2021-12.zip"

# read datapackage
datapackage_tjal_2021_12 <- read_dadosjus(pkgURL = pkgURL)

# inspec
datapackage_tjal_2021_12 %>% enframe()

# List resources
resources(datapackage_tjal_2021_12)

# read table and create a plot
p <- datapackage_tjal_2021_12 %>%
    read_resource("remuneracao") %>% 
    group_by(natureza, categoria) %>%
    summarise(soma_valor = sum(valor), .groups = "drop") %>%
    mutate(
        categoria = reorder(categoria, soma_valor),
        natureza = if_else(natureza == "D", "Descontos", "Recebimentos")
    ) %>%
    ggplot(aes(x = categoria, y = soma_valor, fill = natureza)) +
    geom_col() +
    coord_flip() +
    scale_fill_manual(values = c("Descontos" = "#B361C6", "Recebimentos" = "#2FBB96"))

# save plot
jpeg(file = here("R/plot1.jpeg"))
p
dev.off()

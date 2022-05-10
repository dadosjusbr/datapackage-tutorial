FROM rocker/tidyverse:latest
RUN R -e 'install.packages("remotes")'
RUN R -e 'remotes::install_github("r-lib/remotes", ref = "6c8fdaa")'
RUN R -e 'remotes::install_cran("attempt")'
RUN R -e 'remotes::install_cran("remotes")'
RUN R -e 'remotes::install_cran("dockerfiler")'
RUN R -e 'remotes::install_cran("devtools")'
RUN R -e 'devtools::install_github("frictionlessdata/datapackage-r")'
RUN R -e 'install.packages("frictionless")'
RUN R -e 'install.packages("languageserver")'
RUN R -e 'install.packages("jsonlite")'
RUN R -e 'install.packages("here")'
RUN R -e 'install.packages("lubridate")'
RUN R -e 'install.packages("tidyverse")'
EXPOSE 8787
ENV "PASSWORD"="password"

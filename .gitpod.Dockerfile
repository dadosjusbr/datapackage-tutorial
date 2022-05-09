FROM gitpod/workspace-full

RUN brew install R && brew install gcc@5
RUN R -e 'install.packages("devtools", repo = "https://cloud.r-project.org/")
RUN R -e 'install.packages("tidyverse", repo = "https://cloud.r-project.org/")
RUN R -e 'install.packages("jsonlite", repo = "https://cloud.r-project.org/")
RUN R -e 'devtools::install_github("frictionlessdata/datapackage-r")'
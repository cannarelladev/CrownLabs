FROM ubuntu:22.04

# Prevent apt from prompting the user for input during install/configuration
ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies
RUN apt-get update && apt-get install -y \ 
    wget \
    tar \
    libopenexr-dev \ 
    bzip2 \ 
    build-essential \ 
    zlib1g-dev \ 
    libxmu-dev \ 
    libxi-dev \ 
    libxxf86vm-dev \ 
    libfontconfig1 \ 
    libxrender1 \ 
    libgl1-mesa-glx \ 
    xz-utils \
    tzdata \
    glmark2 && \
  apt-get clean -y && \
  rm -rf /var/lib/apt/lists/*.*

RUN wget -O mangohud.tar.gz https://github.com/flightlessmango/MangoHud/releases/download/v0.7.0/MangoHud-0.7.0.r0.g113912e.tar.gz
RUN tar -xvf mangohud.tar.gz
WORKDIR  /MangoHud
RUN tar xf MangoHud-package.tar
RUN ./mangohud-setup.sh install
RUN chmod 777 /usr/bin/mangohud 

# Define user and user id default arguments
ARG USER=crownlabs
ARG UID=1010

# Define basic default enviroment variables
ENV DISPLAY=:0 \
  USER=${USER} \
  HOME=/home/$USER \
  TZ=Europe/Rome

# Create new user and set a set a valid shell for them
RUN useradd -ms /bin/bash -u ${UID} $USER

# Set permissions on user home
RUN chown -R $USER:$USER $HOME

# Set the user to use
USER $USER


# CMD [ "executable", "parameter", ... ] form throws an error if $PROJECT_DIR is not set
# Therefore, bash-like form is used
CMD ["mangohud", "--dlsym", "glmark2"]

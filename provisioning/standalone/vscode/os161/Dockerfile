# syntax = edrevo/dockerfile-plus

FROM ubuntu:18.04

#Build OS161

ENV DEBIAN_FRONTEND=noninteractive

COPY ./os161/install.sh ./install.sh
RUN chmod 777 install.sh && \
    ./install.sh && \
    rm install.sh

USER crownlabs

COPY --chown=crownlabs ./os161/preliminary.sh ./preliminary.sh
RUN chmod 777 preliminary.sh && \
    ./preliminary.sh

WORKDIR /home/crownlabs/os161

COPY --chown=crownlabs ./os161/binutils.sh ./binutils.sh
RUN chmod 777 binutils.sh && \
    ./binutils.sh

COPY --chown=crownlabs ./os161/gcc.sh ./gcc.sh
RUN chmod 777 gcc.sh && \
    ./gcc.sh

COPY --chown=crownlabs ./os161/gdb-7.8+os161-2.1 ./gdb-7.8+os161-2.1
COPY --chown=crownlabs ./os161/gdb.sh ./gdb.sh
RUN chmod 777 gdb.sh && \
    ./gdb.sh

COPY --chown=crownlabs ./os161/sys161-2.0.8 ./sys161-2.0.8
COPY --chown=crownlabs ./os161/system161.sh ./system161.sh
RUN chmod 777 system161.sh && \
    ./system161.sh

COPY --chown=crownlabs ./os161/generate-symlinks.sh ./generate-symlinks.sh
RUN chmod 777 generate-symlinks.sh && \
    ./generate-symlinks.sh

COPY --chown=crownlabs ./os161/os161-base.sh ./os161-base.sh
RUN chmod 777 os161-base.sh && \
    ./os161-base.sh

COPY --chown=crownlabs ./os161/first-build.sh ./first-build.sh
RUN chmod 777 first-build.sh && \
    ./first-build.sh

COPY --chown=crownlabs ./os161/.gdbinit /home/crownlabs/os161/root/.gdbinit
RUN echo 'set auto-load safe-path /' > /home/crownlabs/.gdbinit


USER root

RUN rm binutils.sh && \
    rm gcc.sh && \
    rm gdb.sh && \
    rm system161.sh && \
    rm generate-symlinks.sh && \
    rm os161-base.sh

#Install vscode

ENV SUDO_FORCE_REMOVE yes
ENV CPPTOOLS_VERSION=1.7.1
ENV CODESERVER_VERSION=4.1.0

RUN apt-get update &&\
    apt-get install -y curl git &&\ 
    curl -fsSL https://code-server.dev/install.sh | sh -s -- --version=${CODESERVER_VERSION} &&\
    apt-get purge -y curl &&\
    apt-get clean

COPY --chown=crownlabs ./os161/config.sh config.sh
RUN chmod 777 config.sh && \
    ./config.sh

COPY --chown=crownlabs ./c-cpp/settings.json /config/data/User/settings.json

RUN apt-get update && apt-get install -y build-essential cmake gdb && \
    apt-get clean && \
    apt-get remove --autoremove --purge -y sudo apt --allow-remove-essential

ADD "https://github.com/microsoft/vscode-cpptools/releases/download/${CPPTOOLS_VERSION}/cpptools-linux.vsix" "./cpptools-linux.vsix"

RUN chown -R crownlabs:crownlabs /config && \
    chown crownlabs:crownlabs cpptools-linux.vsix && \
    chmod 777 cpptools-linux.vsix

COPY --chown=crownlabs ./os161/vscode /home/crownlabs/os161/os161-base-2.0.3/kern/.vscode
RUN chmod 777 -R /home/crownlabs/os161/os161-base-2.0.3/kern/.vscode

COPY --chown=crownlabs ./os161/start.sh vscode-start.sh
RUN chmod 777 vscode-start.sh 

USER crownlabs
RUN code-server --extensions-dir /config/extensions --install-extension cpptools-linux.vsix && \
    code-server --extensions-dir /config/extensions --install-extension formulahendry.code-runner

ENTRYPOINT [ "./vscode-start.sh" ]
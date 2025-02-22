FROM ubuntu:latest

#Make Environment
RUN apt clean
RUN apt update -y
RUN apt install wget -y
RUN wget https://download.oracle.com/graalvm/21/latest/graalvm-jdk-21_linux-aarch64_bin.tar.gz -P /root/
RUN (cd /root && tar xzvf /root/graalvm-jdk-21_linux-aarch64_bin.tar.gz)
RUN wget https://api.papermc.io/v2/projects/paper/versions/1.21.3/builds/82/downloads/paper-1.21.3-82.jar -O /root/server.jar
RUN (cd /root && graalvm-jdk-21.0.6+8.1/bin/java -jar server.jar nogui)
RUN sed -i 's/false/true/g' /root/eula.txt

#Apply Environment
RUN mkdir -p /root/plugins/Skript/scripts/
COPY .github/workflows/plugins/Skript-2.10.1.jar /root/plugins/
COPY .github/workflows/plugins/SkBee-3.8.2.jar /root/plugins/
COPY . /root/plugins/Skript/scripts/

#Runtime
CMD (cd /root && graalvm-jdk-21.0.6+8.1/bin/java -jar server.jar nogui)
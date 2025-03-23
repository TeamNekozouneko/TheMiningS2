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

#Cancelled cache
ADD "https://www.random.org/cgi-bin/randbyte?nbytes=10&format=h" /dev/null

#Apply Environment
RUN mkdir -p /root/plugins/Skript/scripts/
COPY .github/workflows/plugins/ /root/plugins/
COPY . /root/plugins/Skript/scripts/
COPY .github/workflows/worlds/*.tar.gz /root/worlds/
RUN tar xzvf /root/worlds/*.tar.gz -C /root/

#Runtime
CMD (cd /root && graalvm-jdk-21.0.6+8.1/bin/java -jar server.jar nogui)
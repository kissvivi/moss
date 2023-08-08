FROM centos

MAINTAINER jobs_vivi
#定义时区参数
ENV TZ=Asia/Shanghai
#设置时区
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo '$TZ' > /etc/timezone


COPY ./moos /home/moos   
WORKDIR /home/  

RUN chmod +x /home/moos  

ENTRYPOINT ["/home/moos"]  

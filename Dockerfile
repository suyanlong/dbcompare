FROM ubuntu:16.04

WORKDIR /root/chain33

# Use speedup source for Chinese Mainland user,if not you can remove it
RUN mkdir -p $HOME/.config/pip \
    && cp /etc/apt/sources.list  /etc/apt/sources.list.old \
    && sed -i "s/archive.ubuntu.com/mirrors.aliyun.com/g" /etc/apt/sources.list \
    && echo "[global]" > $HOME/.config/pip/pip.conf \
    && echo "index-url = https://mirrors.ustc.edu.cn/pypi/web/simple" >> $HOME/.config/pip/pip.conf \
    && echo "format = columns" >> $HOME/.config/pip/pip.conf

RUN apt-get update \
    && apt-get install -y  python-pip \
    && pip install -U pip \
    && rm -rf /var/lib/apt/lists \
    && rm -rf ~/.cache/pip \
    && apt-get autoremove \
    && apt-get clean \
    && apt-get autoclean
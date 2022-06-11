FROM ubuntu:jammy

WORKDIR /print

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update

RUN apt-get install curl python3 python3-pip python3-venv libbluetooth-dev libopenjp2-7 libtiff5 software-properties-common python3-pil python3-pil.imagetk -y

RUN pip install setuptools==58

COPY requirements.txt .

COPY custom_pkgs custom_pkgs

RUN pip install -r requirements.txt

RUN pip install -e ./custom_pkgs/pybluez

RUN pip install -e ./custom_pkgs/ppa6

COPY main.py main.py

CMD ["python3", "main.py", "http://localhost:8000"]
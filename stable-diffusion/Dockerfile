FROM nvidia/cuda:11.8.0-base-ubuntu22.04

ENV DEBIAN_FRONTEND=noninteractive \
    TORCH_CUDA_VERSION=cu118 \
    TORCH_VERSION=2.0.1 \
    WEB_UI_DIR=/app/stable-diffusion-webui \
    PYTHONUNBUFFERED=1

RUN apt-get update && apt-get install -y \
    wget \
    git \
    libgl1 \
    libglib2.0-0 \
    software-properties-common

RUN add-apt-repository ppa:deadsnakes/ppa -y

RUN apt-get install -y \
    python3.10 \
    python3.10-venv

RUN useradd -m sduser

RUN mkdir -p /app && chown -R sduser:sduser /app

USER sduser

RUN git clone https://github.com/pireu2/stable-diffusion-webui.git ${WEB_UI_DIR}
RUN cd ${WEB_UI_DIR} && python3.10 -m venv venv

WORKDIR ${WEB_UI_DIR}

ENTRYPOINT ["/bin/bash", "webui.sh", "--listen", "--enable-insecure-extension-access", "--api", "--xformers"]
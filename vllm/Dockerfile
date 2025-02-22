FROM nvidia/cuda:12.1.1-base-ubuntu22.04

ENV DEBIAN_FRONTEND=noninteractive \
    PYTHONUNBUFFERED=1 \
    HF_HOME=/app/huggingface \
    MODEL_NAME=deepseek-ai/DeepSeek-R1-Distill-Qwen-14B

RUN apt-get update && apt-get install -y \
    python3.11 \
    python3.11-venv \
    python3.11-dev \
    curl \
    git \
    && rm -rf /var/lib/apt/lists/*

RUN python3.11 -m venv /venv
ENV PATH="/venv/bin:$PATH"

RUN pip install --upgrade pip && \
    pip install torch==2.1.2 torchvision==0.16.2 torchaudio==2.1.2 --index-url https://download.pytorch.org/whl/cu121 && \
    pip install vllm==0.3.3 fastapi==0.110.0 uvicorn==0.27.1 pydantic==2.6.1

RUN mkdir -p ${HF_HOME}

WORKDIR /app
COPY serve.py .

EXPOSE 5000

CMD ["uvicorn", "serve:app", "--host", "0.0.0.0", "--port", "5000"]
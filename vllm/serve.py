from fastapi import FastAPI
from pydantic import BaseModel
from vllm import LLM, SamplingParams
import os

app = FastAPI(title="DeepSeek R1-14B API")
MODEL_NAME = os.getenv("MODEL_NAME", "deepseek-ai/DeepSeek-R1-Distill-Qwen-14B")

@app.on_event("startup")
async def load_model():
    print(f"Loading model {MODEL_NAME} with vLLM")
    app.state.llm = LLM(
        model=MODEL_NAME,
        dtype="float16",
        tensor_parallel_size=int(os.getenv("TENSOR_PARALLEL_SIZE", 1)),
        gpu_memory_utilization=float(os.getenv("GPU_MEMORY_UTILIZATION", 0.95)),
        trust_remote_code=True
    )
    print("Model loaded successfully")

class GenerationRequest(BaseModel):
    prompt: str
    max_tokens: int = int(os.getenv("MAX_TOKENS", 1024))
    temperature: float = 0.6
    top_p: float = 0.95

@app.post("/generate")
async def generate(request: GenerationRequest):
    sampling_params = SamplingParams(
        max_tokens=request.max_tokens,
        temperature=request.temperature,
        top_p=request.top_p
    )
    outputs = llm.generate([request.prompt], sampling_params)
    return {
        "response": outputs[0].outputs[0].text,
        "model": MODEL_NAME,
        "tokens_used": len(outputs[0].outputs[0].token_ids)
    }
"use client"

import { useState } from "react"
import { ArrowLeft } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Card, CardContent } from "@/components/ui/card"
import { Label } from "@/components/ui/label"

export default function ImageGeneratorDashboard({ onBack, isLightThemed }) {

  const [prompt, setPrompt] = useState("")
  const [batchSize, setBatchSize] = useState("4")
  const [isGenerating, setIsGenerating] = useState(false)

  const handleGenerate = () => {
    let handledProp = prompt.trim()
    console.log(`Prompt has been clicked, here is prompt "${handledProp}"`)
    setIsGenerating(true)

    setTimeout(() => {
      setIsGenerating(false)
    }, 2000)
  }

  return (
    <div className={`flex flex-col min-h-screen ${isLightThemed ? "bg-gray-100" : "bg-gray-0"}`}>
      <div className="bg-white border-b p-4 flex items-center">
        <Button variant="ghost" className="mr-4" onClick={onBack}>
          <ArrowLeft className="h-5 w-5 mr-2" />
          Back
        </Button>
      </div>


      <div className="flex-1 p-6 space-y-6 ">

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">

          <Card className="h-full">
            <CardContent className="pt-6 space-y-4">
              <div className="space-y-2">
                <Label htmlFor="prompt" className="text-base font-medium">
                  Prompt:
                </Label>
                <Input
                  id="prompt"
                  placeholder="Enter your image generation prompt..."
                  className="min-h-[120px] resize-none"
                  value={prompt}
                  onChange={(e) => setPrompt(e.target.value)}
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="batchSize" className="text-base font-medium">
                  Batch size:
                </Label>
                <Input
                  id="batchSize"
                  type="number"
                  min="1"
                  max="10"
                  className="w-24"
                  value={batchSize}
                  onChange={(e) => setBatchSize(e.target.value)}
                />
              </div>
            </CardContent>
          </Card>

          <Card className="h-full flex items-center justify-center">
            <CardContent className="w-full flex items-center justify-center">
              <Button
                size="lg"
                className="w-3/4 h-16 text-lg"
                onClick={handleGenerate}
                disabled={isGenerating || !prompt.trim()}
              >
                {isGenerating ? "Generating..." : "Generate"}
              </Button>
            </CardContent>
          </Card>
        </div>

        <Card className="mt-6">
          <CardContent className="pt-6 flex flex-col items-center">
            <h3 className="text-xl font-medium mb-4">Image Preview</h3>
            <div className="w-full max-w-2xl aspect-square border-2 border-dashed border-gray-300 rounded-md flex items-center justify-center bg-gray-50">
              {isGenerating ? (
                <div className="text-center">
                  <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-2"></div>
                  <p className="text-gray-500">Generating image...</p>
                </div>
              ) : (
                <p className="text-gray-500">Generated image will appear here</p>
              )}
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}


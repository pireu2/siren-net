"use client"

import { useState } from "react"
import { ArrowLeft,Download, Share2  } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Card, CardContent,CardFooter } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { useApi } from "./ApiContext";

export default function ImageGeneratorDashboard({ onBack, isLightThemed }) {

  const [prompt, setPrompt] = useState("")
  const [isGenerating, setIsGenerating] = useState(false)
  const [imageData, setImageData] = useState(null)
  const { getStableDiffusionImage } = useApi()


  const handleDownload = () => {
    if (imageData) {
      const link = document.createElement("a")
      link.href = `data:image/png;base64,${imageData}`
      link.download = "generated-image.png"
      document.body.appendChild(link)
      link.click()
      document.body.removeChild(link)
    }
  }

  const handleShare = async () => {
    if (imageData && navigator.share) {
      try {
        //base64 to blob
        const response = await fetch(`data:image/png;base64,${imageData}`)
        const blob = await response.blob()
        const file = new File([blob], "generated-image.png", { type: "image/png" })

        await navigator.share({
          title: "Generated Image",
          text: "Check out this generated image!",
          files: [file],
        })
      } catch (error) {
        console.error("Error sharing:", error)
        navigator.clipboard.writeText(`data:image/png;base64,${imageData}`)
      }
    } else {
      navigator.clipboard.writeText(`data:image/png;base64,${imageData}`)
    }
  }


  const handleGenerate = () => {
    let handledProp = prompt.trim()
    setIsGenerating(true)

    try {
    getStableDiffusionImage(handledProp).then((data) => {
  if (data && data.images) {
    setImageData(data.images[0]);
  } else if (data && data.error) {
    console.error("Error generating image:", data.error);
    alert(`Error: ${data.error}`);
  } else {
    console.error("No image data received");
  }
  setIsGenerating(false);
});
  } catch (error) {
    console.error("Error generating response:", error);
    setIsGenerating(false);
  }
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
          ) : imageData ? (
            <div className="image-preview-container w-full h-full flex items-center justify-center">
              <img
                src={`data:image/png;base64,${imageData}`}
                alt="Generated Image"
                className="w-full h-auto rounded-md shadow-md"
              />
            </div>
          ) : null}
        </div>
      </CardContent>

          {imageData && !isGenerating && (
            <CardFooter className="flex justify-center gap-3 pb-6">
              <Button
                onClick={handleDownload}
                size="lg"
                className="flex items-center gap-2 bg-gradient-to-r from-violet-600 to-purple-700 hover:from-violet-700 hover:to-purple-800 text-white shadow-lg hover:shadow-xl transition-all duration-200 transform hover:scale-105"
              >
                <Download className="w-5 h-5" />
                Download
              </Button>

              <Button
                onClick={handleShare}
                variant="outline"
                size="lg"
                className="flex items-center gap-2 border-2 border-black hover:border-violet-600 bg-white hover:bg-gradient-to-r hover:from-violet-50 hover:to-purple-50 text-black hover:text-violet-700 shadow-md hover:shadow-lg transition-all duration-200 transform hover:scale-105"
              >
                <Share2 className="w-5 h-5" />
                Share
              </Button>
            </CardFooter>
          )}
        </Card>
      </div>
    </div>
  )
}


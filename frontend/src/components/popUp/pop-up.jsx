"use client"

import { useState, useEffect } from "react"
import { X } from 'lucide-react'
import { cn } from "@/lib/utils"

export default function Popup({
  message,
  isOpen,
  setOpenState,
  duration = 10000,
  className,
}) {

  useEffect(() => {
    let timer
    
    if (isOpen && duration > 0) {
      timer = setTimeout(() => {
        setOpenState(false)
      }, duration)
    }
    
    return () => {
      if (timer) clearTimeout(timer)
    }
  }, [isOpen, duration])

  const handleClose = () => {
    setOpenState(false)
  }

  if (!isOpen) return null

  return (
    <div className="fixed inset-x-0 z-50 flex justify-center pointer-events-none">
      <div 
        className={cn(
          "max-w-sm w-full mx-4 bg-red-500/80 backdrop-blur-sm text-white rounded-lg shadow-lg transition-all duration-300 ease-in-out pointer-events-auto top-4",
          className
        )}
      >
        <div className="relative p-4 flex items-center">
          <div className="flex-1 pr-8">{message}</div>
          <button 
            onClick={handleClose}
            className="absolute right-2 top-2 p-1 rounded-full hover:bg-red-200/150 transition-colors"
            aria-label="Close"
          >
            <X className="h-4 w-4" />
          </button>
        </div>
      </div>
    </div>
  )
}

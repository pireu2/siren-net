"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { ArrowLeft} from "lucide-react"

const teamMembers = [
  {
    name: "Duica Sebastian",
    role: "Lead Backend Engineer",
    image: "src/images/sebi.jpg",
    id: 1
  },
  {
    name: "Oltean Simion",
    role: "Lead Frontend Architect",
    image: "src/images/egg.jpg",
    id: 2
  },
]

export default function TeamPage({onBack}) {

  function PanelComponent()
  {
    return(
      <div className="min-h-screen bg-gradient-to-br flex items-center justify-center p-4">
      <div className="bg-white rounded-lg shadow-2xl p-8 max-w-4xl w-full">
        <h1 className="text-3xl font-bold text-center mb-8 text-gray-800">Our Amazing Team</h1>
        <div className="flex flex-wrap justify-center gap-12">
          {teamMembers.map(member => (
            <div
              key={member.id}
              className="flex flex-col items-center"
            >
              <div className="relative">
                <div className= 'w-48 h-48 rounded-full overflow-hidden'>
                  <img
                    src={member.image}
                    alt={member.name}
                    className="w-full h-full object-cover"
                  />
                </div>
              </div>
              <h2 className="mt-4 text-xl font-semibold text-gray-800">{member.name}</h2>
              <p className="text-gray-600 italic">{member.role}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
    );
  }

  return (
    <div className="flex flex-col min-h-screen bg-gray-50">
      <div className="bg-white border-b p-4 flex items-center" onClick = {onBack}>
        <Button variant="ghost" className="mr-4">
          <ArrowLeft className="h-5 w-5 mr-2" />
          Back
        </Button>
      </div>

      <div className="flex-1 p-6 space-y-6">
        <PanelComponent/>
      </div>
    </div>
    
  )
}


"use client"

import { useState } from "react"
import { ArrowLeft, ChevronDown, ChevronUp, SortAsc, MessageSquare } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from "@/components/ui/collapsible"
import { Textarea } from "@/components/ui/textarea"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"

// mock data for clients
const mockClients = [
  { id: 1, name: "Acme Corp", importance: "High" },
  { id: 2, name: "Globex Industries", importance: "Medium" },
  { id: 3, name: "Wayne Enterprises", importance: "High" },
  { id: 4, name: "Stark Industries", importance: "Low" },
  { id: 5, name: "Umbrella Corporation", importance: "Medium" },
]

// mock data for conversations
const mockConversations = [
  { id: 1, date: "2024-02-25", title: "Initial Project Discussion", summary: "Discussed project scope and timeline." },
  {
    id: 2,
    date: "2024-02-20",
    title: "Budget Review",
    summary: "Reviewed budget constraints and resource allocation.",
  },
  {
    id: 3,
    date: "2024-02-15",
    title: "Technical Requirements",
    summary: "Defined technical specifications and requirements.",
  },
  {
    id: 4,
    date: "2024-02-10",
    title: "Stakeholder Meeting",
    summary: "Met with key stakeholders to align on project goals.",
  },
  { id: 5, date: "2024-02-05", title: "Kickoff Meeting", summary: "Project kickoff and team introductions." },
]

export default function AiPromptingDashboard({onBack}) {
  const [selectedClient, setSelectedClient] = useState("")
  const [isContextOpen, setIsContextOpen] = useState(false)
  const [clientSummary, setClientSummary] = useState("")
  const [sortByImportance, setSortByImportance] = useState(false)
  const [sortByDate, setSortByDate] = useState(false)


  // sort clients by importance if needed
  const sortedClients = [...mockClients].sort((a, b) => {
    if (!sortByImportance) return 0

    const importanceOrder = { High: 3, Medium: 2, Low: 1 }
    return importanceOrder[b.importance] - importanceOrder[a.importance]
  })

  // sort conversations by date if needed
  const sortedConversations = [...mockConversations].sort((a, b) => {
    if (!sortByDate) return 0
    return new Date(b.date).getTime() - new Date(a.date).getTime()
  })

  return (
    <div className="flex flex-col min-h-screen bg-gray-50">
      {/* Top slim rectangle */}
      <div className="bg-white border-b p-4 flex items-center" onClick = {onBack}>
        <Button variant="ghost" className="mr-4">
          <ArrowLeft className="h-5 w-5 mr-2" />
          Back
        </Button>
      </div>

      {/* Bottom chunky rectangle */}
      <div className="flex-1 p-6 space-y-6">
        {/* Clients section */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between pb-2">
            <CardTitle className="text-xl">Clients</CardTitle>
            <Button
              variant="outline"
              size="sm"
              onClick={() => setSortByImportance(!sortByImportance)}
              className="flex items-center gap-1"
            >
              <SortAsc className="h-4 w-4" />
              Sort by Importance
            </Button>
          </CardHeader>
          <CardContent>
            <Select value={selectedClient} onValueChange={setSelectedClient}>
              <SelectTrigger className="w-full">
                <SelectValue placeholder="Select a client" />
              </SelectTrigger>
              <SelectContent>
                {sortedClients.map((client) => (
                  <SelectItem key={client.id} value={client.id.toString()}>
                    <div className="flex items-center justify-between w-full">
                      <span>{client.name}</span>
                      <Badge
                        variant={
                          client.importance === "High"
                            ? "destructive"
                            : client.importance === "Medium"
                              ? "default"
                              : "secondary"
                        }
                      >
                        {client.importance}
                      </Badge>
                    </div>
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </CardContent>
        </Card>

        {/* Context section */}
        <Card>
          <CardHeader className="pb-2">
            <Collapsible open={isContextOpen} onOpenChange={setIsContextOpen} className="w-full">
              <CollapsibleTrigger asChild>
                <div className="flex items-center justify-between cursor-pointer w-full">
                  <CardTitle className="text-xl">Context</CardTitle>
                  <div className="flex items-center gap-2">
                    <Button
                      variant="outline"
                      size="sm"
                      onClick={(e) => {
                        e.stopPropagation()
                        setSortByDate(!sortByDate)
                      }}
                      className="flex items-center gap-1"
                    >
                      <SortAsc className="h-4 w-4" />
                      Sort by Date
                    </Button>
                    {isContextOpen ? <ChevronUp className="h-4 w-4" /> : <ChevronDown className="h-4 w-4" />}
                  </div>
                </div>
              </CollapsibleTrigger>
              <CollapsibleContent className="mt-4">
                <div className="space-y-3">
                  {sortedConversations.map((conversation) => (
                    <div key={conversation.id} className="border rounded-md p-3 bg-gray-50">
                      <div className="flex justify-between items-center mb-1">
                        <h4 className="font-medium">{conversation.title}</h4>
                        <span className="text-sm text-gray-500">
                          {new Date(conversation.date).toLocaleDateString()}
                        </span>
                      </div>
                      <p className="text-sm text-gray-600">{conversation.summary}</p>
                    </div>
                  ))}
                </div>
              </CollapsibleContent>
            </Collapsible>
          </CardHeader>
        </Card>

        {/* Summary of client tendencies */}
        <Card>
          <CardHeader className="pb-2">
            <CardTitle className="text-xl">Summary of Client Tendencies</CardTitle>
          </CardHeader>
          <CardContent>
            <Textarea
              placeholder="Enter a summary of client behaviors and tendencies..."
              className="min-h-[120px]"
              value={clientSummary}
              onChange={(e) => setClientSummary(e.target.value)}
            />
          </CardContent>
        </Card>

        {/* Parameters section */}
        <Card>
          <CardHeader className="pb-2">
            <CardTitle className="text-xl">Parameters</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="p-4 border border-dashed rounded-md text-center text-gray-500">
              Parameters section (empty for now)
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}


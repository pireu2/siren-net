"use client"
import { getCookie } from 'react-use-cookie';
import { useState } from "react"
import { ArrowLeft,ArrowRight, ChevronDown, ChevronUp, SortAsc, MessageSquare } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from "@/components/ui/collapsible"
import { Textarea } from "@/components/ui/textarea"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import  { useEffect } from "react"
import { filterProps } from 'framer-motion';
import {useRef} from "react"
import { cn } from "@/lib/utils"

export default function AiPromptingDashboard({onBack}) {
  const [selectedClient, setSelectedClient] = useState("")
  const [selectedAgent, setSelectedAgent] = useState("")
  const [isContextOpen, setIsContextOpen] = useState(false);
  const [clientSummary, setClientSummary] = useState("")
  const [sortByImportance, setSortByImportance] = useState(false)
  const [sortByDate, setSortByDate] = useState(false)
  const [Clients, setClients] = useState([])
  const [Agents, setAgents] = useState([])
  const [sortedConversations, setSortedConversations] = useState([]);
  const agentRef = useRef(0)
  const [clientSelectOpen, setClientSelectOpen] = useState(false);
  const [isGenerating, setIsGenerating] = useState(false);
  const [response, setResponse] = useState("");


  const handleGenerateResponse = () => {
    if (!clientSummary.trim()) return;
    
    setIsGenerating(true);
    setResponse(""); 
    
    setTimeout(() => {
      setIsGenerating(false);
      setResponse("Based on the client's tendencies, I recommend approaching them with a supportive tone. Focus on their key concerns about [topic] and emphasize how our solution addresses their specific needs around [pain point].");
    }, 2000);
  }



  const getConversations = async () => {
      if (!selectedAgent || !selectedClient) {
        setSortedConversations([]); 
        return;
      }
      const token = getCookie('token');
      if (!token) {
        return;
      }
      try {
        const response = await fetch(
          `/messages/agent/${selectedAgent}/client/${selectedClient}`,
          {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
              "Authorization": `Bearer ${token}`,
            },
          }
        );
        const data = await response.json();
        console.log("Conversations data:", data);
        const mappedConversations = data.map(msg => ({
          id: msg.ID,
          date: msg.CreatedAt,
          title: msg.Title || "Message", 
          summary: msg.Summary || msg.Content || "", 
          type: msg.Type,
        }));
        
        setSortedConversations(mappedConversations.sort((a, b) => {
          if (!sortByDate) return 0;
          return new Date(b.date).getTime() - new Date(a.date).getTime();
        }));
      } catch (error) {
        console.error("Error fetching conversations:", error);
        setSortedConversations([]);
      }
    };

    useEffect(() => { 
      const token = getCookie('token');
      if (!token) {
        return;
      }
 
      const getClients = async () => { 
      try {
        const agentId = selectedAgent || agentRef.current;
    if (!agentId) return;

        const response = await fetch(`/clients/agent/${agentId}`, {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
          },
        });
        
        const data = await response.json(); 
        const mappedClients = data.map(client => ({
          id: client.ID, 
          name: client.Name,
          importance: client.Score 
        }));
        setClients(mappedClients);
      } 
      catch (error) {
        console.error("Error making client request:", error);
      }
    };

    const getAgents = async () => { 
      try {
        const response = await fetch("/agents", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
          },
        });
        
        const data = await response.json(); 
        const mappedAgents = data.map(agent => ({
          id: agent.ID,
          date: agent.CreatedAt,
          title: agent.Name,
          summary: agent.Characteristics
        }));
        setAgents(mappedAgents);
      } 
      catch (error) {
        console.error("Error making agent request:", error);
      }
    };
  
    getClients();
    getAgents();
    if (isContextOpen && selectedAgent && selectedClient) {
      getConversations();
    }

  }, [selectedAgent, selectedClient,isContextOpen]);

  const sortedClients = [...Clients].sort((a, b) => {
  if (!sortByImportance) return 0;
  return b.importance - a.importance;
});


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
              onClick={() => {setSortByImportance(!sortByImportance);}}
              className="flex items-center gap-1"
            >
              <SortAsc className="h-4 w-4" />
              Sort by Importance
            </Button>
          </CardHeader>
          <CardContent>
            <div className="flex gap-4">
              <Select 
              value={selectedClient} 
              onValueChange={setSelectedClient} 
              open={clientSelectOpen}
              onOpenChange={setClientSelectOpen}
              className="w-full">
                <SelectTrigger>
                  <SelectValue placeholder="Select a client" />
                </SelectTrigger>
                <SelectContent>
                  {sortedClients.map((client) => (
                    <SelectItem key={client.id} value={client.id.toString()}>
                      <div className="flex items-center just</SelectItem>ify-between w-full">
                        <span>{client.name}</span>
                        <Badge
                          variant={
                            client.importance >= 5
                              ? "destructive"
                              : client.importance >= 3
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
              
              <Select value={selectedAgent} onValueChange={setSelectedAgent} className="w-full">
                <SelectTrigger>
                  <SelectValue placeholder="Select an agent" />
                </SelectTrigger>
                <SelectContent>
                  {Agents.map((agent) => (
                    <SelectItem key={agent.id} value={agent.id.toString()}>
                      <div className="flex items-center justify-between w-full">
                        <span>{agent.title}</span>
                      </div>
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </CardContent>
        </Card>

        {/* Context section */}
        <Card>
          <CardHeader className="pb-2">
            <Collapsible
                open={isContextOpen}
                onOpenChange={(open) => {
                  setIsContextOpen(open);
                  console.log("Context section toggled:", selectedAgent, selectedClient, open);
                  if (open && selectedAgent && selectedClient) {
                    getConversations();
                  }
                }}
                className="w-full"
              >
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

                        if (isContextOpen && selectedAgent && selectedClient) {
                          getConversations();
                        }
                        setIsContextOpen(!isContextOpen);
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
                        <h4 className="font-medium flex items-center">
                          {conversation.type === "CLIENT_TO_AGENT" ? (
                            <>
                              Client
                              <ArrowRight className="inline h-4 w-4 mx-1" />
                              Agent
                            </>
                          ) : (
                            <>
                              Client
                              <ArrowLeft className="inline h-4 w-4 mx-1" />
                              Agent
                            </>
                          )}
                        </h4>
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
          <CardContent className="space-y-4">
            <Textarea
              placeholder="Enter a summary of client behaviors and tendencies..."
              className="min-h-[120px]"
              value={clientSummary}
              onChange={(e) => setClientSummary(e.target.value)}
            />
            <div className="flex justify-center">
              <Button
                size="lg"
                className="w-3/4 h-12 text-lg"
                onClick={handleGenerateResponse}
                disabled={isGenerating || !clientSummary.trim()}
              >
                {isGenerating ? "Generating..." : "Generate Response"}
              </Button>
            </div>
          </CardContent>
        </Card>


        {/* Response section (formerly Parameters) */}
<Card>
  <CardHeader className="pb-2">
    <CardTitle className="text-xl">Response</CardTitle>
  </CardHeader>
  <CardContent>
    {isGenerating ? (
      <div className="p-4 border rounded-md text-center">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-2"></div>
        <p className="text-gray-500">Generating response...</p>
      </div>
    ) : response ? (
      <div className="p-4 border rounded-md">
        <p className="text-gray-700">{response}</p>
      </div>
    ) : (
      <div className="p-4 border border-dashed rounded-md text-center text-gray-500">
        Response will appear here
      </div>
    )}
  </CardContent>
</Card>
      </div>
    </div>
  )
}


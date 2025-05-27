"use client"

import { useState } from "react"
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
  LineChart,
  Line,
  ScatterChart,
  Scatter,
  Cell,
} from "recharts"
import { MessageSquare, Clock, DollarSign, ChevronDown, ChevronRight, TrendingUp, Users, Activity } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Progress } from "@/components/ui/progress"
import { ArrowLeft } from "lucide-react"
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarProvider,
  SidebarInset,
  SidebarHeader,
} from "@/components/ui/sidebar"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { cn } from "@/lib/utils"
import { useEffect } from "react";
import { useApi } from "@/components/ai-prompting/ApiContext"


// const timelineData = [
//   { time: "09:00", sender: "agent", type: "normal", message: "Good morning! How can I help?" },
//   { time: "09:15", sender: "client", type: "normal", message: "Hi, I have a question about pricing" },
//   { time: "09:18", sender: "agent", type: "normal", message: "I'd be happy to help with that" },
//   { time: "09:45", sender: "client", type: "transaction", message: "I'd like to proceed with the purchase" },
//   { time: "10:00", sender: "agent", type: "transaction", message: "Great! Let me process that for you" },
//   { time: "10:30", sender: "client", type: "emotional", message: "Thank you so much for your help!" },
// ]

export default function AIAgentDashboard({onBack}) {
const [agents, setAgents] = useState([]);
const [selectedAgent, setSelectedAgent] = useState(agents[0] || { id: null, name: "Select an agent", clientCount: 0, status: "active" });
const [expandedClient, setExpandedClient] = useState(null);
const [symmetryData, setSymmetryData] = useState([]);
const [messageGapData, setMessageGapData] = useState([]);
const [transactionData, setTransactionData] = useState([]);
const [totalRevenue, setTotalRevenue] = useState(0);
const [clientPerformanceData, setClientPerformanceData] = useState([]);
const [timelineData, setTimelineData] = useState([]);
const { getClients, getAgents, getConversations, getTransactions } = useApi();

        async function fetchAgents() {
            try {
                const agentsData = await getAgents();
                const mappedAgents = await Promise.all(
                    agentsData.map(async (agent) => {
                        const agentClients = await getClients(agent.ID);
                        return {
                            id: agent.ID,
                            name: agent.Name,
                            status: agent.status || "active",
                            clientCount: Array.isArray(agentClients) ? agentClients.length : 0,
                        };
                    })
                );
                setAgents(mappedAgents);
                if (mappedAgents.length > 0 && !selectedAgent.id) {
                    setSelectedAgent(mappedAgents[0]);
                    fetchDataForAgent(mappedAgents[0].id);
                }
                
            } catch (err) {
                console.error("Error fetching agents or clients:", err);
            }
        }


        async function fetchDataForAgent(agentId) {
            try {
            const clientData = await getClients(agentId);
            console.log("Client Data fetched:", clientData);
            var symmetry = [];
            var timeGapsByDay = [];
            var transactionArr = [];
            var ClientPerformance = [];
            let transactionX = 0; 
            let totalRevenueLocal = 0;
            let timeLineDataLocal = [];
            let MaxScore = -1;
            let bestClientData = null;

                for(var client of clientData)
                {
                    var clientTransaction = 0;
                    await getTransactions(agentId, client.ID).then((transactions) => 
                    {
                        clientTransaction = transactions.length;
                        totalRevenueLocal += transactions.reduce((sum, transaction) => sum + (transaction.Amount || 0), 0);
                    })

                    if(client.Score > MaxScore) 
                        {
                            MaxScore = client.Score;
                            bestClientData = client;
                        };

                    await getConversations(agentId, client.ID).then((conversationsClient) => {
                        let agentToClient = 0;
                        let clientToAgent = 0;

                        const getDayAbbreviation = (dateObj) => {
                        const days = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
                        return days[dateObj.getDay()];
                        };

                        const sortedConvs = [...conversationsClient].sort(
                          (a, b) => new Date(a.CreatedAt) - new Date(b.CreatedAt)
                        );

                        const gapsByDay = {};

                        for (let i = 1; i < sortedConvs.length; i++) {
                          const prev = new Date(sortedConvs[i - 1].CreatedAt);
                          const curr = new Date(sortedConvs[i].CreatedAt);
                          const dayOfWeek = getDayAbbreviation(curr);
                          
                          const gapMs = curr - prev;
                          const gapHours = gapMs / (1000 * 60 * 60);
                          if (!gapsByDay[dayOfWeek]) {
                            gapsByDay[dayOfWeek] = [];
                        }
                        gapsByDay[dayOfWeek].push(gapHours);
                        
                        }


                        conversationsClient.forEach(conv => {
                            if (conv.Type === "AGENT_TO_CLIENT") agentToClient++;
                            else if (conv.Type === "CLIENT_TO_AGENT") clientToAgent++;
                        });
                        let ratio = clientToAgent === 0 ? 0 : agentToClient / clientToAgent;
                        symmetry.push({ client: client.Name, ratio});

                        transactionArr.push({
                        client: client.Name,
                        ratio: (clientTransaction / conversationsClient.length).toFixed(3), 
                        amount: clientTransaction,
                        x: transactionX++,
                        y: (clientTransaction / conversationsClient.length).toFixed(3) 
                        });

                        ClientPerformance.push({
                            id: client.ID,
                            name: client.Name,
                            score: client.Score, 
                            symmetryRatio: ratio,
                            avgMsgGap: (Object.values(gapsByDay)
                            .flat()
                            .reduce((sum, gap) => sum + gap, 0) / 
                            Object.values(gapsByDay).flat().length).toFixed(2) || 0,
                            txRatio: (clientTransaction / conversationsClient.length).toFixed(3),
                            status: clientTransaction > 1 ? "excellent" : clientTransaction > 2 ? "not important" : "needs attention",
                        });

                       for (const day in gapsByDay) {
                        const dayGaps = gapsByDay[day];
                        const avgGap = dayGaps.reduce((a, b) => a + b, 0) / dayGaps.length;
                        
                        timeGapsByDay.push({
                            day: day,
                            avgGap: avgGap.toFixed(3)
                        });
                        }
                    });
                }
                var formattedTimelineData = [];
                if (bestClientData) {
                const bestClientConvs = await getConversations(agentId, bestClientData.ID);
                formattedTimelineData = bestClientConvs.map((conv) => {
                    const time = new Date(conv.CreatedAt).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
                    return {
                        time,
                        sender: conv.Type === "AGENT_TO_CLIENT" ? "agent" : "client",
                        type: conv.Type === "TRANSACTION" ? "transaction" : conv.Type === "EMOTIONAL" ? "emotional" : "normal",
                        message: conv.Content || "",
                    };
                });
                }


                const groupedByDay = {};
                timeGapsByDay.forEach(({ day, avgGap }) => {
                  if (!groupedByDay[day]) groupedByDay[day] = [];
                  groupedByDay[day].push(Number(avgGap));
                });

                const averagedByDay = Object.entries(groupedByDay).map(([day, gaps]) => ({
                  day,
                  avgGap: gaps.reduce((a, b) => a + b, 0) / gaps.length,
                }));
                console.log("performance data:", ClientPerformance);
                setClientPerformanceData(ClientPerformance);
                setTransactionData(transactionArr);
                setMessageGapData(averagedByDay);
                setSymmetryData(symmetry);
                console.log("timeLineDataLocal:", timeLineDataLocal);
                setTotalRevenue(totalRevenueLocal);
                setTimelineData(formattedTimelineData);
            
            } catch (err) {
                console.error("Error fetching agents or clients:", err);
            }
        }
    

    
    useEffect(() => {
        fetchAgents();
    }, []);


  const getScoreBadgeVariant = (score) => {
    if (score >= 4) return "default"
    if (score >= 2) return "secondary"
    return "destructive"
  }


    if(!selectedAgent.id)
    {
        return null;
    }
    return (
    
    <SidebarProvider>
      <div className="flex flex-col min-h-screen bg-gray-50">
        {/* Top slim rectangle */}
      <div className="bg-white border-b px-4 py-2 flex items-center" onClick={onBack}>
    <Button variant="ghost" className="mr-4">
      <ArrowLeft className="h-5 w-5 mr-2" />
      Back
    </Button>
  </div>
        
        <div className="flex flex-1">
        {/* Sidebar */}
        <Sidebar className="border-r">
          <SidebarHeader className="p-4 border-b">
            <h2 className="text-lg font-semibold">My Agents</h2>
          </SidebarHeader>
          <SidebarContent>
            <SidebarGroup>
              <SidebarGroupLabel>Active Agents</SidebarGroupLabel>
              <SidebarGroupContent>
                <SidebarMenu>
                  {agents.map((agent) => (
                    <SidebarMenuItem key={agent.id}>
                      <SidebarMenuButton
                        onClick={() => {setSelectedAgent(agent);fetchDataForAgent(agent.id);}}
                        isActive={selectedAgent.id === agent.id}
                        className="w-full justify-start"
                      >
                        <div className="flex items-center justify-between w-full">
                          <span className="truncate">{agent.name}</span>
                          <div className="flex items-center gap-2">
                            <Badge variant={agent.status === "active" ? "default" : "secondary"}>
                              {agent.clientCount}
                            </Badge>
                            <div
                              className={cn(
                                "w-2 h-2 rounded-full",
                                agent.status === "active" ? "bg-green-500" : "bg-gray-400",
                              )}
                            />
                          </div>
                        </div>
                      </SidebarMenuButton>
                    </SidebarMenuItem>
                  ))}
                </SidebarMenu>
              </SidebarGroupContent>
            </SidebarGroup>
          </SidebarContent>
        </Sidebar>

        {/* Main Content */}
        <SidebarInset className="flex-1">
          <div className="p-6 space-y-6">
            {/* Dashboard Summary */}
            <Card>
                <CardHeader className="flex flex-col items-center justify-center">
                    <CardTitle className="flex items-center gap-2 justify-center">
                        <Activity className="h-5 w-5" />
                        <span className="text-center">Dashboard Summary - {selectedAgent.name}</span>
                    </CardTitle>
                </CardHeader>
                <CardContent>
                    <div className="flex justify-center gap-12">
                        <div className="text-center">
                            <div className="text-2xl font-bold text-blue-600">{selectedAgent.clientCount}</div>
                            <div className="text-sm text-gray-500">Active Clients</div>
                        </div>
                        <div className="text-center">
                            <div className="text-2xl font-bold text-purple-600">{totalRevenue}$</div>
                            <div className="text-sm text-gray-500">Total Revenue</div>
                        </div>
                    </div>
                </CardContent>
            </Card>

            {/* KPI Cards */}
            <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
              {/* Symmetry Ratio */}
              <Card>
                <CardHeader className="pb-2">
                  <CardTitle className="flex items-center gap-2 text-lg">
                    <MessageSquare className="h-5 w-5" />
                    Symmetry Ratio
                  </CardTitle>
                </CardHeader>
                <CardContent>
                  <ResponsiveContainer width="100%" height={200}>
                    <BarChart data={symmetryData}>
                      <CartesianGrid strokeDasharray="3 3" />
                      <XAxis dataKey="client" />
                      <YAxis />
                      <Tooltip />
                      <Bar dataKey="ratio" fill="#3b82f6" />
                    </BarChart>
                  </ResponsiveContainer>
                </CardContent>
              </Card>

              {/* Average Message Gaps */}
              <Card>
                <CardHeader className="pb-2">
                  <CardTitle className="flex items-center gap-2 text-lg">
                    <Clock className="h-5 w-5" />
                    Avg Message Gaps
                  </CardTitle>
                </CardHeader>
                <CardContent>
                  <ResponsiveContainer width="100%" height={200}>
                    <LineChart data={messageGapData}>
                      <CartesianGrid strokeDasharray="3 3" />
                      <XAxis dataKey="day" />
                      <YAxis />
                      <Tooltip />
                      <Line type="monotone" dataKey="avgGap" stroke="#10b981" strokeWidth={2} />
                    </LineChart>
                  </ResponsiveContainer>
                </CardContent>
              </Card>

              {/* Transaction Ratio */}
              <Card>
                <CardHeader className="pb-2">
                  <CardTitle className="flex items-center gap-2 text-lg">
                    <DollarSign className="h-5 w-5" />
                    Transaction Ratio
                  </CardTitle>
                </CardHeader>
                <CardContent>
                  <ResponsiveContainer width="100%" height={200}>
                    <ScatterChart data={transactionData}>
                    <CartesianGrid strokeDasharray="3 3" />
                    <XAxis dataKey="x" />
                    <YAxis dataKey="y" />
                    <Tooltip
                        formatter={(value, name, props) => {
                        if (name === "y") return [`${(value * 100).toFixed(1)}%`, "Tx Ratio"];
                        return [value, name];
                        }}
                        labelFormatter={(label, payload) => {
                        if (payload && payload.length > 0) {
                            return payload[0].payload.client;
                        }
                        return `Client ${label}`;
                        }}
                    />
                    <Scatter dataKey="y" fill="#8b5cf6">
                        {transactionData.map((entry, index) => (
                        <Cell key={`cell-${index}`} fill="#8b5cf6" />
                        ))}
                    </Scatter>
                    </ScatterChart>
                  </ResponsiveContainer>
                </CardContent>
              </Card>
            </div>

            {/* Client Performance Table */}
            <Card>
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  <Users className="h-5 w-5" />
                  Client Performance Table
                </CardTitle>
              </CardHeader>
              <CardContent>
                <Table>
                  <TableHeader>
                    <TableRow>
                      <TableHead>Name</TableHead>
                      <TableHead>Score</TableHead>
                      <TableHead>Symmetry Ratio</TableHead>
                      <TableHead>Avg Msg Gap</TableHead>
                      <TableHead>Tx Ratio</TableHead>
                      <TableHead>Actions</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {clientPerformanceData.map((client) => (
                      <TableRow key={client.id}>
                        <TableCell className="font-medium">{client.name}</TableCell>
                        <TableCell>
                          <Badge variant={getScoreBadgeVariant(client.score)}>{client.score}</Badge>
                        </TableCell>
                        <TableCell>
                          <div className="w-full">
                            <Progress value={client.symmetryRatio * 50} className="w-16" />
                            <span className="text-xs text-gray-500 ml-2">{client.symmetryRatio.toFixed(1)}</span>
                          </div>
                        </TableCell>
                        <TableCell>{client.avgMsgGap}h</TableCell>
                        <TableCell>
                          <Badge variant="outline">{(client.txRatio * 100).toFixed(1)}%</Badge>
                        </TableCell>
                        <TableCell>
                          <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => setExpandedClient(expandedClient === client.id ? null : client.id)}
                          >
                            {expandedClient === client.id ? (
                              <ChevronDown className="h-4 w-4" />
                            ) : (
                              <ChevronRight className="h-4 w-4" />
                            )}
                          </Button>
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </CardContent>
            </Card>

            {/* Timeline Graph + Message Explorer */}
            <Card>
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  <TrendingUp className="h-5 w-5" />
                  Timeline Graph of Most Active Client
                </CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  {/* Timeline visualization */}
                  <div className="border rounded-lg p-4 bg-gray-50">
                    <h4 className="font-medium mb-4">Conversation Flow</h4>
                    <div className="space-y-2">
                      {timelineData.map((item, index) => (
                        <div key={index} className="flex items-center gap-4">
                          <div className="text-xs text-gray-500 w-12">{item.time}</div>
                          <div
                            className={cn(
                              "w-3 h-3 rounded-full",
                              item.sender === "agent" ? "bg-blue-500" : "bg-green-500",
                            )}
                          />
                          <div className="flex-1">
                            <div className="text-sm">
                              <span
                                className={cn(
                                  "font-medium",
                                  item.sender === "agent" ? "text-blue-600" : "text-green-600",
                                )}
                              >
                                {item.sender === "agent" ? "Agent" : "Client"}:
                              </span>
                              <span className="ml-2">{item.message}</span>
                            </div>
                          </div>
                        </div>
                      ))}
                    </div>
                  </div>

                  {/* Legend */}
                  <div className="flex items-center gap-6 text-sm">
                    <div className="flex items-center gap-2">
                      <div className="w-3 h-3 rounded-full bg-blue-500" />
                      <span>Agent</span>
                    </div>
                    <div className="flex items-center gap-2">
                      <div className="w-3 h-3 rounded-full bg-green-500" />
                      <span>Client</span>
                    </div>
                
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </SidebarInset>
       </div>                 

      </div>
    </SidebarProvider>
  )
}

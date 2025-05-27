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

import ApiContext from './ApiContext';
import { useContext } from "react";

// Mock data
const agents = [
  { id: 1, name: "Sales Agent Alpha", status: "active", clientCount: 12 },
  { id: 2, name: "Support Agent Beta", status: "active", clientCount: 8 },
  { id: 3, name: "Marketing Agent Gamma", status: "inactive", clientCount: 15 },
  { id: 4, name: "Customer Success Delta", status: "active", clientCount: 6 },
]

const symmetryData = [
  { client: "Client A", ratio: 0.8 },
  { client: "Client B", ratio: 1.2 },
  { client: "Client C", ratio: 0.6 },
  { client: "Client D", ratio: 1.5 },
  { client: "Client E", ratio: 0.9 },
]

const messageGapData = [
  { day: "Mon", avgGap: 2.5 },
  { day: "Tue", avgGap: 3.2 },
  { day: "Wed", avgGap: 1.8 },
  { day: "Thu", avgGap: 4.1 },
  { day: "Fri", avgGap: 2.9 },
  { day: "Sat", avgGap: 5.2 },
  { day: "Sun", avgGap: 6.1 },
]

const transactionData = [
  { client: "Client A", ratio: 0.15, amount: 5000, x: 1, y: 0.15 },
  { client: "Client B", ratio: 0.25, amount: 8000, x: 2, y: 0.25 },
  { client: "Client C", ratio: 0.08, amount: 2000, x: 3, y: 0.08 },
  { client: "Client D", ratio: 0.35, amount: 12000, x: 4, y: 0.35 },
  { client: "Client E", ratio: 0.18, amount: 6500, x: 5, y: 0.18 },
]

const clientPerformanceData = [
  {
    id: 1,
    name: "Client A",
    score: 85,
    symmetryRatio: 0.8,
    avgMsgGap: 2.5,
    txRatio: 0.15,
    status: "excellent",
  },
  {
    id: 2,
    name: "Client B",
    score: 72,
    symmetryRatio: 1.2,
    avgMsgGap: 3.2,
    txRatio: 0.25,
    status: "good",
  },
  {
    id: 3,
    name: "Client C",
    score: 58,
    symmetryRatio: 0.6,
    avgMsgGap: 1.8,
    txRatio: 0.08,
    status: "needs attention",
  },
  {
    id: 4,
    name: "Client D",
    score: 91,
    symmetryRatio: 1.5,
    avgMsgGap: 4.1,
    txRatio: 0.35,
    status: "excellent",
  },
  {
    id: 5,
    name: "Client E",
    score: 67,
    symmetryRatio: 0.9,
    avgMsgGap: 2.9,
    txRatio: 0.18,
    status: "good",
  },
]

const timelineData = [
  { time: "09:00", sender: "agent", type: "normal", message: "Good morning! How can I help?" },
  { time: "09:15", sender: "client", type: "normal", message: "Hi, I have a question about pricing" },
  { time: "09:18", sender: "agent", type: "normal", message: "I'd be happy to help with that" },
  { time: "09:45", sender: "client", type: "transaction", message: "I'd like to proceed with the purchase" },
  { time: "10:00", sender: "agent", type: "transaction", message: "Great! Let me process that for you" },
  { time: "10:30", sender: "client", type: "emotional", message: "Thank you so much for your help!" },
]

export default function AIAgentDashboard({onBack}) {
  const [selectedAgent, setSelectedAgent] = useState(agents[0])
  const [expandedClient, setExpandedClient] = useState(null)
  const { getClients, getAgents, getConversations } = useContext(ApiContext);

    useEffect(() => { 
         
      }, []);


  const getScoreBadgeVariant = (score) => {
    if (score >= 80) return "default"
    if (score >= 60) return "secondary"
    return "destructive"
  }


  const getStatusColor = (status) => {
    switch (status) {
      case "excellent":
        return "text-green-600"
      case "good":
        return "text-blue-600"
      case "needs attention":
        return "text-orange-600"
      default:
        return "text-gray-600"
    }
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
                        onClick={() => setSelectedAgent(agent)}
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
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  <Activity className="h-5 w-5" />
                  Dashboard Summary - {selectedAgent.name}
                </CardTitle>
              </CardHeader>
              <CardContent>
                <div className="grid grid-cols-3 gap-4">
                  <div className="text-center">
                    <div className="text-2xl font-bold text-blue-600">{selectedAgent.clientCount}</div>
                    <div className="text-sm text-gray-500">Active Clients</div>
                  </div>
                  <div className="text-center">
                    <div className="text-2xl font-bold text-green-600">94%</div>
                    <div className="text-sm text-gray-500">Response Rate</div>
                  </div>
                  <div className="text-center">
                    <div className="text-2xl font-bold text-purple-600">$45K</div>
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
                        formatter={(value, name) => {
                          if (name === "y") return [`${((value) * 100).toFixed(1)}%`, "Tx Ratio"]
                          return [value, name]
                        }}
                        labelFormatter={(label) => `Client ${label}`}
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
                  Timeline Graph + Message Explorer
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
                          <div
                            className={cn(
                              "w-2 h-2 rounded-full",
                              item.type === "normal"
                                ? "bg-gray-400"
                                : item.type === "transaction"
                                  ? "bg-yellow-500"
                                  : "bg-red-500",
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
                    <div className="flex items-center gap-2">
                      <div className="w-2 h-2 rounded-full bg-gray-400" />
                      <span>Normal</span>
                    </div>
                    <div className="flex items-center gap-2">
                      <div className="w-2 h-2 rounded-full bg-yellow-500" />
                      <span>Transaction</span>
                    </div>
                    <div className="flex items-center gap-2">
                      <div className="w-2 h-2 rounded-full bg-red-500" />
                      <span>Emotional</span>
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

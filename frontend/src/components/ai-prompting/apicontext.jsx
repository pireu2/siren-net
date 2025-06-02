import React, { createContext, useContext } from "react";
import { getCookie } from "react-use-cookie";

const ApiContext = createContext();

export function ApiProvider({ children }) {
  const token = getCookie('token');

  const getClients = async (agentId) => {
    if (!agentId) return [];
    try {
      const response = await fetch(`/clients/agent/${agentId}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
      });
      return await response.json();
    } catch (error) {
      console.error("Error making client request:", error);
      return [];
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
      return await response.json();
    } catch (error) {
      console.error("Error making agent request:", error);
      return [];
    }
  };

  const getConversations = async (agentId, clientId) => {
    if (!agentId || !clientId) return [];
    try {
      const response = await fetch(
        `/messages/agent/${agentId}/client/${clientId}`,
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
          },
        }
      );
      return await response.json();
    } catch (error) {
      console.error("Error fetching conversations:", error);
      return [];
    }
  };


  const getDeepSeekResponse = async (promptText) => {
    try {
      const response = await fetch(
        `/llm/ask`,
        {
          method: "POST",
          body: JSON.stringify({
            prompt: promptText
          }),
        }
      );
      return await response.json();
    } catch (error) {
      console.error("Error fetching conversations:", error);
      return [];
    }
  };

  const getTransactions = async (agentId, clientId) => {
  if (!agentId || !clientId) return [];
  
  try {
    const response = await fetch(
      `/transactions/agent/${agentId}/client/${clientId}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
      }
    );
    
    // Check response status first
    if (!response.ok) {
      console.error(`Transaction API error: ${response.status} ${response.statusText}`);
      return [];
    }
    
    // Try to parse as JSON safely
    const contentType = response.headers.get("content-type");
    if (contentType && contentType.includes("application/json")) {
      return await response.json();
    } else {
      // If not JSON, log the issue
      const text = await response.text();
      console.error("API returned non-JSON response:", text.substring(0, 100) + "...");
      return [];
    }
  } catch (error) {
    console.error("Error fetching transactions:", error);
    return [];
  }
};


  

  return (
    <ApiContext.Provider value={{ getClients, getAgents, getConversations, getTransactions,getDeepSeekResponse }}>
      {children}
    </ApiContext.Provider>
  );
}

export default ApiContext;
export function useApi() {
  return useContext(ApiContext);
}
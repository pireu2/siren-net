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

  return (
    <ApiContext.Provider value={{ getClients, getAgents, getConversations }}>
      {children}
    </ApiContext.Provider>
  );
}

export default ApiContext;
export function useApi() {
  return useContext(ApiContext);
}
import React, { useState } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import AuthShuffle from './AuthSwitcher';
import Dashboard from './dashboard-container/Dashboard.jsx'
import AuthProvider,{ AuthContext } from "./auth/auth-handler";
import { useEffect} from "react";
import { getCookie } from 'react-use-cookie';

function App() {
  return (
    <AuthProvider>
      <AppRoutes/>
    </AuthProvider>
  );
}


function AppRoutes()
{
  const token = getCookie('token');
  const [loading, setLoading] = useState(true);


  useEffect(() => {
    if (loading) {
      setLoading(false);
    }
  }, []);

  const condition = token ? <Dashboard /> : <AuthShuffle />;

  if (loading) {
    return null;
  }

  return (
    <>
    <BrowserRouter>
      <Routes>
          <Route path="/" element={condition} />
          <Route path="/dashboard" element={condition} />
      </Routes>
    </BrowserRouter>
    </>
  );
}

export default App;

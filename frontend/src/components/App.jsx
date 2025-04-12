import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import AuthShuffle from './AuthSwitcher';
import Dashboard from './dashboard-container/Dashboard.jsx'
import AuthProvider,{ AuthContext } from "./auth/auth-handler";
import { useContext} from "react";

function App() {
  return (
    <AuthProvider>
      <AppRoutes/>
    </AuthProvider>
  );
}


function AppRoutes()
{
  const { userToken } = useContext(AuthContext);
  const condition = userToken ? <Dashboard /> : <AuthShuffle />;
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

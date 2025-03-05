import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { LoginForm } from './login-form';
import Register from './Register.jsx';
import AuthShuffle from './AuthSwitcher';
import Dashboard from './dashboard-container/Dashboard.jsx'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<AuthShuffle />} />
        <Route path="/dashboard" element={<Dashboard />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;

import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { LoginForm } from './login-form.jsx';
import Register from './Register.jsx';
import AuthShuffle from './AuthSwitcher.jsx';
import Dashboard from './dashboard-container/Dashboard.jsx'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        {/* <Route path="/" element={<AuthShuffle />} /> */}
        <Route path="/" element={<Dashboard />} />
        <Route path="/register" element={<Register />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;

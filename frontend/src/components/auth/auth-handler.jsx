import { createContext, useEffect } from "react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

export const AuthContext  = createContext();


export default function AuthProvider({children})
{
    const [user, setUser] = useState(null);


    const [isOpen, setOpen] = useState(false);
    const[message , setMessage]= useState('');
    
      
    useEffect(() =>
        {
            const token = localStorage.getItem('token');
            if(token)
            {
                console.log("It has a token!");
                setUser({token});
            }
        },
    []);


    const register = async (formData) => {
        console.log("Form Data:", formData);
    
        try {
          const response = await fetch("/auth/register", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(formData),
          });
    
          const data = await response.json();
          if (!response.ok) {
            setMessage(`${data.error.charAt(0).toUpperCase() + data.error.slice(1)}.`);
            setOpen(true);
          }
        } 
        
        catch (error) {
          console.error("Error:", error);
        }
    
      };

    const login = (token) =>
    {
        localStorage.setItem('token',token);
        setUser({token});
    }

    const logout = () =>
        {
            localStorage.removeItem('token');
            setUser(null);
            window.location.href = "/";
        }


    return(
        <>
        <AuthContext.Provider value={{user, login, logout, register, setMessage, isOpen}}>
            {children}
        </AuthContext.Provider>
        </>
    );
}
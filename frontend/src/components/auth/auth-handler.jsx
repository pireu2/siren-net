import { createContext, useEffect } from "react";
import { useState } from "react";
import useCookie from 'react-use-cookie';

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

    const login = async (formData) => {
      try {
        const response = await fetch("auth/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(formData),
        });
  
        const data = await response.json();

        if (response.ok) {
          const token = data.token;
          localStorage.setItem('token',token);
          setUser({token});
          window.location.reload();
        }
        else {
          console.log(data);
          setOpen(true);
          setMessage(`${data.error.charAt(0).toUpperCase() + data.error.slice(1)}.`);
        }
      } 
      
      catch (error) {
        console.error("Error:", error);
      }
  
    };



    const logout = () =>
        {
            localStorage.removeItem('token');
            setUser(null);
            window.location.href = "/";
        }


    return(
        <>
        <AuthContext.Provider 
        value={{user, login, logout, register, setMessage, isOpen,setOpen,message}}>
            {children}
        </AuthContext.Provider>
        </>
    );
}
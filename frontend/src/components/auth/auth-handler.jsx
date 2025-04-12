import { createContext, useEffect } from "react";
import { useState } from "react";
import useCookie,{ getCookie } from 'react-use-cookie';

export const AuthContext  = createContext();


export default function AuthProvider({children})
{
    const [userToken, setUserToken, removeUserToken] = useCookie('token', 0);

    const [isOpen, setOpen] = useState(false);
    const[message , setMessage]= useState('');
    
      
    useEffect(() =>
        {
            const token = getCookie('token');
            if(token)
            {
                console.log("It has a token!");
                setUserToken(token);
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
            return ("Error:" + error);
          }
        } 
        
        catch (error) {
          return ("Error:" + error);
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
          alert(`${formData.username} has logged in!`);
          const token = data.token;
          setUserToken(token);
          window.location.href = "/dashboard";
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
            removeUserToken();
            window.location.href = "/";
        }


    return(
        <>
        <AuthContext.Provider 
        value={{login, logout, register, setMessage, isOpen,userToken,setOpen,message}}>
            {children}
        </AuthContext.Provider>
        </>
    );
}
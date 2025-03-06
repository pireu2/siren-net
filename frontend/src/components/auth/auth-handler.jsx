import { createContext, useEffect } from "react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

export const AuthContext  = createContext();


export default function AuthProvider({children})
{
    const [user, setUser] = useState(null);

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


    const login = (token) =>
    {
        localStorage.setItem('token',token);
        setUser({token});
    }

    const logout = () =>
        {
            localStorage.removeItem('token');
            setUser(null);
        }


    return(
        <>
        <AuthContext.Provider value={{user,login,logout}}>
            {children}
        </AuthContext.Provider>
        </>
    );
}
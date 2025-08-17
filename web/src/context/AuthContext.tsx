import React, { createContext, useContext, useEffect, useState } from 'react'
import { api } from '../lib/api'

type User = { id: string; email: string } | null
type Ctx = {
  user: User
  login: (email: string, password: string) => Promise<void>
  logout: () => Promise<void>
}
const AuthContext = createContext<Ctx | null>(null)

export const AuthProvider: React.FC<{children: React.ReactNode}> = ({ children }) => {
  const [user, setUser] = useState<User>(null)

  useEffect(() => {
    api<User>('/api/auth/me').then(setUser).catch(() => setUser(null))
  }, [])

  async function login(email: string, password: string) {
    await api('/api/auth/login', { method: 'POST', body: JSON.stringify({ email, password }) })
    const me = await api<User>('/api/auth/me')
    setUser(me)
  }

  async function logout() {
    await api('/api/auth/logout', { method: 'POST' })
    setUser(null)
  }

  return <AuthContext.Provider value={{ user, login, logout }}>{children}</AuthContext.Provider>
}

export function useAuth() {
  const ctx = useContext(AuthContext)
  if (!ctx) throw new Error('useAuth must be used within AuthProvider')
  return ctx
}
